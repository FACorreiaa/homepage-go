---
title: "Thread by @pointfreeco"
source: "https://x.com/pointfreeco/status/2044831732378316990"
author:
  - "[[@pointfreeco]]"
published: 2026-04-16
created: 2026-04-21
description: "Swift’s concurrency system has serious smarts: 6 “awaits” are squashed down to just 3 jobs enqueued when the scheduler sees it can avoid a h"
tags:
  - "clippings"
---
**Point-Free** @pointfreeco [2026-04-16](https://x.com/pointfreeco/status/2044831732378316990)

Swift’s concurrency system has serious smarts: 6 “awaits” are squashed down to just 3 jobs enqueued when the scheduler sees it can avoid a hop. We can verify by creating a custom executor that logs whenever the actor enqueues a job:

![An Xcode window highlighting 6 `await` but showing only 3 `Job enqueued` printed to the console.](https://pbs.twimg.com/media/HGCyf7JbMAAhWOh?format=jpg&name=large)

---

**Point-Free** @pointfreeco [2026-04-16](https://x.com/pointfreeco/status/2044831734467039309)

We dive deep into the topic in this week’s video. Watch today:

---

**Abilash S** @zeyrie\_dev [2026-04-17](https://x.com/zeyrie_dev/status/2044991675228168495)

I think I should start beige watching the Isolation series. Can’t procrastinate anymore 😭