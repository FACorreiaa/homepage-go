// Starts the home-page showreel only when it is about to be seen, and stops it
// when it is not. The markup ships with preload="none" and no autoplay, so a
// visitor who never scrolls this far never downloads it.
//
// Reduced motion and Save-Data keep the poster; the toggle is the way in.
const reel = document.querySelector('[data-showreel]');

if (reel) {
  const video = reel.querySelector('video');
  const toggle = reel.querySelector('[data-showreel-toggle]');
  const reduce = window.matchMedia('(prefers-reduced-motion: reduce)').matches;
  const saveData = navigator.connection?.saveData === true;
  const auto = !reduce && !saveData;

  let loaded = false;
  let visible = false;
  // A visitor who pauses it has decided; scrolling back must not overrule them.
  let held = !auto;

  const load = () => {
    if (loaded) return;
    loaded = true;
    video.preload = 'auto';
    video.load();
  };

  const sync = () => {
    const playing = !video.paused;
    reel.classList.toggle('is-paused', !playing);
    toggle.setAttribute('aria-label', playing ? 'Pause showreel' : 'Play showreel');
  };

  const play = () => {
    load();
    // Autoplay can still be refused (Low Power Mode on iOS). The poster stays
    // up and the toggle offers the play the browser would not.
    video.play().catch(sync);
  };

  video.addEventListener('playing', () => reel.classList.add('is-playing'));
  video.addEventListener('play', sync);
  video.addEventListener('pause', sync);

  toggle.hidden = false;
  toggle.addEventListener('click', () => {
    if (video.paused) {
      held = false;
      play();
    } else {
      held = true;
      video.pause();
    }
  });

  const observe = () => new IntersectionObserver(
    (entries) => {
      for (const entry of entries) {
        visible = entry.isIntersecting;
        if (visible && !held) play();
        else if (!visible && !video.paused) video.pause();
      }
    },
    // Start fetching a little before it arrives so the first frame is ready.
    { rootMargin: '300px 0px' },
  ).observe(reel);

  // The reel sits close to the fold, so without this it would start fetching
  // alongside the page's own first-paint resources.
  if (document.readyState === 'complete') observe();
  else window.addEventListener('load', observe, { once: true });

  sync();
}
