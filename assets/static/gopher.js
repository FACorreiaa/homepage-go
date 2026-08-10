// The gopher, and the four characters it plays.
//
// One mesh, loaded once, cloned per persona. The personas are placement,
// rotation and props — the model is a 3D-print model with no rig, so nothing
// here bends its arms.
//
// The mesh is the CC0 model from github.com/cattaka/go-gopher-model, converted
// by assets/static/models/convert.py. It arrives one unit tall with its feet at
// y=0, centred on X/Z and facing +Z, so the scene sizes it with one scalar and
// never guesses at scale. See models/LICENSE.md.

import { GLTFLoader } from 'three/examples/jsm/loaders/GLTFLoader.js';

const MODEL_URL = '/assets/static/models/gopher.glb';

// The three meshes in the GLB, named by convert.py. Body follows the theme;
// eyes and pupils are fixed, because a gopher whose eyes invert in light mode
// is a different animal.
const FIXED = { white: 0xf2f4f6, dark: 0x15181d };

let template = null;

/**
 * Loads the mesh once and hands back a template to clone. Normals are computed
 * here rather than shipped: three does it in milliseconds and it keeps about a
 * third of the file out of the repo.
 */
export async function loadGopher(THREE) {
  if (template) return template;

  const gltf = await new GLTFLoader().loadAsync(MODEL_URL);
  gltf.scene.traverse((object) => {
    if (object.isMesh) object.geometry.computeVertexNormals();
  });
  template = gltf.scene;
  return template;
}

/** Paints a clone's three meshes and tags the body so the theme can repaint it. */
function paint(THREE, root, palette) {
  root.traverse((object) => {
    if (!object.isMesh) return;
    const fixed = FIXED[object.name];
    // Opaque on purpose. Transparency here sorts per object, not per triangle,
    // so at 0.9 the far side of the body bled through the near side as muddy
    // patches across the belly. Restraint comes from size and placement.
    const material = new THREE.MeshLambertMaterial({ color: fixed ?? palette.go });
    // Only the body follows a token; see FIXED above.
    if (!fixed) material.userData.role = 'go';
    object.material = material;
  });
}

/**
 * The pale belly, which the mesh does not have and which is the most
 * recognisable thing about the gopher after its eyes. A flattened sphere
 * bulging just proud of the lower front, tinted toward white so it reads in
 * both themes without being a second colour to maintain.
 */
function buildBelly(THREE, palette) {
  // readPalette derives 'go-belly' so the theme repaint reaches this too.
  const colour = palette['go-belly'] ?? palette.go.clone().lerp(new THREE.Color(0xffffff), 0.62);
  const material = new THREE.MeshLambertMaterial({ color: colour });
  material.userData.role = 'go-belly';
  const belly = new THREE.Mesh(new THREE.SphereGeometry(0.235, 24, 18), material);
  belly.scale.set(1.0, 0.94, 0.52);
  belly.position.set(0, 0.315, 0.25);
  return belly;
}

/**
 * Motion marks. Both reference images put speed lines behind the gopher, and
 * the Go logo itself carries them — it is the cheapest way to say "running"
 * for a model that cannot move its legs.
 */
function buildSpeedLines(THREE, palette) {
  const group = new THREE.Group();
  const material = new THREE.MeshLambertMaterial({ color: palette.go });
  material.userData.role = 'go';
  const rows = [
    { y: 0.72, len: 0.5, x: -0.62 },
    { y: 0.52, len: 0.72, x: -0.72 },
    { y: 0.32, len: 0.44, x: -0.6 },
  ];
  for (const row of rows) {
    const bar = new THREE.Mesh(new THREE.BoxGeometry(row.len, 0.035, 0.035), material);
    bar.position.set(row.x, row.y, -0.1);
    group.add(bar);
  }
  return group;
}

/**
 * A laptop, from two boxes. The screen is a canvas texture drawn at runtime —
 * no image bytes, and deliberately the site's own mark rather than the Go logo,
 * which is a Google trademark and not covered by the gopher's licence.
 */
