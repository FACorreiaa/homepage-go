---
title: Introducing Norviq: A Portfolio Tracker
summary: Building a personal stock portfolio tracker with Vapor and SwiftUI.
category: Go & Backend
date: 2026-05-08
---

# Introducing Norviq: A Portfolio Tracker

**Norviq** (formerly known as StockPlan) is my take on a modern portfolio tracker built for active investors. It's a full-stack Swift project that pairs a **Vapor** backend with a native **SwiftUI** mobile app.

## Why Norviq?
Most portfolio trackers are either too simple or too complex. Norviq focuses on the *investing workflow*:
- **Due Diligence Drafts**: Capture your thesis, risks, and catalysts as you research.
- **Target Scenarios**: Define base, bear, and bull targets with specific rationales.
- **Cross-Device Sync**: Your notes and positions stay in sync across iOS and macOS via secure JWT authentication.

## The Stack
Building with Swift across the entire stack has massive benefits. Sharing models between the backend and the mobile app ensures type safety and reduces duplication. I've optimized the backend to run on budget VPS instances (like Hetzner's $5/month tiers), maintaining high performance with low resource usage.

## MVP Strategy
My approach with Norviq was "CSV first, Broker API second." By implementing CSV imports early, I was able to validate the data model and deliver value immediately without the complexity of OAuth and rate limits from broker APIs.

Norviq is more than just a tracker; it's a testament to the power of a unified Swift stack for building production-ready products.
