// The landing page's 3D world.
//
// Scrolling flies the camera along a single path through the stack the hero
// copy describes — device, server, cluster, horizon. Each DOM section is a
// station on that path; ScrollTrigger scrubs a 0..1 along the curve between
// them. The language graph lives in home-graph.js. The only mesh bytes are the
// shared gopher GLB cloned into station personas.
//
// Progressive enhancement, and aggressively so. The server-rendered page is the
// real landing page: it converts, it is what a crawler reads, and it is what
// most visitors get. This file is decoration that has to earn its bytes, so it
// declines to load itself whenever that trade looks bad — see shouldRun below.
//
// Nothing here may change layout. The canvas is fixed, behind the content, and
// aria-hidden; if every line of this file failed the page would be unchanged
// apart from a flat background.

import { STATIONS, STATION_ORDER, createBearingPath, clamp01 } from './home-path.js';

const stage = document.querySelector('[data-world]');

const LINE_SEGMENTS = 400;
const MOTE_COUNT = 2400;
const WORLD_UP = { x: 0, y: 1, z: 0 };

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

function luminance(color) {
  return color.r * 0.2126 + color.g * 0.7152 + color.b * 0.0722;
}

function isLightPalette(palette) {
  return luminance(palette.bg) > 0.5;
}

function fogDensity(palette) {
  // Denser on light pages so the field does not dirty the type. Looser in dark
  // so three or four stations stay in frame at once.
  return isLightPalette(palette) ? 0.022 : 0.013;
}

