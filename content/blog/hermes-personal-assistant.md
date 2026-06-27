---
title: Building Hermes into My Personal Assistant
summary: How I started building Hermes as a self-hosted AI memory layer and personal assistant.
category: Building in Public
date: 2026-05-08
---

# Building Hermes into My Personal Assistant

The journey of turning an AI into a true personal assistant begins with memory. Most AI interactions today are ephemeral—you talk, it responds, and then it forgets. To build a true assistant, I needed a way for the AI to remember, learn, and grow with me.

Enter **Hermes**.

Hermes started as a simple private gateway hosted on a VPS, acting as a coordinator for multiple platforms. But it quickly evolved into something more: a **self-improving memory layer**.

## The Vision
The vision for Hermes is to turn any input—links, screenshots, photos, notes, HealthKit data, or Maps locations—into a "living" second brain. Instead of just storing data, Hermes automatically compiles it into a clean wiki. It learns habits, detects patterns, and surfaces insights when they are most meaningful.

## How it Works
Using a `kb-compile` engine, Hermes takes raw Markdown vault inputs and processes them into a structured knowledge base. With semantic search powered by `pgvector` and `PostGIS`, I can query my own life's data with natural language.

It's **Obsidian with persistent AI memory**—fully self-hosted, private, and under my control.

## Why Self-Hosted?
Privacy is not an afterthought; it's the foundation. By hosting Hermes on my own infrastructure, I ensure that my personal assistant truly belongs to me. No third-party data mining, no platform lock-in—just a private, intelligent partner.
