// The landing page's background graph: the languages and adjacent stack from
// the CV, laid out by hand and connected by claimed relationships.
//
// This is decoration for home-world.js. Positions are authored and stable —
// nothing here runs a force simulation. The camera ride stays in home-world.js;
// this file only reads how far along it is so a constellation can brighten as
// the camera reaches it.
//
// Nodes sit in the gutters of each station's X/Y plane. The ride travels -Z,
// so a station is a volume the camera flies through, not a floor it looks down on.

import { STATIONS } from './home-path.js';

const CLEAR_R = 2.8;

/**
 * Labeled nodes. x is the gutter (negative left). yOff is height. z is a small
 * authored offset along the path, applied on top of the station's z.
 * Hubs are the three languages the site is hired for.
 */
const NODES = [
  // device — sparse arc in the gutters, not on the headline
  { id: 'go', label: 'Go', tier: 'hub', station: 'device', x: -6.4, z: -2.2, yOff: 0.55 },
  { id: 'swift', label: 'Swift', tier: 'hub', station: 'device', x: 6.2, z: -2.4, yOff: 0.45 },
  { id: 'typescript', label: 'TypeScript', tier: 'hub', station: 'device', x: 6.0, z: -4.2, yOff: -1.35 },

  // server — dense field in the gutters around the fork panels
  { id: 'csharp', label: 'C#', tier: 'primary', station: 'server', x: 8.4, z: -1.6, yOff: 0.2 },
  { id: 'rust', label: 'Rust', tier: 'primary', station: 'server', x: 8.8, z: -4.4, yOff: 0.9 },
  { id: 'sql', label: 'SQL', tier: 'primary', station: 'server', x: 0.8, z: -7.2, yOff: -0.6 },
  { id: 'react', label: 'React', tier: 'primary', station: 'server', x: -8.6, z: -2.0, yOff: 0.45 },
  { id: 'swiftui', label: 'SwiftUI', tier: 'primary', station: 'server', x: -8.0, z: -5.0, yOff: -0.25 },
  { id: 'vapor', label: 'Vapor', tier: 'primary', station: 'server', x: -8.8, z: 0.8, yOff: 0.8 },
  { id: 'grpc', label: 'gRPC', tier: 'primary', station: 'server', x: 7.2, z: -6.2, yOff: -0.55 },
  { id: 'html', label: 'HTML', tier: 'secondary', station: 'server', x: -9.4, z: -3.8, yOff: 0.35 },
  { id: 'css', label: 'CSS', tier: 'secondary', station: 'server', x: 9.4, z: -3.4, yOff: -0.7 },
  { id: 'tailwind', label: 'Tailwind', tier: 'secondary', station: 'server', x: 8.8, z: 1.0, yOff: 0.5 },
  { id: 'postgres', label: 'PostgreSQL', tier: 'secondary', station: 'server', x: 3.2, z: -8.2, yOff: -1.2 },

  // cluster — frameworks in the margins around the featured heading
  { id: 'reactnative', label: 'React Native', tier: 'secondary', station: 'cluster', x: -7.0, z: -3.4, yOff: 0.15 },
  { id: 'angular', label: 'Angular', tier: 'secondary', station: 'cluster', x: 7.1, z: -2.8, yOff: 0.85 },
  { id: 'hummingbird', label: 'Hummingbird', tier: 'secondary', station: 'cluster', x: -6.4, z: 1.2, yOff: -0.4 },
  { id: 'graphql', label: 'GraphQL', tier: 'secondary', station: 'cluster', x: 6.2, z: -5.6, yOff: 0.5 },
  { id: 'redis', label: 'Redis', tier: 'secondary', station: 'cluster', x: -5.8, z: -6.2, yOff: -0.7 },
  { id: 'mongodb', label: 'MongoDB', tier: 'secondary', station: 'cluster', x: 5.6, z: 1.4, yOff: 0.25 },

  // horizon — quiet ops constellation, no floor grid
  { id: 'docker', label: 'Docker', tier: 'primary', station: 'horizon', x: -6.2, z: -3.8, yOff: 0.1 },
  { id: 'linux', label: 'Linux', tier: 'secondary', station: 'horizon', x: 1.4, z: -7.2, yOff: -0.55 },
  { id: 'kubernetes', label: 'Kubernetes', tier: 'primary', station: 'horizon', x: 6.4, z: -4.2, yOff: 0.4 },
  { id: 'mssql', label: 'MSSQL', tier: 'secondary', station: 'horizon', x: -4.8, z: 1.6, yOff: -0.25 },

  // horizon — the products, a quieter constellation of their own
  { id: 'norviq', label: 'Norviq', tier: 'product', station: 'horizon', x: -7.4, z: 2.2, yOff: 1.2 },
  { id: 'luminavault', label: 'LuminaVault', tier: 'product', station: 'horizon', x: 7.6, z: 1.6, yOff: 0.9 },
  { id: 'hermes', label: 'Hermes', tier: 'product', station: 'horizon', x: 5.2, z: 3.0, yOff: -1.6 },
];

