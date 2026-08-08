// The landing page's 3D world.
//
// Scrolling descends the stack the hero copy describes — device, server,
// cluster, horizon — as one camera move through a single scene. Each DOM
// section is a station; ScrollTrigger scrubs the camera between them.
//
// Progressive enhancement, and aggressively so. The server-rendered page is the
// real landing page: it converts, it is what a crawler reads, and it is what
// most visitors get. This file is decoration that has to earn its 850KB, so it
// declines to load itself whenever that trade looks bad — see shouldRun below.
//
// Nothing here may change layout. The canvas is fixed, behind the content, and
// aria-hidden; if every line of this file failed the page would be unchanged
// apart from a flat background.

const stage = document.querySelector('[data-world]');

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
 * Smooth scrolling. Deliberately separate from the WebGL scene: Lenis is 18KB
 * and improves the page on its own, so it runs even if three.js never loads.
 * Kept off touch, where the platform's own scrolling is better than ours.
 */
function startLenis() {
  if (typeof Lenis !== 'function') return null;

  const lenis = new Lenis({ autoRaf: true });
  // The scroll-to-top button in the layout looks for this.
  window.__lenis = lenis;
  return lenis;
}

/**
 * Builds the scene. Imported lazily so three.js is never on the critical path —
 * the hero has painted long before this runs.
 */
async function startWorld() {
  const THREE = await import('three');

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

  const scene = new THREE.Scene();
  const camera = new THREE.PerspectiveCamera(38, 1, 0.1, 200);
  camera.position.set(0, 0, 8);

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

  // Nothing is in the scene yet — this draws a transparent frame. The scaffold
  // ships before the geometry so the bail-outs can be verified in isolation.
  function render() {
    renderer.render(scene, camera);
  }
  render();

  window.__world = { THREE, renderer, scene, camera, render };
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
