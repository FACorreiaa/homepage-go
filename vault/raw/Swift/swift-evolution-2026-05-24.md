---
source: https://github.com/swiftlang/swift-evolution
ingested_at: 2026-05-24T08:11:19Z
type: web
status: uncompiled
category: Swift
subcategory: swift-evolution
topic: Swift, language-design, proposals
recent_commits: 10
featured_proposals: 3
---

# Swift Evolution Proposals — 2026-05-24

## Recent Commits


- **2026-05-20** — Mark ST-0024 as implemented (#3303) (SE: N/A, PR: 3303)

- **2026-05-19** — rename replace(removing:...) to replaceSubrange(_:...) (#3302) (SE: N/A, PR: 3302)

- **2026-05-19** — [Proposal] Literal expressions (#3258) (SE: N/A, PR: 3258)

- **2026-05-18** — Update proposal status to Active Review (#3298) (SE: N/A, PR: 3298)

- **2026-05-18** — [SE-0526] Revision for cooperative cancellation and deadline composition (#3289) (SE: SE-0526, SE-0526, PR: 3289)

- **2026-05-18** — Merge pull request #3296 from glessard/0527-cleanup (SE: SE-0527, PR: 3296)

- **2026-05-15** — Amend SE-0519 to include `Sendable` and `BitwiseCopyable` conformances for `Ref` and `MutableRef` (SE: SE-0519, PR: N/A)

- **2026-05-14** — [SE-0527] tweak declaration formatting (SE: SE-0527, PR: N/A)

- **2026-05-14** — [SE-0527] fix some wording issues. (SE: SE-0527, PR: N/A)

- **2026-05-14** — [SE-0527] fix typos and text-editing mistakes (SE: SE-0527, PR: N/A)



## Featured Proposals


### [0519-borrow-inout-types.md](https://github.com/swiftlang/swift-evolution/blob/main/proposals/0519-borrow-inout-types.md)

**Status:** Unknown | **SE:** 0519


This document was renamed as part of the evolution process. Please see [SE-0519: `Ref` and `MutableRef` types for safe, first-class references](0519-ref-mutableref-types.md).


### [withDeadline](https://github.com/swiftlang/swift-evolution/blob/main/proposals/0526-deadline.md)

**Status:** Unknown | **SE:** 0526


# withDeadline

* Proposal: [SE-0526](0526-deadline.md)
* Authors: [Franz Busch](https://github.com/FranzBusch), [Philippe Hausler](https://github.com/phausler), [Konrad Malawski](https://github.com/ktoso)
* Status: **Active Review (May 18th...May 31st, 2026)**
* Implementation: https://github.com/swiftlang/swift/pull/88323
* Review Manager: [Freddy Kellison-Linn](https://github.com/Jumhyn)
* Review: ([pitch](https://forums.swift.org/t/pitch-withdeadline/85262)) ([review](https://forums.swift.org/t/se-0526-withdeadline/85850)) ([returned for revision](https://forums.swift.org/t/returned-for-revision-se-0526-withdeadline/86438))

## Summary of changes

This proposal introduces `withDeadline`, a function that executes asynchronous
operations with a composable absolute time limit. The function accepts a 
continuous clock instant representing the deadline by which the operation must 
complet


### [RigidArray and UniqueArray](https://github.com/swiftlang/swift-evolution/blob/main/proposals/0527-rigidarray-uniquearray.md)

**Status:** Unknown | **SE:** 0527


# RigidArray and UniqueArray

* Proposal: [SE-0527](0527-rigidarray-uniquearray.md)
* Authors: [Karoy Lorentey](https://github.com/lorentey), [Alejandro Alonso](https://github.com/Azoy)
* Review Manager: [Steve Canon](https://github.com/stephentyrone)
* Status: **Active Review (April 13...27, 2026)**
* Implementation: [swiftlang/swift#87521](https://github.com/swiftlang/swift/pull/87521)
* Review: ([pitch](https://forums.swift.org/t/pitch-rigidarray-and-uniquearray/85455))
          ([review](https://forums.swift.org/t/se-0527-rigidarray-and-uniquearray/85985))

[swift-collections]: https://github.com/apple/swift-collections

## Summary of changes

We propose to introduce two new array types to the Swift Standard Library,
`RigidArray` and `UniqueArray`, that are both capable of storing noncopyable
elements.

## Motivation

Swift 5.9 introduced noncopyable struct and enum types into the l

