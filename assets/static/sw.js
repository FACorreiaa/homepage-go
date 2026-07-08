// Service Worker for FC Software Studio
const CACHE_NAME = 'fc-studio-v5';
const PRECACHE = [
  '/',
  '/assets/css/output.css',
  '/assets/static/vendor/alpine/alpine.min.js',
  '/assets/static/vendor/htmx/htmx.min.js',
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

  const path = new URL(req.url).pathname;
  const isRevalidatedAsset = path === '/assets/css/output.css' ||
    path === '/assets/static/sw.js' ||
    path === '/assets/static/manifest.json';
  if (isRevalidatedAsset) {
    event.respondWith(networkFirst(req));
    return;
  }

  const isAsset = path.startsWith('/assets/');
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

function networkFirst(req) {
  return fetch(req)
    .then((res) => {
      const copy = res.clone();
      caches.open(CACHE_NAME).then((cache) => cache.put(req, copy)).catch(() => {});
      return res;
    })
    .catch(() => caches.match(req));
}

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
