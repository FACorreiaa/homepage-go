---
source: https://www.swift.org/swift-evolution/
ingested_at: 2026-05-16T08:08:24Z
type: web
status: uncompiled
category: Swift
subcategory: Swift Evolution
title: "Swift Evolution Proposals - Active Reviews & Recent Acceptances"
---
# Swift Evolution Proposals

**Ingested:** 2026-05-16T08:08:24Z
**Source:** [Swift Evolution](https://www.swift.org/swift-evolution/)
**Total Proposals:** 530

## Active Reviews

- **SE-0529**: Add `FilePath` to the Standard Library — Authors: Michael Ilseman, Saleem Abdulrasool

## Recently Accepted

- **SE-0527**: RigidArray and UniqueArray — Authors: Karoy Lorentey, Alejandro Alonso
- **SE-0530**: Async Result Support — Authors: Konrad 'ktoso' Malawski, Matt Massicotte
- **SE-0525**: Safe loading API for `RawSpan` — Authors: Guillaume Lessard
- **SE-0522**: Source-Level Control Over Compiler Warnings — Authors: Artem Chikin, Doug Gregor, Holly Borla
- **SE-0521**: Improved Syntax for Optionals of Opaque and Existential Types — Authors: Tony Allevato
- **SE-0520**: Discardable result use in Task initializers — Authors: Konrad 'ktoso' Malawski
- **SE-0517**: UniqueBox — Authors: Alejandro Alonso
- **SE-0515**: Allow `reduce` to produce noncopyable results — Authors: Ben Cohen
- **SE-0511**: SwiftPM Add Target Plugin Command — Authors: Gage Halverson
- **SE-0510**: Introduce `Dictionary.mapValuesWithKeys` — Authors: Diana Ma
- **SE-0506**: Advanced Observation Tracking — Authors: Philippe Hausler
- **SE-0503**: Suppressed Default Conformances on Associated Types With Defaults — Authors: Kavon Farvardin
- **SE-0501**: HTML Coverage Report — Authors: Sam Khouri
- **SE-0500**: Improving package creation with custom templates: SwiftPM Template Initialization — Authors: John Bute
- **SE-0484**: Allow Additional Arguments to `@dynamicMemberLookup` Subscripts — Authors: Itai Ferber
- **SE-0474**: Yielding accessors — Authors: Ben Cohen, Nate Chandler, Joe Groff
- **SE-0455**: SwiftPM @testable build setting — Authors: Jake Petroules
- **SE-0454**: Custom Allocator for Toolchain — Authors: Saleem Abdulrasool
- **SE-0448**: Regex lookbehind assertions — Authors: Jacob Hearst, Michael Ilseman
- **SE-0419**: Swift Backtrace API — Authors: Alastair Houghton

## Implemented

- **SE-0288**: Adding `isPower(of:)` to `BinaryInteger`
- **SE-0524**: Add `withTemporaryAllocation` using `Output(Raw)Span`
- **SE-0523**: Hashable conformance for `UnownedTaskExecutor`
- **SE-0519**: `Ref` and `MutableRef` types for safe, first-class references
- **SE-0518**: `~Sendable` for explicitly marking non-`Sendable` types
- **SE-0514**: `Hashable` Conformance for `Dictionary.Keys`, `CollectionOfOne` and `EmptyCollection`
- **SE-0512**: Document that `Mutex.withLockIfAvailable(_:)` cannot spuriously fail
- **SE-0509**: Software Bill of Materials (SBOM) Generation for Swift Package Manager
- **SE-0508**: Array expression trailing closures
- **SE-0507**: Borrow and Mutate Accessors

## All Proposals by Status

### Active Reviews (1)
- SE-0529: Add `FilePath` to the Standard Library (Michael Ilseman, Saleem Abdulrasool)

### Accepted (23)
- SE-0527: RigidArray and UniqueArray (Karoy Lorentey, Alejandro Alonso)
- SE-0530: Async Result Support (Konrad 'ktoso' Malawski, Matt Massicotte)
- SE-0525: Safe loading API for `RawSpan` (Guillaume Lessard)
- SE-0522: Source-Level Control Over Compiler Warnings (Artem Chikin, Doug Gregor, Holly Borla)
- SE-0521: Improved Syntax for Optionals of Opaque and Existential Types (Tony Allevato)
- SE-0520: Discardable result use in Task initializers (Konrad 'ktoso' Malawski)
- SE-0517: UniqueBox (Alejandro Alonso)
- SE-0515: Allow `reduce` to produce noncopyable results (Ben Cohen)
- SE-0511: SwiftPM Add Target Plugin Command (Gage Halverson)
- SE-0510: Introduce `Dictionary.mapValuesWithKeys` (Diana Ma)
- SE-0506: Advanced Observation Tracking (Philippe Hausler)
- SE-0503: Suppressed Default Conformances on Associated Types With Defaults (Kavon Farvardin)
- SE-0501: HTML Coverage Report (Sam Khouri)
- SE-0500: Improving package creation with custom templates: SwiftPM Template Initialization (John Bute)
- SE-0484: Allow Additional Arguments to `@dynamicMemberLookup` Subscripts (Itai Ferber)
- SE-0474: Yielding accessors (Ben Cohen, Nate Chandler, Joe Groff)
- SE-0455: SwiftPM @testable build setting (Jake Petroules)
- SE-0454: Custom Allocator for Toolchain (Saleem Abdulrasool)
- SE-0448: Regex lookbehind assertions (Jacob Hearst, Michael Ilseman)
- SE-0419: Swift Backtrace API (Alastair Houghton)
- SE-0342: Statically link Swift runtime libraries by default on supported platforms (neonichu, tomerd)
- SE-0321: Package Registry Service - Publish Endpoint (Whitney Imura, Mattt Zmuda)
- SE-0528: `Continuation` — Safe and Performant Async Continuations (Fabian Fett, Konrad Malawski)

### Other Status
- SE-0001: Allow (most) keywords as argument labels [Returned]
- SE-0526: withDeadline [Returned]
- SE-0516: Borrowing Sequence [Returned]
- SE-0513: API to get the path to the current executable [Returned]
- SE-0505: Delayed Enqueuing for Executors [Returned]
- SE-0490: Environment Constrained Shared Libraries [Returned]
- SE-0479: Method and Initializer Key Paths [Returned]
- SE-0478: Default actor isolation typealias [Returned]
- SE-0406: Backpressure support for AsyncStream [Returned]
- SE-0403: Package Manager Mixed Language Target Support [Returned]
- SE-0385: Custom Reflection Metadata [Returned]
- SE-0379: Swift Opt-In Reflection Metadata [Returned]
- SE-0359: Build-Time Constant Values [Returned]
- SE-0330: Conditionals in Collections [Returned]
- SE-0318: Package Creation [Returned]
- SE-0312: Add `indexed()` and `Collection` conformances for `enumerated()` and `zip(_:_:)` [Returned]
- SE-0283: Tuples Conform to `Equatable`, `Comparable`, and `Hashable` [Returned]
- SE-0265: Offset-Based Access to Indices, Elements, and Slices [Returned]
- SE-0259: Approximate Equality for Floating Point [Returned]
- SE-0257: Eliding commas from multiline expression lists [Returned]

---

*Ingested as part of the scheduled Swift/iOS market intelligence monitoring.*
