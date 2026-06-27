---
title: "Precise Control Flags over Compiler Warnings – available from Swift 6.1"
source: "https://www.hackingwithswift.com/swift/6.1/diagnostic-groups"
author:
  - "[[Paul Hudson @twostraws]]"
published:
created: 2026-04-13
description: "Precise Control Flags over Compiler Warnings example code from Swift 6.1"
tags:
  - "clippings"
---
[Paul Hudson](https://www.hackingwithswift.com/about) [@twostraws](https://twitter.com/twostraws)

[SE-0443](https://github.com/swiftlang/swift-evolution/blob/main/proposals/0443-warning-control-flags.md) introduces fine-grained control over how the Swift compiler issues warnings and errors, adding more flexibility around the "Suppress warnings" and "Treat warnings as errors" options we had previously.

The first step to using this new feature is to add adjust your target's build settings to include `-print-diagnostic-groups` in the "Swift Compiler – Custom Flags" section in Xcode. If you're building on the command line, use something like `swiftc -print-diagnostic-groups main.swift`.

Once you do that, many warnings and errors will give you a little extra information at the end. For example, code like this in a macOS Sequoia target triggers a deprecation warning:

```swift
@available(macOS, deprecated: 15.0, renamed: "sequoiaFunction")
func sonomaFunction() { }

sonomaFunction()
```

That deprecation warning would always have been there, but now there's an extra piece of information at the end: **'sonomaFunction()' was deprecated in macOS 15.0: renamed to 'sequoiaFunction' \[DeprecatedDeclaration\]** – that `[DeprecatedDeclaration]` part is the *diagnostic group* this warning belongs to, which is our entry point into controlling how that type of message is applied.

Now we know the diagnostic group we're dealing with, we can modify our Swift compiler flags to treat that type of warning differently. For example, if you add the `-Werror DeprecatedDeclaration` flag it means "upgrade deprecation warnings to errors." **Important:** In Xcode these flags must be added one at a time, so your final compiler flags list will be `-print-diagnostic-groups`, `-Werror`, `DeprecatedDeclaration`.

The opposite of `-Werror` is `-Wwarning`, which tells Swift that a particular warning should remain a warning even when "Treat warnings as errors" is enabled.

**Important:** The order these flags are applied affects the end result. If you say "treat all warnings as errors except this one," that's what you'll get, but if you say "ensure deprecations are only a warning, but treat all warnings as errors," then deprecations will be upgraded to errors. Because Xcode applies its build settings in its own order, you might find you need to use `-warnings-as-errors` manually in Other Swift Flags, positioning it exactly where you want it.

You can find some [documentation on Swift's diagnostic groups here](https://docs.swift.org/compiler/documentation/diagnostics), although it's an area that's evolving quickly and that documentation already references things not available in Swift 6.1.

- [Allow trailing comma in comma-separated lists](https://www.hackingwithswift.com/swift/6.1/trailing-commas)
- [Metatype Keypaths](https://www.hackingwithswift.com/swift/6.1/metatype-key-paths)
- [Allow TaskGroup's ChildTaskResult Type To Be Inferred](https://www.hackingwithswift.com/swift/6.1/childtaskresult-inference)
- [Allow nonisolated to prevent global actor inference](https://www.hackingwithswift.com/swift/6.1/prevent-global-actor-inference)
- [Member import visibility](https://www.hackingwithswift.com/swift/6.1/member-import-visiblity)
- [Formalize ‘language mode’ terminology](https://www.hackingwithswift.com/swift/6.1/language-mode)
- [Swift Testing: Range-based confirmations](https://www.hackingwithswift.com/swift/6.1/swift-testing-range-confirmations)
- [Swift Testing: Return errors from expect(throws:)](https://www.hackingwithswift.com/swift/6.1/swift-testing-errors-expect)
- [Swift Testing: Test Scoping Traits](https://www.hackingwithswift.com/swift/6.1/swift-testing-test-scopes)

![Unknown user](https://www.hackingwithswift.com/img/unknown_user.png)

You are not logged in