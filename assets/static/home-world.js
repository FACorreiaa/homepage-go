// The landing page's 3D world.
//
// Scrolling descends the stack the hero copy describes — device, server,
// cluster, horizon — as one camera move down a single scene. Each DOM section
// is a station; ScrollTrigger scrubs the camera between them.
//
// Progressive enhancement, and aggressively so. The server-rendered page is the
// real landing page: it converts, it is what a crawler reads, and it is what
// most visitors get. This file is decoration that has to earn its bytes, so it
// declines to load itself whenever that trade looks bad — see shouldRun below.
//
// Nothing here may change layout. The canvas is fixed, behind the content, and
// aria-hidden; if every line of this file failed the page would be unchanged
// apart from a flat background. Everything is procedural — no textures, no
// models, no image bytes at all.

const stage = document.querySelector('[data-world]');

// Where each station sits on the camera's descent, and how the camera frames it.
// The stations stack down -Y in the order the copy introduces them. Each one
// carries its own camera height and aim rather than sharing a single tilt: a
// grid seen from its own altitude is edge-on and reads as nothing, while the
// device wants to be looked at almost level.
//
//   groupY — where the geometry lives
//   camY/camZ — where the camera sits at that station
//   lookY — the height the camera aims at, which is what sets the tilt
const STATIONS = {
  device: { groupY: 0, camY: 0, camZ: 9, lookY: -1 },
  server: { groupY: -34, camY: -28, camZ: 12, lookY: -37 },
  cluster: { groupY: -76, camY: -76, camZ: 11, lookY: -79 },
  horizon: { groupY: -122, camY: -114, camZ: 16, lookY: -131 },
};

const GRID_COLS = 20; // 400 instanced nodes
const POD_COUNT = 8000;

/**
 * Whether the scene is worth loading at all. Four ways to say no, and the page
 * is complete under all of them.
 */
function shouldRun() {
  if (!stage) return false;

  // Someone who asked for less motion is not asking for a camera ride.
  if (window.matchMedia('(prefers-reduced-motion: reduce)').matches) return false;

  // Below this the layout is a single column and the canvas would sit behind
  // text on a device most likely to be on a metered connection and a battery.
  if (!window.matchMedia('(min-width: 900px)').matches) return false;

  // An explicit request not to spend bytes.
  if (navigator.connection?.saveData) return false;

  return true;
}

/** Best-effort idle callback; Safari still lacks requestIdleCallback. */
function whenIdle(fn) {
  if (typeof requestIdleCallback === 'function') {
    requestIdleCallback(fn, { timeout: 2000 });
  } else {
    setTimeout(fn, 200);
  }
}

/**
 * Smooth scrolling. Deliberately independent of the WebGL scene: Lenis is 18KB
 * and improves the page on its own, so it runs even if three.js never loads.
 * Touch is left alone — the platform's own scrolling is better than ours.
 */
function startLenis() {
  if (typeof Lenis !== 'function') return null;

  const hasGsap = typeof gsap === 'object' && typeof ScrollTrigger === 'function';

  // With GSAP present, Lenis is driven from the gsap ticker instead of its own
  // rAF. One clock for scroll, tweens and rendering keeps the camera from
  // lagging the text by a frame.
  const lenis = new Lenis({ autoRaf: !hasGsap });
  if (hasGsap) {
    lenis.on('scroll', ScrollTrigger.update);
    gsap.ticker.add((time) => lenis.raf(time * 1000));
    gsap.ticker.lagSmoothing(0);
  }

  // The scroll-to-top button in the layout looks for this.
  window.__lenis = lenis;
  return lenis;
}

/** Reads the site's design tokens. They are oklch(), so the browser resolves them. */
function readPalette(tokenColor) {
  return {
    fg: tokenColor('--foreground', '#111111'),
    signal: tokenColor('--signal', '#22c55e'),
    trail: tokenColor('--trail', '#d99a2b'),
    go: tokenColor('--go', '#00add8'),
  };
}

/**
 * Marks which design token a material is painted from, so the theme toggle can
 * repaint the scene by walking it rather than by remembering child indices.
 */
function tag(material, role) {
  material.userData.role = role;
  return material;
}

/**
 * Station one — the device. A single slab alone in the void, ringed by an
 * orbit. The whole stack starts as one object on someone's desk.
 */
function buildDevice(THREE, palette) {
  const group = new THREE.Group();
  // Offset right: the hero headline owns the left half of the viewport, so the
  // one discrete object in the scene sits in the space the text leaves.
  group.position.set(4.6, STATIONS.device.groupY - 0.6, 0);

  const slab = new THREE.Mesh(
    new THREE.BoxGeometry(1.15, 2.3, 0.1),
    tag(new THREE.MeshBasicMaterial({ color: palette.fg, transparent: true, opacity: 0.14 }), 'fg'),
  );
  group.add(slab);

  const edges = new THREE.LineSegments(
    new THREE.EdgesGeometry(slab.geometry),
    tag(new THREE.LineBasicMaterial({ color: palette.fg, transparent: true, opacity: 0.85 }), 'fg'),
  );
  group.add(edges);

  // A screen, lit. This is the only emissive thing at this altitude.
  const screen = new THREE.Mesh(
    new THREE.PlaneGeometry(0.92, 2.0),
    tag(new THREE.MeshBasicMaterial({ color: palette.signal, transparent: true, opacity: 0.1 }), 'signal'),
  );
  screen.position.z = 0.055;
  group.add(screen);

  // The orbit ring reads as latitude — this thing is somewhere real.
  const ring = new THREE.Line(
    new THREE.BufferGeometry().setFromPoints(
      Array.from({ length: 97 }, (_, i) => {
        const a = (i / 96) * Math.PI * 2;
        return new THREE.Vector3(Math.cos(a) * 2.5, 0, Math.sin(a) * 2.5);
      }),
    ),
    tag(new THREE.LineBasicMaterial({ color: palette.fg, transparent: true, opacity: 0.28 }), 'fg'),
  );
  ring.rotation.x = 0.42;
  group.add(ring);

  group.userData.tick = (t) => {
    group.rotation.y = Math.sin(t * 0.25) * 0.35;
    ring.rotation.z = t * 0.06;
    screen.material.opacity = 0.08 + Math.sin(t * 1.6) * 0.03;
  };

  return group;
}

/**
 * Station two — Go. The gopher, built out of primitives rather than loaded from
 * a model: this site ships no image or mesh bytes, and a recognisable gopher is
 * mostly spheres anyway. It sits in a grid of server nodes, because this is the
 * altitude where the CV says Go and Vapor on the server.
 */
