---
title: "Automatic Reference Counting"
url: "https://docs.swift.org/swift-book/documentation/the-swift-programming-language/automaticreferencecounting/"
author: "Apple (Swift Team)"
date: "2026-05-13"
ingested_at: "2026-05-13T16:27:39Z"
tags: [swift, ARC, memory-management, reference-counting, weak-references, closures]
topic: Swift
status: uncompiled
---

# Automatic Reference Counting

Swift uses _Automatic Reference Counting_ (ARC) to track and manage your app's memory usage. In most cases, this means that memory management "just works" in Swift, and you don't need to think about memory management yourself. ARC automatically frees up the memory used by class instances when those instances are no longer needed.

Reference counting applies only to instances of classes. Structures and enumerations are value types, not reference types, and aren't stored and passed by reference.

## How ARC Works

Every time you create a new instance of a class, ARC allocates a chunk of memory to store information about that instance. This memory holds information about the type of the instance, together with the values of any stored properties associated with that instance.

Additionally, when an instance is no longer needed, ARC frees up the memory used by that instance so that the memory can be used for other purposes instead. This ensures that class instances don't take up space in memory when they're no longer needed.

To make sure that instances don't disappear while they're still needed, ARC tracks how many properties, constants, and variables are currently referring to each class instance. ARC will not deallocate an instance as long as at least one active reference to that instance still exists.

To make this possible, whenever you assign a class instance to a property, constant, or variable, that property, constant, or variable makes a _strong reference_ to the instance.

## Strong Reference Cycles Between Class Instances

It's possible to write code in which an instance of a class _never_ gets to a point where it has zero strong references. This can happen if two class instances hold a strong reference to each other, such that each instance keeps the other alive. This is known as a _strong reference cycle_.

You resolve strong reference cycles by defining some of the relationships between classes as weak or unowned references instead of as strong references.

## Resolving Strong Reference Cycles

Swift provides two ways to resolve strong reference cycles when you work with properties of class type: weak references and unowned references.

Weak and unowned references enable one instance in a reference cycle to refer to the other instance _without_ keeping a strong hold on it.

Use a **weak reference** when the other instance has a shorter lifetime — that is, when the other instance can be deallocated first. Use an **unowned reference** when the other instance has the same lifetime or a longer lifetime.

### Weak References

A _weak reference_ is a reference that doesn't keep a strong hold on the instance it refers to, and so doesn't stop ARC from disposing of the referenced instance. ARC automatically sets a weak reference to `nil` when the instance that it refers to is deallocated. Weak references are always declared as variables, rather than constants, of an optional type.

### Unowned References

Like a weak reference, an _unowned reference_ doesn't keep a strong hold on the instance it refers to. Unlike a weak reference, however, an unowned reference is used when the other instance has the same lifetime or a longer lifetime. Unlike a weak reference, an unowned reference is expected to always have a value — ARC never sets an unowned reference's value to `nil`.

### Unowned Optional References

You can mark an optional reference to a class as unowned. An unowned optional reference doesn't keep a strong hold on the instance of the class that it wraps, and so it doesn't prevent ARC from deallocating the instance. You're responsible for ensuring that it always refers to a valid object or is set to `nil`.

### Unowned References and Implicitly Unwrapped Optional Properties

A third scenario exists in which _both_ properties should always have a value, and neither property should ever be `nil` once initialization is complete. In this scenario, combine an unowned property on one class with an implicitly unwrapped optional property on the other class.

This enables both properties to be accessed directly (without optional unwrapping) once initialization is complete, while still avoiding a reference cycle. The `capitalCity` property of `Country` is declared as an implicitly unwrapped optional (`City!`), meaning it has a default value of `nil` but can be accessed without need to unwrap after initialization.


## Strong Reference Cycles for Closures

A strong reference cycle can also occur if you assign a closure to a property of a class instance, and the body of that closure captures the instance. This capture might occur because the closure's body accesses a property of the instance, such as `self.someProperty`, or because the closure calls a method on the instance, such as `self.someMethod()`.

Swift provides an elegant solution: a _closure capture list_.

### Defining a Capture List

Each item in a capture list is a pairing of the `weak` or `unowned` keyword with a reference to a class instance (such as `self`) or a variable initialized with some value. Place the capture list before a closure's parameter list and return type:

```swift
someClosure { [weak self, unowned delegate] in
    // closure body
}
```

If a closure doesn't specify a parameter list or return type, place the capture list at the very start, followed by `in`:

```swift
someClosure { [unowned self] in
    // closure body
}
```

### Weak and Unowned in Capture Lists

Define a capture as an **unowned reference** when the closure and the instance it captures will always refer to each other, and will always be deallocated at the same time. Define a capture as a **weak reference** when the captured reference may become `nil` at some point in the future.
