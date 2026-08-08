#!/usr/bin/env python3
"""Rasterise Natural Earth land polygons into the globe's land mask.

The /stats globe builds its dot-matrix earth by sampling this mask, so the
image is a build input rather than artwork. Regenerate it with:

    curl -sSLo /tmp/ne_110m_land.zip \\
        https://naciscdn.org/naturalearth/110m/physical/ne_110m_land.zip
    python3 scripts/build-land-mask.py /tmp/ne_110m_land.zip

Source data is Natural Earth 1:110m physical land, which is public domain.
Requires Pillow; nothing at runtime depends on this script.
"""

import struct
import sys
import zipfile
from pathlib import Path

from PIL import Image, ImageDraw

# Rasterise large, then box-filter down. The extra resolution keeps small
# islands from disappearing when the mask is reduced.
SUPERSAMPLE = 4
OUT_WIDTH = 1024
OUT_HEIGHT = 512
OUTPUT = Path("assets/static/globe/land-mask.png")

POLYGON = 5


def read_polygons(shp: bytes):
    """Yield each ring of every polygon record as a list of (lon, lat)."""
    offset = 100  # fixed-size shapefile header
    total = len(shp)

    while offset < total:
        _, content_len = struct.unpack_from(">ii", shp, offset)
        offset += 8
        record_end = offset + content_len * 2

        (shape_type,) = struct.unpack_from("<i", shp, offset)
        if shape_type != POLYGON:
            offset = record_end
            continue

        num_parts, num_points = struct.unpack_from("<ii", shp, offset + 36)
        parts_at = offset + 44
        points_at = parts_at + num_parts * 4

        parts = list(struct.unpack_from("<%di" % num_parts, shp, parts_at))
        coords = struct.unpack_from("<%dd" % (num_points * 2), shp, points_at)
        parts.append(num_points)

        for i in range(num_parts):
            start, end = parts[i], parts[i + 1]
            ring = [(coords[j * 2], coords[j * 2 + 1]) for j in range(start, end)]
            if len(ring) >= 3:
                yield ring

        offset = record_end


def signed_area(ring):
    total = 0.0
    for i in range(len(ring)):
        x1, y1 = ring[i]
        x2, y2 = ring[(i + 1) % len(ring)]
        total += x1 * y2 - x2 * y1
    return total / 2.0


def to_pixels(ring, width, height):
    """Equirectangular projection: lon -180..180 → 0..w, lat 90..-90 → 0..h."""
    return [
        ((lon + 180.0) / 360.0 * width, (90.0 - lat) / 180.0 * height)
        for lon, lat in ring
    ]


def main():
    if len(sys.argv) != 2:
        sys.exit("usage: build-land-mask.py <ne_110m_land.zip>")

    with zipfile.ZipFile(sys.argv[1]) as archive:
        name = next(n for n in archive.namelist() if n.endswith(".shp"))
        shp = archive.read(name)

    width, height = OUT_WIDTH * SUPERSAMPLE, OUT_HEIGHT * SUPERSAMPLE
    canvas = Image.new("L", (width, height), 0)
    draw = ImageDraw.Draw(canvas)

    outer = holes = 0
    for ring in read_polygons(shp):
        pixels = to_pixels(ring, width, height)
        # Shapefile spec: outer rings wind clockwise, holes counter-clockwise.
        if signed_area(ring) < 0:
            draw.polygon(pixels, fill=255)
            outer += 1
        else:
            draw.polygon(pixels, fill=0)
            holes += 1

    mask = canvas.resize((OUT_WIDTH, OUT_HEIGHT), Image.LANCZOS)
    mask = mask.point(lambda v: 255 if v >= 96 else 0).convert("1")

    OUTPUT.parent.mkdir(parents=True, exist_ok=True)
    mask.save(OUTPUT, optimize=True, bits=1)

    land = sum(1 for p in mask.convert("L").tobytes() if p)
    print(
        "wrote %s — %d rings (%d holes), %.1f%% land, %d bytes"
        % (OUTPUT, outer, holes, 100.0 * land / (OUT_WIDTH * OUT_HEIGHT), OUTPUT.stat().st_size)
    )


if __name__ == "__main__":
    main()
