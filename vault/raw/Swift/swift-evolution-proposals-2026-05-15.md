---
source: Swift Evolution Proposals 2026-05-15
ingested_at: 2026-05-15T08:04:06Z
type: market-intelligence
category: swift
subcategory: swift-evolution
status: uncompiled
---
Title: Swift.org

URL Source: http://swift.org/swift-evolution/

Markdown Content:
Docs
Community
Packages
Blog
Install (6.3.2)
Swift Evolution

Anyone with a good idea can help shape the future features and direction of the language. To reach the best possible solution to a problem, we discuss and iterate on ideas in a public forum. Once a proposal is refined and approved, it becomes a release goal, and is tracked as a feature of an upcoming version of Swift.

To support this process, the Swift Evolution repository collects the goals for the upcoming major and minor releases (as defined by the core team) as well as proposals for changes to Swift. The Swift evolution process document details how ideas are proposed, discussed, reviewed, and eventually accepted into upcoming releases.

Below is a list of all the current and upcoming proposal reviews.

530 proposals
Active Review
SE-0529Add `FilePath` to the Standard Library
Authors:Michael Ilseman, Saleem Abdulrasool
Review Manager:
John McCall
Status:
Active Review
Scheduled:
April 22 – May 4
Active Review
SE-0527RigidArray and UniqueArray
Authors:Karoy Lorentey, Alejandro Alonso
Review Manager:
Steve Canon
Implementation:
swift#87521
Status:
Active Review
Scheduled:
April 13 – 27
Accepted
SE-0530Async Result Support
Authors:Konrad 'ktoso' Malawski, Matt Massicotte
Review Manager:
Doug Gregor
Implementation:
swift#88465
Accepted
SE-0525Safe loading API for `RawSpan`
Author:Guillaume Lessard
Review Manager:
Xiaodi Wu
Implementation:
swift#88304
Accepted
SE-0522Source-Level Control Over Compiler Warnings
Authors:Artem Chikin, Doug Gregor, Holly Borla
Review Manager:
Tony Allevato
Implementation:
swift#85036, swift-syntax#3174
Accepted
SE-0521Improved Syntax for Optionals of Opaque and Existential Types
Author:Tony Allevato
Review Manager:
Freddy Kellison-Linn
Implementation:
swift#87115, swift-syntax#3268
Accepted
SE-0520Discardable result use in Task initializers
Author:Konrad 'ktoso' Malawski
Review Manager:
Holly Borla
Implementation:
swift#87439
Accepted
SE-0517UniqueBox
Author:Alejandro Alonso
Review Manager:
Ben Cohen
Implementation:
swift#86336
Accepted
SE-0515Allow `reduce` to produce noncopyable results
Author:Ben Cohen
Review Manager:
Joe Groff
Implementation:
swift#85716
Accepted
SE-0511SwiftPM Add Target Plugin Command
Author:Gage Halverson
Review Manager:
Mikaela Caron
Bug:
swiftlang/swift-package-manager#8169
Implementation:
swift-package-manager#8432
Accepted
SE-0510Introduce `Dictionary.mapValuesWithKeys`
Author:Diana Ma
Review Manager:
Steve Canon
Implementation:
swift#86268
Accepted
SE-0506Advanced Observation Tracking
Author:Philippe Hausler
Review Manager:
Steve Canon
Accepted
SE-0503Suppressed Default Conformances on Associated Types With Defaults
Author:Kavon Farvardin
Review Manager:
Ben Cohen
Accepted
SE-0501HTML Coverage Report
Author:Sam Khouri
Review Manager:
David Cummings
Implementation:
swift-package-manager#9076
Accepted
SE-0500Improving package creation with custom templates: SwiftPM Template Initialization
Author:John Bute
Review Manager:
Franz Busch
Implementation:
swift-package-manager#9211
Accepted
SE-0484Allow Additional Arguments to `@dynamicMemberLookup` Subscripts
Author:Itai Ferber
Review Manager:
Xiaodi Wu
Implementation:
swift#81148
Accepted
SE-0474Yielding accessors
Authors:Ben Cohen, Nate Chandler, Joe Groff
Review Manager:
Steve Canon
Accepted
SE-0455SwiftPM @testable build setting
Author:Jake Petroules
Review Manager:
Tony Allevato
Implementation:
swift-package-manager#8004
Accepted
SE-0454Custom Allocator for Toolchain
Author:Saleem Abdulrasool
Review Manager:
Alastair Houghton
Implementation:
swift#76563
Accepted
SE-0448Regex lookbehind assertions
Authors:Jacob Hearst, Michael Ilseman
Review Manager:
Steve Canon
Accepted
SE-0419Swift Backtrace API
Author:Alastair Houghton
Review Manager:
Steve Canon
Accepted
SE-0342Statically link Swift runtime libraries by default on supported platforms
Authors:neonichu, tomerd
Review Manager:
Ted Kremenek
Implementation:
swift-package-manager#3905
Accepted
SE-0321Package Registry Service - Publish Endpoint
Authors:Whitney Imura, Mattt Zmuda
Review Manager:
Tom Doron
Implementation:
swift-package-manager#3671
Accepted
SE-0528`Continuation` — Safe and Performant Async Continuations
Authors:Fabian Fett, Konrad Malawski
Review Manager:
Joe Groff
Implementation:
swift#88182
Status:
Accepted with revisions
Previewing
SE-0288Adding `isPower(of:)` to `BinaryInteger`
Author:Ding Ye
Review Manager:
Joe Groff
Preview:
Standard Library Preview
Implementation:
swift#24766
Implemented
SE-0524Add `withTemporaryAllocation` using `Output(Raw)Span`
Author:Max Desiatov
Review Manager:
Doug Gregor
Implemented In:
Swift 6.4
Implementation:
swift#85866
Implemented
SE-0523Hashable conformance for `UnownedTaskExecutor`
Authors:Fabian Fett, Konrad Malawski
Review Manager:
John McCall
Implemented In:
Swift 6.4
Implementation:
swift#73968
Implemented
SE-0519`Ref` and `MutableRef` types for safe, first-class references
Authors:Joe Groff, Alejandro Alonso
Review Manager:
Doug Gregor
Implemented In:
Swift 6.4
Implementation:
swift#87222
Implemented
SE-0518`~Sendable` for explicitly marking non-`Sendable` types
Author:Pavel Yaskevich
Review Manager:
Xiaodi Wu
Implemented In:
Swift 6.4
Implementation:
swift#84777, swift#85105
Implemented
SE-0514`Hashable` Conformance for `Dictionary.Keys`, `CollectionOfOne` and `EmptyCollection`
Author:Clinton Nkwocha
Review Manager:
Steve Canon
Implemented In:
Swift 6.4
Implementation:
swift#86899
Implemented
SE-0512Document that `Mutex.withLockIfAvailable(_:)` cannot spuriously fail
Author:Jonathan Grynspan
Review Manager:
John McCall
Implemented In:
Swift 6.0
Implementation:
swift#85497
Implemented
SE-0509Software Bill of Materials (SBOM) Generation for Swift Package Manager
Author:Ev Cheng
Review Manager:
Franz Busch
Implemented In:
Swift 6.4
Implementation:
swift-package-manager#9633
Implemented
SE-0508Array expression trailing closures
Author:Cal Stephens
Review Manager:
Xiaodi Wu
Implemented In:
Swift 6.4
Implementation:
swift#86244
Implemented
SE-0507Borrow and Mutate Accessors
Authors:Meghana Gupta, Tim Kientzle
Review Manager:
Doug Gregor
Implemented In:
Swift 6.4
Implemented
SE-0504Task Cancellation Shields
Author:Konrad 'ktoso' Malawski
Review Manager:
John McCall
Implemented In:
Swift 6.4
Implementation:
swift#85637
Implemented
SE-0502Exclude private initialized properties from memberwise initializer
Authors:Hamish Knight, Holly Borla
Review Manager:
Tony Allevato
Implemented In:
Swift 6.4
Implementation:
swift#84514, swift#87028
Implemented
SE-0499Support ~Copyable, ~Escapable in simple standard library protocols
Author:Ben Cohen
Review Manager:
Holly Borla
Implemented In:
Swift 6.4
Implementation:
swift#85079
Implemented
SE-0498Expose demangle function in Runtime module
Authors:Konrad 'ktoso' Malawski, Alejandro Alonso
Review Manager:
Steve Canon
Implemented In:
Swift 6.4
Implementation:
swift#84788
Implemented
SE-0497Controlling function definition visibility in clients
Author:Doug Gregor
Review Manager:
Becca Royal-Gordon
Implemented In:
Swift 6.3
Implemented
SE-0496`@inline(always)` attribute
Author:Arnold Schwaighofer
Review Manager:
Tony Allevato
Implemented In:
Swift 6.3
Implementation:
swift#84178
Implemented
SE-0495C compatible functions and enums
Author:Alexis Laferrière
Review Manager:
Steve Canon
Implemented In:
Swift 6.3
Implemented
SE-0494Add `isTriviallyIdentical(to:)` Methods for Quick Comparisons to Concrete Types
Authors:Rick van Voorden, Karoy Lorentey
Review Manager:
John McCall
Implemented In:
Swift 6.4
Implementation:
swift#82055, swift#82438, swift#82439, swift#84998
Implemented
SE-0493Support `async` calls in `defer` bodies
Author:Freddy Kellison-Linn
Review Manager:
Holly Borla
Implemented In:
Swift 6.4
Implementation:
swift#83891
Implemented
SE-0492Section Placement Control
Author:Kuba Mracek
Review Manager:
Doug Gregor
Implemented In:
Swift 6.3
Implemented
SE-0491Module selectors for name disambiguation
Author:Becca Royal-Gordon
Review Manager:
Freddy Kellison-Linn
Bug:
swiftlang/swift#53580
Implemented In:
Swift 6.3
Implementation:
swift#34556
Implemented
SE-0489Improve `EncodingError` and `DecodingError`'s printed descriptions
Author:Zev Eisenberg
Review Manager:
Xiaodi Wu
Implemented In:
Swift 6.3
Implemented
SE-0488Apply the extracting() slicing pattern more widely
Author:Guillaume Lessard
Review Manager:
Tony Allevato
Implemented In:
Swift 6.2
Implemented
SE-0487Nonexhaustive enums
Authors:Pavel Yaskevich, Franz Busch, Cory Benfield
Review Manager:
Ben Cohen
Bug:
apple/swift#55110
Implemented In:
Swift 6.2.3
Implementation:
swift#80503
Implemented
SE-0486Migration tooling for Swift features
Authors:Anthony Latsis, Pavel Yaskevich
Review Manager:
Franz Busch
Implemented In:
Swift 6.2
Implemented
SE-0485OutputSpan: delegate initialization of contiguous memory
Author:Guillaume Lessard
Review Manager:
Doug Gregor
Implemented In:
Swift 6.2
Implementation:
swift#81637
Implemented
SE-0483`InlineArray` Type Sugar
Authors:Hamish Knight, Ben Cohen
Review Manager:
Holly Borla
Implemented In:
Swift 6.2
Implemented
SE-0482Binary Static Library Dependencies
Authors:Daniel Grumberg, Max Desiatov, Franz Busch
Review Manager:
Kuba Mracek
Bug:
Swift Package Manger Issue
Implemented In:
Swift 6.2
Implementation:
swift-package-manager#8639, swift-package-manager#8741
Implemented
SE-0481`weak let`
Author:Mykola Pokhylets
Review Manager:
John McCall
Implemented In:
Swift 6.3
Implementation:
swift#80440
Upcoming Feature Flag:
ImmutableWeakCaptures
Implemented
SE-0480Warning Control Settings for SwiftPM
Author:Dmitrii Galimzianov
Review Managers:
John McCall, Franz Busch
Implemented In:
Swift 6.2
Implementation:
swift-package-manager#8315
Implemented
SE-0477Default Value in String Interpolations
Author:Nate Cook
Review Manager:
Xiaodi Wu
Implemented In:
Swift 6.2
Implementation:
swift#80547
Implemented
SE-0476Controlling the ABI of a function, initializer, property, or subscript
Author:Becca Royal-Gordon
Review Manager:
Holly Borla
Implemented In:
Swift 6.2
Implemented
SE-0475Transactional Observation of Values
Author:Philippe Hausler
Review Manager:
Freddy Kellison-Linn
Implemented In:
Swift 6.2
Implemented
SE-0473Clock Epochs
Author:Philippe Hausler
Review Manager:
John McCall
Implemented In:
Swift 6.3
Implementation:
swift#80409
Implemented
SE-0472Starting tasks synchronously from caller context
Author:Konrad 'ktoso' Malawski
Review Manager:
Tony Allevato
Implemented In:
Swift 6.2
Implemented
SE-0471Improved Custom SerialExecutor isolation checking for Concurrency Runtime
Author:Konrad 'ktoso' Malawski
Review Manager:
Doug Gregor
Implemented In:
Swift 6.2
Implemented
SE-0470Global-actor isolated conformances
Author:Doug Gregor
Review Manager:
Xiaodi Wu

