#!/usr/bin/env bash
# Turns the masters in out/ into the files the home page serves, in
# assets/static/promo/. Run after `npm run render -- all`.
#
# The budget is the point: this plays for every visitor, so each cut must stay
# small enough to be invisible on a phone connection. AV1 does the heavy
# lifting; H.264 is the fallback for Safari on hardware without an AV1 decoder.
set -euo pipefail
export SVT_LOG=1 # errors only
cd "$(dirname "$0")"
dst=../assets/static/promo
mkdir -p "$dst"

# Poster: the "Four products" wall, the frame that says the most on its own.
POSTER_T=11.7

cut() { # name master scale
  local name=$1 src=$2 vf=$3
  ffmpeg -v error -y -i "$src" -vf "$vf" -an -c:v libsvtav1 -preset 4 -crf 42 \
    -svtav1-params tune=0 -pix_fmt yuv420p -g 150 "$dst/reel-$name.webm"
  ffmpeg -v error -y -i "$src" -vf "$vf" -an -c:v libx264 -preset veryslow -crf 27 \
    -profile:v high -pix_fmt yuv420p -g 150 -movflags +faststart "$dst/reel-$name.mp4"
  # Homebrew's ffmpeg has no libwebp; cwebp (brew install webp) does the poster.
  ffmpeg -v error -y -ss "$POSTER_T" -i "$src" -vf "$vf" -frames:v 1 "out/poster-$name.png"
  cwebp -quiet -q 72 "out/poster-$name.png" -o "$dst/poster-$name.webp"
}

cut desktop out/reel-desktop-1920x1080.mp4 "scale=1920:1080"
cut mobile out/reel-iphone-1170x2532.mp4 "scale=720:1558:flags=lanczos"

ls -lh "$dst"
