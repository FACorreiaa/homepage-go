---
author: 0xlelouch_
date: 2026-06-24
source: X (Twitter)
url: https://x.com/0xlelouch_/status/2069795951695290634
metrics: N/A
tags: [System Design, Messaging, Tech, Interview]
classification: keyword
---

# WhatsApp Message Delivery System Design

Detailed interview answer covering requirements (1:1/group chats, near real-time, at least-once with idempotency, per-conversation ordering, offline/multi-device support), data model (Message table keyed by thread+seq, delivery state per device), architecture (regional API gateway → chat service for seq assignment + durable log, fanout to per-device inboxes, WebSocket push for online, polling for offline), scaling (threadId sharding, LSM KV stores like Cassandra/Scylla for inboxes), tradeoffs (at-least-once over exactly-once, per-thread ordering, pointers over full bodies in inboxes), and failure handling (retries with clientMsgId, leader election for seq allocators, backpressure on hot groups).

## Relevance
Strong backend/systems design content relevant to messaging infrastructure (e.g., Norviq chat features, Hermes agent coordination, or any real-time delivery system). Educational value for iOS/server architecture discussions.

## Notes
- E2E encryption out of scope (focus on routing/metadata).
- Push notifications best-effort; inbox is source of truth.
- Hot groups (50k members) require throttling and parallel fanout.
- Read receipts/last seen are separate signals.