function buildServer(THREE, palette) {
  const group = new THREE.Group();
  group.position.y = STATIONS.server.groupY;

  // --- the gopher ---------------------------------------------------------
  const gopher = new THREE.Group();
  // Out past the 6xl content column, in the page gutter, like the device above.
  gopher.position.set(6.6, 1.2, 1.5);
  gopher.scale.setScalar(0.78);

  const skin = tag(
    new THREE.MeshBasicMaterial({ color: palette.go, transparent: true, opacity: 0.38, wireframe: true }),
    'go',
  );
  // Fixed, deliberately untagged: the gopher looks like the gopher in both
  // themes. Tagging these as 'fg' turned the eyes black in light mode.
  const white = new THREE.MeshBasicMaterial({ color: 0xf2f4f6 });
  const dark = new THREE.MeshBasicMaterial({ color: 0x15181d });

  const sphere = (r, seg = 18) => new THREE.SphereGeometry(r, seg, Math.max(8, seg / 2));

  // Body: one tall rounded mass. The gopher is mostly torso.
  const body = new THREE.Mesh(sphere(1.5, 24), skin);
  body.scale.set(0.82, 1, 0.72);
  gopher.add(body);

  // Ears, small and high on the sides.
  for (const side of [-1, 1]) {
    const ear = new THREE.Mesh(sphere(0.3, 12), skin);
    ear.position.set(side * 1.06, 1.12, 0);
    ear.scale.set(1, 1, 0.5);
    gopher.add(ear);
  }

  // Eyes: the feature that makes it read as the gopher and not a rodent.
  // Each gets an outline ring, because near-white eyes on the light theme's
  // near-white page are otherwise just two floating pupils.
  const outlineMat = new THREE.LineBasicMaterial({ color: 0x15181d, transparent: true, opacity: 0.55 });
  const ringPoints = (radius) =>
    Array.from({ length: 41 }, (_, i) => {
      const a = (i / 40) * Math.PI * 2;
      return new THREE.Vector3(Math.cos(a) * radius, Math.sin(a) * radius, 0);
    });

  for (const side of [-1, 1]) {
    const eye = new THREE.Mesh(sphere(0.42, 16), white);
    eye.position.set(side * 0.44, 0.62, 0.92);
    gopher.add(eye);

    const outline = new THREE.Line(
      new THREE.BufferGeometry().setFromPoints(ringPoints(0.42)),
      outlineMat,
    );
    outline.position.set(side * 0.44, 0.62, 1.24);
    gopher.add(outline);

    const pupil = new THREE.Mesh(sphere(0.19, 12), dark);
    pupil.position.set(side * 0.3, 0.62, 1.22);
    gopher.add(pupil);
  }

  // Snout and the two front teeth.
  const snout = new THREE.Mesh(sphere(0.26, 14), white);
  snout.position.set(0, 0.16, 1.06);
  snout.scale.set(1.25, 0.8, 0.7);
  gopher.add(snout);

  for (const side of [-1, 1]) {
    const tooth = new THREE.Mesh(new THREE.BoxGeometry(0.12, 0.2, 0.06), white);
    tooth.position.set(side * 0.08, -0.1, 1.16);
    gopher.add(tooth);
  }

  // Arms and feet.
  for (const side of [-1, 1]) {
    const arm = new THREE.Mesh(sphere(0.24, 12), skin);
    arm.position.set(side * 1.2, -0.25, 0.3);
    arm.scale.set(0.8, 1.5, 0.8);
    gopher.add(arm);

    const foot = new THREE.Mesh(sphere(0.3, 12), skin);
    foot.position.set(side * 0.5, -1.5, 0.34);
    foot.scale.set(1.1, 0.55, 1.35);
    gopher.add(foot);
  }

  group.add(gopher);

  // --- the rack it stands in ----------------------------------------------
  const count = GRID_COLS * GRID_COLS;
  const mesh = new THREE.InstancedMesh(
    new THREE.BoxGeometry(0.42, 0.42, 0.42),
    tag(new THREE.MeshBasicMaterial({ color: palette.fg, transparent: true, opacity: 0.16, wireframe: true }), 'fg'),
    count,
  );

  const dummy = new THREE.Object3D();
  const offsets = new Float32Array(count);
  let i = 0;
  for (let x = 0; x < GRID_COLS; x++) {
    for (let z = 0; z < GRID_COLS; z++) {
      const px = (x - GRID_COLS / 2) * 1.15;
      const pz = (z - GRID_COLS / 2) * 1.15;
      dummy.position.set(px, 0, pz);
      dummy.updateMatrix();
      mesh.setMatrixAt(i, dummy.matrix);
      // Phase by distance from centre so the wave travels outward.
      offsets[i] = Math.hypot(px, pz) * 0.22;
      i++;
    }
  }
  mesh.instanceMatrix.needsUpdate = true;
  group.add(mesh);

  // Traffic between nodes. Straight segments, because this is a network.
  const linePositions = [];
  for (let n = 0; n < 26; n++) {
    const a = Math.floor(Math.random() * count);
    const b = Math.floor(Math.random() * count);
    const ax = ((a / GRID_COLS | 0) - GRID_COLS / 2) * 1.15;
    const az = ((a % GRID_COLS) - GRID_COLS / 2) * 1.15;
    const bx = ((b / GRID_COLS | 0) - GRID_COLS / 2) * 1.15;
    const bz = ((b % GRID_COLS) - GRID_COLS / 2) * 1.15;
    linePositions.push(ax, 0.3, az, bx, 0.3, bz);
  }
  const traffic = new THREE.LineSegments(
    new THREE.BufferGeometry().setAttribute('position', new THREE.Float32BufferAttribute(linePositions, 3)),
    tag(new THREE.LineBasicMaterial({ color: palette.signal, transparent: true, opacity: 0.3 }), 'signal'),
  );
  group.add(traffic);

  const dummyTick = new THREE.Object3D();
  group.userData.tick = (t) => {
    let k = 0;
    for (let x = 0; x < GRID_COLS; x++) {
      for (let z = 0; z < GRID_COLS; z++) {
        const px = (x - GRID_COLS / 2) * 1.15;
        const pz = (z - GRID_COLS / 2) * 1.15;
        dummyTick.position.set(px, Math.sin(t * 0.9 - offsets[k]) * 0.22, pz);
        dummyTick.rotation.y = t * 0.08;
        dummyTick.updateMatrix();
        mesh.setMatrixAt(k, dummyTick.matrix);
        k++;
      }
    }
    mesh.instanceMatrix.needsUpdate = true;
    traffic.material.opacity = 0.2 + Math.sin(t * 2.2) * 0.09;

    // The gopher looks around and bobs. Small amplitudes — it is background.
    gopher.rotation.y = -0.35 + Math.sin(t * 0.4) * 0.35;
    gopher.position.y = 1.2 + Math.sin(t * 1.1) * 0.12;
  };

  return group;
}

