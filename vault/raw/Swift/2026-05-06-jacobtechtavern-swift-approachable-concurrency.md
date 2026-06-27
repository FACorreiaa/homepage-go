---
title: "Swift 6.2 Approachable Concurrency: 3 Tiers Explained"
url: "https://x.com/jacobtechtavern/status/2052057063325335912"
author: "Jacob Bartlett (@jacobtechtavern)"
date: 2026-05-06
ingested_at: 2026-05-13T16:22:46Z
tags:
  - swift
  - concurrency
  - approachable-concurrency
  - swift6
  - apple
  - progressive-disclosure
topic: Swift
status: uncompiled
---

# Swift 6.2 Approachable Concurrency: 3 Tiers Explained

Swift Concurrency is really bl**dy complicated, which is great for me, financially.

Swift's whole deal is making it really hard to write broken code: type-safe nullability, automatic memory management, and compiler guarantees around data races.

This is the whole reason your codebase suddenly grew 14,000,001 compiler warnings around non-Sendable classes. The Swift team were not being jerks, they wanted you to think about data isolation.

In a rare spurt of Apple humility, the team acknowledged this unintentional complexity in the Approachable Concurrency manifesto:
Swift's built-in support for concurrency has three goals:
1. Extend memory safety guarantees to low-level data races.
2. Progressive disclosure, making basic use of concurrency easy.
3. Make advanced uses of concurrency to improve performance natural.

Swift meets the first goal, but it comes at the cost of the second, and can be frustrating to adopt.

Progressive disclosure is a core Apple design philosophy: make basic options easy to access, but reveal advanced use cases when needed. Think Dock → Finder → bash shell for working in your file system. The same approach applies to language design.

Easy to get started.

Straightforward to level up when needed.

Thus the Swift 6.2 Approachable Concurrency story began.

Learn the 3 tiers of Swift 6.2 approachable concurrency here: https://blog.jacobstechtavern.com/p/3-levels-of-concurrency
