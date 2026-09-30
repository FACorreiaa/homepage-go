// Renders reel.html to video, one seeked frame at a time.
//
//   node render.mjs desktop            master MP4 for one format
//   node render.mjs all                every master
//   node render.mjs desktop --stills 2,5.5,12   PNG stills for review
//
// Motion blur is real, not a filter: each output frame is the average of SUB
// rendered subframes spread across the frame interval (a 360° shutter), so a
// whip pan smears the way a camera would and a still title stays razor sharp.
import { chromium } from 'playwright-core';
import { spawn } from 'node:child_process';
import { mkdirSync } from 'node:fs';
import { fileURLToPath, pathToFileURL } from 'node:url';
import path from 'node:path';

const here = path.dirname(fileURLToPath(import.meta.url));
const FORMATS = {
  desktop: { w: 1920, h: 1080 },
  shorts: { w: 1080, h: 1920 },
  iphone: { w: 1170, h: 2532 },
};
const FPS = 30;
const SUB = +(process.env.SUB || 4);
const CHROME = process.env.CHROME || '/Applications/Google Chrome.app/Contents/MacOS/Google Chrome';

const args = process.argv.slice(2);
const which = args[0] === 'all' ? Object.keys(FORMATS) : [args[0] || 'desktop'];
const stillsArg = args.includes('--stills') ? args[args.indexOf('--stills') + 1] : null;
const out = path.join(here, 'out');
mkdirSync(out, { recursive: true });

const browser = await chromium.launch({ executablePath: CHROME, args: ['--allow-file-access-from-files', '--force-color-profile=srgb', '--hide-scrollbars'] });

async function open(name) {
  const { w, h } = FORMATS[name];
  const page = await browser.newPage({ viewport: { width: w, height: h }, deviceScaleFactor: 1 });
  page.on('pageerror', e => console.error(`[${name}] page error:`, e.message));
  await page.goto(pathToFileURL(path.join(here, 'reel.html')).href + `?w=${w}&h=${h}&render=1`);
  await page.evaluate(() => window.__ready);
  return page;
}

async function renderStills(name, times) {
  const page = await open(name);
  for (const t of times) {
    await page.evaluate(t => window.__seek(t), t);
    const file = path.join(out, `${name}-${String(t).replace('.', '_')}.png`);
    await page.screenshot({ path: file });
    console.log(file);
  }
  await page.close();
}

async function renderVideo(name) {
  const { w, h } = FORMATS[name];
  const page = await open(name);
  const dur = await page.evaluate(() => window.__dur);
  const total = Math.round(dur * FPS * SUB);
  const file = path.join(out, `reel-${name}-${w}x${h}.mp4`);
  const ff = spawn('ffmpeg', [
    '-v', 'error', '-y',
    '-f', 'image2pipe', '-framerate', String(FPS * SUB), '-c:v', 'png', '-i', '-',
    '-vf', `tmix=frames=${SUB},select='eq(mod(n\\,${SUB})\\,${SUB - 1})',setpts=N/${FPS}/TB,format=yuv420p`,
    '-r', String(FPS), '-c:v', 'libx264', '-preset', 'slow', '-crf', '14',
    '-colorspace', 'bt709', '-color_primaries', 'bt709', '-color_trc', 'bt709',
    '-movflags', '+faststart', file,
  ], { stdio: ['pipe', 'inherit', 'inherit'] });
  const done = new Promise((res, rej) => ff.on('close', c => (c === 0 ? res() : rej(new Error(`ffmpeg exited ${c}`)))));
  const started = Date.now();
  for (let i = 0; i < total; i++) {
    await page.evaluate(t => window.__seek(t), i / (FPS * SUB));
    const png = await page.screenshot({ type: 'png' });
    if (!ff.stdin.write(png)) await new Promise(r => ff.stdin.once('drain', r));
    if (i % (FPS * SUB) === 0) console.log(`[${name}] ${(i / (FPS * SUB)).toFixed(0)}s / ${dur}s  (${((Date.now() - started) / 1000).toFixed(0)}s elapsed)`);
  }
  ff.stdin.end();
  await done;
  await page.close();
  console.log(file);
}

if (stillsArg) await Promise.all(which.map(n => renderStills(n, stillsArg.split(',').map(Number))));
else await Promise.all(which.map(renderVideo));
await browser.close();
