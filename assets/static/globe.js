// Dot-matrix globe for /stats.
//
// Draws this server's own traffic: land is sampled from a land mask into a
// point cloud, the day/night terminator is computed from the current time, and
// each visit arrives as a marker plus an arc into Porto.
//
// Progressive enhancement only. Without WebGL, without JS, or if any of this
// throws, the server-rendered tables below the canvas remain the real content.

import * as THREE from 'three';
import { OrbitControls } from 'three/examples/jsm/controls/OrbitControls.js';
import { tokenColor } from './three-utils.js';

const GLOBE_RADIUS = 1;
const HOME = { lat: 41.15, lon: -8.61 }; // Porto — where the arcs land.
const MASK_URL = '/assets/static/globe/land-mask.png';

const MAX_MARKERS = 2000;
const MAX_ARCS = 40;
const FRESH_SECONDS = 60; // how long a visit reads as "here now"
const METRIC_REFRESH_MS = 60000;

const DEG = Math.PI / 180;

const stage = document.querySelector('[data-globe]');
const status = document.getElementById('globe-status');

const reducedMotion = window.matchMedia('(prefers-reduced-motion: reduce)');

function fail(message) {
  if (status) status.textContent = message;
}

/** Standard three.js globe mapping. Land and markers must share it. */
function toVec3(lat, lon, radius) {
  const phi = (90 - lat) * DEG;
  const theta = (lon + 180) * DEG;
  return new THREE.Vector3(
    -radius * Math.sin(phi) * Math.cos(theta),
    radius * Math.cos(phi),
    radius * Math.sin(phi) * Math.sin(theta),
  );
}

/**
 * Subsolar point, accurate to a couple of degrees. That is far below what a
 * terminator drawn at this scale can show, and it avoids pulling in an
 * ephemeris for a decoration that only has to be honest, not precise.
 */
function sunDirection(date) {
  const start = Date.UTC(date.getUTCFullYear(), 0, 0);
  const dayOfYear = (Date.UTC(date.getUTCFullYear(), date.getUTCMonth(), date.getUTCDate()) - start) / 86400000;
  const utcHours = date.getUTCHours() + date.getUTCMinutes() / 60;

  const lat = 23.44 * Math.sin(((2 * Math.PI) / 365) * (dayOfYear - 81));
  const lon = -15 * (utcHours - 12);

  return toVec3(lat, lon, 1).normalize();
}

async function loadLandPoints() {
  // Decoded via createImageBitmap rather than an Image element. HTMLImageElement
  // .decode() never settles for this mask when the element is not in the
  // document — not a rejection, a hang — which stalled the whole globe behind a
  // permanent "loading map…" with no error anywhere. createImageBitmap decodes
  // off the main thread and, crucially, actually rejects when it fails.
  const response = await fetch(MASK_URL);
  if (!response.ok) throw new Error(`land mask ${response.status}`);
  const image = await createImageBitmap(await response.blob());

  const canvas = document.createElement('canvas');
  canvas.width = image.width;
  canvas.height = image.height;

  const ctx = canvas.getContext('2d', { willReadFrequently: true });
  ctx.drawImage(image, 0, 0);
  image.close();
  const { data } = ctx.getImageData(0, 0, canvas.width, canvas.height);

  const positions = [];
  const rowStep = 4;

  for (let py = 0; py < canvas.height; py += rowStep) {
    const lat = 90 - (py / canvas.height) * 180;

    // Columns converge at the poles, so widen the column step by 1/cos(lat)
    // to keep the dots roughly evenly spaced on the sphere.
    const spread = Math.max(0.12, Math.cos(lat * DEG));
    const colStep = Math.max(rowStep, Math.round(rowStep / spread));

    for (let px = 0; px < canvas.width; px += colStep) {
      if (data[(py * canvas.width + px) * 4] < 128) continue;

      const lon = (px / canvas.width) * 360 - 180;
      const point = toVec3(lat, lon, GLOBE_RADIUS);
      positions.push(point.x, point.y, point.z);
    }
  }

  const geometry = new THREE.BufferGeometry();
  geometry.setAttribute('position', new THREE.Float32BufferAttribute(positions, 3));
  return geometry;
}

// Shared GLSL: how strongly a point faces the camera. Used to fade the far
// side of the sphere so the globe reads as a solid body without an opaque
// shell hiding the page background behind it.
const FACING_GLSL = `
  float facingOf(vec4 mvPosition, vec3 worldNormal) {
    vec3 viewNormal = normalize(normalMatrix * worldNormal);
    return smoothstep(-0.15, 0.35, dot(viewNormal, normalize(-mvPosition.xyz)));
  }
`;

