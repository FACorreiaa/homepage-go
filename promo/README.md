# Showreel

The 15-second reel on the home page, as code. `reel.html` is a GSAP timeline;
`render.mjs` seeks it frame by frame in headless Chrome and pipes the frames to
ffmpeg, so every render is identical and every frame is editable text.

```sh
npm install                       # playwright-core only; uses your installed Chrome
node render.mjs desktop --stills 4,11.7   # PNG stills in out/ for review
node render.mjs all               # masters in out/ (~10 min)
./encode.sh                       # web cuts into ../assets/static/promo/
```

Open `reel.html?w=1080&h=1920` in a browser to watch it live, or add `&t=7.2`
to park on one frame.

| Cut     | Size      | Where it goes                                  |
| ------- | --------- | ---------------------------------------------- |
| desktop | 1920×1080 | home page (≥768px), YouTube, LinkedIn          |
| iphone  | 1170×2532 | home page (<768px, encoded at 720×1558)        |
| shorts  | 1080×1920 | YouTube Shorts, Reels, TikTok                  |

Each output frame averages four subframes (`SUB=4`), which is real motion
blur. `SUB=1` renders four times faster for drafts.

The reel names the four live products. When that list changes, edit the
product sections in `reel.html`, re-render, re-encode.
