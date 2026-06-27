---
source: https://x.com/jacobtechtavern/status/2058507557157712218?s=48
ingested_at: 2026-05-24T11:34:13Z
type: web
author: Jacob Bartlett
handle: "@jacobtechtavern"
status: uncompiled
---

# How do you get off the main thread in Swift 6.2?

**Jacob Bartlett** (@jacobtechtavern) on Swift 6.2 concurrency.

## Key Point

Use `Task { @concurrent … }` — the slightly less sexy younger brother of the raw `@concurrent` function attribute.

## Behavior

It creates a new unstructured task that explicitly opts out of main actor isolation, scheduling the work on Swift's global concurrent executor.

## Code Example

```swift
func calculateScoreTaskConcurrent(label: String, iterations: Int) async -> Int {
    let task = Task { @concurrent in
        var total = 0
        for value in 0..<iterations {
            total = (total &+ value) % 97_531
        }
        return total
    }
    return await task.value
}
```

The caller continues (or returns) immediately, with the new task body running separately. This is why we need to suspend at the return, awaiting the result.

## Thread Behavior

We sit on thread one prior to the task body and hop off the main actor once inside the `Task` body.

## Link

- [3 Levels of Swift 6.2 Approachable Concurrency](https://blog.jacobstechtavern.com/p/3-levels-of-)