// Point sizes are given in world units and projected here. uProjScale is
// drawingBufferHeight / (2 * tan(fov/2)); without it gl_PointSize is in raw
// device pixels and every dot collapses to a speck whatever the camera does.
const POINT_SIZE_GLSL = `
  float projectedSize(float worldSize, float viewZ) {
    return worldSize * uProjScale / -viewZ;
  }
`;

const LAND_DOT_SIZE = 0.0060;

function makeLandMaterial(colors) {
  return new THREE.ShaderMaterial({
    transparent: true,
    depthWrite: false,
    uniforms: {
      uSun: { value: new THREE.Vector3(1, 0, 0) },
      uLand: { value: colors.land },
      uSize: { value: LAND_DOT_SIZE },
      uProjScale: { value: 1000 },
    },
    vertexShader: `
      uniform float uSize;
      uniform float uProjScale;
      uniform vec3 uSun;
      varying float vFacing;
      varying float vDay;
      ${FACING_GLSL}
      ${POINT_SIZE_GLSL}
      void main() {
        vec3 worldNormal = normalize(position);
        vec4 mvPosition = modelViewMatrix * vec4(position, 1.0);
        vFacing = facingOf(mvPosition, worldNormal);
        // The terminator: a soft band where the sun grazes the surface.
        vDay = smoothstep(-0.09, 0.09, dot(worldNormal, uSun));
        gl_PointSize = projectedSize(uSize, mvPosition.z);
        gl_Position = projectionMatrix * mvPosition;
      }
    `,
    fragmentShader: `
      uniform vec3 uLand;
      varying float vFacing;
      varying float vDay;
      void main() {
        vec2 offset = gl_PointCoord - vec2(0.5);
        if (dot(offset, offset) > 0.25) discard;
        // The night floor stays high enough that the unlit half still reads as
        // part of a sphere rather than missing geometry.
        float alpha = mix(0.42, 0.95, vDay) * mix(0.07, 1.0, vFacing);
        gl_FragColor = vec4(uLand, alpha);
      }
    `,
  });
}

function makeMarkerMaterial(colors) {
  return new THREE.ShaderMaterial({
    transparent: true,
    depthWrite: false,
    // Normal blending, not additive: additive adds light, which is invisible
    // against the white background of the light theme.
    blending: THREE.NormalBlending,
    uniforms: {
      uTime: { value: 0 },
      uSignal: { value: colors.signal },
      uTrail: { value: colors.trail },
      uProjScale: { value: 1000 },
      uFresh: { value: FRESH_SECONDS },
    },
    vertexShader: `
      attribute float aBirth;
      uniform float uTime;
      uniform float uFresh;
      uniform float uProjScale;
      varying float vFacing;
      varying float vFresh;
      ${FACING_GLSL}
      ${POINT_SIZE_GLSL}
      void main() {
        vec3 worldNormal = normalize(position);
        vec4 mvPosition = modelViewMatrix * vec4(position, 1.0);
        vFacing = facingOf(mvPosition, worldNormal);
        vFresh = clamp(1.0 - (uTime - aBirth) / uFresh, 0.0, 1.0);
        // A brief pop on arrival, then settle to the resting size.
        float age = uTime - aBirth;
        float pop = 1.0 + 1.6 * exp(-age * 3.5);
        float size = (0.018 + 0.020 * vFresh) * pop;
        gl_PointSize = projectedSize(size, mvPosition.z);
        gl_Position = projectionMatrix * mvPosition;
      }
    `,
    fragmentShader: `
      uniform vec3 uSignal;
      uniform vec3 uTrail;
      varying float vFacing;
      varying float vFresh;
      void main() {
        vec2 offset = gl_PointCoord - vec2(0.5);
        float d = dot(offset, offset);
        if (d > 0.25) discard;
        float core = smoothstep(0.25, 0.0, d);
        vec3 color = mix(uTrail, uSignal, vFresh);
        float alpha = core * mix(0.45, 1.0, vFresh) * mix(0.06, 1.0, vFacing);
        gl_FragColor = vec4(color, alpha);
      }
    `,
  });
}

