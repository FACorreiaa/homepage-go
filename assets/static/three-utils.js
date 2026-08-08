// Colour helpers shared by every three.js scene on this site.
//
// They exist because the site's design tokens are oklch(), which THREE.Color
// cannot parse. Rather than duplicating the palette in JS — where it would
// drift from input.css the first time a token changes — the browser is asked to
// resolve each token and the result is read back as RGB.

import * as THREE from 'three';

/**
 * Resolves any CSS colour — including the oklch() tokens this site uses, which
 * THREE.Color cannot parse — by letting the browser do it on a 1x1 canvas.
 */
export function resolveColor(cssValue, fallback) {
  try {
    const ctx = document.createElement('canvas').getContext('2d');
    ctx.fillStyle = fallback;
    ctx.fillStyle = cssValue;
    ctx.fillRect(0, 0, 1, 1);
    const [r, g, b] = ctx.getImageData(0, 0, 1, 1).data;
    return new THREE.Color(r / 255, g / 255, b / 255);
  } catch {
    return new THREE.Color(fallback);
  }
}

/** Reads a design token off :root and resolves it to a THREE.Color. */
export function tokenColor(name, fallback) {
  const value = getComputedStyle(document.documentElement).getPropertyValue(name).trim();
  return resolveColor(value || fallback, fallback);
}
