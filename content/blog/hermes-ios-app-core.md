---
title: How Hermes Inspired HermesVault for iOS
summary: The journey of building a native iOS app with Hermes as its core memory layer.
category: iOS Development
date: 2026-05-08
---

# How Hermes Inspired HermesVault for iOS

After building the backend for Hermes, I realized that the biggest friction point was *capture*. To truly use a second brain, you need to be able to feed it information instantly, from anywhere. This realization led to the creation of **HermesVault Client**.

## Native Polish for Private Data
While many self-hosted tools struggle with user experience, I wanted HermesVault to feel like a premium, native iOS app. I chose **SwiftUI** and **SwiftData** to build a modern, responsive interface that integrates deeply with the Apple ecosystem.

## Deep System Integration
One of the most powerful features of HermesVault is its ability to capture data from across the system:
- **Vision-based OCR**: Extract text from screenshots automatically.
- **HealthKit Integration**: Feed sleep and recovery data into the Hermes learning loop.
- **Native Media Capture**: Seamlessly upload photos and voice memos.

## The Hybrid Path
HermesVault follows a hybrid path. It combines the privacy of a self-hosted stack with the polish of a native app. The backend (running Hummingbird 2 and Postgres) provides the intelligence, while the iOS client provides the gateway to my daily life.

By putting Hermes at the core of the iOS experience, I've transformed my phone from a consumption device into a powerful capture tool for my digital second brain.
