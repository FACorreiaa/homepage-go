// Service Worker for FC Software Studio
// Network-first: always serve fresh content when online, fall back to cache
// offline. Cache-first (the previous strategy) silently served stale CSS/HTML
// after every deploy until the cache name was bumped — avoid that.
const CACHE_NAME = 'fc-studio-v2';
const PRECACHE = [
  '/',
  '/assets/css/output.css',
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

// Fetch - network-first for same-origin GETs, cache as offline fallback
self.addEventListener('fetch', (event) => {
  const req = event.request;
  if (req.method !== 'GET') return;
  if (new URL(req.url).origin !== self.location.origin) return;

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