// Claimed relationships, not random pairs. Cross-station edges are what make
// the descent read as one network rather than four separate diagrams.
const EDGES = [
  ['go', 'grpc'],
  ['go', 'postgres'],
  ['go', 'rust'],
  ['swift', 'swiftui'],
  ['swift', 'vapor'],
  ['vapor', 'hummingbird'],
  ['typescript', 'react'],
  ['typescript', 'angular'],
  ['typescript', 'html'],
  ['typescript', 'graphql'],
  ['react', 'reactnative'],
  ['html', 'css'],
  ['css', 'tailwind'],
  ['csharp', 'mssql'],
  ['csharp', 'mongodb'],
  ['sql', 'postgres'],
  ['sql', 'mssql'],
  ['grpc', 'graphql'],
  ['postgres', 'redis'],
  ['docker', 'linux'],
  ['linux', 'kubernetes'],
  ['swiftui', 'vapor'],
  ['norviq', 'hermes'],
  ['luminavault', 'hermes'],
  ['hermes', 'kubernetes'],
  ['norviq', 'go'],
];

// A subset of the edges carry a travelling packet. Short list on purpose.
const PACKET_EDGES = [
  ['go', 'grpc'],
  ['swift', 'swiftui'],
  ['typescript', 'react'],
  ['react', 'reactnative'],
  ['grpc', 'graphql'],
  ['postgres', 'redis'],
  ['docker', 'linux'],
  ['linux', 'kubernetes'],
  ['swift', 'vapor'],
  ['html', 'css'],
  ['csharp', 'mongodb'],
  ['go', 'postgres'],
  ['norviq', 'hermes'],
  ['hermes', 'kubernetes'],
];

const TIER = {
  hub: { nodeR: 0.22, sat: 2, labelH: 0.4, opacity: 0.92, edgeOpacity: 0.4, role: 'go' },
  primary: { nodeR: 0.16, sat: 2, labelH: 0.34, opacity: 0.8, edgeOpacity: 0.32, role: 'fg' },
  product: { nodeR: 0.18, sat: 3, labelH: 0.36, opacity: 0.84, edgeOpacity: 0.34, role: 'trail' },
  secondary: { nodeR: 0.12, sat: 1, labelH: 0.28, opacity: 0.58, edgeOpacity: 0.28, role: 'fg' },
  satellite: { nodeR: 0.04, sat: 0, labelH: 0, opacity: 0.2, edgeOpacity: 0, role: 'fg' },
};

function hash01(n) {
  const x = Math.sin(n * 127.1 + 311.7) * 43758.5453;
  return x - Math.floor(x);
}

function placedAt(node) {
  // x is the gutter. yOff is height. Old z (depth toward the previous camera)
  // becomes a small offset along the path so a station is a volume, not a card.
  const x = node.x;
  const y = node.yOff * 1.8;
  const z = STATIONS[node.station].z + node.z * 0.45;
  const r = Math.hypot(x, y);
  const scale = r < CLEAR_R ? (CLEAR_R + 0.4) / Math.max(r, 0.001) : 1;
  return { x: x * scale, y: y * scale, z };
}

function stationNear(station, progress) {
  const t = STATIONS[station]?.t ?? 0;
  const d = Math.abs((progress ?? 0) - t);
  return Math.max(0, 1 - d / 0.22);
}

function cssOf(color) {
  return `#${color.getHexString()}`;
}

function tag(material, role) {
  material.userData.role = role;
  return material;
}

function drawLabel(canvas, text, fill, weight) {
  const size = 28;
  const ctx = canvas.getContext('2d');
  const font = `${weight} ${size}px "Geist Mono", ui-monospace, monospace`;
  ctx.font = font;
  const padX = 18;
  const padY = 10;
  const textW = Math.ceil(ctx.measureText(text).width);
  const w = textW + padX * 2;
  const h = size + padY * 2;
  canvas.width = w * 2;
  canvas.height = h * 2;
  ctx.scale(2, 2);
  ctx.font = font;
  ctx.textAlign = 'center';
  ctx.textBaseline = 'middle';
  ctx.fillStyle = fill;
  ctx.fillText(text, w / 2, h / 2);
  return { w, h };
}