/** Reads the site's design tokens. They are oklch(), so the browser resolves them. */
function readPalette(THREE, tokenColor) {
  const go = tokenColor('--go', '#00add8');
  const fg = tokenColor('--foreground', '#111111');
  const bg = tokenColor('--background', '#111111');
  return {
    bg,
    fg,
    signal: tokenColor('--signal', '#22c55e'),
    trail: tokenColor('--trail', '#d99a2b'),
    go,
    // The gopher's pale belly. Derived rather than its own token: it is only
    // ever a tint of the body, and computing it here means the theme repaint
    // reaches it like every other role.
    'go-belly': go.clone().lerp(new THREE.Color(0xffffff), 0.62),
    hairline: fg,
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

function createMoteTexture(THREE) {
  const size = 64;
  const canvas = document.createElement('canvas');
  canvas.width = canvas.height = size;
  const ctx = canvas.getContext('2d');
  const gradient = ctx.createRadialGradient(size / 2, size / 2, 0, size / 2, size / 2, size / 2);
  gradient.addColorStop(0, 'rgba(255,255,255,1)');
  gradient.addColorStop(0.4, 'rgba(255,255,255,0.5)');
  gradient.addColorStop(1, 'rgba(255,255,255,0)');
  ctx.fillStyle = gradient;
  ctx.fillRect(0, 0, size, size);
  const texture = new THREE.CanvasTexture(canvas);
  texture.colorSpace = THREE.SRGBColorSpace;
  return texture;
}

function hash01(n) {
  const x = Math.sin(n * 127.1 + 311.7) * 43758.5453;
  return x - Math.floor(x);
}

/**
 * The travelled/ahead hairline, two flanking rails, station stars, viewfinder
 * brackets, and the dust that gives the volume its depth. All of it is one
 * group so a theme repaint can walk it without remembering child indices.
 */
function buildGalaxy(THREE, palette, curve) {
  const group = new THREE.Group();
  const light = isLightPalette(palette);

  const linePoints = curve.getPoints(LINE_SEGMENTS);
  const lineGeo = new THREE.BufferGeometry().setFromPoints(linePoints);
  const lineColors = new Float32Array((LINE_SEGMENTS + 1) * 3);
  lineGeo.setAttribute('color', new THREE.BufferAttribute(lineColors, 3));
  const lineMat = new THREE.LineBasicMaterial({
    vertexColors: true,
    transparent: true,
    opacity: light ? 0.55 : 0.9,
  });
  const line = new THREE.Line(lineGeo, lineMat);
  line.frustumCulled = false;
  group.add(line);

  let lineBoundary = -1;
  function paintLine(progress) {
    const next = Math.round(clamp01(progress) * LINE_SEGMENTS);
    if (next === lineBoundary) return;
    lineBoundary = next;
    for (let i = 0; i <= LINE_SEGMENTS; i++) {
      const c = i <= next ? palette.signal : palette.hairline;
      lineColors[i * 3] = c.r;
      lineColors[i * 3 + 1] = c.g;
      lineColors[i * 3 + 2] = c.b;
    }
    lineGeo.attributes.color.needsUpdate = true;
  }
  paintLine(0);

  const railVerts = [];
  const point = new THREE.Vector3();
  const tangent = new THREE.Vector3();
  const side = new THREE.Vector3();
  const worldUp = new THREE.Vector3(0, 1, 0);
  const railSteps = 200;
  const railOffset = 3.2;
  for (const sign of [-1, 1]) {
    for (let i = 0; i < railSteps; i++) {
      for (const t of [i / railSteps, (i + 1) / railSteps]) {
        curve.getPointAt(t, point);
        curve.getTangentAt(t, tangent);
        side.crossVectors(tangent, worldUp).normalize().multiplyScalar(railOffset * sign);
        railVerts.push(point.x + side.x, point.y + side.y, point.z + side.z);
      }
    }
  }
  const rails = new THREE.LineSegments(
    new THREE.BufferGeometry().setAttribute('position', new THREE.Float32BufferAttribute(railVerts, 3)),
    tag(
      new THREE.LineBasicMaterial({
        color: palette.hairline,
        transparent: true,
        opacity: light ? 0.22 : 0.4,
      }),
      'fg',
    ),
  );
  rails.frustumCulled = false;
  group.add(rails);

  const fractions = STATION_ORDER.map((name) => STATIONS[name].t);
  const starGeo = new THREE.OctahedronGeometry(0.22, 0);
  const starMat = new THREE.MeshBasicMaterial({ transparent: true, opacity: light ? 0.4 : 0.72 });
  const stars = new THREE.InstancedMesh(starGeo, starMat, fractions.length);
  stars.instanceMatrix.setUsage(THREE.StaticDrawUsage);
  stars.frustumCulled = false;
  const matrix = new THREE.Matrix4();
  const position = new THREE.Vector3();
  const quaternion = new THREE.Quaternion();
  const scale = new THREE.Vector3(1, 1, 1);
  const euler = new THREE.Euler();
  fractions.forEach((t, i) => {
    curve.getPointAt(clamp01(t), position);
    // Sit below the path so a waypoint does not land on the headline.
    position.y -= 1.35;
    euler.set(0.4, i * 0.7, 0.2);
    quaternion.setFromEuler(euler);
    matrix.compose(position, quaternion, scale);
    stars.setMatrixAt(i, matrix);
    stars.setColorAt(i, palette.hairline);
  });
  stars.instanceMatrix.needsUpdate = true;
  group.add(stars);

  let litCount = -1;
  function paintStars(progress) {
    const lit = fractions.filter((t) => t <= progress + 0.02).length;
    if (lit === litCount) return;
    litCount = lit;
    for (let i = 0; i < fractions.length; i++) {
      stars.setColorAt(i, i < lit ? palette.signal : palette.hairline);
    }
    stars.instanceColor.needsUpdate = true;
  }
  paintStars(0);

  const bracketVerts = [];
  const up = new THREE.Vector3();
  const HALF = 1.55;
  const ARM = 0.48;
  fractions.forEach((t) => {
    curve.getPointAt(clamp01(t), point);
    curve.getTangentAt(t, tangent);
    side.crossVectors(tangent, worldUp).normalize();
    up.crossVectors(side, tangent).normalize();
    for (const sx of [-1, 1]) {
      for (const sy of [-1, 1]) {
        const cx = point.x + side.x * HALF * sx + up.x * HALF * sy;
        const cy = point.y + side.y * HALF * sx + up.y * HALF * sy;
        const cz = point.z + side.z * HALF * sx + up.z * HALF * sy;
        bracketVerts.push(
          cx,
          cy,
          cz,
          cx - side.x * ARM * sx,
          cy - side.y * ARM * sx,
          cz - side.z * ARM * sx,
          cx,
          cy,
          cz,
          cx - up.x * ARM * sy,
          cy - up.y * ARM * sy,
          cz - up.z * ARM * sy,
        );
      }
    }
  });
  const brackets = new THREE.LineSegments(
    new THREE.BufferGeometry().setAttribute('position', new THREE.Float32BufferAttribute(bracketVerts, 3)),
    tag(
      new THREE.LineBasicMaterial({
        color: palette.hairline,
        transparent: true,
        opacity: light ? 0.28 : 0.55,
      }),
      'fg',
    ),
  );
  brackets.frustumCulled = false;
  group.add(brackets);

  const motePos = new Float32Array(MOTE_COUNT * 3);
  const moteCol = new Float32Array(MOTE_COUNT * 3);
  function paintMotes() {
    const lightNow = isLightPalette(palette);
    for (let i = 0; i < MOTE_COUNT; i++) {
      const useGo = hash01(i * 3 + 1) < 0.22;
      const c = lightNow ? palette.fg : useGo ? palette.go : palette.signal;
      moteCol[i * 3] = c.r;
      moteCol[i * 3 + 1] = c.g;
      moteCol[i * 3 + 2] = c.b;
    }
    moteGeo.attributes.color.needsUpdate = true;
    moteMat.size = lightNow ? 0.1 : 0.16;
    moteMat.opacity = lightNow ? 0.38 : 0.88;
    moteMat.blending = lightNow ? THREE.NormalBlending : THREE.AdditiveBlending;
    moteMat.needsUpdate = true;
    lineMat.opacity = lightNow ? 0.55 : 0.9;
    rails.material.opacity = lightNow ? 0.22 : 0.4;
    brackets.material.opacity = lightNow ? 0.22 : 0.45;
    starMat.opacity = lightNow ? 0.4 : 0.72;
  }

  for (let i = 0; i < MOTE_COUNT; i++) {
    const t = hash01(i + 2);
    curve.getPointAt(t, point);
    const angle = hash01(i * 7) * Math.PI * 2;
    const radius = 2.4 + Math.sqrt(hash01(i * 11)) * 12;
    motePos[i * 3] = point.x + Math.cos(angle) * radius;
    motePos[i * 3 + 1] = point.y + Math.sin(angle) * radius * 0.55;
    motePos[i * 3 + 2] = point.z + (hash01(i * 13) - 0.5) * 6;
  }
  const moteGeo = new THREE.BufferGeometry();
  moteGeo.setAttribute('position', new THREE.BufferAttribute(motePos, 3));
  moteGeo.setAttribute('color', new THREE.BufferAttribute(moteCol, 3));
  const moteMat = new THREE.PointsMaterial({
    size: 0.16,
    sizeAttenuation: true,
    vertexColors: true,
    map: createMoteTexture(THREE),
    transparent: true,
    opacity: 0.88,
    depthWrite: false,
    blending: THREE.AdditiveBlending,
  });
  const motes = new THREE.Points(moteGeo, moteMat);
  motes.frustumCulled = false;
  group.add(motes);
  paintMotes();

  group.userData.tick = (_time, progress) => {
    paintLine(progress);
    paintStars(progress);
  };

  group.userData.repaint = (next) => {
    palette = next;
    lineBoundary = -1;
    litCount = -1;
    rails.material.color.copy(next.hairline);
    brackets.material.color.copy(next.hairline);
    paintLine(0);
    paintStars(0);
    paintMotes();
  };

  // Theme toggle calls repaint(next) then the caller paints the current
  // progress; keep a hook so home-world can refresh from ride.t.
  group.userData.refresh = (progress) => {
    lineBoundary = -1;
    litCount = -1;
    paintLine(progress);
    paintStars(progress);
  };

  return group;
}

/**
 * Station one — the device. A gate on the path and the coder just off it.
 * The whole stack starts as one object on someone's desk.
 */
function buildDevice(THREE, palette, character) {
  const group = new THREE.Group();
  group.position.set(1.35, -1.7, STATIONS.device.z);

  const ring = new THREE.Line(
    new THREE.BufferGeometry().setFromPoints(
      Array.from({ length: 97 }, (_, i) => {
        const a = (i / 96) * Math.PI * 2;
        return new THREE.Vector3(Math.cos(a) * 2.4, Math.sin(a) * 2.4, 0);
      }),
    ),
    tag(new THREE.LineBasicMaterial({ color: palette.fg, transparent: true, opacity: 0.28 }), 'fg'),
  );
  group.add(ring);

  if (character) {
    character.position.set(0, 0, 0.35);
    group.add(character);
  }

  group.userData.tick = (t) => {
    ring.rotation.z = t * 0.06;
    character?.userData.tick?.(t);
  };

  return group;
}

/**
 * Station two — Go on the server. The runner stands in a hole the language
 * graph leaves open. The rack itself is the graph, not a cube lattice.
 */
function buildServer(THREE, palette, character) {
  const group = new THREE.Group();
  group.position.set(0.7, -1.5, STATIONS.server.z);

  if (character) {
    character.position.set(0, 0, 0);
    character.userData.baseY = 0;
    group.add(character);
  }

  group.userData.tick = (t) => {
    character?.userData.tick?.(t);
  };

  return group;
}

/**
 * Station three — the cluster. The language graph continues through here; this
 * group only keeps the helm and the professor. The helm sits as a gate the
 * camera flies through.
 */
function buildCluster(THREE, palette, character) {
  const group = new THREE.Group();
  group.position.set(0.4, -1.2, STATIONS.cluster.z);

  const helm = new THREE.Group();
  const SPOKES = 7;
  const outer = 2.6;
  const inner = 0.95;

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

  const spokeEnds = [];
  for (let i = 0; i < SPOKES; i++) {
    const a = (i / SPOKES) * Math.PI * 2 + Math.PI / 2;
    const dir = new THREE.Vector3(Math.cos(a), Math.sin(a), 0);
    spokeEnds.push(dir.clone().multiplyScalar(inner), dir.clone().multiplyScalar(outer * 1.24));
  }
  helm.add(new THREE.LineSegments(new THREE.BufferGeometry().setFromPoints(spokeEnds), helmLine));
  helm.position.set(0, 0.2, -1.4);
  helm.scale.setScalar(0.78);
  group.add(helm);

  if (character) {
    character.position.set(1.6, -0.9, 0.4);
    group.add(character);
  }

  group.userData.tick = (t) => {
    helm.rotation.z = -t * 0.09;
    character?.userData.tick?.(t);
  };

  return group;
}

/**
 * Station four — the horizon. The graph thins into a constellation here. The
 * ninja stands in the remaining space: production shipped, still watching.
 */
function buildHorizon(THREE, palette, character) {
  const group = new THREE.Group();
  group.position.set(1.3, -1.6, STATIONS.horizon.z);

  if (character) {
    character.position.set(0, 0, 0);
    character.userData.baseY = 0;
    group.add(character);
  }

  group.userData.tick = (t) => {
    character?.userData.tick?.(t);
  };

  return group;
}

/**
 * A chase camera, offset from the line rather than sitting on it — you cannot
 * see a path you are standing in the middle of.
 *
 * The offset is built in the path's own frame (behind along the tangent, out
 * along its side) rather than along the world axes. The lateral distance
 * widens as the journey goes on so more of the line already travelled stays
 * in frame, without spinning the camera around to look back.
 */
function makePlacer(THREE, camera, curve) {
  const at = new THREE.Vector3();
  const ahead = new THREE.Vector3();
  const tangent = new THREE.Vector3();
  const side = new THREE.Vector3();
  const worldUp = new THREE.Vector3(WORLD_UP.x, WORLD_UP.y, WORLD_UP.z);

  return function placeCamera(p) {
    const t = clamp01(p);
    curve.getPointAt(t, at);
    curve.getPointAt(clamp01(t + 0.09), ahead);
    curve.getTangentAt(t, tangent);
    side.crossVectors(tangent, worldUp).normalize();

    const out = 1.15 + t * 2.6;
    const up = 0.85 + t * 0.85;
    const back = 6.4;

    camera.position
      .copy(at)
      .addScaledVector(side, out)
      .addScaledVector(worldUp, up)
      .addScaledVector(tangent, -back);
    camera.lookAt(ahead);
  };
}

/**
 * Scrubs a 0..1 along the curve. One trigger per leg rather than one timeline
 * across the whole page: a single timeline has to bake each leg's share of the
 * scroll from the stations' offsetTop at build time, and those go stale the
 * moment the document reflows.
 */
function wireScrollTriggers(gsap, ScrollTrigger, ride) {
  const order = ['cover', 'device', 'server', 'cluster', 'horizon'];
  const sections = order
    .map((name) => ({ name, el: document.querySelector(`[data-world-station="${name}"]`) }))
    .filter((s) => s.el);

  if (sections.length < 2) return;

  for (let i = 1; i < sections.length; i++) {
    const from = STATIONS[sections[i - 1].name];
    const to = STATIONS[sections[i].name];

    gsap
      .timeline({
        scrollTrigger: {
          trigger: sections[i - 1].el,
          start: 'top top',
          endTrigger: sections[i].el,
          end: 'top top',
          scrub: 1.1,
          invalidateOnRefresh: true,
        },
      })
      .fromTo(
        ride,
        { t: from.t },
        { t: to.t, ease: 'none', duration: 1, immediateRender: false },
        0,
      );
  }

  wireCoverVeil(gsap, sections);
}

/**
 * Fades the canvas up as the cover scrolls away.
 *
 * The cover is meant to read as flat type on a flat background, so the scene
 * has to be invisible there — but it still has to be running, because building
 * it at the moment it becomes visible would stutter. So it renders behind an
 * opacity of 0 and is revealed on the way out.
 */
function wireCoverVeil(gsap, sections) {
  const cover = sections.find((s) => s.name === 'cover');
  if (!cover) return;

  gsap.fromTo(
    stage,
    { opacity: 0 },
    {
      opacity: 1,
      ease: 'none',
      scrollTrigger: {
        trigger: cover.el,
        start: 'top top',
        end: 'bottom top',
        scrub: true,
        invalidateOnRefresh: true,
      },
    },
  );
}

/**
 * Builds the scene. three.js is imported lazily so 720KB never touches the
 * critical path — the hero has painted long before this runs.
 */
async function startWorld() {
  const THREE = await import('three');
  const { tokenColor } = await import('./three-utils.js');
  const { loadGopher, makeGopher } = await import('./gopher.js');
  const { buildGraph } = await import('./home-graph.js');
  try {
    await document.fonts?.ready;
  } catch {
    /* canvas labels fall back to ui-monospace */
  }

  const canvas = document.createElement('canvas');
  canvas.className = 'world-canvas';
  canvas.setAttribute('aria-hidden', 'true');

  let renderer;
  try {
    renderer = new THREE.WebGLRenderer({ canvas, antialias: true, alpha: true, powerPreference: 'high-performance' });
  } catch {
    return; // No WebGL. The page is already complete without it.
  }

  renderer.setPixelRatio(Math.min(window.devicePixelRatio, 1.5));
  renderer.setClearColor(0x000000, 0);
  stage.appendChild(canvas);

  // Claim the page background only now that there is definitely a scene to put
  // there. .public-shell paints a gradient that ends in an opaque --background,
  // which would otherwise cover the canvas everywhere below the fold. Every
  // bail-out above this line leaves that gradient exactly as it is.
  document.documentElement.classList.add('world-on');

  let palette = readPalette(THREE, tokenColor);

  const scene = new THREE.Scene();
  scene.fog = new THREE.FogExp2(palette.bg, fogDensity(palette));
  const camera = new THREE.PerspectiveCamera(48, 1, 0.1, 320);

  const curve = createBearingPath(THREE);
  const ride = { t: 0 };
  const placeCamera = makePlacer(THREE, camera, curve);
  placeCamera(0);

  scene.add(new THREE.AmbientLight(0xffffff, 1.55));
  const key = new THREE.DirectionalLight(0xffffff, 1.8);
  key.position.set(4, 6, 6);
  scene.add(key);
  const rim = new THREE.DirectionalLight(0x99ddff, 0.8);
  rim.position.set(-5, 2, 4);
  scene.add(rim);
  // A follow light so stations 100 units down the path are not left on ambient
  // alone. Parenting it to the camera keeps the current constellation readable.
  const follow = new THREE.PointLight(0xffffff, 1.1, 48);
  camera.add(follow);
  scene.add(camera);

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

  const galaxy = buildGalaxy(THREE, palette, curve);
  scene.add(galaxy);

  const graph = buildGraph(THREE, palette);
  scene.add(graph);

  function resize() {
    const w = window.innerWidth;
    const h = window.innerHeight;
    if (!w || !h) return;
    camera.aspect = w / h;
    camera.updateProjectionMatrix();
    renderer.setSize(w, h, false);
  }
  resize();
  window.addEventListener('resize', resize, { passive: true });

  const clock = new THREE.Clock();
  function render() {
    if (document.hidden) return;
    const t = clock.getElapsedTime();
    placeCamera(ride.t);
    for (const g of groups) g.userData.tick?.(t);
    galaxy.userData.tick?.(t, ride.t);
    graph.userData.tick?.(t, camera, ride.t);
    renderer.render(scene, camera);
  }

  if (typeof gsap === 'object' && typeof ScrollTrigger === 'function') {
    wireScrollTriggers(gsap, ScrollTrigger, ride);
    gsap.ticker.add(render);
  } else {
    renderer.setAnimationLoop(render);
  }

  const themeObserver = new MutationObserver(() => {
    palette = readPalette(THREE, tokenColor);
    applyPalette(scene, graph, galaxy, palette, ride.t);
  });
  themeObserver.observe(document.documentElement, { attributes: true, attributeFilter: ['class'] });

  window.__world = { THREE, renderer, scene, camera, groups, graph, galaxy, ride, render };
}

/** Repaints existing materials from a freshly resolved palette, by role. */
function applyPalette(scene, graph, galaxy, palette, progress) {
  scene.traverse((object) => {
    const mat = object.material;
    if (!mat?.userData) return;
    const role = mat.userData.role;
    if (role && palette[role]) mat.color.set(palette[role]);
    const emissiveRole = mat.userData.emissiveRole;
    if (emissiveRole && palette[emissiveRole] && mat.emissive) {
      mat.emissive.set(palette[emissiveRole]);
    }
  });
  if (scene.fog && palette.bg) {
    scene.fog.color.copy(palette.bg);
    scene.fog.density = fogDensity(palette);
  }
  galaxy?.userData.repaint?.(palette);
  galaxy?.userData.refresh?.(progress ?? 0);
  graph?.userData.repaint?.(palette);
}

function boot() {
  if (!shouldRun()) return;
  startLenis();
  whenIdle(() => {
    startWorld().catch(() => {
      document.querySelector('.world-canvas')?.remove();
    });
  });
}

boot();
