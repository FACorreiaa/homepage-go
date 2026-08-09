# gopher.glb

## The model

`gopher.glb` is derived from the Go Gopher 3D model at
<https://github.com/cattaka/go-gopher-model>, released under
**CC0 1.0 Universal** (public domain dedication). No attribution is legally
required for the model itself; it is recorded here because knowing where an
asset came from matters more than the licence minimum.

The widely-linked original of that repo, `justinribeiro/go-gopher-model`, is
now a 404 — `cattaka`'s is the surviving copy.

## The character

The Go Gopher mascot was created by **Renée French** and is licensed under the
[Creative Commons 4.0 Attribution License](https://creativecommons.org/licenses/by/4.0/).
That licence **does** require attribution, so the site credits her in the
footer. See <https://go.dev/brand>.

## What is deliberately not here

The **Go logo** and the stylised "GO" wordmark are trademarks of Google and are
*not* covered by either licence above. From <https://go.dev/brand>:

> Do not use the Go Logo or Go as a stylized form without permission.

Two of the reference images this scene was built from contain it. Neither is
reproduced: the laptop screen in the scene shows this site's own FC mark.

## Regenerating

`convert.py` in this directory is the one-time conversion, kept for the next
time the mesh needs rebuilding. It is **not** part of any build — the Docker
image has no Python and no Node.

```
python3 -m venv .venv
./.venv/bin/pip install trimesh numpy fast-simplification scipy pillow
curl -LO https://raw.githubusercontent.com/cattaka/go-gopher-model/master/go_gopher_high.obj
mv go_gopher_high.obj gopher.obj
./.venv/bin/python convert.py
```

Notes for whoever runs it next:

- The `3d-printable/*.stl` files in that repo are **laid out for a print bed,
  not assembled** — the "eyes" file is a flat tray of parts sitting at the feet.
  Use the high-res OBJ, which is the assembled model.
- That OBJ declares one object with one grey material, but splits into 19
  connected components. `convert.py` groups them into `body`, `white` and
  `dark` so the scene can colour a cyan body, white eyes and teeth, and dark
  pupils and nose without shipping a texture.
- Normals are deliberately **not** exported. three computes them on load, which
  costs nothing and keeps about a third of the file out of the repo.
- Output is normalised: one unit tall, feet at y=0, centred on X/Z, facing +Z.