function labelColor(node, palette) {
  if (node.tier === 'hub') return palette.go;
  if (node.tier === 'product') return palette.trail;
  return palette.fg;
}

function makeLabel(THREE, node, palette) {
  const canvas = document.createElement('canvas');
  const color = labelColor(node, palette);
  const weight = node.tier === 'hub' || node.tier === 'product' ? 560 : 450;
  const { w, h } = drawLabel(canvas, node.label, cssOf(color), weight);
  const texture = new THREE.CanvasTexture(canvas);
  texture.colorSpace = THREE.SRGBColorSpace;
  texture.needsUpdate = true;
  const material = new THREE.SpriteMaterial({
    map: texture,
    transparent: true,
    depthWrite: false,
    opacity: TIER[node.tier].opacity,
  });
  const sprite = new THREE.Sprite(material);
  const worldH = TIER[node.tier].labelH;
  sprite.scale.set(worldH * (w / h), worldH, 1);
  sprite.position.set(node.x, node.y + 0.28, node.z);
  sprite.renderOrder = 2;
  sprite.userData.label = {
    canvas,
    text: node.label,
    tier: node.tier,
    weight,
    baseOpacity: TIER[node.tier].opacity,
  };
  return sprite;
}

function placeInstance(dummy, mesh, i, x, y, z, scale, rot = null) {
  dummy.position.set(x, y, z);
  dummy.scale.setScalar(scale);
  if (rot) dummy.rotation.set(rot.x, rot.y, rot.z);
  else dummy.rotation.set(0, 0, 0);
  dummy.updateMatrix();
  mesh.setMatrixAt(i, dummy.matrix);
}

function restPose(seed, tier) {
  const hx = hash01(seed * 5);
  const hz = hash01(seed * 7);
  if (tier === 'primary' || tier === 'secondary' || tier === 'product') {
    return { x: 0.42 + hx * 0.28, z: hz * 0.5 - 0.25 };
  }
  return { x: hx * 0.35 - 0.17, z: hz * 0.35 - 0.17 };
}

function spinRates(seed) {
  return {
    x: 0.06 + hash01(seed * 19) * 0.08,
    y: 0.06 + hash01(seed * 23) * 0.08,
    z: 0.06 + hash01(seed * 29) * 0.08,
  };
}

function tierGeometry(THREE, tier) {
  if (tier === 'hub' || tier === 'product') return new THREE.OctahedronGeometry(1, 0);
  if (tier === 'primary') return new THREE.OctahedronGeometry(1, 0);
  if (tier === 'secondary') return new THREE.OctahedronGeometry(1, 0);
  return new THREE.TetrahedronGeometry(1, 0);
}

function makePhongMaterial(THREE, palette, tier) {
  const spec = TIER[tier];
  const role = spec.role;
  const mat = new THREE.MeshPhongMaterial({
    color: palette[role],
    transparent: true,
    opacity: spec.opacity,
    shininess: 28,
  });
  if (tier === 'hub' || tier === 'product') {
    mat.emissive.copy(palette[role]);
    mat.emissiveIntensity = 0.18;
    mat.userData.emissiveRole = role;
  }
  return tag(mat, role);
}

function makeEdgeMaterial(THREE, palette, tier) {
  const spec = TIER[tier];
  const role = spec.role;
  return tag(
    new THREE.LineBasicMaterial({
      color: palette[role],
      transparent: true,
      opacity: spec.edgeOpacity,
    }),
    role,
  );
}

/** One labeled node: scaled group, rest-tilted pivot, lit fill + hairline edges. */
function createLabeledNode(THREE, node, palette) {
  const spec = TIER[node.tier];
  const geo = tierGeometry(THREE, node.tier);
  const nodeGroup = new THREE.Group();
  nodeGroup.position.set(node.x, node.y, node.z);
  nodeGroup.scale.setScalar(spec.nodeR);

  const rest = restPose(node.seed, node.tier);
  const spin = spinRates(node.seed);
  const phase = hash01(node.seed * 31) * Math.PI * 2;

  const pivot = new THREE.Group();
  pivot.rotation.set(rest.x, 0, rest.z);
  const fill = new THREE.Mesh(geo, makePhongMaterial(THREE, palette, node.tier));
  pivot.add(fill);
  pivot.add(new THREE.LineSegments(new THREE.EdgesGeometry(geo, 15), makeEdgeMaterial(THREE, palette, node.tier)));
  nodeGroup.add(pivot);

  return { group: nodeGroup, pivot, fill, node, spin, phase, rest };
}