function buildLaptop(THREE, palette) {
  const laptop = new THREE.Group();

  const shell = new THREE.MeshLambertMaterial({ color: palette.fg, transparent: true, opacity: 0.55 });
  shell.userData.role = 'fg';

  const base = new THREE.Mesh(new THREE.BoxGeometry(0.62, 0.03, 0.42), shell);
  base.position.y = 0.015;
  laptop.add(base);

  const lid = new THREE.Mesh(new THREE.BoxGeometry(0.62, 0.4, 0.025), shell);
  lid.position.set(0, 0.2, -0.2);
  lid.rotation.x = -0.28;
  laptop.add(lid);

  const canvas = document.createElement('canvas');
  canvas.width = 256;
  canvas.height = 160;
  const ctx = canvas.getContext('2d');
  ctx.fillStyle = '#0d1117';
  ctx.fillRect(0, 0, 256, 160);
  ctx.fillStyle = '#00add8';
  ctx.font = 'bold 76px ui-monospace, monospace';
  ctx.textAlign = 'center';
  ctx.textBaseline = 'middle';
  ctx.fillText('FC', 128, 84);

  const screen = new THREE.Mesh(
    new THREE.PlaneGeometry(0.54, 0.34),
    new THREE.MeshBasicMaterial({ map: new THREE.CanvasTexture(canvas) }),
  );
  screen.position.set(0, 0.2, -0.185);
  screen.rotation.x = -0.28;
  laptop.add(screen);

  return laptop;
}

/** Round spectacles: two rings, a bridge and two temples. */
function buildGlasses(THREE, palette) {
  const glasses = new THREE.Group();
  const wire = new THREE.MeshLambertMaterial({ color: palette.fg });
  wire.userData.role = 'fg';

  const ring = new THREE.TorusGeometry(0.115, 0.012, 8, 24);
  for (const side of [-1, 1]) {
    const lens = new THREE.Mesh(ring, wire);
    lens.position.set(side * 0.125, 0, 0);
    glasses.add(lens);

    const temple = new THREE.Mesh(new THREE.CylinderGeometry(0.008, 0.008, 0.22, 6), wire);
    temple.position.set(side * 0.225, 0, -0.1);
    temple.rotation.set(Math.PI / 2, 0, 0);
    glasses.add(temple);
  }

  const bridge = new THREE.Mesh(new THREE.CylinderGeometry(0.008, 0.008, 0.03, 6), wire);
  bridge.rotation.z = Math.PI / 2;
  glasses.add(bridge);

  return glasses;
}

/** A pointer stick, held out at an angle. */
function buildPointer(THREE, palette) {
  const material = new THREE.MeshLambertMaterial({ color: palette.fg });
  material.userData.role = 'fg';
  const group = new THREE.Group();
  const stick = new THREE.Mesh(new THREE.CylinderGeometry(0.0075, 0.0075, 0.62, 6), material);
  // Rotate about the lower end so it pivots from the paw, not its middle.
  stick.position.y = 0.31;
  group.add(stick);
  group.rotation.z = -0.55;
  return group;
}


/**
 * Ninja gear from the IMG_5927 reference: hood, open-eye wrap, trailing scarf,
 * torso gi bands, and a thin katana. Same primitive discipline as the glasses
 * and pointer — the mesh has no rig, so costume is placement, not deformation.
 * Dark gear tracks foreground so light mode does not leave a pure-black hole.
 */
function ninjaCloth(THREE, palette, opacity = 0.92) {
  const cloth = new THREE.MeshLambertMaterial({ color: palette.fg, transparent: true, opacity });
  cloth.userData.role = 'fg';
  return cloth;
}

