"""One-time conversion: the CC0 gopher OBJ -> a web-ready GLB in three parts.

Not part of any build. Run once, commit the GLB, keep this script in the
scratchpad for the next time the model needs regenerating.

The 3d-printable STLs in that repo are laid out for a print bed, not assembled —
the "eyes" file is a flat tray of parts sitting at the feet. The high-res OBJ is
the assembled model, and although it declares a single object with one grey
material, it separates cleanly into 19 connected components. Those get grouped
into three meshes so the scene can colour a cyan body, white eyes and teeth, and
dark pupils and nose without shipping a single texture byte.

Normals are deliberately NOT exported. three computes them on load, which costs
nothing and keeps roughly a third of the file out of the repo.
"""

import sys
from collections import defaultdict

import numpy as np
import trimesh

SRC = "gopher.obj"
OUT = "gopher.glb"

# Decimation targets per group. The body carries the silhouette so it keeps the
# most; the face parts are small on screen and read fine much lighter.
TARGETS = {"body": 9000, "white": 2600, "dark": 1600}


def classify(part):
    """Which of the three materials this component belongs to.

    Rules are written against measured geometry rather than component order,
    because trimesh's ordering has ties (eight 768-face pieces) and would not
    survive a library upgrade.
    """
    faces = len(part.faces)
    centre = part.bounds.mean(axis=0)
    size = part.bounds[1] - part.bounds[0]

    if faces > 20000:
        return "body"  # the whole torso, ears, arms, feet

    on_face = centre[2] > 40  # everything below is a facial feature

    # Eyeballs: the two big spheres.
    if on_face and size[0] > 30 and size[1] > 30:
        return "white"
    # Teeth: the pair low on the muzzle.
    if on_face and centre[1] < 115:
        return "white"
    # Specular glints: tiny flat flecks sitting on the pupils.
    if on_face and max(size) < 4:
        return "white"
    # Pupils: flat discs in front of the eyeballs.
    if on_face and size[2] < 6 and centre[1] > 130:
        return "dark"
    # Nose: the frontmost lump.
    if centre[2] > 68:
        return "dark"

    # Muzzle, brows and whiskers stay body-coloured.
    return "body"


def main():
    mesh = trimesh.load(SRC, force="mesh")
    mesh.merge_vertices()
    parts = mesh.split(only_watertight=False)
    print(f"source: {len(mesh.faces)} faces in {len(parts)} components")

    groups = defaultdict(list)
    for part in parts:
        groups[classify(part)].append(part)

    merged = {}
    for name, members in groups.items():
        combined = trimesh.util.concatenate(members)
        print(f"  {name:<6} {len(members):2d} components, {len(combined.faces):6d} faces")
        merged[name] = combined

    # Normalise: centre on X/Z, feet at y=0, exactly one unit tall, so the scene
    # sizes the gopher with a single scalar and never guesses at scale.
    everything = trimesh.util.concatenate(list(merged.values()))
    lo, hi = everything.bounds
    centre = (lo + hi) / 2.0
    height = hi[1] - lo[1]
    print(f"\nraw bounds {lo.round(1)} .. {hi.round(1)}")

    for name, part in merged.items():
        target = TARGETS[name]
        if len(part.faces) > target:
            before = len(part.faces)
            part = part.simplify_quadric_decimation(face_count=target)
            print(f"  {name:<6} decimated {before} -> {len(part.faces)}")
        part.apply_translation([-centre[0], -lo[1], -centre[2]])
        part.apply_scale(1.0 / height)
        merged[name] = part

    scene = trimesh.Scene()
    for name, part in merged.items():
        part.visual = trimesh.visual.ColorVisuals(part)
        scene.add_geometry(part, geom_name=name, node_name=name)

    glb = scene.export(file_type="glb")
    with open(OUT, "wb") as fh:
        fh.write(glb)

    final = trimesh.util.concatenate(list(merged.values()))
    total = sum(len(p.faces) for p in merged.values())
    print(f"\nnormalised {final.bounds[0].round(3)} .. {final.bounds[1].round(3)}")
    print(f"total {total} faces -> {OUT} {len(glb) / 1024:.1f} KB")
    return 0


if __name__ == "__main__":
    sys.exit(main())