/**
 * Builds the graph as one group spanning the whole descent. home-world.js adds
 * it to the scene once; station groups only keep the gophers and the helm.
 */
export function buildGraph(THREE, palette) {
  const group = new THREE.Group();

  const placed = NODES.map((node, i) => {
    const p = placedAt(node);
    return { ...node, ...p, seed: i + 1 };
  });

  const byId = new Map(placed.map((n) => [n.id, n]));

  const satellites = [];
  placed.forEach((node) => {
    const count = TIER[node.tier].sat;
    for (let s = 0; s < count; s++) {
      const h = hash01(node.seed * 17 + s * 9);
      const h2 = hash01(node.seed * 3 + s * 13 + 5);
      const h3 = hash01(node.seed * 11 + s * 7 + 2);
      const yaw = h * Math.PI * 2;
      const reach = 0.7 + h2 * 0.7;
      satellites.push({
        x: node.x + Math.cos(yaw) * reach,
        y: node.y + (h3 - 0.5) * 0.7,
        z: node.z + Math.sin(yaw) * reach,
        parent: node.id,
        seed: node.seed * 10 + s,
      });
    }
  });

  // A thin dust ring at the horizon so that station does not fall off a cliff.
  // Unlabeled, attached to the nearest ops/product node. In XY: a halo the
  // camera flies toward, not a floor ring it looks down on.
  const horizonAnchors = placed.filter((n) => n.station === 'horizon');
  for (let i = 0; i < 10; i++) {
    const h = hash01(90 + i);
    const a = (i / 10) * Math.PI * 2 + 0.35;
    const r = 8.2 + h * 2.4;
    const x = Math.cos(a) * r;
    const y = Math.sin(a) * r * 0.55;
    const z = STATIONS.horizon.z + (hash01(40 + i) - 0.5) * 3.2;
    let parent = horizonAnchors[0];
    let best = Infinity;
    for (const n of horizonAnchors) {
      const d = (n.x - x) ** 2 + (n.y - y) ** 2;
      if (d < best) {
        best = d;
        parent = n;
      }
    }
    satellites.push({ x, y, z, parent: parent.id, seed: 200 + i });
  }

  const dummy = new THREE.Object3D();
  const satGeo = tierGeometry(THREE, 'satellite');

  const labeledNodes = placed
    .filter((n) => n.tier === 'hub' || n.tier === 'primary' || n.tier === 'secondary' || n.tier === 'product')
    .map((node) => {
      const hw = createLabeledNode(THREE, node, palette);
      group.add(hw.group);
      return hw;
    });

  const satMesh = new THREE.InstancedMesh(
    satGeo,
    tag(new THREE.MeshLambertMaterial({ color: palette.fg, transparent: true, opacity: TIER.satellite.opacity }), 'fg'),
    satellites.length,
  );
  satellites.forEach((n, i) => placeInstance(dummy, satMesh, i, n.x, n.y, n.z, TIER.satellite.nodeR));
  satMesh.instanceMatrix.needsUpdate = true;

  group.add(satMesh);

  const structure = [];
  const signal = [];
  const signalSet = new Set(PACKET_EDGES.map(([a, b]) => `${a}|${b}`));

  const pushSeg = (bucket, ax, ay, az, bx, by, bz) => {
    bucket.push(ax, ay, az, bx, by, bz);
  };

  EDGES.forEach(([a, b]) => {
    const na = byId.get(a);
    const nb = byId.get(b);
    if (!na || !nb) return;
    const bucket = signalSet.has(`${a}|${b}`) || signalSet.has(`${b}|${a}`) ? signal : structure;
    pushSeg(bucket, na.x, na.y, na.z, nb.x, nb.y, nb.z);
  });

  satellites.forEach((sat) => {
    const parent = byId.get(sat.parent);
    if (!parent) return;
    pushSeg(structure, sat.x, sat.y, sat.z, parent.x, parent.y, parent.z);
  });

  const structureLines = new THREE.LineSegments(
    new THREE.BufferGeometry().setAttribute('position', new THREE.Float32BufferAttribute(structure, 3)),
    tag(new THREE.LineBasicMaterial({ color: palette.fg, transparent: true, opacity: 0.26 }), 'fg'),
  );
  const signalLines = new THREE.LineSegments(
    new THREE.BufferGeometry().setAttribute('position', new THREE.Float32BufferAttribute(signal, 3)),
    tag(new THREE.LineBasicMaterial({ color: palette.signal, transparent: true, opacity: 0.42 }), 'signal'),
  );
  group.add(structureLines, signalLines);

  const labels = placed.map((node) => {
    const sprite = makeLabel(THREE, node, palette);
    group.add(sprite);
    return { sprite, node };
  });

  const packetCount = PACKET_EDGES.length;
  const packetPos = new Float32Array(packetCount * 3);
  const packetGeo = new THREE.BufferGeometry();
  packetGeo.setAttribute('position', new THREE.BufferAttribute(packetPos, 3));
  const packetMat = tag(
    new THREE.PointsMaterial({
      color: palette.signal,
      size: 0.09,
      transparent: true,
      opacity: 0.85,
      depthWrite: false,
      sizeAttenuation: true,
    }),
    'signal',
  );
  const packets = new THREE.Points(packetGeo, packetMat);
  packets.frustumCulled = false;
  group.add(packets);

  const packetRoutes = PACKET_EDGES.map(([a, b], i) => {
    const na = byId.get(a);
    const nb = byId.get(b);
    return {
      a: na,
      b: nb,
      phase: hash01(i + 4),
      speed: 0.07 + hash01(i + 21) * 0.05,
    };
  }).filter((r) => r.a && r.b);

  const ndc = new THREE.Vector3();
  const world = new THREE.Vector3();

  group.userData.tick = (t, camera, progress) => {
    packetMat.opacity = 0.7 + Math.sin(t * 1.8) * 0.12;
    signalLines.material.opacity = 0.18 + Math.sin(t * 1.4) * 0.07;

    labeledNodes.forEach(({ pivot, fill, node, spin, phase, rest }) => {
      pivot.rotation.x = rest.x + Math.sin(t * spin.x + phase) * 0.06;
      pivot.rotation.y = t * spin.y + phase;
      pivot.rotation.z = rest.z + Math.cos(t * spin.z + phase) * 0.05;
      if (fill?.material?.emissive) {
        const near = stationNear(node.station, progress);
        fill.material.emissiveIntensity = 0.08 + near * 0.32;
      }
    });

    satellites.forEach((sat, i) => {
      const yaw = t * 0.05 + hash01(sat.seed) * Math.PI * 2;
      dummy.position.set(sat.x, sat.y, sat.z);
      dummy.rotation.set(hash01(sat.seed * 3) * 0.4, yaw, hash01(sat.seed * 5) * 0.4);
      dummy.scale.setScalar(TIER.satellite.nodeR);
      dummy.updateMatrix();
      satMesh.setMatrixAt(i, dummy.matrix);
    });
    satMesh.instanceMatrix.needsUpdate = true;

    const attr = packetGeo.getAttribute('position');
    packetRoutes.forEach((route, i) => {
      const u = (t * route.speed + route.phase) % 1;
      const s = u < 0.5 ? u * 2 : (1 - u) * 2;
      attr.setXYZ(
        i,
        route.a.x + (route.b.x - route.a.x) * s,
        route.a.y + (route.b.y - route.a.y) * s,
        route.a.z + (route.b.z - route.a.z) * s,
      );
    });
    attr.needsUpdate = true;

    if (!camera) return;
    labels.forEach(({ sprite, node }) => {
      world.set(node.x, node.y, node.z);
      const dist = camera.position.distanceTo(world);
      ndc.copy(world).project(camera);
      const offCenter = Math.hypot(ndc.x * 1.05, ndc.y * 1.15);
      // Hide names that land on the copy column; keep the gutters fully on.
      const onCopy = Math.abs(ndc.x) < 0.36 && ndc.y > -0.22 && ndc.y < 0.52;
      const mid = onCopy ? 0.08 : offCenter < 0.18 ? 0.55 : 1;
      const near = THREE.MathUtils.smoothstep(dist, 2.4, 5.5);
      const far = 1 - THREE.MathUtils.smoothstep(dist, 22, 40);
      const along = 0.45 + stationNear(node.station, progress) * 0.55;
      const behind = ndc.z > 1 ? 0 : 1;
      sprite.material.opacity = sprite.userData.label.baseOpacity * mid * near * far * along * behind;
      sprite.visible = sprite.material.opacity > 0.03;
    });
  };

  group.userData.repaint = (next) => {
    labels.forEach(({ sprite, node }) => {
      const meta = sprite.userData.label;
      const color = labelColor(node, next);
      drawLabel(meta.canvas, meta.text, cssOf(color), meta.weight);
      sprite.material.map.needsUpdate = true;
    });
  };

  return group;
}