function buildNinjaMask(THREE, palette) {
  const mask = new THREE.Group();
  const cloth = ninjaCloth(THREE, palette, 0.93);

  // Hood cap over the crown. Flattened so the ears still poke free.
  const hood = new THREE.Mesh(new THREE.SphereGeometry(0.28, 16, 12, 0, Math.PI * 2, 0, Math.PI * 0.55), cloth);
  hood.scale.set(1.05, 0.72, 1.0);
  hood.position.set(0, 0.16, -0.02);
  mask.add(hood);

  // Side panels of the eye band — gap on the midline so mesh pupils stay visible.
  for (const side of [-1, 1]) {
    const panel = new THREE.Mesh(new THREE.BoxGeometry(0.2, 0.14, 0.3), cloth);
    panel.position.set(side * 0.16, 0.02, 0.04);
    mask.add(panel);
  }

  // Thin bridge above the eyes so the two panels still read as one wrap.
  const brow = new THREE.Mesh(new THREE.BoxGeometry(0.18, 0.05, 0.28), cloth);
  brow.position.set(0, 0.08, 0.03);
  mask.add(brow);

  // Lower face wrap — gi collar / mouth cover without burying the belly.
  const jaw = new THREE.Mesh(new THREE.BoxGeometry(0.44, 0.13, 0.3), cloth);
  jaw.position.set(0, -0.13, 0.05);
  mask.add(jaw);

  // Scarf knot + tail trailing off the left side of the head (reference side).
  const knot = new THREE.Mesh(new THREE.BoxGeometry(0.12, 0.1, 0.1), cloth);
  knot.position.set(-0.28, 0.04, -0.02);
  knot.rotation.z = 0.25;
  mask.add(knot);

  const tail = new THREE.Mesh(new THREE.BoxGeometry(0.3, 0.08, 0.06), cloth);
  tail.position.set(-0.42, 0.0, -0.06);
  tail.rotation.z = 0.4;
  tail.rotation.y = 0.35;
  mask.add(tail);

  const tailTip = new THREE.Mesh(new THREE.BoxGeometry(0.16, 0.06, 0.05), cloth);
  tailTip.position.set(-0.58, -0.06, -0.1);
  tailTip.rotation.z = 0.75;
  mask.add(tailTip);

  return mask;
}

/**
 * Cross-chest gi bands + a waist belt. Sells the full-body dark silhouette of
 * the reference without recolouring the GLB body mesh.
 */
function buildNinjaGi(THREE, palette) {
  const gi = new THREE.Group();
  const cloth = ninjaCloth(THREE, palette, 0.9);

  // Diagonal wrap, viewer-left shoulder down to right hip.
  const sashA = new THREE.Mesh(new THREE.BoxGeometry(0.12, 0.55, 0.42), cloth);
  sashA.position.set(-0.02, 0.42, 0.08);
  sashA.rotation.z = 0.55;
  gi.add(sashA);

  // Crossing band the other way — thinner so the pale belly still peeks.
  const sashB = new THREE.Mesh(new THREE.BoxGeometry(0.1, 0.5, 0.4), cloth);
  sashB.position.set(0.02, 0.4, 0.06);
  sashB.rotation.z = -0.5;
  gi.add(sashB);

  // Torso panel behind the sashes so the midsection reads dark, not bare cyan.
  const torso = new THREE.Mesh(new THREE.BoxGeometry(0.48, 0.38, 0.36), cloth);
  torso.position.set(0, 0.38, 0.02);
  gi.add(torso);

  // Waist belt.
  const belt = new THREE.Mesh(new THREE.BoxGeometry(0.5, 0.08, 0.4), cloth);
  belt.position.set(0, 0.18, 0.04);
  gi.add(belt);

  return gi;
}

/** Thin katana: blade, guard, handle. Held upright from the right paw. */
function buildKatana(THREE, palette) {
  const katana = new THREE.Group();
  const steel = new THREE.MeshLambertMaterial({ color: palette.fg, transparent: true, opacity: 0.75 });
  steel.userData.role = 'fg';
  const grip = new THREE.MeshLambertMaterial({ color: palette.fg });
  grip.userData.role = 'fg';

  const blade = new THREE.Mesh(new THREE.BoxGeometry(0.018, 0.72, 0.01), steel);
  blade.position.y = 0.42;
  katana.add(blade);

  const guard = new THREE.Mesh(new THREE.BoxGeometry(0.1, 0.016, 0.04), grip);
  guard.position.y = 0.06;
  katana.add(guard);

  const handle = new THREE.Mesh(new THREE.CylinderGeometry(0.014, 0.016, 0.14, 8), grip);
  handle.position.y = -0.02;
  katana.add(handle);

  // Tip cap so the blade does not read as a flat bar end-on.
  const tip = new THREE.Mesh(new THREE.ConeGeometry(0.012, 0.05, 6), steel);
  tip.position.y = 0.8;
  katana.add(tip);

  return katana;
}

/**
 * Builds one gopher in a given role. Returns a group whose userData.tick, if
 * present, animates it — the caller drives it from the scene clock.
 *
 * @param {'coder'|'runner'|'professor'|'ninja'} persona
 */
