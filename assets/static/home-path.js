// The landing world's bearing path.
//
// Scrolling used to drop the camera down -Y through four stacked diagrams.
// The ride is now a flight along -Z: one curve, a station at each section, and
// the language graph sitting in the gutters as a constellation. home-world.js
// owns the camera. home-graph.js places nodes against the same table so the
// two files cannot drift.
//
// Cover and device share t=0. The cover is a veil over a still world; travel
// begins on the device -> server leg.

export const PATH_LENGTH = 122;

export const STATIONS = {
  cover: { z: 0, t: 0 },
  device: { z: 0, t: 0 },
  server: { z: -34, t: 34 / PATH_LENGTH },
  cluster: { z: -76, t: 76 / PATH_LENGTH },
  horizon: { z: -122, t: 1 },
};

export const STATION_ORDER = ['device', 'server', 'cluster', 'horizon'];

/**
 * A slow sway rather than a straight run down -Z. A straight corridor has no
 * parallax, and the world stops reading as a space. Amplitude is kept inside
 * the copy column so gutter nodes at |x| ≈ 6–9 stay in the flanks.
 */
export function createBearingPath(THREE) {
  const points = [];
  const segments = 24;
  for (let i = 0; i <= segments; i++) {
    const u = i / segments;
    points.push(
      new THREE.Vector3(Math.sin(i * 0.55) * 3.2, Math.cos(i * 0.4) * 1.6, -u * PATH_LENGTH),
    );
  }
  return new THREE.CatmullRomCurve3(points, false, 'catmullrom', 0.4);
}

export function clamp01(v) {
  return Math.min(1, Math.max(0, v));
}
