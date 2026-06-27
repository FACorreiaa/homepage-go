---
source: "https://x.com/azamsharp/status/2055664932335538593"
author: "Mohammad Azam (@azamsharp)"
published: 2026-05-16
topic: Swift
tags:
  - swift
  - ios
  - storekit2
  - app-store-server-notifications
  - subscriptions
  - in-app-purchase
  - expressjs
  - tutorial
metrics:
  comments: 1
  reposts: 2
  likes: 8
  views: 391
---

# StoreKit 2 — App Store Server Notifications (Coming Soon)

Azam Sharp teasing a comprehensive tutorial on App Store Server Notifications V2.

Article preview: `2026-05-16.storekit2-app-store-server-notifications.md`

## What it covers
1. **Lightweight ExpressJS (Node.js) backend** endpoint to receive Apple notifications
2. **Decoding signed payloads (JWS)** — verifying authenticity of App Store data
3. **Transaction & renewal data management** — update database, monitor refunds
4. **Alerting pipeline** — trigger Slack notifications, analytics, etc.

## Relevance
Directly applicable to **Norviq** (subscription support for a stock plan app), **Loci** (itinerary app with potential premium tiers), and **ObsidianBrain** iOS app — all would benefit from StoreKit 2 + App Store Server Notifications for subscription lifecycle management.

Azam notes StoreKit 2 is a "pleasant" alternative to third-party services but doesn't handle events outside the app (renewals, refunds, billing issues) out of the box — server notifications fill that gap.
