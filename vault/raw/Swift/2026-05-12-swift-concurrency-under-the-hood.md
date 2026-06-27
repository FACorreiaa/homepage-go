---
title: "Swift Concurrency Under the Hood"
url: ""
author: "Fernando"
date: 2026-05-12
ingested_at: 2026-05-12T13:05:00Z
tags: [swift, concurrency, vapor, ios, internals]
topic: Swift
status: uncompiled
---

# Swift Concurrency Under the Hood

## The Thread Pool (Cooperative Thread Pool)

When you await something in Swift concurrency, you're not spawning a thread. You're yielding execution back to Swift's cooperative thread pool. This pool has a fixed number of threads:

- iOS/macOS: Typically one thread per core, plus a small overflow
- Linux (Vapor/Hummingbird): Same model, but initialized differently. On Linux the pool size defaults to max(cpuCount, 2) unless overridden with `SWIFT_CONCURRENCY_WORKER_THREADS`

The key word is **cooperative**. There's no preemption. If your code runs on a pool thread without awaiting (a CPU-bound loop, a heavy computation), that thread is stuck. No other task on that thread can make progress. This is the #1 concurrency bug in Swift: **blocking the pool**.

Here's what actually happens when you write:

```swift
func fetchUser(id: String) async throws -> User {
    let (data, response) = try await URLSession.shared
        .data(for: URLRequest(url: url))
    return try decoder.decode(User.self, from: data)
}
```

1. The caller hits `await` and gets suspended — its state (registers, locals, the continuation point) is saved in a heap-allocated async frame
2. The pool thread that was running the caller pops back to the pool dispatcher and picks up the next ready task
3. When the network request completes (handled by libDispatch / libuv on the I/O side, NOT a pool thread), the completion callback enqueues the resumed task back onto the pool
4. A pool thread picks it up, restores the async frame, and continues from after the `await`

On iOS, this pool sits on top of GCD. Under the hood, Swift concurrency is essentially a sophisticated wrapper around dispatch queues with structured concurrency guarantees. The async functions you call (URLSession.data, Actor methods) eventually submit continuations to GCD.

On Linux, there's no GCD. SwiftNIO provides the event loop layer instead. The cooperative pool is built on top of NIO's event loops via `NIOSingletons.posixThreadPool`. Vapor's request handler runs on one of these pool threads, and when you await a database query, the thread goes back to the pool while NIO handles the actual I/O on its own event loop.

## The Actor Reentrancy Problem

This catches everyone:

```swift
actor Cache {
    private var items: [String: Data] = [:]

    func fetch(_ key: String) async throws -> Data {
        if let cached = items[key] {
            return cached
        }
        // This await means ANYONE can reenter this actor
        // while we're suspended here. Another caller might
        // have already populated items[key] by the time
        // we resume.
        let data = try await fetchFromNetwork(key)
        items[key] = data
        return data
    }
}
```

When you `await` inside an actor method, the actor's isolation is released during suspension. Other tasks can enter the actor while you're suspended. This is actor reentrancy. It's not a bug — it's by design to prevent deadlocks — but if you don't expect it, you get duplicate work or inconsistent state.

The fix: check again after the await (double-check pattern).

```swift
if let cached = items[key] { return cached } // double-check after await
```

## Under the Hood: iOS vs Server Differences

### iOS (GCD-backed)

```
User Task → Swift Concurrency → GCD (libdispatch) → pthreads
                                       ↓
                                  CoreFoundation
                                  CFNetwork (for URLSession)
                                  RunLoop integration
```

On iOS, Swift concurrency is deeply integrated with the run loop. MainActor is bound to the main dispatch queue, which is bound to the main run loop. UI updates on MainActor are processed through the run loop's source/event system.

Key detail: iOS does NOT give you unlimited threads. The pool is small (4-8 threads on modern iPhones). If you saturate it, everything stalls. This is why `Task.detached` is dangerous on iOS — it creates a thread outside the pool, and if you overuse it, you're fighting the system's own thread management.

### Server/Vapor (NIO-backed)

```
HTTP Request → EventLoop → Swift Concurrency Pool → Posix ThreadPool → epoll
                         ↓
                  NIOChannel (socket)
                  NIOByteBuffer pipeline
```

On Linux, everything goes through epoll. The event loop watches sockets, NIO handles the byte parsing, and when you await in Vapor, the request handler suspends on the pool while NIO manages the actual socket I/O.

