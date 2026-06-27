---
author: 0xlelouch_
date: 2026-06-09
source: X (Twitter)
url: https://x.com/0xlelouch_/status/2064385653958176907
tags: [Go]
classification: keyword
---

# Go Interview Production Questions

Most Go interviews test syntax. Production experience shows up in the boring questions: How do you stop a goroutine leak? When do you close a channel, and who owns it? What does a stuck HTTP handler look like in pprof? How do you ship timeouts and context cancellation through 4 layers without making a mess?

I like asking about real incidents: a memory climb from an unbounded buffer, a deadlock from copying a mutex, a hot loop because time.After in select allocates, a service that “works” until GOMAXPROCS changes in a container. Good answers mention tracing, budgets, backpressure, and the exact debug tools they used (pprof, trace, expvar, race detector).

## Relevance

Directly aligns with mission directive to harden Go skills for server-side development (alongside Swift/Vapor stack on Norviq and related projects). Focus on real production debugging patterns.