// The gopher, and the three characters it plays.
//
// One mesh, loaded once, cloned three times. The personas are placement,
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
function paint(THREE, root, palette, opacity) {
  root.traverse((object) => {
    if (!object.isMesh) return;
    const fixed = FIXED[object.name];
    const material = new THREE.MeshLambertMaterial({
      color: fixed ?? palette.go,
      transparent: opacity < 1,
      opacity,
    });
    // Only the body follows a token; see FIXED above.
    if (!fixed) material.userData.role = 'go';
    object.material = material;
  });
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
  const stick = new THREE.Mesh(new THREE.CylinderGeometry(0.008, 0.008, 1.1, 6), material);
  stick.rotation.z = 0.72;
  stick.rotation.x = -0.1;
  return stick;
}

/**
 * Builds one gopher in a given role. Returns a group whose userData.tick, if
 * present, animates it — the caller drives it from the scene clock.
 *
 * @param {'coder'|'runner'|'professor'} persona
 */
export function makeGopher(THREE, palette, persona, { scale = 2.6, opacity = 1 } = {}) {
  const group = new THREE.Group();

  const model = template.clone(true);
  paint(THREE, model, palette, opacity);
  model.scale.setScalar(scale);
  group.add(model);

  if (persona === 'coder') {
    const laptop = buildLaptop(THREE, palette);
    laptop.scale.setScalar(scale);
    // On the ground in front of the gopher, angled up toward it.
    laptop.scale.multiplyScalar(0.72);
    laptop.position.set(0.02 * scale, 0.03 * scale, 0.46 * scale);
    laptop.rotation.y = Math.PI - 0.5;
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
    // Held out to the side, clear of the head rather than across it.
    pointer.position.set(0.62 * scale, 0.34 * scale, 0.16 * scale);
    group.add(pointer);

    group.userData.tick = (t) => {
      group.rotation.y = 0.25 + Math.sin(t * 0.3) * 0.18;
      pointer.rotation.z = 0.72 + Math.sin(t * 0.9) * 0.1; // gesturing
    };
  } else {
    group.userData.tick = (t) => {
      group.rotation.y = -0.35 + Math.sin(t * 0.4) * 0.35;
      group.position.y = group.userData.baseY + Math.sin(t * 1.1) * 0.1;
    };
  }

  return group;
}