Key difference: the server's bottleneck is I/O, not UI. You can handle thousands of concurrent HTTP requests because most of them are suspended on I/O at any given moment. The 4-8 pool threads are sufficient because they're constantly cycling through ready tasks while waiting threads are blocked in epoll.

## Daemon / Long-Running Concurrency Patterns

### Pattern 1: TaskGroup as a Bounded Worker Pool

```swift
func processQueue() async throws {
    try await withThrowingTaskGroup(of: Void.self) { group in
        while let job = await queue.dequeue() {
            if group.taskCount >= maxConcurrent {
                try await group.next() // wait for one to finish
            }
            group.addTask {
                try await process(job)
            }
        }
        try await group.waitForAll()
    }
}
```

This is how Vapor itself handles concurrent requests under the hood — bounded parallelism. Without the `taskCount >= maxConcurrent` guard, you'd spawn unlimited tasks and OOM.

### Pattern 2: AsyncStream for Producer/Consumer

```swift
let (stream, continuation) = AsyncStream<Event>.makeStream(
    bufferingPolicy: .bufferingNewest(1000)
)

// Producer (could be a WebSocket, file watch, timer, etc.)
Task {
    for try await line in fileHandle.bytes.lines {
        continuation.yield(.event(line))
    }
    continuation.finish()
}

// Consumer
Task {
    for await event in stream {
        await processor.handle(event)
    }
}
```

AsyncStream is Swift's replacement for Combine subjects or NotificationCenter observers. It has backpressure: if the consumer is slow, the buffer fills and depending on policy, new events are dropped or the producer suspends.

### Pattern 3: Actor-Based State Machine

```swift
actor ServiceState {
    private var isRunning = false
    private var task: Task<Void, Never>?

    func startWork() {
        guard !isRunning else { return }
        isRunning = true

        task = Task {
            while await isRunning {
                await performOneCycle()
                try? await Task.sleep(for: .seconds(cycleInterval))
            }
        }
    }

    func stopWork() {
        isRunning = false
        task?.cancel()
    }
}
```

This is how you build a daemon that can be cleanly stopped. The actor owns the Task, the Task checks the actor's state each cycle, and cancellation cooperates properly because `Task.sleep` is cancellation-aware (resumes immediately with CancellationError).

## App-Side Patterns: Making iOS Faster and More Responsive

### Pattern 1: Prefetching on Appear

```swift
struct ContentView: View {
    @State private var data: [Item] = []

    var body: some View {
        List(data) { item in
            ItemRow(item: item)
        }
        .task {
            // .task runs on MainActor, but you await away
            // immediately. UI stays responsive.
            if let result = await try? repository.fetchItems() {
                data = result
            }
        }
    }
}
```

`.task` is better than `.onAppear { Task { ... } }` because it's lifecycle-aware: if the view disappears before the task completes, the task is automatically cancelled. No orphaned network requests.

### Pattern 2: Concurrent Fetch with Structured Concurrency

```swift
func loadDashboard() async throws -> DashboardData {
    // Run 3 independent fetches concurrently
    async let userData = try userLoader.fetch()
    async let notifications = try notificationLoader.fetch()
    async let feed = try feedLoader.fetch()

    // All three run in parallel. The function returns
    // when the SLOWEST one completes.
    return DashboardData(
        user: try await userData,
        notifications: try await notifications,
        feed: try await feed
    )
}
```

`async let` is structured — if one throws, all siblings are cancelled. This matters for correctness: you don't want a successful feed fetch to persist while userData failed and the whole screen should error out.

### Pattern 3: MainActor Isolation for UI

```swift
@MainActor
class ViewModel: ObservableObject {
    @Published var items: [Item] = []
    @Published var state: LoadingState = .loading

    func load() async {
        state = .loading
        // We're on MainActor here
        do {
            let items = try await repository.fetchItems()
            // After the await, we're back on MainActor
            // automatically — no DispatchQueue.main.async needed
            self.items = items
            state = .loaded
        } catch {
            state = .error(error)
        }
    }
}
```

`@MainActor` on the class guarantees all methods run on the main thread. When you await in a @MainActor method, Swift handles the thread hop back to main automatically. Zero dispatch code needed.

### Pattern 4: AsyncImage with Sendable Closures

```swift
AsyncImage(url: url) { phase in
    switch phase {
    case .success(let image):
        image.resizable()
    case .failure(let error):
        ErrorView(error)
    default:
        ProgressView()
    }
}
```

SwiftUI's AsyncImage uses Swift concurrency under the hood. The image download runs on a background pool thread, and the phase updates happen on MainActor through SwiftUI's internal transaction system.

