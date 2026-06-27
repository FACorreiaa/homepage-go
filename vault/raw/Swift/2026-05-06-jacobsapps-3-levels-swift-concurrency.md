---
title: "The 3 Levels of Swift Concurrency"
url: "https://blog.jacobstechtavern.com/p/3-levels-of-concurrency"
author: "Jacob (Jacob's Tech Tavern)"
date: "2026-05-06"
ingested_at: "2026-05-13T16:27:39Z"
tags: [swift, concurrency, actor-isolation, approachable-concurrency, swift6]
topic: Swift
status: uncompiled
---

The 3 Levels of Swift Concurrency

Swift Concurrency is really bl**dy complicated, which is great for me, financially.

Swift's whole deal is making it **really hard to write broken code**: type-safe nullability, automatic memory management, and _compiler guarantees around data races_.

This is the whole reason your codebase suddenly grew 14,000,001 compiler warnings around non-Sendable classes. The Swift team were not being jerks, they wanted you to think about data isolation.

In a rare spurt of Apple humility, the team acknowledged this unintentional complexity in the [Approachable Concurrency manifesto](https://github.com/swiftlang/swift-evolution/blob/main/visions/approachable-concurrency.md):

Swift's built-in support for concurrency has three goals:

1. Extend memory safety guarantees to low-level data races.
2. Progressive disclosure, making basic use of concurrency easy.
3. Make advanced uses of concurrency to improve performance natural.

Swift meets the first goal, but it comes at the cost of the second, and can be frustrating to adopt.

Progressive disclosure is a core Apple design philosophy: _make basic options easy to access, but reveal advanced use cases when needed_. Think Dock → Finder → bash shell for working in your file system. The same approach applies to language design.

Easy to get started. Straightforward to level up when needed.

Thus the Swift 6.2 Approachable Concurrency story began.

Your concurrency journey now has 3 levels:

1. Code on the main actor by default. It's all single threaded.
2. Basic use of **async** functions, running on the main actor.
3. Explicit concurrency off-main actor via **@concurrent**.

Fundamentally, the new world revolves around 2 new settings in Xcode:

* Approachable Concurrency (**Yes**/No)
* Default Actor Isolation (**MainActor**/nonisolated)

The 3 levels don't _strictly_ have anything to do with concurrency: they're more about how tightly you're coupled to the main actor. From Xcode 26, new projects are opted into approachable concurrency and default MainActor isolation.

## Level 1: Main Actor by Default

In understanding Approachable Concurrency, you have to understand the incentives and motivations of Apple Inc., the business.

The _target audience_ for most Apple frameworks and features is the marginal developer. The one picking between SwiftUI and React Native, and most at risk of defecting to (grits teeth) Android or (shudder) _the web_ if their needs aren't met.

90% of budding indie devs don't give a sh*t about performance, because a workout tracker, AI recipe planner, or fart app doesn't need parallelism. It needs to work out of the box and find product-market fit.

By default, Swift now puts these devs in baby mode to protect them from themselves.

The safest possible starting point is a single-threaded app. And most apps don't actually _need_ much more than that.

Serial main actor execution is the first-class mode in Swift 6.2.

Concurrent mutation of shared state, and data races more broadly, are rendered impossible. This serialised code is also easy to reason about: everything runs on the main actor.

## Level 2: Async Functions on the Main Actor

We introduce async operations, but all our code, and critically, our data models, remain single-threaded, isolated to the main actor.

## Level 3: Explicit Concurrency Off-Main Actor

Via **@concurrent** for when you need actual parallelism.

---

Sample project: [The3Levels](https://github.com/jacobsapps/The3Levels)

More detail: [Swift Concurrency Course](https://blog.jacobstechtavern.com/p/swift-concurrency-course)
