// Service Worker for FC Software Studio
const CACHE_NAME = 'fc-studio-v3';
const PRECACHE = [
  '/',
  '/assets/css/output.css',
  '/assets/static/site-motion.js',
  '/assets/static/vendor/gsap/gsap.min.js',
  '/assets/static/vendor/gsap/ScrollTrigger.min.js',
  '/assets/static/manifest.json',
];

// Install - precache core shell, activate immediately
self.addEventListener('install', (event) => {
  self.skipWaiting();
  event.waitUntil(
    caches.open(CACHE_NAME).then((cache) => cache.addAll(PRECACHE)).catch(() => {})
  );
});

// Activate - take control now and drop old caches
self.addEventListener('activate', (event) => {
  event.waitUntil(
    caches.keys()
      .then((names) => Promise.all(
        names.map((name) => (name !== CACHE_NAME ? caches.delete(name) : null))
      ))
      .then(() => self.clients.claim())
  );
});

// Fetch - cache assets aggressively, keep HTML/API network-first for freshness.
self.addEventListener('fetch', (event) => {
  const req = event.request;
  if (req.method !== 'GET') return;
  if (new URL(req.url).origin !== self.location.origin) return;

  const isAsset = new URL(req.url).pathname.startsWith('/assets/');
  if (isAsset) {
    event.respondWith(cacheFirst(req));
    return;
  }

  event.respondWith(
    fetch(req)
      .then((res) => {
        const copy = res.clone();
        caches.open(CACHE_NAME).then((cache) => cache.put(req, copy)).catch(() => {});
        return res;
      })
      .catch(() => caches.match(req))
  );
});

function cacheFirst(req) {
  return caches.match(req).then((cached) => {
    if (cached) return cached;
    return fetch(req).then((res) => {
      const copy = res.clone();
      caches.open(CACHE_NAME).then((cache) => cache.put(req, copy)).catch(() => {});
      return res;
    });
  });
}