## Server-Side Patterns: APIs Handling More Work

### Pattern 1: Connection Pooling + Concurrency Limiting

```swift
// Vapor route handler
func register(routes: RoutesBuilder) {
    routes.group("api") { api in
        api.get("items") { req async throws -> [Item] in
            // This handler runs on the Swift concurrency pool.
            // Multiple HTTP requests each get their own task.
            // The limiting factor is your database connection pool,
            // NOT the number of pool threads.
            let items = try await db.query(Item.self).all()
            return items
        }
    }
}
```

Vapor's NIO-based HTTP server can handle thousands of concurrent connections because the I/O is event-driven. Each request handler is a lightweight async task that suspends on DB queries. The connection pool (e.g., SQLKit's pool) is what limits concurrency, not the thread pool.

### Pattern 2: Background Tasks for Heavy Work

```swift
// Offload CPU-heavy work to a detached task
// so it doesn't block the cooperative pool
func generateReport(req: Request) async throws -> HTTPStatus {
    let task = Task.detached(priority: .utility) {
        // Runs on a thread OUTSIDE the pool
        // Won't block other request handlers
        try await HeavyProcessor().generate()
    }

    // Return immediately, let the detached task run
    return .accepted
}
```

Task.detached breaks free of structured concurrency and the pool thread constraints. Use it for CPU-bound work (image processing, report generation, ML inference) that would otherwise starve the pool of threads serving other requests. Caveat: you need to handle cancellation manually, and you can't inherit actor isolation or task-local values from the parent.

### Pattern 3: Batch Processing with TaskGroup

```swift
func processBatch(items: [ItemID]) async throws -> [Result<Item, Error>] {
    try await withThrowingTaskGroup(of: (Int, Result<Item, Error>).self) { group in
        for (index, id) in items.enumerated() {
            group.addTask { (index, Result { try await fetchAndProcess(id) }) }
        }

        var results = Array<Item?>(repeating: nil, count: items.count)
        for try await (index, result) in group {
            results[index] = result
        }
        return results.compactMap { $0 }
    }
}
```

Preserves ordering of results while processing concurrently. The index tuple ensures results go back in the right position even though tasks complete out of order.

### Pattern 4: Continuations for Legacy Code

When you have callback-based APIs that haven't been updated to async:

```swift
func legacyCallbackAPI(completion: @escaping (Result<Data, Error>) -> Void) {
    // ... old callback code
}

func wrapped() async throws -> Data {
    try await withCheckedThrowingContinuation { continuation in
        legacyCallbackAPI { result in
            switch result {
            case .success(let data):
                continuation.resume(returning: data)
            case .failure(let error):
                continuation.resume(throwing: error)
            }
        }
    }
}
```

`withCheckedThrowingContinuation` is the bridge. The "checked" variant validates at runtime that you resume exactly once — if you don't, the program crashes. This catches bugs immediately rather than hanging silently. For performance-critical paths, `withUnsafeThrowingContinuation` skips the check.

### Pattern 5: Async Result Support (SE-0530)

```swift
// Before SE-0530, you'd write:
let result: Result<User, Error>
do {
    result = .success(try await fetchUser())
} catch {
    result = .failure(error)
}

// With SE-0530:
let result = await Result { try await fetchUser() }
```

Small quality-of-life change, but it eliminates duplication in error-handling paths across a large codebase. Vapor route handlers that return Result types benefit directly.

## The Big Picture

| Aspect | iOS | Server (Vapor) |
|--------|-----|----------------|
| Thread pool | 4-8 threads, GCD-backed | max(cpuCount, 2), posixThreadPool |
| Primary bottleneck | MainActor contention, memory | I/O, DB connection pool |
| Task.detached | Dangerous (thread explosion) | Good for CPU-heavy workloads |
| Actor isolation | Critical for UI thread safety | Critical for shared state |
| Key pattern | async let for parallel prefetch | TaskGroup for bounded concurrency |
| I/O model | GCD + CFNetwork run loops | NIO + epoll event loop |
| Memory pressure | Aggressive — kill at ~500MB | Not bounded by default |

The common trap on both platforms is the same: blocking the cooperative pool with synchronous work. If you call any synchronous function from an async context that takes more than a few milliseconds (JSON encoding a large object, database access via a non-async driver, file I/O), you've stolen a pool thread from every other task. The fix is always the same: `Task.detached` for CPU work, or an async-compatible library for I/O.
