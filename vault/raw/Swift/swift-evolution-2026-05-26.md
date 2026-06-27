---
source: swift.org/swift-evolution
ingested_at: 2026-05-26T08:06:17Z
type: market-intelligence
category: Swift
subcategory: swift-evolution
status: uncompiled
---

Title: Swift.org

URL Source: http://www.swift.org/swift-evolution/

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

531 proposals
Active Review
SE-0531Literal Expressions
Authors:Artem Chikin, Doug Gregor
Review Manager:
Ben Cohen
Implementation:
swift#86500, swift#86934, swift#87006
Status:
Active Review
Scheduled:
May 18 – 29
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
Active Review
SE-0526withDeadline
Authors:Franz Busch, Philippe Hausler, Konrad Malawski
Review Manager:
Freddy Kellison-Linn
Status:
Active Review
Scheduled:
May 18 – 31
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
Implemented In:
Swift 6.2
Upcoming Feature Flag:
InferIsolatedConformances
Implemented
SE-0469Task Naming
Authors:Konrad Malawski, Harjas Monga
Review Manager:
Holly Borla
Implemented In:
Swift 6.2
Implementation:
swift#79600
Implemented
SE-0468`Hashable` conformance for `Async(Throwing)Stream.Continuation`
Author:Mykola Pokhylets
Review Manager:
Freddy Kellison-Linn
Implemented In:
Swift 6.2
Implementation:
swift#79457
Implemented
SE-0467MutableSpan and MutableRawSpan: delegate mutations of contiguous memory
Author:Guillaume Lessard
Review Manager:
Joe Groff
Implemented In:
Swift 6.2
Implementation:
swift#79650, swift#80517
Implemented
SE-0466Control default actor isolation inference
Authors:Holly Borla, Doug Gregor
Review Manager:
Steve Canon
Implemented In:
Swift 6.2
Implemented
SE-0465Standard Library Primitives for Nonescapable Types
Author:Karoy Lorentey
Review Manager:
Doug Gregor
Implemented In:
Swift 6.2
Implemented
SE-0464UTF8Span: Safe UTF-8 Processing Over Contiguous Bytes
Authors:Michael Ilseman, Guillaume Lessard
Review Manager:
Tony Allevato
Implemented In:
Swift 6.2
Implementation:
swift#78531
Implemented
SE-0463Import Objective-C completion handler parameters as `@Sendable`
Author:Holly Borla
Review Manager:
John McCall
Implemented In:
Swift 6.2
Implemented
SE-0462Task Priority Escalation APIs
Author:Konrad 'ktoso' Malawski
Review Manager:
Freddy Kellison-Linn
Implemented In:
Swift 6.2
Implemented
SE-0461Run nonisolated async functions on the caller's actor by default
Authors:Holly Borla, John McCall
Review Manager:
Xiaodi Wu
Implemented In:
Swift 6.2
Upcoming Feature Flag:
NonisolatedNonsendingByDefault
Implemented
SE-0460Explicit Specialization
Author:Ben Cohen
Review Manager:
Steve Canon
Implemented In:
Swift 6.3
Implemented
SE-0459Add `Collection` conformances for `enumerated()`
Author:Alejandro Alonso
Review Manager:
Ben Cohen
Implemented In:
Swift 6.2
Implementation:
swift#78092
Implemented
SE-0458Opt-in Strict Memory Safety Checking
Author:Doug Gregor
Review Manager:
John McCall
Implemented In:
Swift 6.2
Implemented
SE-0457Expose attosecond representation of `Duration`
Author:Philipp Gabriel
Review Manager:
Stephen Canon
Implemented In:
Swift 6.2
Implementation:
swift#78202
Implemented
SE-0456Add `Span`-providing Properties to Standard Library Types
Author:Guillaume Lessard
Review Manager:
Doug Gregor
Implemented In:
Swift 6.2
Implementation:
swift#78561, swift#80116, swift-foundation#1276
Implemented
SE-0453InlineArray, a fixed-size array
Author:Alejandro Alonso
Review Manager:
Freddy Kellison-Linn
Implemented In:
Swift 6.2
Implementation:
swift#76438
Implemented
SE-0452Integer Generic Parameters
Authors:Alejandro Alonso, Joe Groff
Review Manager:
Ben Cohen
Implemented In:
Swift 6.2
Implementation:
swift#75518, swift#78248
Implemented
SE-0451Raw identifiers
Author:Tony Allevato
Review Manager:
Joe Groff
Implemented In:
Swift 6.2
Implementation:
swift#76636, swift-syntax#2857
Implemented
SE-0450Package traits
Authors:Franz Busch, Max Desiatov
Review Manager:
Mishal Shah
Implemented In:
Swift 6.1
Implemented
SE-0449Allow `nonisolated` to prevent global actor inference
Authors:Sima Nerush, Holly Borla
Review Manager:
Tony Allevato
Implemented In:
Swift 6.1
Implemented
SE-0447Span: Safe Access to Contiguous Storage
Authors:Guillaume Lessard, Michael Ilseman, Andrew Trick
Review Manager:
Doug Gregor
Implemented In:
Swift 6.2
Implementation:
swift#76406
Implemented
SE-0446Nonescapable Types
Authors:Andrew Trick, Tim Kientzle
Review Manager:
Joe Groff
Implemented In:
Swift 6.2
Upcoming Feature Flag:
NonescapableTypes
Implemented
SE-0445Improving `String.Index`'s printed descriptions
Author:Karoy Lorentey
Review Manager:
Xiaodi Wu
Implemented In:
Swift 6.1
Implementation:
swift#75433
Implemented
SE-0444Member import visibility
Author:Allan Shortlidge
Review Manager:
Becca Royal-Gordon
Bug:
apple/swift#46493
Implemented In:
Swift 6.1
Implementation:
swift#72974, swift#73063
Upcoming Feature Flag:
MemberImportVisibility
Implemented
SE-0443Precise Control Flags over Compiler Warnings
Authors:Doug Gregor, Dmitrii Galimzianov
Review Manager:
John McCall
Implemented In:
Swift 6.1
Implementation:
swift#74466
Implemented
SE-0442Allow TaskGroup's ChildTaskResult Type To Be Inferred
Author:Richard L Zarth III
Review Manager:
Doug Gregor
Implemented In:
Swift 6.1
Implementation:
swift#74517
Implemented
SE-0441Formalize ‘language mode’ terminology
Author:James Dempsey
Review Manager:
John McCall
Implemented In:
Swift 6.1
Implementation:
swift-package-manager#7620, swift#75564
Implemented
SE-0440DebugDescription Macro
Author:Dave Lee
Review Manager:
Steve Canon
Implemented In:
Swift 6.0
Implementation:
swift#69626
Implemented
SE-0439Allow trailing comma in comma-separated lists
Author:Mateus Rodrigues
Review Manager:
Xiaodi Wu
Implemented In:
Swift 6.1
Implementation:
swift#74522
Implemented
SE-0438Metatype Keypaths
Authors:Amritpan Kaur, Pavel Yaskevich
Review Manager:
Joe Groff
Implemented In:
Swift 6.1
Implementation:
swift#73242
Implemented
SE-0437Noncopyable Standard Library Primitives
Author:Karoy Lorentey
Review Manager:
John McCall
Implemented In:
Swift 6.0
Implemented
SE-0436Objective-C implementations in Swift
Author:Becca Royal-Gordon
Review Manager:
Freddy Kellison-Linn
Implemented In:
Swift 6.1
Implementation:
swift#73309, swift#74801
Implemented
SE-0435Swift Language Version Per Target
Author:Pavel Yaskevich
Review Manager:
Becca Royal-Gordon
Implemented In:
Swift 6.0
Implemented
SE-0434Usability of global-actor-isolated types
Authors:Sima Nerush, Matt Massicotte, Holly Borla
Review Manager:
John McCall
Implemented In:
Swift 6.0
Upcoming Feature Flag:
GlobalActorIsolatedTypesUsability
Implemented
SE-0433Synchronous Mutual Exclusion Lock 🔒
Author:Alejandro Alonso
Review Manager:
Stephen Canon
Implemented In:
Swift 6.0
Implemented
SE-0432Borrowing and consuming pattern matching for noncopyable types
Author:Joe Groff
Review Manager:
Ben Cohen
Implemented In:
Swift 6.0
Implemented
SE-0431`@isolated(any)` Function Types
Author:John McCall
Review Manager:
Doug Gregor
Implemented In:
Swift 6.0
Implemented
SE-0430`sending` parameter and result values
Authors:Michael Gottesman, Holly Borla, John McCall
Review Manager:
Becca Royal-Gordon
Implemented In:
Swift 6.0
Implemented
SE-0429Partial consumption of noncopyable values
Authors:Michael Gottesman, Nate Chandler
Review Manager:
Xiaodi Wu
Implemented In:
Swift 6.0
Implemented
SE-0428Resolve DistributedActor protocols
Authors:Konrad 'ktoso' Malawski, Pavel Yaskevich
Review Manager:
Freddy Kellison-Linn
Implemented In:
Swift 6.0
Implemented
SE-0427Noncopyable Generics
Authors:Kavon Farvardin, Tim Kientzle, Slava Pestov
Review Managers:
Holly Borla, Ben Cohen
Implemented In:
Swift 6.0
Implemented
SE-0426BitwiseCopyable
Authors:Kavon Farvardin, Guillaume Lessard, Nate Chandler, Tim Kientzle
Review Manager:
Tony Allevato
Implemented In:
Swift 6.0
Implemented
SE-0425128-bit Integer Types
Author:Stephen Canon
Review Manager:
Doug Gregor
Implemented In:
Swift 6.0
Implemented
SE-0424Custom isolation checking for SerialExecutor
Author:Konrad 'ktoso' Malawski
Review Manager:
John McCall
Implemented In:
Swift 6.0
Implemented
SE-0423Dynamic actor isolation enforcement from non-strict-concurrency contexts
Authors:Holly Borla, Pavel Yaskevich
Review Manager:
Ben Cohen
Implemented In:
Swift 6.0
Upcoming Feature Flag:
DynamicActorIsolation
Implemented
SE-0422Expression macro as caller-side default argument
Author:Apollo Zhu
Review Manager:
Doug Gregor
Implemented In:
Swift 6.0
Implemented
SE-0421Generalize effect polymorphism for `AsyncSequence` and `AsyncIteratorProtocol`
Authors:Doug Gregor, Holly Borla
Review Manager:
Freddy Kellison-Linn
Implemented In:
Swift 6.0
Implemented
SE-0420Inheritance of actor isolation
Authors:John McCall, Holly Borla, Doug Gregor
Review Manager:
Xiaodi Wu
Implemented In:
Swift 6.0
Implemented
SE-0418Inferring `Sendable` for methods and key path literals
Authors:Angela Laar, Kavon Farvardin, Pavel Yaskevich
Review Manager:
Becca Royal-Gordon
Implemented In:
Swift 6.0
Upcoming Feature Flag:
InferSendableFromCaptures
Implemented
SE-0417Task Executor Preference
Authors:Konrad 'ktoso' Malawski, John McCall, Franz Busch
Review Manager:
Doug Gregor
Implemented In:
Swift 6.0
Implemented
SE-0416Subtyping for keypath literals as functions
Author:Frederick Kellison-Linn
Review Manager:
John McCall
Implemented In:
Swift 6.0
Implementation:
swift#39612
Implemented
SE-0415Function Body Macros
Author:Doug Gregor
Review Manager:
Tony Allevato
Implemented In:
Swift 6.0
Implemented
SE-0414Region based Isolation
Authors:Michael Gottesman, Joshua Turcotti
Review Manager:
Holly Borla
Implemented In:
Swift 6.0
Upcoming Feature Flag:
RegionBasedIsolation
Implemented
SE-0413Typed throws
Authors:Jorge Revuelta (@minuscorp), Torsten Lehmann, Doug Gregor
Review Manager:
Steve Canon
Implemented In:
Swift 6.0
Implemented
SE-0412Strict concurrency for global variables
Authors:John McCall, Sophia Poirier
Review Manager:
Holly Borla
Implemented In:
Swift 5.10
Upcoming Feature Flag:
GlobalConcurrency
Implemented
SE-0411Isolated default value expressions
Author:Holly Borla
Review Manager:
Doug Gregor
Bug:
apple/swift#58177
Implemented In:
Swift 5.10
Implementation:
swift#68794
Upcoming Feature Flag:
IsolatedDefaultValues
Implemented
SE-0410Low-Level Atomic Operations ⚛︎
Authors:Karoy Lorentey, Alejandro Alonso
Review Manager:
Joe Groff
Bug:
SR-9144
Implemented In:
Swift 6.0
Implementation:
swift#68857
Implemented
SE-0409Access-level modifiers on import declarations
Author:Alexis Laferrière
Review Manager:
Frederick Kellison-Linn
Implemented In:
Swift 6.0
Upcoming Feature Flag:
InternalImportsByDefault
Implemented
SE-0408Pack Iteration
Authors:Sima Nerush, Holly Borla
Review Manager:
Doug Gregor
Implemented In:
Swift 6.0
Implementation:
swift#67594
Implemented
SE-0407Member Macro Conformances
Author:Doug Gregor
Review Manager:
John McCall
Implemented In:
Swift 5.9.2
Implementation:
swift#67758
Implemented
SE-0405String Initializers with Encoding Validation
Author:Guillaume Lessard
Review Manager:
Tony Allevato
Implemented In:
Swift 6.0
Implementation:
swift#68419, swift#68423
Implemented
SE-0404Allow Protocols to be Nested in Non-Generic Contexts
Author:Karl Wagner
Review Manager:
Holly Borla
Implemented In:
Swift 5.10
Implementation:
swift#66247
Implemented
SE-0402Generalize `conformance` macros as `extension` macros
Author:Holly Borla
Review Manager:
John McCall
Implemented In:
Swift 5.9
Implementation:
swift#66967, swift-syntax#1859
Implemented
SE-0401Remove Actor Isolation Inference caused by Property Wrappers
Author:BJ Homer
Review Manager:
Holly Borla
Implemented In:
Swift 5.9
Implementation:
swift#63884
Upcoming Feature Flag:
DisableOutwardActorInference
Implemented
SE-0400Init Accessors
Authors:Holly Borla, Doug Gregor
Review Manager:
Frederick Kellison-Linn
Implemented In:
Swift 5.9
Implemented
SE-0399Tuple of value pack expansion
Authors:Sophia Poirier, Holly Borla
Review Manager:
Xiaodi Wu
Implemented In:
Swift 5.9
Implemented
SE-0398Allow Generic Types to Abstract Over Packs
Authors:Slava Pestov, Holly Borla
Review Manager:
Frederick Kellison-Linn
Implemented In:
Swift 5.9
Implemented
SE-0397Freestanding Declaration Macros
Authors:Doug Gregor, Richard Wei, Holly Borla
Review Manager:
John McCall
Implemented In:
Swift 5.9
Implemented
SE-0396Conform `Never` to `Codable`
Author:Nate Cook
Review Manager:
Tony Allevato
Implemented In:
Swift 5.9
Implementation:
swift#64899
Implemented
SE-0395Observation
Authors:Philippe Hausler, Nate Cook
Review Manager:
Ben Cohen
Implemented In:
Swift 5.9
Implemented
SE-0394Package Manager Support for Custom Macros
Authors:Boris Buegling, Doug Gregor
Review Manager:
Becca Royal-Gordon
Implemented In:
Swift 5.9
Implementation:
swift-package-manager#6185, swift-package-manager#6200
Implemented
SE-0393Value and Type Parameter Packs
Authors:Holly Borla, John McCall, Slava Pestov
Review Manager:
Xiaodi Wu
Implemented In:
Swift 5.9
Implemented
SE-0392Custom Actor Executors
Authors:Konrad 'ktoso' Malawski, John McCall, Kavon Farvardin
Review Manager:
Joe Groff
Implemented In:
Swift 5.9
Implemented
SE-0391Package Registry Publish
Author:Yim Lee
Review Manager:
Tom Doron
Implemented In:
Swift 5.9
Implementation:
swift-package-manager#6101, swift-package-manager#6146, swift-package-manager#6159, swift-package-manager#6169, swift-package-manager#6188, swift-package-manager#6189, swift-package-manager#6215, swift-package-manager#6217, swift-package-manager#6220, swift-package-manager#6229, swift-package-manager#6237
Implemented
SE-0390Noncopyable structs and enums
Authors:Joe Groff, Michael Gottesman, Andrew Trick, Kavon Farvardin
Review Manager:
Stephen Canon
Implemented In:
Swift 5.9
Implemented
SE-0389Attached Macros
Authors:Doug Gregor, Holly Borla, Richard Wei
Review Manager:
Tony Allevato
Implemented In:
Swift 5.9
Implemented
SE-0388Convenience Async[Throwing]Stream.makeStream methods
Author:Franz Busch
Review Manager:
Becca Royal-Gordon
Implemented In:
Swift 5.9
Implementation:
swift#62968
Implemented
SE-0387Swift SDKs for Cross-Compilation
Authors:Max Desiatov, Saleem Abdulrasool, Evan Wilde
Review Manager:
Mishal Shah
Implemented In:
Swift 6.1
Implementation:
swift-package-manager#5911, swift-package-manager#5922, swift-package-manager#6023, swift-package-manager#6186
Implemented
SE-0386New access modifier: `package`
Authors:Ellie Shin, Alexis Laferriere
Review Manager:
John McCall
Implemented In:
Swift 5.9
Implementation:
swift#62700, swift#62704, swift#62652, swift#62652
Implemented
SE-0384Importing Forward Declared Objective-C Interfaces and Protocols
Author:Nuri Amari
Review Manager:
Tony Allevato
Implemented In:
Swift 5.9
Implementation:
swift#61606
Upcoming Feature Flag:
ImportObjcForwardDeclarations
Implemented
SE-0383Deprecate @UIApplicationMain and @NSApplicationMain
Author:Robert Widmann
Review Manager:
John McCall
Implemented In:
Swift 5.10
Implementation:
swift#62151
Upcoming Feature Flag:
DeprecateApplicationMain
Implemented
SE-0382Expression Macros
Author:Doug Gregor
Review Manager:
Xiaodi Wu
Implemented In:
Swift 5.9
Implemented
SE-0381DiscardingTaskGroups
Authors:Cory Benfield, Konrad Malawski
Review Manager:
Doug Gregor
Implemented In:
Swift 5.9
Implementation:
swift#62361
Implemented
SE-0380`if` and `switch` expressions
Authors:Ben Cohen, Hamish Knight
Review Manager:
Holly Borla
Implemented In:
Swift 5.9
Implementation:
swift#62178
Implemented
SE-0378Package Registry Authentication
Author:Yim Lee
Review Manager:
Tom Doron
Implemented In:
Swift 5.8
Implementation:
swift-package-manager#5838
Implemented
SE-0377`borrowing` and `consuming` parameter ownership modifiers
Authors:Michael Gottesman, Joe Groff
Review Manager:
John McCall
Implemented In:
Swift 5.9
Implemented
SE-0376Function Back Deployment
Author:Allan Shortlidge
Review Manager:
Frederick Kellison-Linn
Implemented In:
Swift 5.8
Implementation:
swift#41271, swift#41348, swift#41416, swift#41612
Implemented
SE-0375Opening existential arguments to optional parameters
Author:Doug Gregor
Review Manager:
Xiaodi Wu
Implemented In:
Swift 5.8
Implementation:
swift#61321
Implemented
SE-0374Add sleep(for:) to Clock
Authors:Brandon Williams, Stephen Celis
Review Manager:
Steve Canon
Implemented In:
Swift 5.9
Implementation:
swift#61222
Implemented
SE-0373Lift all limitations on variables in result builders
Author:Pavel Yaskevich
Review Manager:
John McCall
Implemented In:
Swift 5.8
Implementation:
swift#60839
Implemented
SE-0372Document Sorting as Stable
Author:Nate Cook
Review Manager:
Tony Allevato
Implemented In:
Swift 5.8
Implementation:
swift#60936
Implemented
SE-0371Isolated synchronous deinit
Author:Mykola Pokhylets
Review Manager:
Frederick Kellison-Linn
Implemented In:
Swift 6.2
Implementation:
swift#60057
Implemented
SE-0370Pointer Family Initialization Improvements and Better Buffer Slices
Author:Guillaume Lessard
Review Manager:
John McCall
Implemented In:
Swift 5.8
Implementation:
swift#41608
Implemented
SE-0369Add CustomDebugStringConvertible conformance to AnyKeyPath
Author:Ben Pious
Review Manager:
Xiaodi Wu
Implemented In:
Swift 5.8
Implementation:
swift#60133
Implemented
SE-0368StaticBigInt
Author:Ben Rimmington
Review Managers:
Doug Gregor, Holly Borla
Implemented In:
Swift 5.8
Implementation:
swift#40722, swift#62733
Implemented
SE-0367Conditional compilation for attributes
Author:Doug Gregor
Review Manager:
Joe Groff
Implemented In:
Swift 5.8
Implementation:
swift#60208
Implemented
SE-0366`consume` operator to end the lifetime of a variable binding
Authors:Michael Gottesman, Andrew Trick, Joe Groff
Review Manager:
Holly Borla
Implemented In:
Swift 5.9
Implemented
SE-0365Allow implicit `self` for `weak self` captures, after `self` is unwrapped
Author:Cal Stephens
Review Manager:
Saleem Abdulrasool
Implemented In:
Swift 5.8
Implementation:
swift#40702, swift#61520
Implemented
SE-0364Warning for Retroactive Conformances of External Types
Author:Harlan Haskins
Review Manager:
Steve Canon
Implemented In:
Swift 6.0
Implemented
SE-0363Unicode for String Processing
Authors:Nate Cook, Alejandro Alonso
Review Manager:
Ben Cohen
Implemented In:
Swift 5.7
Implemented
SE-0362Piecemeal adoption of upcoming language improvements
Author:Doug Gregor
Review Manager:
Holly Borla
Implemented In:
Swift 5.8
Implementation:
swift#59055, swift-package-manager#5632
Implemented
SE-0361Extensions on bound generic types
Author:Holly Borla
Review Manager:
John McCall
Implemented In:
Swift 5.7
Implementation:
swift#41172
Implemented
SE-0360Opaque result types with limited availability
Author:Pavel Yaskevich
Review Manager:
Joe Groff
Implemented In:
Swift 5.7
Implementation:
swift#42072, swift#42104, swift#42167, swift#42456
Implemented
SE-0358Primary Associated Types in the Standard Library
Author:Karoy Lorentey
Review Manager:
John McCall
Implemented In:
Swift 5.7
Implementation:
swift#41843
Implemented
SE-0357Regex-powered string processing algorithms
Authors:Tina Liu, Michael Ilseman, Nate Cook, Tim Vermeulen
Review Manager:
Ben Cohen
Implemented In:
Swift 5.7
Implemented
SE-0356Swift Snippets
Author:Ashley Garland
Review Manager:
Tom Doron
Implemented In:
Swift 5.7
Implementation:
swift-docc#61, swift-package-manager#3694, swift-package-manager#3732, swift-docc-symbolkit#10, swift-docc-symbolkit#15, swift-docc-plugin#7
Implemented
SE-0355Regex Syntax and Run-time Construction
Authors:Hamish Knight, Michael Ilseman
Review Manager:
Ben Cohen
Implemented In:
Swift 5.7
Implemented
SE-0354Regex Literals
Authors:Hamish Knight, Michael Ilseman, David Ewing
Review Manager:
Ben Cohen
Implemented In:
Swift 5.7
Implementation:
swift#42119, swift#58835
Upcoming Feature Flag:
BareSlashRegexLiterals
Implemented
SE-0353Constrained Existential Types
Author:Robert Widmann
Review Manager:
Joe Groff
Implemented In:
Swift 5.7
Implemented
SE-0352Implicitly Opened Existentials
Author:Doug Gregor
Review Manager:
Joe Groff
Implemented In:
Swift 5.7
Implementation:
swift#41996
Upcoming Feature Flag:
ImplicitOpenExistentials
Implemented
SE-0351Regex builder DSL
Authors:Richard Wei, Michael Ilseman, Nate Cook, Alejandro Alonso
Review Manager:
Ben Cohen
Implemented In:
Swift 5.7
Implemented
SE-0350Regex Type and Overview
Author:Michael Ilseman
Review Manager:
Ben Cohen
Implemented In:
Swift 5.7
Implemented
SE-0349Unaligned Loads and Stores from Raw Memory
Authors:Guillaume Lessard, Andrew Trick
Review Manager:
John McCall
Implemented In:
Swift 5.7
Implementation:
swift#41033
Implemented
SE-0348`buildPartialBlock` for result builders
Author:Richard Wei
Review Manager:
Ben Cohen
Implemented In:
Swift 5.7
Implementation:
swift#41576
Implemented
SE-0347Type inference from default expressions
Author:Pavel Yaskevich
Review Manager:
Doug Gregor
Implemented In:
Swift 5.7
Implementation:
swift#41436
Implemented
SE-0346Lightweight same-type requirements for primary associated types
Authors:Pavel Yaskevich, Holly Borla, Slava Pestov
Review Manager:
John McCall
Implemented In:
Swift 5.7
Implemented
SE-0345`if let` shorthand for shadowing an existing optional variable
Author:Cal Stephens
Review Manager:
Doug Gregor
Implemented In:
Swift 5.7
Implementation:
swift#40694
Implemented
SE-0344Distributed Actor Runtime
Authors:Konrad 'ktoso' Malawski, Pavel Yaskevich, Doug Gregor, Kavon Farvardin, Dario Rexin, Tomer Doron
Review Manager:
Joe Groff
Implemented In:
Swift 5.7
Implemented
SE-0343Concurrency in Top-level Code
Author:Evan Wilde
Review Manager:
Saleem Abdulrasool
Implemented In:
Swift 5.7
Implementation:
swift#40963, swift#40998, swift#41061
Implemented
SE-0341Opaque Parameter Declarations
Author:Doug Gregor
Review Manager:
Ben Cohen
Implemented In:
Swift 5.7
Implementation:
swift#40993
Implemented
SE-0340Unavailable From Async Attribute
Author:Evan Wilde
Review Manager:
Joe Groff
Implemented In:
Swift 5.7
Implementation:
swift#40769
Implemented
SE-0339Module Aliasing For Disambiguation
Author:Ellie Shin
Review Manager:
John McCall
Implemented In:
Swift 5.7
Implementation:
swift#40899, swift-package-manager#4023
Implemented
SE-0338Clarify the Execution of Non-Actor-Isolated Async Functions
Author:John McCall
Review Manager:
Doug Gregor
Implemented In:
Swift 5.7
Implemented
SE-0337Incremental migration to concurrency checking
Authors:Doug Gregor, Becca Royal-Gordon
Review Manager:
Ben Cohen
Implemented In:
Swift 5.6
Implementation:
swift#40680
Upcoming Feature Flag:
StrictConcurrency
Implemented
SE-0336Distributed Actor Isolation
Authors:Konrad 'ktoso' Malawski, Pavel Yaskevich, Doug Gregor, Kavon Farvardin
Review Manager:
Joe Groff
Implemented In:
Swift 5.7
Implemented
SE-0335Introduce existential `any`
Author:Holly Borla
Review Manager:
Doug Gregor
Implemented In:
Swift 5.6
Implementation:
swift#40282
Upcoming Feature Flag:
ExistentialAny
Implemented
SE-0334Pointer API Usability Improvements
Authors:Guillaume Lessard, Andrew Trick
Review Manager:
Ben Cohen
Bugs:
rdar://64342031, SR-11156, rdar://53272880, rdar://22541346
Implemented In:
Swift 5.7
Implementation:
swift#39639
Implemented
SE-0333Expand usability of `withMemoryRebound`
Authors:Guillaume Lessard, Andrew Trick
Review Manager:
Ben Cohen
Bugs:
SR-11082, SR-11087
Implemented In:
Swift 5.7
Implementation:
swift#39529
Implemented
SE-0332Package Manager Command Plugins
Author:Anders Bertelrud
Review Manager:
Tom Doron
Implemented In:
Swift 5.6
Implementation:
swift-package-manager#3855
Implemented
SE-0331Remove Sendable conformance from unsafe pointer types
Author:Andrew Trick
Review Manager:
Doug Gregor
Implemented In:
Swift 5.6
Implementation:
swift#39218
Implemented
SE-0329Clock, Instant, and Duration
Author:Philippe Hausler
Review Manager:
John McCall
Implemented In:
Swift 5.7
Implementation:
swift#40609
Implemented
SE-0328Structural opaque result types
Authors:Benjamin Driscoll, Holly Borla
Review Manager:
Ben Cohen
Implemented In:
Swift 5.7
Implementation:
swift#38392
Implemented
SE-0327On Actors and Initialization
Authors:Kavon Farvardin, John McCall, Konrad Malawski
Review Manager:
Doug Gregor
Implemented In:
Swift 5.10
Implemented
SE-0326Enable multi-statement closure parameter/result type inference
Author:Pavel Yaskevich
Review Manager:
Ben Cohen
Implemented In:
Swift 5.7
Implementation:
swift#38577, swift#40397, swift#41730
Implemented
SE-0325Additional Package Plugin APIs
Author:Anders Bertelrud
Review Manager:
Tom Doron
Implemented In:
Swift 5.6
Implementation:
swift-package-manager#3758
Implemented
SE-0324Relax diagnostics for pointer arguments to C functions
Authors:Andrew Trick, Pavel Yaskevich
Review Manager:
Saleem Abdulrasool
Bug:
SR-10246
Implemented In:
Swift 5.6
Implementation:
swift#37956
Implemented
SE-0323Asynchronous Main Semantics
Author:Evan Wilde
Review Manager:
Doug Gregor
Implemented In:
Swift 5.5.2
Implementation:
swift#38604
Implemented
SE-0322Temporary uninitialized buffers
Author:Jonathan Grynspan
Review Manager:
Joe Groff
Implemented In:
Swift 5.6
Implementation:
swift#37666
Implemented
SE-0320Allow coding of non `String` / `Int` keyed `Dictionary` into a `KeyedContainer`
Author:Morten Bek Ditlevsen
Review Manager:
Tom Doron
Implemented In:
Swift 5.6
Implementation:
swift#34458
Implemented
SE-0319Conform Never to Identifiable
Author:Kyle Macomber
Review Manager:
Tom Doron
Implemented In:
Swift 5.5
Implementation:
swift#38103
Implemented
SE-0317`async let` bindings
Authors:John McCall, Joe Groff, Doug Gregor, Konrad 'ktoso' Malawski
Review Manager:
Ben Cohen
Implemented In:
Swift 5.5
Implemented
SE-0316Global actors
Authors:John McCall, Doug Gregor
Review Manager:
Joe Groff
Implemented In:
Swift 5.5
Implemented
SE-0315Type placeholders (formerly, "Placeholder types")
Author:Frederick Kellison-Linn
Review Manager:
Joe Groff
Implemented In:
Swift 5.6
Implementation:
swift#36740
Implemented
SE-0314`AsyncStream` and `AsyncThrowingStream`
Authors:Philippe Hausler, Tony Parker, Ben D. Jones, Nate Cook
Review Manager:
Doug Gregor
Implemented In:
Swift 5.5
Implementation:
swift#36921
Implemented
SE-0313Improved control over actor isolation
Authors:Doug Gregor, Chris Lattner
Review Manager:
Ted Kremenek
Implemented In:
Swift 5.5
Implemented
SE-0311Task Local Values
Author:Konrad 'ktoso' Malawski
Review Manager:
John McCall
Implemented In:
Swift 5.5
Implemented
SE-0310Effectful Read-only Properties
Author:Kavon Farvardin
Review Manager:
Doug Gregor
Implemented In:
Swift 5.5
Implementation:
swift#36430, swift#36670, swift#37225
Implemented
SE-0309Unlock existentials for all protocols
Authors:Anthony Latsis, Filip Sakel, Suyash Srijan
Review Manager:
Joe Groff
Implemented In:
Swift 5.7
Implementation:
swift#33767, swift#39492, swift#41198
Implemented
SE-0308`#if` for postfix member expressions
Author:Rintaro Ishizaki
Review Manager:
Saleem Abdulrasool
Implemented In:
Swift 5.5
Implementation:
swift#35097
Implemented
SE-0307Allow interchangeable use of `CGFloat` and `Double` types
Author:Pavel Yaskevich
Review Manager:
Ted Kremenek
Implemented In:
Swift 5.5
Implementation:
swift#34401
Implemented
SE-0306Actors
Authors:John McCall, Doug Gregor, Konrad Malawski, Chris Lattner
Review Manager:
Joe Groff
Implemented In:
Swift 5.5
Implemented
SE-0305Package Manager Binary Target Improvements
Authors:Anders Bertelrud, Tom Doron
Review Manager:
Tom Doron
Implemented In:
Swift 5.6
Implemented
SE-0304Structured concurrency
Authors:John McCall, Joe Groff, Doug Gregor, Konrad Malawski
Review Manager:
Ben Cohen
Implemented In:
Swift 5.5
Implemented
SE-0303Package Manager Extensible Build Tools
Authors:Anders Bertelrud, Konrad 'ktoso' Malawski, Tom Doron
Review Manager:
Tom Doron
Implemented In:
Swift 5.6
Implemented
SE-0302`Sendable` and `@Sendable` closures
Authors:Chris Lattner, Doug Gregor
Review Manager:
John McCall
Implemented In:
Swift 5.7
Implementation:
swift#35264
Implemented
SE-0301Package Editor Commands
Author:Owen Voorhees
Review Manager:
Tom Doron
Implemented In:
Swift 6.0
Implementation:
swift-package-manager#3034
Implemented
SE-0300Continuations for interfacing async tasks with synchronous code
Authors:John McCall, Joe Groff, Doug Gregor, Konrad Malawski
Review Manager:
Ben Cohen
Implemented In:
Swift 5.5
Implemented
SE-0299Extending Static Member Lookup in Generic Contexts
Authors:Pavel Yaskevich, Sam Lazarus, Matt Ricketson
Review Manager:
Joe Groff
Implemented In:
Swift 5.5
Implementation:
swift#34523
Implemented
SE-0298Async/Await: Sequences
Authors:Tony Parker, Philippe Hausler
Review Manager:
Doug Gregor
Implemented In:
Swift 5.5
Implementation:
swift#35224
Implemented
SE-0297Concurrency Interoperability with Objective-C
Author:Doug Gregor
Review Manager:
Chris Lattner
Implemented In:
Swift 5.5
Implemented
SE-0296Async/await
Authors:John McCall, Doug Gregor
Review Manager:
Ben Cohen
Implemented In:
Swift 5.5
Implemented
SE-0295Codable synthesis for enums with associated values
Author:Dario Rexin
Review Manager:
Saleem Abdulrasool
Implemented In:
Swift 5.5
Implementation:
swift#34855
Implemented
SE-0294Declaring executable targets in Package Manifests
Author:Anders Bertelrud
Review Manager:
Tom Doron
Bug:
SR-13924
Implemented In:
Swift 5.4
Implementation:
swift-package-manager#3045
Implemented
SE-0293Extend Property Wrappers to Function and Closure Parameters
Authors:Holly Borla, Filip Sakel
Review Manager:
Chris Lattner
Implemented In:
Swift 5.5
Implementation:
swift#34272, swift#36344
Implemented
SE-0292Package Registry Service
Authors:Bryan Clark, Whitney Imura, Mattt Zmuda
Review Manager:
Tom Doron
Implemented In:
Swift 5.7
Implementation:
swift-package-manager#3023
Implemented
SE-0291Package Collections
Authors:Boris Bügling, Yim Lee, Tom Doron
Review Manager:
Tom Doron
Implemented In:
Swift 5.5
Implementation:
swift-package-manager#3030
Implemented
SE-0290Unavailability Condition
Author:Bruno Rocha
Review Manager:
Ted Kremenek
Implemented In:
Swift 5.6
Implementation:
swift#33932
Implemented
SE-0289Result builders
Authors:John McCall, Doug Gregor
Review Manager:
Saleem Abdulrasool
Implemented In:
Swift 5.4
Implemented
SE-0287Extend implicit member syntax to cover chains of member references
Author:Frederick Kellison-Linn
Review Manager:
Doug Gregor
Implemented In:
Swift 5.4
Implementation:
swift#31679
Implemented
SE-0286Forward-scan matching for trailing closures
Author:Doug Gregor
Review Manager:
John McCall
Implemented In:
Swift 5.3
Implementation:
swift#33092
Upcoming Feature Flag:
ForwardTrailingClosures
Implemented
SE-0285Ease the transition to concise magic file strings
Author:Becca Royal-Gordon
Review Manager:
Tom Doron
Implemented In:
Swift 5.3
Implementation:
swift#32700
Implemented
SE-0284Allow Multiple Variadic Parameters in Functions, Subscripts, and Initializers
Author:Owen Voorhees
Review Manager:
Saleem Abdulrasool
Implemented In:
Swift 5.4
Implementation:
swift#29735
Implemented
SE-0282Clarify the Swift memory consistency model ⚛︎
Author:Karoy Lorentey
Review Manager:
Joe Groff
Bug:
SR-9144
Implemented In:
Swift 5.3
Implemented
SE-0281`@main`: Type-Based Program Entry Points
Authors:Nate Cook, Nate Chandler, Matt Ricketson
Review Manager:
Tom Doron
Implemented In:
Swift 5.3
Implementation:
swift#30693
Implemented
SE-0280Enum cases as protocol witnesses
Author:Suyash Srijan
Review Manager:
John McCall
Bug:
SR-3170
Implemented In:
Swift 5.3
Implementation:
swift#28916
Implemented
SE-0279Multiple Trailing Closures
Authors:Kyle Macomber, Pavel Yaskevich, Doug Gregor, John McCall
Review Manager:
Ben Cohen
Implemented In:
Swift 5.3
Implementation:
swift#31052
Implemented
SE-0278Package Manager Localized Resources
Author:David Hart
Review Manager:
Boris Buegling
Implemented In:
Swift 5.3
Implementation:
swift-package-manager#2535, swift-package-manager#2606
Implemented
SE-0277Float16
Author:Stephen Canon
Review Manager:
Ben Cohen
Implemented In:
Swift 5.3
Implementation:
swift#30130
Implemented
SE-0276Multi-Pattern Catch Clauses
Author:Owen Voorhees
Review Manager:
Doug Gregor
Implemented In:
Swift 5.3
Implementation:
swift#27776
Implemented
SE-0274Concise magic file names
Authors:Becca Royal-Gordon, Dave DeLong
Review Manager:
Ben Cohen
Implemented In:
Swift 5.8
Upcoming Feature Flag:
ConciseMagicFile
Implemented
SE-0273Package Manager Conditional Target Dependencies
Author:David Hart
Review Manager:
Boris Buegling
Implemented In:
Swift 5.3
Implementation:
swift-package-manager#2428, swift-package-manager#2598
Implemented
SE-0272Package Manager Binary Dependencies
Authors:Braden Scothern, Daniel Dunbar, Franz Busch
Review Manager:
Boris Bügling
Implemented In:
Swift 5.3
Implementation:
swift-package-manager#2509, swift-package-manager#2511, swift-package-manager#2514, swift-package-manager#2588
Implemented
SE-0271Package Manager Resources
Authors:Anders Bertelrud, Ankit Aggarwal
Review Manager:
Boris Buegling
Implemented In:
Swift 5.3
Implementation:
swift-package-manager#2381, swift-package-manager#2510, swift-package-manager#2520, swift-package-manager#2607
Implemented
SE-0270Add Collection Operations on Noncontiguous Elements
Authors:Nate Cook, Jeremy Schonfeld
Review Manager:
John McCall
Implemented In:
Swift 6.0
Implementation:
swift#69766
Implemented
SE-0269Increase availability of implicit `self` in `@escaping` closures when reference cycles are unlikely to occur
Author:Frederick Kellison-Linn
Review Manager:
Ted Kremenek
Bug:
SR-10218
Implemented In:
Swift 5.3
Implementation:
swift#23934
Implemented
SE-0268Refine `didSet` Semantics
Author:Suyash Srijan
Review Manager:
Ben Cohen
Bug:
SR-5982
Implemented In:
Swift 5.3
Implementation:
swift#26632
Implemented
SE-0267`where` clauses on contextually generic declarations
Author:Anthony Latsis
Review Manager:
Joe Groff
Implemented In:
Swift 5.3
Implementation:
swift#23489
Implemented
SE-0266Synthesized `Comparable` conformance for `enum` types
Author:Diana Ma (taylorswift)
Review Manager:
Ben Cohen
Implemented In:
Swift 5.3
Implementation:
swift#25696
Implemented
SE-0264Standard Library Preview Package
Authors:Ben Cohen, Max Moiseev, Nate Cook
Review Manager:
Ted Kremenek
Implemented In:
Swift
Implementation:
swift-evolution#1089, swift-evolution#1090, swift-evolution#1091
Implemented
SE-0263Add a String Initializer with Access to Uninitialized Storage
Author:David Smith
Review Manager:
Ted Kremenek
Bug:
SR-10288
Implemented In:
Swift 5.3
Implementation:
swift#26007, swift#30106
Implemented
SE-0261Identifiable Protocol
Authors:Matthew Johnson, Kyle Macomber
Review Manager:
Ben Cohen
Implemented In:
Swift 5.1
Implementation:
swift#26022
Implemented
SE-0260Library Evolution for Stable ABIs
Authors:Jordan Rose, Ben Cohen
Review Manager:
John McCall
Implemented In:
Swift 5.1
Implementation:
swift#24185
Implemented
SE-0258Property Wrappers
Authors:Doug Gregor, Joe Groff
Review Manager:
John McCall
Implemented In:
Swift 5.1
Implemented
SE-0255Implicit returns from single-expression functions
Author:Nate Chandler
Review Manager:
Ben Cohen
Implemented In:
Swift 5.1
Implementation:
swift#23251
Implemented
SE-0254Static and class subscripts
Author:Becca Royal-Gordon
Review Manager:
Doug Gregor
Implemented In:
Swift 5.1
Implementation:
swift#23358
Implemented
SE-0253Callable values of user-defined nominal types
Authors:Richard Wei, Dan Zheng
Review Manager:
Chris Lattner
Implemented In:
Swift 5.2
Implementation:
swift#24299
Implemented
SE-0252Key Path Member Lookup
Authors:Doug Gregor, Pavel Yaskevich
Review Manager:
Ted Kremenek
Implemented In:
Swift 5.1
Implementation:
swift#23436
Implemented
SE-0251SIMD additions
Author:Stephen Canon
Review Manager:
John McCall
Implemented In:
Swift 5.1
Implementation:
swift#23421, swift#24136
Implemented
SE-0249Key Path Expressions as Functions
Authors:Stephen Celis, Greg Titus
Review Manager:
Ben Cohen
Implemented In:
Swift 5.2
Implementation:
swift#26054
Implemented
SE-0248String Gaps and Missing APIs
Author:Michael Ilseman
Review Manager:
Ted Kremenek
Bug:
SR-9955
Implemented In:
Swift 5.1
Implementation:
swift#22869
Implemented
SE-0247Contiguous Strings
Author:Michael Ilseman
Review Manager:
Doug Gregor
Bug:
SR-6475
Implemented In:
Swift 5.1
Implementation:
swift#23051
Implemented
SE-0245Add an Array Initializer with Access to Uninitialized Storage
Author:Nate Cook
Review Manager:
Ted Kremenek
Bug:
SR-3087
Implemented In:
Swift 5.1
Implementation:
swift#23134
Implemented
SE-0244Opaque Result Types
Authors:Doug Gregor, Joe Groff
Review Manager:
Ben Cohen
Implemented In:
Swift 5.1
Implementation:
swift#22072
Implemented
SE-0242Synthesize default values for the memberwise initializer
Author:Alejandro Alonso
Review Manager:
Ted Kremenek
Implemented In:
Swift 5.1
Implementation:
swift#19743
Implemented
SE-0241Deprecate String Index Encoded Offsets
Author:Michael Ilseman
Review Manager:
John McCall
Implemented In:
Swift 5.0
Implementation:
swift#22108
Implemented
SE-0240Ordered Collection Diffing
Authors:Scott Perry, Kyle Macomber
Review Managers:
Doug Gregor, Ben Cohen
Implemented In:
Swift 5.1
Implementation:
swift#21845
Implemented
SE-0239Add Codable conformance to Range types
Authors:Dale Buckley, Ben Cohen, Maxim Moiseev
Review Manager:
Ted Kremenek
Implemented In:
Swift 5.0
Implementation:
swift#19532, swift#21857
Implemented
SE-0238Package Manager Target Specific Build Settings
Author:Ankit Aggarwal
Review Manager:
Boris Bügling
Implemented In:
Swift 5.0
Implemented
SE-0237Introduce `withContiguous{Mutable}StorageIfAvailable` methods
Author:Ben Cohen
Review Manager:
Doug Gregor
Implemented In:
Swift 5.0
Implementation:
swift#21138
Implemented
SE-0236Package Manager Platform Deployment Settings
Author:Ankit Aggarwal
Review Manager:
Boris Bügling
Implemented In:
Swift 5.0
Implemented
SE-0235Add Result to the Standard Library
Author:Jon Shier
Review Manager:
Chris Lattner
Implemented In:
Swift 5.0
Implementation:
swift#21073, swift#21225, swift#21378
Implemented
SE-0234Remove `Sequence.SubSequence`
Author:Ben Cohen
Review Manager:
John McCall
Implemented In:
Swift 5.0
Implementation:
swift#20221
Implemented
SE-0233Make `Numeric` Refine a new `AdditiveArithmetic` Protocol
Author:Richard Wei
Review Manager:
Chris Lattner
Implemented In:
Swift 5.0
Implementation:
swift#20422
Implemented
SE-0232Remove Some Customization Points from the Standard Library's `Collection` Hierarchy
Author:Ben Cohen
Review Manager:
Ted Kremenek
Implemented In:
Swift 5.0
Implementation:
swift#19995
Implemented
SE-0230Flatten nested optionals resulting from 'try?'
Author:BJ Homer
Review Manager:
John McCall
Implemented In:
Swift 5.0
Implementation:
swift#16942
Implemented
SE-0229SIMD Vectors
Author:Stephen Canon
Review Manager:
Ben Cohen
Implemented In:
Swift 5.0
Implementation:
swift#20344
Implemented
SE-0228Fix `ExpressibleByStringInterpolation`
Authors:Becca Royal-Gordon, Michael Ilseman
Review Manager:
Doug Gregor
Implemented In:
Swift 5.0
Implementation:
swift#20214
Implemented
SE-0227Identity key path
Author:Joe Groff
Review Manager:
Ben Cohen
Implemented In:
Swift 5.0
Implementation:
swift#18804, swift#19382
Implemented
SE-0226Package Manager Target Based Dependency Resolution
Author:Ankit Aggarwal
Review Manager:
Boris Bügling
Bug:
SR-8658
Implemented In:
Swift 5.2
Implemented
SE-0225Adding `isMultiple` to `BinaryInteger`
Authors:Robert MacEachern, Micah Hansonbrook
Review Manager:
John McCall
Implemented In:
Swift 5.0
Implementation:
swift#18689
Implemented
SE-0224Support 'less than' operator in compilation conditions
Author:Daniel Martín
Review Manager:
Ted Kremenek
Bug:
SR-6852
Implemented In:
Swift 5.0
Implementation:
swift#14503, swift#17960
Implemented
SE-0221Character Properties
Authors:Michael Ilseman, Tony Allevato
Review Manager:
Ben Cohen
Implemented In:
Swift 5.0
Implementation:
swift#20520
Implemented
SE-0220`count(where:)`
Author:Soroush Khanlou
Review Manager:
Ben Cohen
Implemented In:
Swift 6.0
Implementation:
swift#72705
Implemented
SE-0219Package Manager Dependency Mirroring
Author:Ankit Aggarwal
Review Manager:
Boris Bügling
Bug:
SR-8328
Implemented In:
Swift 5.0
Implementation:
swift-package-manager#1776
Implemented
SE-0218Introduce `compactMapValues` to Dictionary
Author:Daiki Matsudate
Review Manager:
Ben Cohen
Implemented In:
Swift 5.0
Implementation:
swift#15017
Implemented
SE-0216Introduce user-defined dynamically "callable" types
Authors:Chris Lattner, Dan Zheng
Review Manager:
John McCall
Implemented In:
Swift 5.0
Implementation:
swift#20305
Implemented
SE-0215Conform `Never` to `Equatable` and `Hashable`
Author:Matt Diephouse
Review Manager:
Ted Kremenek
Implemented In:
Swift 5.0
Implementation:
swift#16857
Implemented
SE-0214Renaming the `DictionaryLiteral` type to `KeyValuePairs`
Authors:Erica Sadun, Chéyo Jiménez
Review Manager:
Chris Lattner
Implemented In:
Swift 5.0
Implementation:
swift#16577
Implemented
SE-0213Literal initialization via coercion
Author:Pavel Yaskevich
Review Manager:
John McCall
Implemented In:
Swift 5.0
Implementation:
swift#17860
Implemented
SE-0212Compiler Version Directive
Author:David Hart
Review Manager:
Ted Kremenek
Implemented In:
Swift 4.2
Implementation:
swift#15977
Implemented
SE-0211Add Unicode Properties to `Unicode.Scalar`
Author:Tony Allevato
Review Manager:
Ben Cohen
Implemented In:
Swift 5.0
Implementation:
swift#15593
Implemented
SE-0210Add an `offset(of:)` method to `MemoryLayout`
Author:Joe Groff
Review Manager:
Doug Gregor
Implemented In:
Swift 4.2
Implementation:
swift#15519
Implemented
SE-0209Package Manager Swift Language Version API Update
Author:Ankit Aggarwal
Review Manager:
Boris Bügling
Bug:
SR-7464
Implemented In:
Swift 4.2
Implementation:
swift-package-manager#1563
Implemented
SE-0208Package Manager System Library Targets
Authors:Ankit Aggarwal, Daniel Dunbar
Review Manager:
Boris Bügling
Bug:
SR-7434
Implemented In:
Swift 4.2
Implementation:
swift-package-manager#1586
Implemented
SE-0207Add an `allSatisfy` algorithm to `Sequence`
Author:Ben Cohen
Review Manager:
Dave Abrahams
Implemented In:
Swift 4.2
Implementation:
swift#15120
Implemented
SE-0206Hashable Enhancements
Authors:Karoy Lorentey, Vincent Esche
Review Manager:
Joe Groff
Implemented In:
Swift 4.2
Implementation:
swift#14913, swift#16009, swift#16073
Implemented
SE-0205`withUnsafePointer(to:_:)` and `withUnsafeBytes(of:_:)` for immutable values
Author:Joe Groff
Review Manager:
Ben Cohen
Implemented In:
Swift 4.2
Implementation:
swift#15608
Implemented
SE-0204Add `last(where:)` and `lastIndex(where:)` Methods
Author:Nate Cook
Review Manager:
John McCall
Bug:
SR-1504
Implemented In:
Swift 4.2
Implementation:
swift#13337
Implemented
SE-0202Random Unification
Author:Alejandro Alonso
Review Manager:
Ben Cohen
Implemented In:
Swift 4.2
Implementation:
swift#12772
Implemented
SE-0201Package Manager Local Dependencies
Author:Ankit Aggarwal
Review Manager:
Boris Bügling
Bug:
SR-7433
Implemented In:
Swift 4.2
Implementation:
swift-package-manager#1583
Implemented
SE-0200Enhancing String Literals Delimiters to Support Raw Text
Authors:John Holdsworth, Becca Royal-Gordon, Erica Sadun
Review Manager:
Doug Gregor
Bug:
SR-6362
Implemented In:
Swift 5.0
Implementation:
swift#17668
Implemented
SE-0199Adding `toggle` to `Bool`
Author:Chris Eidhof
Review Manager:
Ben Cohen
Implemented In:
Swift 4.2
Implementation:
swift#14586
Implemented
SE-0198Playground QuickLook API Revamp
Author:Connor Wakamo
Review Manager:
Ben Cohen
Implemented In:
Swift 4.1
Implementation:
swift#13911, swift-xcode-playground-support#21, swift#14252, swift-corelibs-foundation#1415, swift-xcode-playground-support#20
Implemented
SE-0197Adding in-place `removeAll(where:)` to the Standard Library
Author:Ben Cohen
Review Manager:
John McCall
Implemented In:
Swift 4.2
Implementation:
swift#11576
Implemented
SE-0196Compiler Diagnostic Directives
Author:Harlan Haskins
Review Manager:
Ted Kremenek
Implemented In:
Swift 4.2
Implementation:
swift#14048
Implemented
SE-0195Introduce User-defined "Dynamic Member Lookup" Types
Author:Chris Lattner
Review Manager:
Ted Kremenek
Implemented In:
Swift 4.2
Implementation:
swift#14546
Implemented
SE-0194Derived Collection of Enum Cases
Authors:Jacob Bandes-Storch, Becca Royal-Gordon, Robert Widmann
Review Manager:
Doug Gregor
Bugs:
SR-7151, SR-7152
Implemented In:
Swift 4.2
Implementation:
swift#13655
Implemented
SE-0193Cross-module inlining and specialization
Author:Slava Pestov
Review Manager:
Ted Kremenek
Implemented In:
Swift 4.2
Implementation:
swift#15787
Implemented
SE-0192Handling Future Enum Cases
Author:Jordan Rose
Review Manager:
Ted Kremenek
Implemented In:
Swift 5.0
Implementation:
swift#14945
Implemented
SE-0191Eliminate `IndexDistance` from `Collection`
Author:Ben Cohen
Review Manager:
Doug Gregor
Implemented In:
Swift 4.1
Implementation:
swift#12641
Implemented
SE-0190Target environment platform condition
Authors:Erica Sadun, Graydon Hoare
Review Manager:
Ted Kremenek
Implemented In:
Swift 4.1
Implementation:
swift#12964
Implemented
SE-0189Restrict Cross-module Struct Initializers
Author:Jordan Rose
Review Manager:
Ted Kremenek
Implemented In:
Swift 4.1
Implementation:
swift#12834
Implemented
SE-0188Make Standard Library Index Types Hashable
Author:Nate Cook
Review Manager:
Ben Cohen
Implemented In:
Swift 4.1
Implementation:
swift#12777
Implemented
SE-0187Introduce Sequence.compactMap(_:)
Author:Max Moiseev
Review Manager:
John McCall
Implemented In:
Swift 4.1
Implementation:
swift#12819
Implemented
SE-0186Remove ownership keyword support in protocols
Author:Greg Spiers
Review Manager:
Ted Kremenek
Bug:
SR-479
Implemented In:
Swift 4.1
Implementation:
swift#11744
Implemented
SE-0185Synthesizing `Equatable` and `Hashable` conformance
Author:Tony Allevato
Review Manager:
Chris Lattner
Implemented In:
Swift 4.1
Implementation:
swift#9619
Implemented
SE-0184Unsafe[Mutable][Raw][Buffer]Pointer: add missing methods, adjust existing labels for clarity, and remove deallocation size
Author:Diana Ma (“Taylor Swift”)
Review Manager:
Doug Gregor
Implemented In:
Swift 4.1
Implementation:
swift#12200
Implemented
SE-0183Substring performance affordances
Author:Ben Cohen
Review Manager:
Chris Lattner
Bug:
SR-4933
Implemented In:
Swift 4.0
Implemented
SE-0182String Newline Escaping
Authors:John Holdsworth, David Hart, Adrian Zubarev
Review Manager:
Chris Lattner
Implemented In:
Swift 4.0
Implementation:
swift#11080
Implemented
SE-0181Package Manager C/C++ Language Standard Support
Author:Ankit Aggarwal
Review Manager:
Daniel Dunbar
Implemented In:
Swift 4.0
Implementation:
swift-package-manager#1264
Implemented
SE-0180String Index Overhaul
Author:Dave Abrahams
Review Manager:
Ted Kremenek
Implemented In:
Swift 4.0
Implementation:
swift#9806
Implemented
SE-0179Swift `run` Command
Author:David Hart
Review Manager:
Daniel Dunbar
Implemented In:
Swift 4.0
Implementation:
swift-package-manager#1187
Implemented
SE-0178Add `unicodeScalars` property to `Character`
Author:Ben Cohen
Review Manager:
Ted Kremenek
Implemented In:
Swift 4.0
Implementation:
swift#9675
Implemented
SE-0176Enforce Exclusive Access to Memory
Author:John McCall
Review Manager:
Ben Cohen
Implemented In:
Swift 4.0
Implemented
SE-0175Package Manager Revised Dependency Resolution
Author:Rick Ballard
Review Manager:
Ankit Aggarwal
Implemented In:
Swift 4.0
Implemented
SE-0174Change `RangeReplaceableCollection.filter` to return `Self`
Author:Ben Cohen
Review Manager:
Doug Gregor
Bug:
SR-3444
Implemented In:
Swift 4.2
Implemented
SE-0173Add `MutableCollection.swapAt(_:_:)`
Author:Ben Cohen
Review Manager:
Ted Kremenek
Implemented In:
Swift 4.0
Implementation:
swift#9119
Implemented
SE-0172One-sided Ranges
Authors:Ben Cohen, Dave Abrahams, Becca Royal-Gordon
Review Manager:
Doug Gregor
Implemented In:
Swift 4.0
Implemented
SE-0171Reduce with `inout`
Author:Chris Eidhof
Review Manager:
Ben Cohen
Implemented In:
Swift 4.0
Implemented
SE-0170NSNumber bridging and Numeric types
Author:Philippe Hausler
Review Manager:
Ben Cohen
Implemented In:
Swift 4.0
Implemented
SE-0169Improve Interaction Between `private` Declarations and Extensions
Authors:David Hart, Chris Lattner
Review Manager:
Doug Gregor
Bug:
SR-4616
Implemented In:
Swift 4.0
Implemented
SE-0168Multi-Line String Literals
Authors:John Holdsworth, Becca Royal-Gordon, Tyler Cloutier
Review Manager:
Joe Groff
Bugs:
SR-170, SR-4701, SR-4708, SR-4874
Implemented In:
Swift 4.0
Implementation:
swift#8813
Implemented
SE-0167Swift Encoders
Authors:Itai Ferber, Michael LeHew, Tony Parker
Review Manager:
Doug Gregor
Implemented In:
Swift 4.0
Implementation:
swift#9005
Implemented
SE-0166Swift Archival & Serialization
Authors:Itai Ferber, Michael LeHew, Tony Parker
Review Manager:
Doug Gregor
Implemented In:
Swift 4.0
Implementation:
swift#9004
Implemented
SE-0165Dictionary & Set Enhancements
Author:Nate Cook
Review Manager:
Ben Cohen
Implemented In:
Swift 4.0
Implemented
SE-0164Remove final support in protocol extensions
Author:Brian King
Review Manager:
Doug Gregor
Bug:
SR-1762
Implemented In:
Swift 4.0
Implemented
SE-0163String Revision: Collection Conformance, C Interop, Transcoding
Authors:Ben Cohen, Dave Abrahams
Review Manager:
John McCall
Implemented In:
Swift 4.0
Implemented
SE-0162Package Manager Custom Target Layouts
Author:Ankit Aggarwal
Review Manager:
Rick Ballard
Bug:
SR-29
Implemented In:
Swift 4.0
Implemented
SE-0161Smart KeyPaths: Better Key-Value Coding for Swift
Authors:David Smith, Michael LeHew, Joe Groff
Review Manager:
Doug Gregor
Implemented In:
Swift 4.0
Implemented
SE-0160Limiting `@objc` inference
Author:Doug Gregor
Review Manager:
Chris Lattner
Bug:
SR-4481
Implemented In:
Swift 4.0
Implementation:
swift#8379
Implemented
SE-0158Package Manager Manifest API Redesign
Author:Ankit Aggarwal
Review Manager:
Rick Ballard
Bug:
SR-3949
Implemented In:
Swift 4.0
Implemented
SE-0157Support recursive constraints on associated types
Authors:Douglas Gregor, Erica Sadun, Austin Zheng
Review Manager:
John McCall
Bug:
SR-1445
Implemented In:
Swift 4.1
Implemented
SE-0156Class and Subtype existentials
Authors:David Hart, Austin Zheng
Review Manager:
Doug Gregor
Bug:
SR-4296
Implemented In:
Swift 4.0
Implemented
SE-0155Normalize Enum Case Representation
Authors:Daniel Duan, Joe Groff
Review Manager:
John McCall
Bugs:
SR-4691, SR-12206, SR-12229
Implemented In:
Swift 3.0
Implemented
SE-0154Provide Custom Collections for Dictionary Keys and Values
Author:Nate Cook
Review Manager:
Doug Gregor
Implemented In:
Swift 4.0
Implemented
SE-0152Package Manager Tools Version
Author:Rick Ballard
Review Manager:
Anders Bertelrud
Bug:
SR-3965
Implemented In:
Swift 3.1
Implemented
SE-0151Package Manager Swift Language Compatibility Version
Authors:Daniel Dunbar, Rick Ballard
Review Manager:
Anders Bertelrud
Bug:
SR-3964
Implemented In:
Swift 3.1
Implemented
SE-0150Package Manager Support for branches
Author:Boris Bügling
Review Manager:
Daniel Dunbar
Bug:
SR-666
Implemented In:
Swift 4.0
Implemented
SE-0149Package Manager Support for Top of Tree development
Author:Boris Bügling
Review Manager:
Daniel Dunbar
Bug:
SR-3709
Implemented In:
Swift 4.0
Implemented
SE-0148Generic Subscripts
Author:Chris Eidhof
Review Manager:
Doug Gregor
Bug:
SR-115
Implemented In:
Swift 4.0
Implemented
SE-0147Move UnsafeMutablePointer.initialize(from:) to UnsafeMutableBufferPointer
Author:Ben Cohen
Review Manager:
Doug Gregor
Implemented In:
Swift 3.1
Implementation:
swift#6601
Implemented
SE-0146Package Manager Product Definitions
Author:Anders Bertelrud
Review Manager:
Daniel Dunbar
Bug:
SR-3606
Implemented In:
Swift 4.0
Implemented
SE-0145Package Manager Version Pinning
Authors:Daniel Dunbar, Ankit Aggarwal, Graydon Hoare
Review Manager:
Anders Bertelrud
Implemented In:
Swift 3.1
Implemented
SE-0143Conditional conformances
Author:Doug Gregor
Review Manager:
Joe Groff
Implemented In:
Swift 4.2
Implemented
SE-0142Permit where clauses to constrain associated types
Authors:David Hart, Jacob Bandes-Storch, Doug Gregor
Review Manager:
Doug Gregor
Bug:
SR-4506
Implemented In:
Swift 4.0
Implemented
SE-0141Availability by Swift version
Author:Graydon Hoare
Review Manager:
Doug Gregor
Bug:
SR-2709
Implemented In:
Swift 3.1
Implemented
SE-0140Warn when `Optional` converts to `Any`, and bridge `Optional` As Its Payload Or `NSNull`
Author:Joe Groff
Review Manager:
Doug Gregor
Implemented In:
Swift 3.0.1
Implemented
SE-0139Bridge Numeric Types to `NSNumber` and Cocoa Structs to `NSValue`
Author:Joe Groff
Review Manager:
Doug Gregor
Implemented In:
Swift 3.0.1
Implemented
SE-0138UnsafeRawBufferPointer
Author:Andrew Trick
Review Manager:
Dave Abrahams
Implemented In:
Swift 3.0.1
Implemented
SE-0137Avoiding Lock-In to Legacy Protocol Designs
Authors:Dave Abrahams, Dmitri Gribenko
Review Manager:
John McCall
Implemented In:
Swift 3.0
Implemented
SE-0136Memory layout of values
Author:Xiaodi Wu
Review Manager:
Dave Abrahams
Implemented In:
Swift 3.0
Implementation:
swift#4041
Implemented
SE-0135Package Manager Support for Differentiating Packages by Swift version
Author:Anders Bertelrud
Review Manager:
Daniel Dunbar
Implemented In:
Swift 3.0
Implemented
SE-0134Rename two UTF8-related properties on String
Authors:Xiaodi Wu, Erica Sadun
Review Manager:
Chris Lattner
Implemented In:
Swift 3.0
Implementation:
swift#3816
Implemented
SE-0133Rename `flatten()` to `joined()`
Author:Jacob Bandes-Storch
Review Manager:
Chris Lattner
Implemented In:
Swift 3.0
Implementation:
swift#3809, swift#3838, swift#3839
Implemented
SE-0131Add `AnyHashable` to the standard library
Author:Dmitri Gribenko
Review Manager:
Chris Lattner
Implemented In:
Swift 3.0
Implemented
SE-0130Replace repeating `Character` and `UnicodeScalar` forms of String.init
Author:Roman Levenstein
Review Manager:
Chris Lattner
Implemented In:
Swift 3.0
Implemented
SE-0129Package Manager Test Naming Conventions
Author:Anders Bertelrud
Review Manager:
Daniel Dunbar
Implemented In:
Swift 3.0
Implemented
SE-0128Change failable UnicodeScalar initializers to failable
Author:Xin Tong
Review Manager:
Chris Lattner
Implemented In:
Swift 3.0
Implementation:
swift#3662
Implemented
SE-0127Cleaning up stdlib Pointer and Buffer Routines
Author:Charlie Monroe
Review Manager:
Chris Lattner
Bugs:
SR-1937, SR-1955, SR-1957
Implemented In:
Swift 3.0
Implemented
SE-0125Remove `NonObjectiveCBase` and `isUniquelyReferenced`
Author:Arnold Schwaighofer
Review Manager:
Chris Lattner
Bug:
SR-1962
Implemented In:
Swift 3.0
Implemented
SE-0124`Int.init(ObjectIdentifier)` and `UInt.init(ObjectIdentifier)` should have a `bitPattern:` label
Author:Arnold Schwaighofer
Review Manager:
Chris Lattner
Bug:
SR-2064
Implemented In:
Swift 3.0
Implemented
SE-0121Remove `Optional` Comparison Operators
Author:Jacob Bandes-Storch
Review Manager:
Chris Lattner
Implemented In:
Swift 3.0
Implementation:
swift#3637
Implemented
SE-0120Revise `partition` Method Signature
Authors:Lorenzo Racca, Jeff Hajewski, Nate Cook
Review Manager:
Chris Lattner
Bug:
SR-1965
Implemented In:
Swift 3.0
Implemented
SE-0118Closure Parameter Names and Labels
Authors:Dave Abrahams, Dmitri Gribenko, Maxim Moiseev
Review Manager:
Chris Lattner
Bug:
SR-2072
Implemented In:
Swift 3.0
Implemented
SE-0117Allow distinguishing between public access and public overridability
Authors:Javier Soto, John McCall
Review Manager:
Chris Lattner
Implemented In:
Swift 3.0
Implementation:
swift#3882
Implemented
SE-0116Import Objective-C `id` as Swift `Any` type
Author:Joe Groff
Review Manager:
Chris Lattner
Implemented In:
Swift 3.0
Implemented
SE-0115Rename Literal Syntax Protocols
Author:Matthew Johnson
Review Manager:
Chris Lattner
Bug:
SR-2054
Implemented In:
Swift 3.0
Implemented
SE-0114Updating Buffer "Value" Names to "Header" Names
Author:Erica Sadun
Review Manager:
Chris Lattner
Implemented In:
Swift 3.0
Implementation:
swift#3374
Implemented
SE-0113Add integral rounding functions to FloatingPoint
Author:Karl Wagner
Review Manager:
Chris Lattner
Bug:
SR-2010
Implemented In:
Swift 3.0
Implemented
SE-0112Improved NSError Bridging
Authors:Doug Gregor, Charles Srstka
Review Manager:
Chris Lattner
Implemented In:
Swift 3.0
Implemented
SE-0111Remove type system significance of function argument labels
Author:Austin Zheng
Review Manager:
Chris Lattner
Bug:
SR-2009
Implemented In:
Swift 3.0
Implemented
SE-0110Distinguish between single-tuple and multiple-argument function types
Authors:Vladimir S., Austin Zheng
Review Manager:
Chris Lattner
Bug:
SR-2008
Implemented In:
Swift
Implemented
SE-0109Remove the `Boolean` protocol
Authors:Anton Zhilin, Chris Lattner
Review Manager:
Doug Gregor
Implemented In:
Swift 3.0
Implementation:
swift@76cf339, swift@af30ae3
Implemented
SE-0107UnsafeRawPointer API
Author:Andrew Trick
Review Manager:
Chris Lattner
Implemented In:
Swift 3.0
Implemented
SE-0106Add a `macOS` Alias for the `OSX` Platform Configuration Test
Author:Erica Sadun
Review Manager:
Chris Lattner
Bugs:
SR-1823, SR-1887
Implemented In:
Swift 3.0
Implemented
SE-0104Protocol-oriented integers
Authors:Dave Abrahams, Maxim Moiseev
Review Manager:
Ben Cohen
Bug:
SR-3196
Implemented In:
Swift 4.0
Implemented
SE-0103Make non-escaping closures the default
Author:Trent Nadeau
Review Manager:
Chris Lattner
Bug:
SR-1952
Implemented In:
Swift 3.0
Implemented
SE-0102Remove `@noreturn` attribute and introduce an empty `Never` type
Author:Joe Groff
Review Manager:
Chris Lattner
Bug:
SR-1953
Implemented In:
Swift 3.0
Implemented
SE-0101Reconfiguring `sizeof` and related functions into a unified `MemoryLayout` struct
Authors:Erica Sadun, Dave Abrahams
Review Manager:
Chris Lattner
Implemented In:
Swift 3.0
Implemented
SE-0099Restructuring Condition Clauses
Authors:Erica Sadun, Chris Lattner
Review Manager:
Joe Groff
Implemented In:
Swift 3.0
Implemented
SE-0096Converting `dynamicType` from a property to an operator
Author:Erica Sadun
Review Manager:
Chris Lattner
Bug:
SR-2218
Implemented In:
Swift 3.0
Implemented
SE-0095Replace `protocol<P1,P2>` syntax with `P1 & P2` syntax
Authors:Adrian Zubarev, Austin Zheng
Review Manager:
Chris Lattner
Bug:
SR-1938
Implemented In:
Swift 3.0
Implemented
SE-0094Add sequence(first:next:) and sequence(state:next:) to the stdlib
Authors:Lily Ballard, Erica Sadun
Review Manager:
Chris Lattner
Bug:
SR-1622
Implemented In:
Swift 3.0
Implemented
SE-0093Adding a public `base` property to slices
Author:Max Moiseev
Review Manager:
Dave Abrahams
Implemented In:
Swift 3.0
Implementation:
swift#2929
Implemented
SE-0092Typealiases in protocols and protocol extensions
Authors:David Hart, Doug Gregor
Review Manager:
Chris Lattner
Bug:
SR-1539
Implemented In:
Swift 3.0
Implemented
SE-0091Improving operator requirements in protocols
Authors:Tony Allevato, Doug Gregor
Review Manager:
Chris Lattner
Bug:
SR-2073
Implemented In:
Swift 3.0
Implemented
SE-0089Renaming `String.init<T>(_: T)`
Authors:Austin Zheng, Becca Royal-Gordon
Review Manager:
Chris Lattner
Bug:
SR-1881
Implemented In:
Swift 3.0
Implemented
SE-0088Modernize libdispatch for Swift 3 naming conventions
Author:Matt Wright
Review Manager:
Chris Lattner
Implemented In:
Swift 3.0
Implemented
SE-0086Drop NS Prefix in Swift Foundation
Authors:Tony Parker, Philippe Hausler
Review Manager:
Doug Gregor
Implemented In:
Swift 3.0
Implemented
SE-0085Package Manager Command Names
Authors:Rick Ballard, Daniel Dunbar
Review Manager:
Daniel Dunbar
Implemented In:
Swift 3.0
Implementation:
swift-package-manager#364
Implemented
SE-0082Package Manager Editable Packages
Author:Daniel Dunbar
Review Manager:
Anders Bertelrud
Implemented In:
Swift 3.1
Implemented
SE-0081Move `where` clause to end of declaration
Authors:David Hart, Robert Widmann, Pyry Jahkola
Review Manager:
Chris Lattner
Bug:
SR-1561
Implemented In:
Swift 3.0
Implemented
SE-0080Failable Numeric Conversion Initializers
Author:Matthew Johnson
Review Manager:
Chris Lattner
Bug:
SR-1491
Implemented In:
Swift 3.1
Implemented
SE-0079Allow using optional binding to upgrade `self` from a weak to strong reference
Author:Evan Maloney
Review Manager:
TBD
Implemented In:
Swift 4.2
Implementation:
swift#15306
Implemented
SE-0077Improved operator declarations
Author:Anton Zhilin
Review Manager:
Joe Groff
Implemented In:
Swift 3.0
Implemented
SE-0076Add overrides taking an UnsafePointer source to non-destructive copying methods on UnsafeMutablePointer
Author:Janosch Hildebrand
Review Manager:
Chris Lattner
Bug:
SR-1490
Implemented In:
Swift 3.0
Implemented
SE-0075Adding a Build Configuration Import Test
Author:Erica Sadun
Review Manager:
Chris Lattner
Bug:
SR-1560
Implemented In:
Swift 4.1
Implemented
SE-0072Fully eliminate implicit bridging conversions from Swift
Author:Joe Pamer
Review Manager:
Chris Lattner
Implemented In:
Swift 3.0
Implementation:
swift#2419
Implemented
SE-0071Allow (most) keywords in member references
Author:Doug Gregor
Review Manager:
Chris Lattner
Implemented In:
Swift 3.0
Implemented
SE-0070Make Optional Requirements Objective-C-only
Author:Doug Gregor
Review Manager:
Chris Lattner
Bug:
SR-1395
Implemented In:
Swift 3.0
Implemented
SE-0069Mutability and Foundation Value Types
Author:Tony Parker
Review Manager:
Chris Lattner
Implemented In:
Swift 3.0
Implemented
SE-0068Expanding Swift `Self` to class members and value types
Author:Erica Sadun
Review Manager:
Chris Lattner
Bug:
SR-1340
Implemented In:
Swift 5.1
Implementation:
swift#22863
Implemented
SE-0067Enhanced Floating Point Protocols
Author:Stephen Canon
Review Manager:
Chris Lattner
Implemented In:
Swift 3.0
Implementation:
swift#2453
Implemented
SE-0066Standardize function type argument syntax to require parentheses
Author:Chris Lattner
Review Manager:
Doug Gregor
Implemented In:
Swift 3.0
Implementation:
swift@3d2b5bc
Implemented
SE-0065A New Model for Collections and Indices
Authors:Dmitri Gribenko, Dave Abrahams, Maxim Moiseev
Review Manager:
Chris Lattner
Implemented In:
Swift 3.0
Implementation:
swift#2108
Implemented
SE-0064Referencing the Objective-C selector of property getters and setters
Author:David Hart
Review Manager:
Doug Gregor
Bug:
SR-1239
Implemented In:
Swift 3.0
Implemented
SE-0063SwiftPM System Module Search Paths
Author:Max Howell
Review Manager:
Anders Bertelrud
Implemented In:
Swift 3.0
Implementation:
swift-package-manager#257
Implemented
SE-0062Referencing Objective-C key-paths
Author:David Hart
Review Manager:
Doug Gregor
Bug:
SR-1237
Implemented In:
Swift 3.0
Implemented
SE-0061Add Generic Result and Error Handling to autoreleasepool()
Author:Timothy J. Wood
Review Manager:
Dave Abrahams
Bugs:
SR-842, SR-1394
Implemented In:
Swift 3.0
Implemented
SE-0060Enforcing order of defaulted parameters
Author:Joe Groff
Review Manager:
Chris Lattner
Bug:
SR-1489
Implemented In:
Swift 3.0
Implemented
SE-0059Update API Naming Guidelines and Rewrite Set APIs Accordingly
Author:Dave Abrahams
Review Manager:
Doug Gregor
Implemented In:
Swift 3.0
Implemented
SE-0057Importing Objective-C Lightweight Generics
Author:Doug Gregor
Review Manager:
Chris Lattner
Implemented In:
Swift 3.0
Implemented
SE-0055Make unsafe pointer nullability explicit using Optional
Author:Jordan Rose
Review Manager:
Chris Lattner
Implemented In:
Swift 3.0
Implementation:
swift#1878
Implemented
SE-0054Abolish `ImplicitlyUnwrappedOptional` type
Author:Chris Willmore
Review Manager:
Chris Lattner
Implemented In:
Swift 4.2
Implementation:
swift#14299
Implemented
SE-0053Remove explicit use of `let` from Function Parameters
Author:Nicholas Maccharoli
Review Manager:
Chris Lattner
Implemented In:
Swift 3.0
Implementation:
swift#1812
Implemented
SE-0052Change IteratorType post-nil guarantee
Author:Patrick Pijnappel
Review Manager:
Chris Lattner
Implemented In:
Swift 3.0
Implementation:
swift#1702
Implemented
SE-0049Move @noescape and @autoclosure to be type attributes
Author:Chris Lattner
Review Manager:
Doug Gregor
Bug:
SR-1235
Implemented In:
Swift 3.0
Implemented
SE-0048Generic Type Aliases
Author:Chris Lattner
Review Manager:
Doug Gregor
Implemented In:
Swift 3.0
Implemented
SE-0047Defaulting non-Void functions so they warn on unused results
Authors:Erica Sadun, Adrian Kashivskyy
Review Manager:
Chris Lattner
Bug:
SR-1052
Implemented In:
Swift 3.0
Implemented
SE-0046Establish consistent label behavior across all parameters including first labels
Authors:Jake Carter, Erica Sadun
Review Manager:
Chris Lattner
Bug:
SR-961
Implemented In:
Swift 3.0
Implemented
SE-0045Add prefix(while:) and drop(while:) to the stdlib
Author:Lily Ballard
Review Manager:
Chris Lattner
Bug:
SR-1516
Implemented In:
Swift 3.1
Implemented
SE-0044Import as member
Author:Michael Ilseman
Review Manager:
Doug Gregor
Bug:
SR-1053
Implemented In:
Swift 3.0
Implemented
SE-0043Declare variables in 'case' labels with multiple patterns
Author:Andrew Bennett
Review Manager:
Chris Lattner
Implemented In:
Swift 3.0
Implementation:
swift#1383
Implemented
SE-0040Replacing Equal Signs with Colons For Attribute Arguments
Author:Erica Sadun
Review Manager:
Chris Lattner
Implemented In:
Swift 3.0
Implementation:
swift#1537
Implemented
SE-0039Modernizing Playground Literals
Author:Erica Sadun
Review Manager:
Chris Lattner
Bug:
SR-917
Implemented In:
Swift 3.0
Implemented
SE-0038Package Manager C Language Target Support
Author:Daniel Dunbar
Review Manager:
Rick Ballard
Bug:
SR-821
Implemented In:
Swift 3.0
Implemented
SE-0037Clarify interaction between comments & operators
Author:Jesse Rusak
Review Manager:
Chris Lattner
Bug:
SR-960
Implemented In:
Swift 3.0
Implemented
SE-0036Requiring Leading Dot Prefixes for Enum Instance Member Implementations
Authors:Erica Sadun, Chris Lattner
Review Manager:
Doug Gregor
Bug:
SR-1236
Implemented In:
Swift 3.0
Implemented
SE-0035Limiting `inout` capture to `@noescape` contexts
Author:Joe Groff
Review Manager:
Chris Lattner
Bug:
SR-807
Implemented In:
Swift 3.0
Implemented
SE-0034Disambiguating Line Control Statements from Debugging Identifiers
Author:Erica Sadun
Review Manager:
Chris Lattner
Bug:
SR-840
Implemented In:
Swift 3.0
Implemented
SE-0033Import Objective-C Constants as Swift Types
Author:Jeff Kelley
Review Manager:
John McCall
Implemented In:
Swift 3.0
Implemented
SE-0032Add `first(where:)` method to `Sequence`
Author:Lily Ballard
Review Manager:
Chris Lattner
Bug:
SR-1519
Implemented In:
Swift 3.0
Implemented
SE-0031Adjusting `inout` Declarations for Type Decoration
Authors:Joe Groff, Erica Sadun
Review Manager:
Chris Lattner
Implemented In:
Swift 3.0
Implementation:
swift#1333
Implemented
SE-0029Remove implicit tuple splat behavior from function applications
Author:Chris Lattner
Review Manager:
Joe Groff
Implemented In:
Swift 3.0
Implementation:
swift@8e12008
Implemented
SE-0028Modernizing Swift's Debugging Identifiers
Author:Erica Sadun
Review Manager:
Chris Lattner
Bug:
SR-669
Implemented In:
Swift 2.2
Implemented
SE-0025Scoped Access Level
Author:Ilya Belenkiy
Review Manager:
Doug Gregor
Bug:
SR-1275
Implemented In:
Swift 3.0
Implemented
SE-0023API Design Guidelines
Authors:Dave Abrahams, Doug Gregor, Dmitri Gribenko, Ted Kremenek, Chris Lattner, Alex Migicovsky, Max Moiseev, Ali Ozer, Tony Parker
Review Manager:
Doug Gregor
Implemented In:
Swift 3.0
Implemented
SE-0022Referencing the Objective-C selector of a method
Author:Doug Gregor
Review Manager:
Joe Groff
Implemented In:
Swift 2.2
Implementation:
swift#1170
Implemented
SE-0021Naming Functions with Argument Labels
Author:Doug Gregor
Review Manager:
Joe Groff
Implemented In:
Swift 2.2
Implementation:
swift@ecfde0e
Implemented
SE-0020Swift Language Version Build Configuration
Author:Ashley Garland
Review Manager:
Doug Gregor
Implemented In:
Swift 2.2
Implementation:
swift@c32fb8e
Implemented
SE-0019Swift Testing
Authors:Max Howell, Daniel Dunbar, Mattt Thompson
Review Manager:
Rick Ballard
Bug:
SR-592
Implemented In:
Swift 3.0
Implemented
SE-0017Change `Unmanaged` to use `UnsafePointer`
Author:Jacob Bandes-Storch
Review Manager:
Chris Lattner
Bug:
SR-1485
Implemented In:
Swift 3.0
Implemented
SE-0016Add initializers to Int and UInt to convert from UnsafePointer and UnsafeMutablePointer
Author:Michael Buckley
Review Manager:
Chris Lattner
Bug:
SR-1115
Implemented In:
Swift 3.0
Implemented
SE-0015Tuple comparison operators
Author:Lily Ballard
Review Manager:
Dave Abrahams
Implemented In:
Swift 2.2
Implementation:
swift#408
Implemented
SE-0014Constraining `AnySequence.init`
Author:Max Moiseev
Review Manager:
Doug Gregor
Bug:
SR-474
Implemented In:
Swift 2.2
Implemented
SE-0011Replace `typealias` keyword with `associatedtype` for associated type declarations
Author:Loïc Lecrenier
Review Manager:
Doug Gregor
Bug:
SR-511
Implemented In:
Swift 2.2
Implemented
SE-0008Add a Lazy flatMap for Sequences of Optionals
Author:Oisin Kidney
Review Manager:
Doug Gregor
Bug:
SR-361
Implemented In:
Swift 3.0
Implemented
SE-0007Remove C-style for-loops with conditions and incrementers
Author:Erica Sadun
Review Manager:
Doug Gregor
Bugs:
SR-226, SR-227
Implemented In:
Swift 3.0
Implemented
SE-0006Apply API Guidelines to the Standard Library
Authors:Dave Abrahams, Dmitri Gribenko, Maxim Moiseev
Review Manager:
Doug Gregor
Implemented In:
Swift 3.0
Implemented
SE-0005Better Translation of Objective-C APIs Into Swift
Authors:Doug Gregor, Dave Abrahams
Review Manager:
Doug Gregor
Implemented In:
Swift 3.0
Implemented
SE-0004Remove the `++` and `--` operators
Author:Chris Lattner
Implemented In:
Swift 3.0
Implementation:
swift@8e12008
Implemented
SE-0003Removing `var` from Function Parameters
Author:Ashley Garland
Review Manager:
Joe Pamer
Implemented In:
Swift 3.0
Implementation:
swift@8a5ed40
Implemented
SE-0002Removing currying `func` declaration syntax
Author:Joe Groff
Implemented In:
Swift 3.0
Implementation:
swift@983a674
Implemented
SE-0001Allow (most) keywords as argument labels
Author:Doug Gregor
Bug:
SR-344
Implemented In:
Swift 2.2
Returned
SE-0516Borrowing Sequence
Authors:Nate Cook, Ben Cohen
Review Manager:
Holly Borla
Implementation:
swift#86811, swift#87483
Status:
Returned for Revision
Returned
SE-0513API to get the path to the current executable
Author:Jonathan Grynspan
Review Manager:
Tony Allevato
Implementation:
swift#85496
Status:
Returned for Revision
Returned
SE-0505Delayed Enqueuing for Executors
Author:Alastair Houghton
Review Manager:
Freddy Kellison-Linn
Status:
Returned for Revision
Returned
SE-0490Environment Constrained Shared Libraries
Author:tayloraswift
Review Manager:
Alastair Houghton
Bug:
SR-5714
Implementation:
swift-package-manager#8249
Status:
Returned for Revision
Returned
SE-0479Method and Initializer Key Paths
Authors:Amritpan Kaur, Pavel Yaskevich
Review Manager:
Becca Royal-Gordon
Implementation:
swift#78823, swift-syntax#2950, swift-foundation#1179
Status:
Returned for Revision
Returned
SE-0478Default actor isolation typealias
Author:Holly Borla
Review Manager:
Steve Canon
Implementation:
swift#80572
Status:
Returned for Revision
Returned
SE-0406Backpressure support for AsyncStream
Author:Franz Busch
Review Manager:
Xiaodi Wu
Implementation:
swift#66488
Status:
Returned for Revision
Returned
SE-0403Package Manager Mixed Language Target Support
Author:Nick Cooke
Review Manager:
Saleem Abdulrasool
Implementation:
swift-package-manager#5919
Status:
Returned for Revision
Returned
SE-0385Custom Reflection Metadata
Authors:Pavel Yaskevich, Holly Borla, Alejandro Alonso, Stuart Montgomery
Review Manager:
Doug Gregor
Implementation:
swift#62426, swift#62738, swift#62818, swift#62850, swift#62920, swift#63057
Status:
Returned for Revision
Returned
SE-0379Swift Opt-In Reflection Metadata
Author:Max Ovtsin
Review Manager:
Joe Groff
Implementation:
swift#34199
Status:
Returned for Revision
Returned
SE-0359Build-Time Constant Values
Authors:Artem Chikin, Ben Cohen, Xi Ge
Review Manager:
Doug Gregor
Status:
Returned for Revision
Returned
SE-0330Conditionals in Collections
Author:John Holdsworth
Review Manager:
Ted Kremenek
Bug:
SR-8743
Implementation:
swift#19347
Status:
Returned for Revision
Returned
SE-0318Package Creation
Author:Miguel Perez
Review Manager:
Tom Doron
Implementation:
swift-package-manager#3514
Status:
Returned for Revision
Returned
SE-0312Add `indexed()` and `Collection` conformances for `enumerated()` and `zip(_:_:)`
Author:Tim Vermeulen
Review Manager:
Ben Cohen
Implementation:
swift#36851
Status:
Returned for Revision
Returned
SE-0283Tuples Conform to `Equatable`, `Comparable`, and `Hashable`
Author:Alejandro Alonso
Review Manager:
Saleem Abdulrasool
Implementation:
swift#28833
Status:
Returned for Revision
Returned
SE-0265Offset-Based Access to Indices, Elements, and Slices
Author:Michael Ilseman
Review Manager:
John McCall
Implementation:
swift#24296
Status:
Returned for Revision
Returned
SE-0259Approximate Equality for Floating Point
Author:Stephen Canon
Review Manager:
Ben Cohen
Implementation:
swift#23839
Status:
Returned for Revision
Returned
SE-0257Eliding commas from multiline expression lists
Authors:Nate Chandler, Matthew Johnson
Review Manager:
Ted Kremenek
Implementation:
swift#22714
Status:
Returned for Revision
Returned
SE-0250Swift Code Style Guidelines and Formatter
Authors:Tony Allevato, Dave Abrahams
Review Manager:
Ted Kremenek
Status:
Returned for Revision
Returned
SE-0177Add clamp(to:) to the stdlib
Author:Nicholas Maccharoli
Review Manager:
TBD
Status:
Returned for Revision
Returned
SE-0090Remove `.self` and freely allow type references in expressions
Authors:Joe Groff, Tanner Nelson
Review Manager:
Chris Lattner
Status:
Returned for Revision
Returned
SE-0078Implement a rotate algorithm, equivalent to std::rotate() in C++
Authors:Nate Cook, Sergey Bolshedvorsky
Review Manager:
Chris Lattner
Status:
Returned for Revision
Returned
SE-0018Flexible Memberwise Initialization
Author:Matthew Johnson
Review Manager:
Chris Lattner
Status:
Returned for Revision
Rejected
SE-0275Allow more characters (like whitespaces and punctuations) for escaped identifiers
Author:Alfredo Delli Bovi
Review Manager:
Joe Groff
Implementation:
swift#28966
Rejected
SE-0256Introduce `{Mutable}ContiguousCollection` protocol
Author:Ben Cohen
Review Manager:
Ted Kremenek
Implementation:
swift#23616
Rejected
SE-0246Generic Math(s) Functions
Author:Stephen Canon
Review Manager:
John McCall
Implementation:
swift#23140
Rejected
SE-0243Integer-convertible character literals
Authors:Diana Ma (“Taylor Swift”), Chris Lattner, John Holdsworth
Review Manager:
Ben Cohen
Implementation:
swift#21873
Rejected
SE-0231Optional Iteration
Author:Anthony Latsis
Review Manager:
Joe Groff
Implementation:
swift#19207
Rejected
SE-0222Lazy CompactMap Sequence
Authors:TellowKrinkle, Johannes Weiß
Review Manager:
John McCall
Implementation:
swift#14841
Rejected
SE-0217Introducing the `!!` "Unwrap or Die" operator to the Swift Standard Library
Authors:Ben Cohen, Dave DeLong, Paul Cantrell, Erica Sadun, and several other folk
Review Manager:
Joe Groff
Rejected
SE-0203Rename Sequence.elementsEqual
Author:Xiaodi Wu
Review Manager:
Ted Kremenek
Bug:
SR-6102
Implementation:
swift#12884
Rejected
SE-0159Fix Private Access Levels
Author:David Hart
Review Manager:
Doug Gregor
Rejected
SE-0153Compensate for the inconsistency of `@NSCopying`'s behaviour
Author:Torin Kwok
Review Manager:
Doug Gregor
Bug:
SR-4538
Rejected
SE-0144Allow Single Dollar Sign as a Valid Identifier
Author:Ankur Patel
Review Manager:
Chris Lattner
Rejected
SE-0132Rationalizing Sequence end-operation names
Authors:Becca Royal-Gordon, Dave Abrahams
Review Manager:
Chris Lattner
Rejected
SE-0123Disallow coercion to optionals in operator arguments
Authors:Mark Lacey, Doug Gregor, Jacob Bandes-Storch
Review Manager:
Chris Lattner
Rejected
SE-0122Use colons for subscript declarations
Author:James Froggatt
Review Manager:
Chris Lattner
Rejected
SE-0119Remove access modifiers from extensions
Author:Adrian Zubarev
Review Manager:
Chris Lattner
Rejected
SE-0108Remove associated type inference
Authors:Douglas Gregor, Austin Zheng
Review Manager:
Chris Lattner
Rejected
SE-0105Removing Where Clauses from For-In Loops
Author:Erica Sadun
Review Manager:
Chris Lattner
Rejected
SE-0098Lowercase `didSet` and `willSet` for more consistent keyword casing
Author:Erica Sadun
Review Manager:
Chris Lattner
Rejected
SE-0097Normalizing naming for "negative" attributes
Author:Erica Sadun
Review Manager:
Chris Lattner
Rejected
SE-0087Rename `lazy` to `@lazy`
Author:Anton3
Review Manager:
Chris Lattner
Rejected
SE-0084Allow trailing commas in parameter lists and tuples
Authors:Grant Paul, Erica Sadun
Review Manager:
Chris Lattner
Rejected
SE-0083Remove bridging conversion behavior from dynamic casts
Author:Joe Groff
Review Manager:
Chris Lattner
Rejected
SE-0074Implementation of Binary Search functions
Authors:Lorenzo Racca, Jeff Hajewski, Nate Cook
Review Manager:
Chris Lattner
Rejected
SE-0073Marking closures as executing exactly once
Authors:Félix Cloutier, Gwendal Roué
Review Manager:
Chris Lattner
Rejected
SE-0058Allow Swift types to provide custom Objective-C representations
Authors:Russ Bishop, Doug Gregor
Review Manager:
Joe Groff
Rejected
SE-0056Allow trailing closures in `guard` conditions
Author:Chris Lattner
Review Manager:
Doug Gregor
Rejected
SE-0042Flattening the function type of unapplied method references
Author:Joe Groff
Review Manager:
Chris Lattner
Bug:
SR-1051
Rejected
SE-0041Updating Protocol Naming Conventions for Conversions
Authors:Matthew Johnson, Erica Sadun
Review Manager:
Chris Lattner
Rejected
SE-0027Expose code unit initializers on String
Author:Zachary Waldowski
Review Manager:
Doug Gregor
Rejected
SE-0026Abstract classes and methods
Author:David Scrève
Review Manager:
Joe Groff
Rejected
SE-0024Optional Value Setter `??=`
Author:James Campbell
Review Manager:
Doug Gregor
Rejected
SE-0013Remove Partial Application of Non-Final Super Methods (Swift 2.2)
Author:Ashley Garland
Review Manager:
Doug Gregor
Rejected
SE-0012Add `@noescape` to public library API
Author:Jacob Bandes-Storch
Review Manager:
Philippe Hausler
Rejected
SE-0010Add StaticString.UnicodeScalarView
Author:Lily Ballard
Review Manager:
Doug Gregor
Rejected
SE-0009Require self for accessing instance members
Author:David Hart
Review Manager:
Doug Gregor
Withdrawn
SE-0262Demangle Function
Author:Alejandro Alonso
Review Manager:
Joe Groff
Implementation:
swift#25314
Withdrawn
SE-0223Accessing an Array's Uninitialized Buffer
Author:Nate Cook
Review Manager:
Joe Groff
Bug:
SR-3087
Implementation:
swift#17389
Withdrawn
SE-0126Refactor Metatypes, repurpose `T.self` and `Mirror`
Authors:Adrian Zubarev, Anton Zhilin
Review Manager:
Chris Lattner
Withdrawn
SE-0100Add sequence-based initializers and merge methods to Dictionary
Author:Nate Cook
Review Manager:
TBD
Withdrawn
SE-0051Conventionalizing `stride` semantics
Author:Erica Sadun
Review Manager:
N/A
Withdrawn
SE-0050Decoupling Floating Point Strides from Generic Implementations
Authors:Erica Sadun, Xiaodi Wu
Review Manager:
Chris Lattner
Withdrawn
SE-0030Property Behaviors
Author:Joe Groff
Review Manager:
Doug Gregor
Docs
Community
Packages
Blog
Install
Tools
Xcode
Visual Studio Code
Emacs
Neovim
Cursor
Other Editors
Community
Overview
Swift Evolution
Diversity
Mentorship
Contributing
Governance
Code of Conduct
License
Security
Color scheme preference
Light
Dark
Auto

Copyright © 2026 Apple Inc. All rights reserved.

Swift and the Swift logo are trademarks of Apple Inc.

Privacy Policy
Cookies
API