function start(landGeometry) {
  const colors = {
    land: tokenColor('--foreground', '#111111'),
    signal: tokenColor('--signal', '#22c55e'),
    trail: tokenColor('--trail', '#d99a2b'),
  };

  const scene = new THREE.Scene();

  // The globe must fit the vertical field of view: it subtends asin(r/d), so
  // the camera has to sit beyond r/sin(fov/2) ≈ 3.07 or the poles get clipped.
  const camera = new THREE.PerspectiveCamera(38, 1, 0.1, 100);
  camera.position.set(0, 0.5, 3.5);

  let renderer;
  try {
    renderer = new THREE.WebGLRenderer({ antialias: true, alpha: true, powerPreference: 'high-performance' });
  } catch {
    fail('This browser cannot draw the map. The numbers below are the same data.');
    return;
  }
  renderer.setPixelRatio(Math.min(window.devicePixelRatio, 2));
  renderer.setClearColor(0x000000, 0);
  stage.appendChild(renderer.domElement);
  if (status) status.remove();

  const landMaterial = makeLandMaterial(colors);
  scene.add(new THREE.Points(landGeometry, landMaterial));

  // Markers live in one preallocated buffer that wraps around, so a busy day
  // never grows the scene graph.
  const markerPositions = new Float32Array(MAX_MARKERS * 3);
  const markerBirths = new Float32Array(MAX_MARKERS);
  const markerGeometry = new THREE.BufferGeometry();
  markerGeometry.setAttribute('position', new THREE.BufferAttribute(markerPositions, 3));
  markerGeometry.setAttribute('aBirth', new THREE.BufferAttribute(markerBirths, 1));
  markerGeometry.setDrawRange(0, 0);

  const markerMaterial = makeMarkerMaterial(colors);
  scene.add(new THREE.Points(markerGeometry, markerMaterial));

  let markerCount = 0;
  let markerCursor = 0;

  const controls = new OrbitControls(camera, renderer.domElement);
  controls.enableDamping = true;
  controls.dampingFactor = 0.06;
  controls.enablePan = false;
  controls.minDistance = 1.6;
  controls.maxDistance = 5;
  controls.rotateSpeed = 0.45;
  controls.autoRotate = !reducedMotion.matches;
  controls.autoRotateSpeed = 0.32;

  controls.addEventListener('start', () => {
    controls.autoRotate = false;
  });

  const clock = new THREE.Clock();
  const arcs = [];

  function addMarker(lat, lon, ageSeconds) {
    const point = toVec3(lat, lon, GLOBE_RADIUS * 1.004);
    const index = markerCursor;

    markerPositions.set([point.x, point.y, point.z], index * 3);
    // A visit that already happened is born in the past, so one code path
    // covers both the 24h backfill and live arrivals.
    markerBirths[index] = clock.getElapsedTime() - ageSeconds;

    markerCursor = (markerCursor + 1) % MAX_MARKERS;
    markerCount = Math.min(markerCount + 1, MAX_MARKERS);

    markerGeometry.attributes.position.needsUpdate = true;
    markerGeometry.attributes.aBirth.needsUpdate = true;
    markerGeometry.setDrawRange(0, markerCount);

    return point;
  }

  function addArc(from) {
    if (arcs.length >= MAX_ARCS) return;

    const to = toVec3(HOME.lat, HOME.lon, GLOBE_RADIUS * 1.004);
    if (from.distanceTo(to) < 0.02) return; // a local visit has nowhere to travel

    // Lift the control point off the surface so the arc bows outward, more so
    // the further the visitor is from Porto.
    const mid = from.clone().add(to).multiplyScalar(0.5);
    const lift = 1 + 0.42 * from.distanceTo(to);
    mid.normalize().multiplyScalar(GLOBE_RADIUS * lift);

    const curve = new THREE.QuadraticBezierCurve3(from.clone(), mid, to);
    const geometry = new THREE.BufferGeometry().setFromPoints(curve.getPoints(48));
    const material = new THREE.LineBasicMaterial({
      color: colors.signal,
      transparent: true,
      opacity: 0.85,
      depthWrite: false,
    });

    const line = new THREE.Line(geometry, material);
    const instant = reducedMotion.matches;
    geometry.setDrawRange(0, instant ? 49 : 0);
    scene.add(line);

    arcs.push({ line, geometry, material, born: clock.getElapsedTime(), instant });
  }

  function updateArcs(elapsed) {
    for (let i = arcs.length - 1; i >= 0; i--) {
      const arc = arcs[i];
      const age = elapsed - arc.born;

      if (!arc.instant) {
        arc.geometry.setDrawRange(0, Math.min(49, Math.floor((age / 1.1) * 49)));
      }
      arc.material.opacity = 0.85 * Math.max(0, 1 - Math.max(0, age - 1.4) / 2.2);

      if (age > 3.8) {
        scene.remove(arc.line);
        arc.geometry.dispose();
        arc.material.dispose();
        arcs.splice(i, 1);
      }
    }
  }

  function resize() {
    const { clientWidth, clientHeight } = stage;
    if (!clientWidth || !clientHeight) return;

    camera.aspect = clientWidth / clientHeight;
    camera.updateProjectionMatrix();
    renderer.setSize(clientWidth, clientHeight, false);

    // Keep dot sizes consistent in world units across viewport and DPR changes.
    const projScale = renderer.domElement.height / (2 * Math.tan((camera.fov * DEG) / 2));
    landMaterial.uniforms.uProjScale.value = projScale;
    markerMaterial.uniforms.uProjScale.value = projScale;
  }
  resize();
  const resizeObserver = new ResizeObserver(resize);
  resizeObserver.observe(stage);

  // Only draw while the canvas is actually on screen and the tab is in front.
  let onScreen = true;
  const visibilityObserver = new IntersectionObserver(
    ([entry]) => {
      onScreen = entry.isIntersecting;
      if (onScreen) loop();
    },
    { threshold: 0 },
  );
  visibilityObserver.observe(stage);

  let frame = 0;
  let sunUpdatedAt = -Infinity;

  function render() {
    const elapsed = clock.getElapsedTime();

    // The sun moves a quarter of a degree per minute. Recomputing it once a
    // minute is well past what anyone can see.
    if (elapsed - sunUpdatedAt > 60) {
      landMaterial.uniforms.uSun.value.copy(sunDirection(new Date()));
      sunUpdatedAt = elapsed;
    }

    markerMaterial.uniforms.uTime.value = elapsed;
    updateArcs(elapsed);
    controls.update();
    renderer.render(scene, camera);
  }

  function loop() {
    cancelAnimationFrame(frame);
    if (!onScreen || document.hidden) return;
    render();
    frame = requestAnimationFrame(loop);
  }

  document.addEventListener('visibilitychange', () => {
    if (!document.hidden) loop();
  });

  reducedMotion.addEventListener('change', (event) => {
    controls.autoRotate = !event.matches;
  });

  // Repaint on theme change — the tokens differ between light and dark.
  const themeObserver = new MutationObserver(() => {
    colors.land = tokenColor('--foreground', '#111111');
    colors.signal = tokenColor('--signal', '#22c55e');
    colors.trail = tokenColor('--trail', '#d99a2b');
    landMaterial.uniforms.uLand.value.copy(colors.land);
    markerMaterial.uniforms.uSignal.value.copy(colors.signal);
    markerMaterial.uniforms.uTrail.value.copy(colors.trail);
  });
  themeObserver.observe(document.documentElement, { attributes: true, attributeFilter: ['class'] });

  loop();

  // ---- data ----------------------------------------------------------------

  function applyMetrics(payload) {
    for (const [key, value] of Object.entries(payload)) {
      const el = document.querySelector(`[data-metric="${key}"]`);
      if (el && typeof value === 'number') el.textContent = String(value);
    }
  }

  async function backfill() {
    try {
      const response = await fetch('/api/stats');
      if (!response.ok) return;

      const payload = await response.json();
      applyMetrics(payload);

      const now = Date.now();
      for (const visit of payload.visits ?? []) {
        const ageSeconds = Math.max(0, (now - Date.parse(visit.at)) / 1000);
        addMarker(visit.lat, visit.lon, ageSeconds);
      }
    } catch {
      // An empty globe is a fair rendering of "no data".
    }
  }

  const stream = new EventSource('/api/stats/stream');
  stream.addEventListener('visit', (event) => {
    try {
      const visit = JSON.parse(event.data);
      addArc(addMarker(visit.lat, visit.lon, 0));

      const views = document.querySelector('[data-metric="views24h"]');
      if (views) views.textContent = String((parseInt(views.textContent, 10) || 0) + 1);
    } catch {
      // A malformed event costs one dot, not the page.
    }
  });

  const metricTimer = setInterval(() => {
    fetch('/api/stats')
      .then((response) => (response.ok ? response.json() : null))
      .then((payload) => payload && applyMetrics(payload))
      .catch(() => {});
  }, METRIC_REFRESH_MS);

  backfill();

  window.addEventListener('pagehide', () => {
    stream.close();
    clearInterval(metricTimer);
    cancelAnimationFrame(frame);
    resizeObserver.disconnect();
    visibilityObserver.disconnect();
    themeObserver.disconnect();
    controls.dispose();
    landGeometry.dispose();
    landMaterial.dispose();
    markerGeometry.dispose();
    markerMaterial.dispose();
    renderer.dispose();
  });
}

if (stage) {
  loadLandPoints()
    .then(start)
    .catch(() => fail('Could not draw the map. The numbers below are the same data.'));
}