/**
 * Station three — the cluster. Drop through the grid into the pods themselves:
 * a field of points, too many to count, which is the point.
 */
function buildCluster(THREE, palette) {
  const group = new THREE.Group();
  group.position.y = STATIONS.cluster.groupY;

  const positions = new Float32Array(POD_COUNT * 3);
  const seeds = new Float32Array(POD_COUNT);
  for (let i = 0; i < POD_COUNT; i++) {
    // Cylindrical shell, so the camera falls through the middle of it.
    const angle = Math.random() * Math.PI * 2;
    const radius = 4 + Math.random() * 13;
    positions[i * 3] = Math.cos(angle) * radius;
    positions[i * 3 + 1] = (Math.random() - 0.5) * 34;
    positions[i * 3 + 2] = Math.sin(angle) * radius;
    seeds[i] = Math.random();
  }

  const geometry = new THREE.BufferGeometry();
  geometry.setAttribute('position', new THREE.Float32BufferAttribute(positions, 3));
  geometry.setAttribute('aSeed', new THREE.Float32BufferAttribute(seeds, 1));

  // Additive so density reads as brightness — a thick part of the field glows
  // without needing a bloom pass, which would double the frame cost.
  const material = new THREE.ShaderMaterial({
    transparent: true,
    depthWrite: false,
    blending: THREE.AdditiveBlending,
    uniforms: {
      uTime: { value: 0 },
      uCold: { value: palette.trail },
      uHot: { value: palette.signal },
      uScale: { value: 300 },
    },
    vertexShader: `
      attribute float aSeed;
      uniform float uTime;
      uniform float uScale;
      varying float vHeat;
      void main() {
        vec3 p = position;
        p.y += sin(uTime * 0.5 + aSeed * 12.0) * 0.5;
        vec4 mv = modelViewMatrix * vec4(p, 1.0);
        // A pod is "hot" in bursts, so the field flickers like real traffic.
        vHeat = smoothstep(0.72, 1.0, fract(aSeed * 7.3 + uTime * 0.12));
        gl_PointSize = (0.012 + vHeat * 0.02) * uScale / max(-mv.z, 0.001);
        gl_Position = projectionMatrix * mv;
      }
    `,
    fragmentShader: `
      uniform vec3 uCold;
      uniform vec3 uHot;
      varying float vHeat;
      void main() {
        // Round the square point sprite off.
        float d = length(gl_PointCoord - vec2(0.5));
        if (d > 0.5) discard;
        float core = smoothstep(0.5, 0.0, d);
        gl_FragColor = vec4(mix(uCold, uHot, vHeat), core * (0.25 + vHeat * 0.6));
      }
    `,
  });

  group.add(new THREE.Points(geometry, material));

  // The helm. Kubernetes' mark is a seven-spoke ship's wheel, and seven is the
  // whole joke, so the geometry is built from a heptagon rather than a circle.
  const helm = new THREE.Group();
  const SPOKES = 7;
  const outer = 3.2;
  const inner = 1.15;

  const heptagon = (radius) =>
    Array.from({ length: SPOKES + 1 }, (_, i) => {
      const a = (i / SPOKES) * Math.PI * 2 + Math.PI / 2;
      return new THREE.Vector3(Math.cos(a) * radius, Math.sin(a) * radius, 0);
    });

  const helmLine = tag(
    new THREE.LineBasicMaterial({ color: palette.signal, transparent: true, opacity: 0.55 }),
    'signal',
  );

  helm.add(new THREE.Line(new THREE.BufferGeometry().setFromPoints(heptagon(outer)), helmLine));
  helm.add(new THREE.Line(new THREE.BufferGeometry().setFromPoints(heptagon(inner)), helmLine));

  // Spokes out to each vertex, each capped with the handle the mark has.
  const spokeEnds = [];
  for (let i = 0; i < SPOKES; i++) {
    const a = (i / SPOKES) * Math.PI * 2 + Math.PI / 2;
    const dir = new THREE.Vector3(Math.cos(a), Math.sin(a), 0);
    spokeEnds.push(dir.clone().multiplyScalar(inner), dir.clone().multiplyScalar(outer * 1.24));
  }
  helm.add(
    new THREE.LineSegments(new THREE.BufferGeometry().setFromPoints(spokeEnds), helmLine),
  );

  // Out in the gutter and set back, like the gopher above: the project cards
  // own the middle of this station.
  helm.position.set(8.2, -2.6, -9);
  helm.scale.setScalar(0.72);
  group.add(helm);


  group.userData.tick = (t) => {
    material.uniforms.uTime.value = t;
    group.rotation.y = t * 0.035;
    // A wheel at the helm turns slowly and steadily. Nothing dramatic.
    helm.rotation.z = -t * 0.09;
  };
  group.userData.material = material;

  return group;
}

