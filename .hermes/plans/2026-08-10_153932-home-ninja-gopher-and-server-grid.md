# Plan (v2): Ninja gopher + server-grid — polish & ship

**Source plan:** `.hermes/plans/2026-08-10_153932-home-ninja-gopher-and-server-grid.md`  
**Status:** Core implementation already exists as **uncommitted WIP** in `assets/static/gopher.js` + `assets/static/home-world.js` (+223 / −57). This revision re-scopes from “build it” → **“close the silhouette gap vs reference, verify QA, ship.”**

**Reference (Image #1):** `~/Downloads/IMG_5927.PNG` (also session `images/image-23866437-…png`)

---

## Goal (unchanged intent, tighter ship bar)

1. Fourth home-world persona: **ninja**, readable at a glance from the reference silhouette — not merely “cyan gopher + props.”
2. Server station background reads as **infra/network**, not equal-weight wireframe cube wallpaper.
3. Preserve progressive enhancement, theme repaint, scroll/camera system, and footer attribution.

---

## Reference read (Image #1) — what “ninja” must communicate

Phone screenshot of a stylized Go gopher:

| Signal | In the PNG | Why it matters in 3D |
|---|---|---|
| **Full dark gi** | Navy wraps the *whole* body; only muzzle, ears, paws, foot tips stay cyan | Silhouette is costume-first. A bare cyan body + headband will read as “gopher with accessories,” not ninja |
| **Hood + scarf** | Cap over crown, knot/tail on viewer-left, ears poke through | Head costume without a hood cap looks incomplete from camera angles |
| **Eye openings** | White angry eyes with dark pupils visible *through* the mask | Solid box wrap that buries mesh eyes loses the character’s face |
| **Cross-chest wrap** | Diagonal torso banding | Cheap primitive that sells “gi” without recoloring the mesh texture |
| **Katana** | Thin blade, dark grip, nearly vertical in right paw | Already in WIP; keep restrained micro-sway |
| **Body colour** | Cyan gopher under dark cloth | Keep body `--go`; dark gear is `fg` overlays — **do not** full-body recolor the GLB |

**Design dials (still redesign-preserve):**

- `DESIGN_VARIANCE: 6` — evolve props + node vocabulary only  
- `MOTION_INTENSITY: 6` — no new perpetual gimmicks  
- `VISUAL_DENSITY: 3` — special nodes sparse; falloff protects fork-panel text  

---

## Current state (truth on disk)

### Already implemented (WIP)

| Item | Where | Notes |
|---|---|---|
| `ninja` persona + JSDoc union | `gopher.js` | Branch after professor |
| `buildNinjaMask` | eye wrap + jaw + scarf tail/tip | Head-only; **no hood cap, no eye holes, no torso gi** |
| `buildKatana` | blade + guard + handle + cone tip | Right paw, mild tick sway |
| Horizon placement R1 | `buildHorizon(..., character)` | `pos (1.4, -4.35, 0.6)`, scale `1.6`, `baseY` set |
| Layered server field | `buildServer` | box / octa / nested-inner / core icosa + center hole + falloff + seeded hash + core-biased traffic |
| Theme roles | materials use `userData.role` | `applyPalette` scene-walk still works |

### Server field math (current thresholds)

~387 nodes after `CLEAR_R = 2.55` hole on 20×20:

| Kind | Approx share | Matches plan? |
|---|---|---|
| box (+ nested outer) | ~48% | Base lattice a bit low vs planned ~70% — **acceptable** if fork panels read |
| octa | ~29% | High vs plan’s residual; OK if still quiet |
| nested (inner cyan octa) | ~13% | On target |
| core | ~11% (~42) | On target (≤50) |

### Gaps vs success criteria

1. **Ninja silhouette incomplete** — reference is full-body gi; WIP is head wraps + sword only. Highest remaining product risk.  
2. **Mask may bury eyes** — solid `BoxGeometry` wraps with no openings; reference *depends* on visible angry eyes.  
3. **Dead code / stale comments** — `group.userData.serverMaterials` is written but never read; `gopher.js` file header still says “three characters / cloned three times”; `GRID_COLS` comment still implies full 400 instances.  
4. **No visual QA evidence yet** — light/dark, fork readability, reduced-motion, mobile bail-out not closed.  
5. **Optional docs ref PNG** not committed.

---

## Proposed approach (v2)

### A. Ninja costume polish (priority 1 — still no new GLB)

Stay on the procedural prop path. Expand gear so the camera at horizon reads **costume**, not bare mesh.

**Must-do props (beyond current mask+katana):**

1. **Hood cap** — flattened hemisphere or box+sphere strip over crown (`y ≈ 0.9–1.0 × scale`), ears left free (do not cover ear tips).  
2. **Eye openings** — either:
   - **Preferred:** thinner eye-band with two cutouts (two short boxes left/right of midline, gap over pupils), *or*
   - two thin torus/ring frames like professor glasses but solid cloth panels with a center gap  
   Goal: mesh white eyes + dark pupils still visible.  
3. **Torso gi wrap** (was “optional”; now **required for ship** given reference):
   - 2–3 dark `BoxGeometry` bands around belly/torso (diagonal cross like the PNG)
   - optional short belt strip at waist  
   - keep pale `go-belly` partially visible or slightly covered — either is fine if silhouette is dark-forward  
4. **Paw cuffs (optional XS)** — tiny dark rings at wrist if arms still scream “naked cyan.”

**Keep as-is:**

- Body mesh stays `--go` (on-brand with other personas).  
- Gear materials: `palette.fg`, opacity ~0.88–0.95 (theme-safe; no pure `#000`).  
- Steel blade: slightly lower opacity or same `fg` — do **not** invent a new palette token unless steel disappears in light mode.  
- Tick: slow yaw + light bob + blade micro-sway only.

**Still out of scope:** baking `ninja-gopher.glb`, PNG billboard, AI mesh-from-image.

**Fallback if props still fail visual QA:** document “readable costume, not illustration clone” as accepted; only then open a later mesh pass.

### B. Placement — locked

**Horizon ninja (R1)** stays. Narrative: production shipped / perimeter watch. Do not replace runner; do not add stations.

Pose check during QA:

- Feet clear of drifting grid (`y` already −4.35)  
- Not clipped by horizon line  
- Readable at `camZ: 16` — if lost, nudge scale `1.6 → 1.75` or move closer on Z before changing camera

### C. Server grid — treat as feature-complete pending QA

Do **not** rebuild again unless QA fails. Tuning knobs only:

| Symptom | First lever | Second lever |
|---|---|---|
| Fork panels still fight lattice | lower box/octa opacity 0.12/0.11 → ~0.09/0.08 | raise falloff aggressiveness |
| Still “cube demo” | bump nested/core visibility slightly | — |
| Too busy / noisy | drop `GRID_COLS` 20 → 16 | increase `CLEAR_R` |
| Runner buried | increase `CLEAR_R` 2.55 → ~3.0 | — |
| Cores invisible | raise core opacity floor | slightly larger icosa |

**Still deferred (phase 2):** product-logo constellation, stack icons.

**Cleanup while touching the file:**

- Remove unused `serverMaterials` array **or** wire it if there was an intended reason (there isn’t — scene traverse already tags materials). Prefer **delete**.  
- Fix comment on `GRID_COLS` to note center hole + multi-mesh counts.

### D. Reference asset hygiene

- Optional: `docs/references/home-world/ninja-gopher-ref.png` (+ one-line README: original site costume props inspired by common gopher costume tropes; mesh remains Cattaka CC0 / Renée French character credit in footer).  
- **Do not** serve from `/assets/static` on the live path.

---

## Step-by-step (remaining work only)

### 1. Costume polish in `gopher.js` (S–M)

1. Extend `buildNinjaMask` → hood cap + eye-gap band (or split into `buildNinjaHood` / keep one group).  
2. Add `buildNinjaGi(THREE, palette)` — torso diagonal wraps + optional belt; attach in `persona === 'ninja'` at belly height (`~0.35–0.55 × scale` region, scale with persona).  
3. Re-check mask Z so it sits proud of pupils without z-fighting (`~0.34–0.38 × scale`).  
4. Update file header comment: four personas, one mesh cloned N times.  
5. Confirm all new mats set `userData.role = 'fg'`.

### 2. Server grid hygiene in `home-world.js` (XS)

1. Delete dead `serverMaterials` assignment.  
2. Touch opacities / `CLEAR_R` / `GRID_COLS` **only if** desktop screenshot QA demands it.  
3. No scroll/station constant changes unless ninja framing forces a tiny `lookY` nudge.

### 3. Visual QA gate (manual, desktop ≥900px)

Light + dark:

- [ ] Device hero contrast intact  
- [ ] Server: fork panels primary; lattice supports, does not wallpaper  
- [ ] Runner still focal at server  
- [ ] Horizon: ninja reads as **costumed** (hood + torso + blade), not bare cyan  
- [ ] Ninja eyes readable through mask  
- [ ] Theme toggle repaints body, belly, gear, wireframes, cores, traffic  
- [ ] `prefers-reduced-motion`: no canvas  
- [ ] `<900px` / saveData: no canvas, layout unchanged  
- [ ] Network: only existing `gopher.glb` + three + modules  
- [ ] No Go trademark wordmark introduced  

### 4. Light automated / static checks

1. `go test ./...` (expect green; no Go changes required).  
2. Grep guard: no `IMG_5927` under served static paths.  
3. Optional: hard-refresh localhost, scroll full story once.

### 5. Ship packaging

- One PR / one commit series covering gopher + home-world (+ optional docs ref).  
- Do not regenerate GLB / do not add Python-Node to Docker.

---

## Files

| File | Status | Remaining |
|---|---|---|
| `assets/static/gopher.js` | WIP ninja head+blade | Hood, eye gaps, torso gi, header comment |
| `assets/static/home-world.js` | WIP grid + horizon wire | Dead-code cleanup; QA-driven knobs only |
| `assets/css/input.css` | untouched | none |
| `ui/pages/home.templ` | untouched | none |
| `ui/modules/footer.templ` | credit OK | none |
| `docs/references/home-world/*` | missing | optional ref PNG + README |
| `assets/static/models/*` | no new GLB | none |

---

## Implementation order

| P | Work | Effort | Depends |
|---|---|---|---|
| 1 | Ninja silhouette polish (hood, eyes, torso gi) | S–M | — |
| 2 | Comment / dead-code hygiene | XS | — |
| 3 | Visual QA + opacity/hole knobs if needed | S | 1 |
| 4 | Optional docs reference PNG | XS | — |
| 5 | Phase 2 logos (explicitly later) | M | grid stable in prod |

1 and 2 can land in the same edit pass.

---

## Risks & tradeoffs (updated)

| Risk | Mitigation |
|---|---|
| Props still look crude vs polished PNG | Ship “readable costume”; do not chase illustration fidelity with boxes |
| Torso wraps z-fight belly | Slight Z offset forward; lower opacity on belly-adjacent bands |
| Too much dark gear in light theme | Always `palette.fg` with opacity; never hardcoded black |
| Over-tuning grid after already solid pass | Prefer costume polish; only turn grid knobs on evidence |
| License optics | Props original; same mesh credit; no third-party ninja IP |

| Choice | Decision |
|---|---|
| Procedural vs new GLB | **Procedural v1** (locked) |
| Horizon vs replace runner | **Horizon** (locked) |
| Logos in v1 | **No** |
| Full body recolor | **No** — overlays only |
| Torso gi | **Yes for ship** (upgraded from optional) |

---

## Open questions (non-blocking defaults)

1. Commit ref PNG under `docs/references/`? → **default yes if small; else leave on disk**  
2. Phase 2 marks (Norviq / Lumina / Loci / Hermes / North)? → **later, after strangers actually see the page**  
3. Accept “prop costume” vs chase PNG fidelity further? → **default accept after one polish pass**

---

## Out of scope (still)

- New stations / hero copy  
- Bloom / postprocessing  
- Mobile WebGL enablement  
- Cluster / helm redesign  
- Mesh generation from PNG  
- Footer / CTA / project cards  

---

## Success criteria (ship gate)

1. Four personas from one GLB; ninja reads **hood + face wrap + torso gi + katana** at horizon.  
2. Server field is layered (lattice + nested + cores + purposeful traffic), not uniform cubes.  
3. Fork panels / headlines remain primary over the rack.  
4. Bail-outs + theme repaint unchanged and verified.  
5. No new heavy assets on the critical path.  
6. Stale “three characters” comments gone; no dead `serverMaterials`.

---

## Execution notes

- Repo: `~/Work/production/apps/facorreia-site-go`  
- Edit static JS; hard-refresh with cache bypass  
- Do not run `convert.py` unless regenerating GLB on purpose  
- Keep the dry comment voice already in these files  
- After approval: implement remaining polish only; do not re-litigate architecture  

---

## Diff from Hermes plan v1

| v1 | v2 |
|---|---|
| Wrote as greenfield build | Acknowledges WIP already on disk |
| Torso wrap optional | Torso wrap + hood + eye gaps required for ship |
| Server grid full redesign | Grid done; QA knobs only |
| Open placement question | Horizon locked |
| Generic QA | Checklist tied to real knobs (`CLEAR_R`, opacities, scale) |
| No dead-code note | Explicit hygiene for `serverMaterials` + comments |
