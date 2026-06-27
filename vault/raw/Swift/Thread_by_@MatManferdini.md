---
title: "Thread by @MatManferdini"
source: "https://x.com/matmanferdini/status/2046286021210382383?s=48"
author:
  - "[[@MatManferdini]]"
published: 2026-04-20
created: 2026-04-21
description: "Asynchronous sequences are a core feature of Swift concurrency. They allow you to process asynchronous events easily using simple `for` loop"
tags:
  - "clippings"
---
**Matteo | Swift, iOS, Best Practices** @MatManferdini [2026-04-20](https://x.com/MatManferdini/status/2046286021210382383)

Asynchronous sequences are a core feature of Swift concurrency. They allow you to process asynchronous events easily using simple \`for\` loops.

However, conforming to the AsyncSequence protocol to create an asynchronous sequence is not straightforward.

That’s where AsyncStream comes to the rescue, allowing you to easily create asynchronous sequences and quickly adapt old APIs to Swift concurrency.