/**
 * Station four — the horizon. Everything recedes to a ground plane and a line.
 * The scene goes quiet where the page asks for a decision.
 */
function buildHorizon(THREE, palette) {
  const group = new THREE.Group();
  group.position.y = STATIONS.horizon.groupY;

  const grid = new THREE.GridHelper(120, 60, palette.fg, palette.fg);
  grid.material.transparent = true;
  grid.material.opacity = 0.14;
  tag(grid.material, 'fg');
  grid.position.y = -6;
  group.add(grid);

  const horizon = new THREE.Line(
    new THREE.BufferGeometry().setFromPoints([
      new THREE.Vector3(-60, -6, -60),
      new THREE.Vector3(60, -6, -60),
    ]),
    tag(new THREE.LineBasicMaterial({ color: palette.signal, transparent: true, opacity: 0.5 }), 'signal'),
  );
  group.add(horizon);

  group.userData.tick = (t) => {
    // The grid drifts toward the viewer, one cell per cycle, so the ground
    // reads as moving without the camera having to.
    grid.position.z = (t * 0.6) % 2;
    horizon.material.opacity = 0.35 + Math.sin(t * 0.8) * 0.15;
  };

  return group;
}

/**
 * Scrubs the camera down the stations. One timeline over the whole document:
 * each station's DOM section owns a slice of it, so the camera and the text
 * arrive together however long the sections turn out to be.
 */
function wireScrollTriggers(gsap, ScrollTrigger, camera, focus) {
  const order = ['device', 'server', 'cluster', 'horizon'];
  const sections = order
    .map((name) => ({ name, el: document.querySelector(`[data-world-station="${name}"]`) }))
    .filter((s) => s.el);

  if (sections.length < 2) return;

  const timeline = gsap.timeline({
    scrollTrigger: {
      trigger: document.body,
      start: 'top top',
      end: 'bottom bottom',
      scrub: 1.1, // Lags the scroll slightly; the camera has weight.
      invalidateOnRefresh: true,
    },
  });

  // Proportional segments, so a longer section is a longer part of the descent.
  const tops = sections.map((s) => s.el.offsetTop);
  const span = Math.max(1, tops[tops.length - 1] - tops[0]);

  for (let i = 1; i < sections.length; i++) {
    const from = STATIONS[sections[i - 1].name];
    const target = STATIONS[sections[i].name];
    const share = (tops[i] - tops[i - 1]) / span;
    const at = i === 1 ? 0 : '>';

    // fromTo, not to: a bare .to() infers its start from wherever the camera
    // happens to be when the segment first renders, which makes the descent
    // depend on history — seek backwards and it starts from the wrong station.
    // Naming both ends makes the path total and reversible.
    timeline.fromTo(
      camera.position,
      { y: from.camY, z: from.camZ },
      { y: target.camY, z: target.camZ, duration: share, ease: 'none' },
      at,
    );
    // Same slot, so the aim arrives with the position rather than trailing it.
    timeline.fromTo(
      focus,
      { y: from.lookY },
      { y: target.lookY, duration: share, ease: 'none' },
      '<',
    );
  }

  return timeline;
}

