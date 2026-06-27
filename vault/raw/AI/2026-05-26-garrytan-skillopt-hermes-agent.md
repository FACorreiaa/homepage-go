---
author: garrytan
date: 2026-05-26
source: X (Twitter)
url: https://x.com/garrytan/status/2059144022778896392
metrics: N/A
tags: [AI]
classification: llm
---

# Garry Tan on SKILLOPT — implementable trivially in Hermes Agent

Garry Tan (Y Combinator CEO) shares a research diagram of **SKILLOPT**, a text-space optimization framework for AI agents. Instead of fine-tuning model weights, SKILLOPT optimizes a "Skill Document" (text prompt/instructions) through bounded edits — treating the prompt space like a loss landscape and walking downhill toward lower error rates.

The key insight: the agent analyzes its own failures, proposes bounded edits to its own instruction manual, validates them, and converges on expert performance through stable optimization — all without any model weight changes.

## Relevance

Directly relevant to Hermes Agent. Garry explicitly says this can be "implemented in OpenClaw/Hermes Agent trivially (use skillify from GBrain with a link to this tweet)". This suggests a future Hermes feature where agents self-optimize their skill documents through iterative self-improvement — a natural evolution of the current skill system.

## Key Concepts

- **Text-Space Optimization**: Treats skill documents as parameters to be optimized
- **Bounded Skill Edits**: Small, controlled changes with validation gates (reject regressions)
- **Cross-model generalization**: Skills transfer between models because they're just text
- **Low-cost**: Much cheaper than weight fine-tuning
- **Skillify + GBrain**: The path to implement this in Hermes today

## Notes

This is huge for the Hermes ecosystem. If skill documents become self-optimizing, the agent continuously improves without human prompt engineering. Worth watching this research and checking if GBrain's skillify is already usable.