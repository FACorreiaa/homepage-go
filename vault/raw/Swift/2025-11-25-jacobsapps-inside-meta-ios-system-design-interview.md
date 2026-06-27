---
title: "Inside Meta’s “iOS System Design” Interview"
url: https://blog.jacobstechtavern.com/p/ios-system-design
author: jacobsapps
date: 2025-11-25
ingested_at: 2026-05-13T17:45:42Z
tags:
  - swift
  - swiftui
  - system-design
  - interviews
topic: Swift
status: uncompiled
---

[![Image 1](https://substackcdn.com/image/fetch/$s_!wWEK!,w_2400,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F5a7d60be-a18a-4f1e-853b-ac700b921064_1680x1200.png)](https://substackcdn.com/image/fetch/$s_!wWEK!,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F5a7d60be-a18a-4f1e-853b-ac700b921064_1680x1200.png)

If you want to maximise your salary as an iOS developer, the second best thing you can do is work for Facebook.

> _The best thing you can do, naturally, is read Jacob’s Tech Tavern._

Salaries at big American tech companies are hard to believe if you don’t actually know anyone that works for them. Total comp packages for a Meta E5 (senior dev) in London range from £199k to over £400k.

This fact alone makes their interview process a worthwhile field of study.

Roles are competitive, but achievable, with intentional practice. I’ve seen multiple friends and colleagues bang their head against the interviews, year after year, until they finally landed.

Big tech interview loops contain 2 major technical components:

*   Algorithmic interviews (a.k.a. “LeetCode™”)

*   The System Design Interview™

You probably don’t need me to explain LeetCode, so I won’t. It’s a grind.

What’s more interesting is the system design component. For most Software Engineering roles, this typically means explaining how you would build the backend infrastructure for a service like Pastebin, Instagram, or Uber.

For iOS roles at Meta, however, they administer an **“iOS-specific system design interview”**. This is really very opaque and scary if you don’t have insider information, so I wrote this piece.

I personally went through Meta’s interview loop in 2024, and accidentally passed their first round _(but ended up dropping out of the process as I’d just landed a cool startup job)_.

Even better, I got an E5 Meta interviewer drunk. I got the inside-track for how this interview works, and how you are rated by Meta.

[![Image 2](https://substackcdn.com/image/fetch/$s_!q5Tl!,w_1456,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2Fce1d5868-9717-40ae-8204-3e28ac626aad_996x1152.png)](https://substackcdn.com/image/fetch/$s_!q5Tl!,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2Fce1d5868-9717-40ae-8204-3e28ac626aad_996x1152.png)

He shall remain nameless to protect the innocent

What’s really interesting is that Meta levels candidates entirely on the behavioural and system design interviews. That means the intern and the staff engineer get the same hiring bar on the LeetCode component.

Nailing the iOS System Design interview, therefore, makes all the difference.

This interview sounds very amorphous until you understand what it actually is.

My man on the inside summarised it in a single sentence that made me realise it was not so complicated after all:
