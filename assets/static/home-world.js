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
// apart from a flat background. Geometry is mostly procedural; the only mesh bytes are the shared
// gopher GLB cloned into station personas.

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

const GRID_COLS = 20; // lattice footprint; center hole + kind filter reduce live instances
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
function readPalette(THREE, tokenColor) {
  const go = tokenColor('--go', '#00add8');
  return {
    fg: tokenColor('--foreground', '#111111'),
    signal: tokenColor('--signal', '#22c55e'),
    trail: tokenColor('--trail', '#d99a2b'),
    go,
    // The gopher's pale belly. Derived rather than its own token: it is only
    // ever a tint of the body, and computing it here means the theme repaint
    // reaches it like every other role.
    'go-belly': go.clone().lerp(new THREE.Color(0xffffff), 0.62),
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
function buildDevice(THREE, palette, character) {
  const group = new THREE.Group();
  // Framed on the camera's look-at (x=0). A light right bias keeps the coder
  // from sitting dead under the hero headline while still reading as centered.
  group.position.set(1.1, STATIONS.device.groupY - 0.6, 0);

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

  // The coder sits inside the ring, at the desk.
  if (character) {
    character.position.set(0, -1.9, 1.2);
    group.add(character);
  }

  group.userData.tick = (t) => {
    group.rotation.y = Math.sin(t * 0.25) * 0.18;
    ring.rotation.z = t * 0.06;
    character?.userData.tick?.(t);
  };

  return group;
}

/**
 * Station two — Go on the server. The runner stands in a field of network
 * nodes: mostly quiet wireframe lattice, with nested frames and small energy
 * cores so the rack reads as infrastructure rather than a cube demo.
 */
function buildServer(THREE, palette, character) {
  const group = new THREE.Group();
  group.position.y = STATIONS.server.groupY;

  // The runner stands in the centre of the rack — same axis the camera looks at.
  if (character) {
    character.position.set(0.4, -1.4, 1.5);
    character.userData.baseY = -1.4;
    group.add(character);
  }

  const spacing = 1.15;
  const cols = GRID_COLS;
  const half = cols / 2;
  // Soft hole so the character is not buried under equal-weight lattice.
  const CLEAR_R = 2.55;

  // Stable hash — layout must not reshuffle every reload.
  const hash01 = (n) => {
    const x = Math.sin(n * 127.1 + 311.7) * 43758.5453;
    return x - Math.floor(x);
  };

  const cells = [];
  for (let x = 0; x < cols; x++) {
    for (let z = 0; z < cols; z++) {
      const px = (x - half) * spacing;
      const pz = (z - half) * spacing;
      const dist = Math.hypot(px, pz);
      if (dist < CLEAR_R) continue;

      const h = hash01(x * 73 + z * 19 + 3);
      // Depth falloff: far ring stays quieter so fork panels stay readable.
      const falloff = THREE.MathUtils.clamp(1.15 - dist * 0.045, 0.42, 1);
      let kind = 'box';
      if (h > 0.88) kind = 'core';
      else if (h > 0.72) kind = 'nested';
      else if (h > 0.48) kind = 'octa';

      cells.push({
        px,
        pz,
        dist,
        kind,
        falloff,
        phase: dist * 0.22,
        spin: (hash01(x * 11 + z * 29) - 0.5) * 0.4,
      });
    }
  }

  // Nested = outer box + inner octa only (not also a full octa cell).
  const boxes = cells.filter((c) => c.kind === 'box' || c.kind === 'nested');
  const octas = cells.filter((c) => c.kind === 'octa');
  const cores = cells.filter((c) => c.kind === 'core');
  const nested = cells.filter((c) => c.kind === 'nested');

  const boxMesh = new THREE.InstancedMesh(
    new THREE.BoxGeometry(0.4, 0.4, 0.4),
    tag(new THREE.MeshBasicMaterial({ color: palette.fg, transparent: true, opacity: 0.12, wireframe: true }), 'fg'),
    Math.max(boxes.length, 1),
  );
  const octaMesh = new THREE.InstancedMesh(
    new THREE.OctahedronGeometry(0.28, 0),
    tag(new THREE.MeshBasicMaterial({ color: palette.fg, transparent: true, opacity: 0.11, wireframe: true }), 'fg'),
    Math.max(octas.length, 1),
  );
  const nestedMesh = new THREE.InstancedMesh(
    new THREE.OctahedronGeometry(0.16, 0),
    tag(new THREE.MeshBasicMaterial({ color: palette.go, transparent: true, opacity: 0.18, wireframe: true }), 'go'),
    Math.max(nested.length, 1),
  );
  const coreMesh = new THREE.InstancedMesh(
    new THREE.IcosahedronGeometry(0.07, 0),
    tag(new THREE.MeshBasicMaterial({ color: palette.go, transparent: true, opacity: 0.55 }), 'go'),
    Math.max(cores.length, 1),
  );

  const dummy = new THREE.Object3D();
  const place = (mesh, list, scaleFor) => {
    list.forEach((cell, i) => {
      const s = scaleFor(cell);
      dummy.position.set(cell.px, 0, cell.pz);
      dummy.scale.setScalar(s);
      dummy.rotation.set(0, cell.spin, 0);
      dummy.updateMatrix();
      mesh.setMatrixAt(i, dummy.matrix);
    });
    mesh.instanceMatrix.needsUpdate = true;
    mesh.count = list.length;
    group.add(mesh);
  };

  place(boxMesh, boxes, (c) => c.falloff);
  place(octaMesh, octas, (c) => c.falloff);
  place(nestedMesh, nested, (c) => c.falloff);
  place(coreMesh, cores, (c) => c.falloff * 0.95);

  // Traffic biased toward core nodes when possible — purposeful links, not pure noise.
  const linePositions = [];
  const coreOrAll = cores.length >= 2 ? cores : cells;
  const pick = (n) => coreOrAll[Math.floor(hash01(n + 9) * coreOrAll.length) % coreOrAll.length];
  for (let n = 0; n < 28; n++) {
    const a = pick(n * 2);
    const b = pick(n * 2 + 1);
    if (!a || !b || a === b) continue;
    linePositions.push(a.px, 0.28, a.pz, b.px, 0.28, b.pz);
  }
  const traffic = new THREE.LineSegments(
    new THREE.BufferGeometry().setAttribute('position', new THREE.Float32BufferAttribute(linePositions, 3)),
    tag(new THREE.LineBasicMaterial({ color: palette.signal, transparent: true, opacity: 0.28 }), 'signal'),
  );
  group.add(traffic);

  const waveDummy = new THREE.Object3D();
  const wave = (mesh, list, t, rotMul, yAmp, extraRot = 0) => {
    list.forEach((cell, i) => {
      const s = cell.falloff * (mesh === coreMesh ? 0.95 + Math.sin(t * 2.1 + cell.phase) * 0.12 : 1);
      waveDummy.position.set(cell.px, Math.sin(t * 0.9 - cell.phase) * yAmp, cell.pz);
      waveDummy.scale.setScalar(Math.max(s, 0.2));
      waveDummy.rotation.set(extraRot * t, cell.spin + t * rotMul, extraRot * t * 0.6);
      waveDummy.updateMatrix();
      mesh.setMatrixAt(i, waveDummy.matrix);
    });
    mesh.instanceMatrix.needsUpdate = true;
  };

  group.userData.tick = (t) => {
    wave(boxMesh, boxes, t, 0.08, 0.2);
    wave(octaMesh, octas, t, 0.11, 0.22, 0.05);
    wave(nestedMesh, nested, t, -0.16, 0.18, 0.12);
    wave(coreMesh, cores, t, 0.2, 0.16, 0.15);
    // Pulse core material opacity as a field, not per-instance (InstancedMesh shares one material).
    coreMesh.material.opacity = 0.42 + Math.sin(t * 2.0) * 0.14;
    nestedMesh.material.opacity = 0.14 + Math.sin(t * 1.3) * 0.05;
    traffic.material.opacity = 0.18 + Math.sin(t * 2.2) * 0.08;
    character?.userData.tick?.(t);
  };

  return group;
}

/**
 * Station three — the cluster. Drop through the grid into the pods themselves:
 * a field of points, too many to count, which is the point.
 */
function buildCluster(THREE, palette, character) {
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

  const pods = new THREE.Points(geometry, material);
  group.add(pods);

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

  // Slightly right of centre and set back so the wheel frames the professor
  // without drifting into the far gutter.
  helm.position.set(1.6, -2.6, -9);
  helm.scale.setScalar(0.72);
  group.add(helm);


  // The professor explains the cluster, standing under the helm.
  if (character) {
    character.position.set(0.6, -4.2, -4.0);
    group.add(character);
  }

  group.userData.tick = (t) => {
    material.uniforms.uTime.value = t;
    // Only the pod field spins; the professor would spin with it otherwise.
    pods.rotation.y = t * 0.035;
    // A wheel at the helm turns slowly and steadily. Nothing dramatic.
    helm.rotation.z = -t * 0.09;
    character?.userData.tick?.(t);
  };
  group.userData.material = material;

  return group;
}

/**
 * Station four — the horizon. Everything recedes to a ground plane and a line.
 * The scene goes quiet where the page asks for a decision. The ninja stands
 * here: production shipped, still watching the perimeter.
 */
function buildHorizon(THREE, palette, character) {
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

  if (character) {
    // Slightly right of centre, feet clear of the drifting grid.
    character.position.set(1.4, -4.35, 0.6);
    character.userData.baseY = -4.35;
    group.add(character);
  }

  group.userData.tick = (t) => {
    // The grid drifts toward the viewer, one cell per cycle, so the ground
    // reads as moving without the camera having to.
    grid.position.z = (t * 0.6) % 2;
    horizon.material.opacity = 0.35 + Math.sin(t * 0.8) * 0.15;
    character?.userData.tick?.(t);
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

  // One trigger per leg rather than one timeline across the whole page.
  //
  // A single timeline has to bake each leg's share of the scroll from the
  // stations' offsetTop at build time, and those go stale the moment the
  // document reflows — fonts landing and the scroll reveals firing were enough
  // to leave the camera 10 units short of the cluster while its copy was
  // already on screen. invalidateOnRefresh re-measures the trigger, not the
  // durations inside it.
  //
  // Binding each leg to its own pair of station elements means the mapping is
  // the DOM's, not a snapshot of it, and it re-measures for free on resize.
  for (let i = 1; i < sections.length; i++) {
    const from = STATIONS[sections[i - 1].name];
    const to = STATIONS[sections[i].name];

    const leg = gsap.timeline({
      scrollTrigger: {
        trigger: sections[i - 1].el,
        start: 'top top',
        endTrigger: sections[i].el,
        end: 'top top',
        scrub: 1.1, // Lags the scroll slightly; the camera has weight.
        invalidateOnRefresh: true,
      },
    });

    // fromTo, not to: a bare .to() infers its start from wherever the camera
    // happens to be when the leg first renders, so scrolling up would start the
    // leg from the wrong station. Naming both ends makes each leg total.
    leg.fromTo(
      camera.position,
      { y: from.camY, z: from.camZ },
      { y: to.camY, z: to.camZ, ease: 'none', duration: 1 },
      0,
    );
    // Same slot, so the aim arrives with the position rather than trailing it.
    leg.fromTo(focus, { y: from.lookY }, { y: to.lookY, ease: 'none', duration: 1 }, 0);
  }
}

/**
 * Builds the scene. three.js is imported lazily so 720KB never touches the
 * critical path — the hero has painted long before this runs.
 */
async function startWorld() {
  const THREE = await import('three');
  const { tokenColor } = await import('./three-utils.js');
  const { loadGopher, makeGopher } = await import('./gopher.js');

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

  let palette = readPalette(THREE, tokenColor);

  const scene = new THREE.Scene();
  const camera = new THREE.PerspectiveCamera(42, 1, 0.1, 300);
  camera.position.set(0, STATIONS.device.camY, STATIONS.device.camZ);
  // Tweened alongside the camera so the tilt changes between stations too.
  const focus = new THREE.Vector3(0, STATIONS.device.lookY, 0);

  // Lights. Until now every material in this scene was MeshBasicMaterial and
  // there were no lights at all, which is exactly why the old procedural gopher
  // read as a flat silhouette. The model is Lambert-shaded and needs these.
  scene.add(new THREE.AmbientLight(0xffffff, 1.7));
  const key = new THREE.DirectionalLight(0xffffff, 2.4);
  key.position.set(4, 6, 6);
  scene.add(key);
  const rim = new THREE.DirectionalLight(0x99ddff, 1.1);
  rim.position.set(-5, 2, -4);
  scene.add(rim);

  // One mesh, cloned per persona. If it fails to load the stations still build
  // — they simply have no character in them.
  let characters = { coder: null, runner: null, professor: null, ninja: null };
  try {
    await loadGopher(THREE);
    characters = {
      coder: makeGopher(THREE, palette, 'coder', { scale: 1.7 }),
      runner: makeGopher(THREE, palette, 'runner', { scale: 1.7 }),
      professor: makeGopher(THREE, palette, 'professor', { scale: 1.7 }),
      ninja: makeGopher(THREE, palette, 'ninja', { scale: 1.6 }),
    };
  } catch {
    /* no gopher; the scene is still a scene */
  }

  const groups = [
    buildDevice(THREE, palette, characters.coder),
    buildServer(THREE, palette, characters.runner),
    buildCluster(THREE, palette, characters.professor),
    buildHorizon(THREE, palette, characters.ninja),
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
    palette = readPalette(THREE, tokenColor);
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