export function makeGopher(THREE, palette, persona, { scale = 2.6 } = {}) {
  const group = new THREE.Group();

  const model = template.clone(true);
  paint(THREE, model, palette);
  model.add(buildBelly(THREE, palette));
  model.scale.setScalar(scale);
  group.add(model);

  if (persona === 'coder') {
    const laptop = buildLaptop(THREE, palette);
    laptop.scale.setScalar(scale);
    // On the ground in front of the gopher, angled up toward it.
    laptop.scale.multiplyScalar(1.05);
    laptop.position.set(-0.02 * scale, -0.02 * scale, 0.82 * scale);
    // Screen toward the viewer, gopher behind it — the arrangement in the
    // reference. Turning it the physically-correct way round means the mark on
    // it is never seen, and a blank lid is not worth drawing.
    laptop.rotation.y = 0.22;
    group.add(laptop);
    // Leaning in over the keyboard.
    model.rotation.x = 0.16;
    group.userData.tick = (t) => {
      group.rotation.y = -0.55 + Math.sin(t * 0.35) * 0.1;
      model.position.y = Math.sin(t * 2.4) * 0.006 * scale; // typing
    };
  } else if (persona === 'professor') {
    const glasses = buildGlasses(THREE, palette);
    glasses.scale.setScalar(scale);
    // Measured off the source mesh: the eyeballs centre at y=0.77 of the
    // model's height and z=0.27, with the pupils standing proud at z=0.35. The
    // lenses sit just in front of the pupils.
    glasses.position.set(0, 0.77 * scale, 0.37 * scale);
    group.add(glasses);

    const pointer = buildPointer(THREE, palette);
    pointer.scale.setScalar(scale);
    // At the right paw, raised — the reference gopher holds it up, not out.
    // The body is only 0.35 half-wide, so anything past that floats free.
    pointer.position.set(0.325 * scale, 0.4 * scale, 0.26 * scale);
    group.add(pointer);

    group.userData.tick = (t) => {
      group.rotation.y = 0.25 + Math.sin(t * 0.3) * 0.18;
      pointer.rotation.z = -0.55 + Math.sin(t * 0.9) * 0.12; // gesturing
    };
  } else if (persona === 'ninja') {
    const mask = buildNinjaMask(THREE, palette);
    mask.scale.setScalar(scale);
    // Eye height matches the professor glasses; mask sits a touch proud of the pupils.
    mask.position.set(0, 0.77 * scale, 0.34 * scale);
    group.add(mask);

    const gi = buildNinjaGi(THREE, palette);
    gi.scale.setScalar(scale);
    // Torso props are authored in model-local units with y≈0 at the feet of the
    // unit mesh; scaling the whole group matches the body scale.
    group.add(gi);

    const katana = buildKatana(THREE, palette);
    katana.scale.setScalar(scale);
    // Right paw, blade up — the reference holds it nearly vertical.
    katana.position.set(0.34 * scale, 0.28 * scale, 0.22 * scale);
    katana.rotation.z = -0.18;
    katana.rotation.x = 0.08;
    group.add(katana);

    // Slight lean into a ready stance.
    model.rotation.x = 0.06;
    group.userData.tick = (t) => {
      group.rotation.y = 0.15 + Math.sin(t * 0.28) * 0.14;
      const base = group.userData.baseY ?? 0;
      group.position.y = base + Math.sin(t * 1.05) * 0.04;
      katana.rotation.z = -0.18 + Math.sin(t * 0.7) * 0.06;
    };
  } else {
    const speed = buildSpeedLines(THREE, palette);
    speed.scale.setScalar(scale);
    group.add(speed);
    group.userData.tick = (t) => {
      // The marks pulse rather than translate; a trailing line that moves reads
      // as debris, one that flickers reads as speed.
      speed.children.forEach((bar, i) => {
        bar.scale.x = 0.75 + Math.sin(t * 3.2 + i * 1.1) * 0.35;
      });
      group.rotation.y = -0.35 + Math.sin(t * 0.4) * 0.35;
      group.position.y = group.userData.baseY + Math.sin(t * 1.1) * 0.1;
    };
  }

  return group;
}
