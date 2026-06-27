---
title: "Thread by @jacobtechtavern"
source: "https://x.com/jacobtechtavern/status/2043587500053000620?s=48"
author:
  - "[[@jacobtechtavern]]"
published: 2026-04-13
created: 2026-04-13
description: "Tasks in Swift Concurrency are very similar to NSThreads! Both help manage processes, but live at different levels of abstraction. NSThread"
tags:
  - "clippings"
---
**Jacob Bartlett** @jacobtechtavern [2026-04-13](https://x.com/jacobtechtavern/status/2043587500053000620)

Tasks in Swift Concurrency are very similar to NSThreads! Both help manage processes, but live at different levels of abstraction.

NSThread is an operating-system-level entity managed by the kernel. Each thread is a 512kB behemoth with its own call stack, making it expensive to context-switch between threads. The OS has to pause a thread, store its call stack somewhere safe, and resume it later. This overhead adds up in concurrent applications.

Tasks are Swift Concurrency’s lightweight “unit of execution”. They are managed by the Swift Runtime and are scheduled on the Cooperative Thread Pool, which itself aims to maintain 1 thread per core. Asynchronous work on tasks store context in continuations, a.k.a. async frames, and they are designed to yield execution to avoid deadlock and blocking.

Read my full blog post here: https://blog.jacobstechtavern.com/p/tasks-vs-threads…

![Image](https://pbs.twimg.com/media/HFxIAG1WsAAi4aT?format=jpg&name=large) ![Image](https://pbs.twimg.com/media/HFxH_-qW4AAIvlb?format=jpg&name=large) ![Image](https://pbs.twimg.com/media/HFxH_9qXgAACMpj?format=jpg&name=large)