/**
 * Builds the scene. three.js is imported lazily so 720KB never touches the
 * critical path — the hero has painted long before this runs.
 */
async function startWorld() {
  const THREE = await import('three');
  const { tokenColor } = await import('./three-utils.js');

  const canvas = document.createElement('canvas');
  canvas.className = 'world-canvas';
  canvas.setAttribute('aria-hidden', 'true');

  let renderer;
  try {
    renderer = new THREE.WebGLRenderer({ canvas, antialias: true, alpha: true, powerPreference: 'high-performance' });
  } catch {
    return; // No WebGL. The page is already complete without it.
  }

  renderer.setPixelRatio(Math.min(window.devicePixelRatio, 2));
  renderer.setClearColor(0x000000, 0);
  stage.appendChild(canvas);

  // Claim the page background only now that there is definitely a scene to put
  // there. .public-shell paints a gradient that ends in an opaque --background,
  // which would otherwise cover the canvas everywhere below the fold. Every
  // bail-out above this line leaves that gradient exactly as it is.
  document.documentElement.classList.add('world-on');

  let palette = readPalette(tokenColor);

  const scene = new THREE.Scene();
  const camera = new THREE.PerspectiveCamera(42, 1, 0.1, 300);
  camera.position.set(0, STATIONS.device.camY, STATIONS.device.camZ);
  // Tweened alongside the camera so the tilt changes between stations too.
  const focus = new THREE.Vector3(0, STATIONS.device.lookY, 0);

  const groups = [
    buildDevice(THREE, palette),
    buildServer(THREE, palette),
    buildCluster(THREE, palette),
    buildHorizon(THREE, palette),
  ];
  groups.forEach((g) => scene.add(g));

  function resize() {
    const w = window.innerWidth;
    const h = window.innerHeight;
    if (!w || !h) return;
    camera.aspect = w / h;
    camera.updateProjectionMatrix();
    renderer.setSize(w, h, false);
    // Point size is given in world units and projected here, so the pods stay
    // the same physical size across viewport and DPR changes.
    const cluster = groups[2];
    cluster.userData.material.uniforms.uScale.value =
      renderer.domElement.height / (2 * Math.tan((camera.fov * Math.PI) / 360));
  }
  resize();
  window.addEventListener('resize', resize, { passive: true });

  const clock = new THREE.Clock();
  function render() {
    if (document.hidden) return;
    const t = clock.getElapsedTime();
    for (const g of groups) g.userData.tick?.(t);
    camera.lookAt(focus);
    renderer.render(scene, camera);
  }

  if (typeof gsap === 'object' && typeof ScrollTrigger === 'function') {
    wireScrollTriggers(gsap, ScrollTrigger, camera, focus);
    gsap.ticker.add(render);
  } else {
    // No GSAP: no camera move, but the scene still breathes.
    renderer.setAnimationLoop(render);
  }

  // The theme toggle swaps every token underneath us. Re-resolve rather than
  // keeping a second copy of the palette in JS, which would drift from
  // input.css the first time a colour changes there.
  const themeObserver = new MutationObserver(() => {
    palette = readPalette(tokenColor);
    applyPalette(scene, groups, palette);
  });
  themeObserver.observe(document.documentElement, { attributes: true, attributeFilter: ['class'] });

  window.__world = { THREE, renderer, scene, camera, groups, render };
}

/** Repaints existing materials from a freshly resolved palette, by role. */
function applyPalette(scene, groups, palette) {
  scene.traverse((object) => {
    const role = object.material?.userData?.role;
    if (role && palette[role]) object.material.color.set(palette[role]);
  });

  // The pod field is a shader, so its colours are uniforms rather than a
  // material colour and the traversal above cannot reach them.
  const uniforms = groups[2].userData.material.uniforms;
  uniforms.uCold.value = palette.trail;
  uniforms.uHot.value = palette.signal;
}

function boot() {
  if (!shouldRun()) return;
  startLenis();
  whenIdle(() => {
    startWorld().catch(() => {
      // A scene that fails to build leaves the page exactly as it was.
      document.querySelector('.world-canvas')?.remove();
    });
  });
}

boot();
