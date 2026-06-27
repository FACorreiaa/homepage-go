---
source: manual
ingested_at: 2026-05-08T04:55:00Z
type: note
status: uncompiled
title: Hermes Routing Policy
---

# Hermes Routing Policy

This is the default model-routing policy for Hermes jobs.

## Principles

1. Use the cheapest model that can do the job well.
2. Escalate only when the task is ambiguous, high-stakes, long-horizon, or quality-sensitive.
3. Never route reminders, trivial formatting, or tiny maintenance tasks through frontier models.
4. Use strong premium models for synthesis, planning, and user-facing writing.
5. Use explicit fallback chains so rate limits do not break workflows.

## Global Defaults

- **Default internal worker:** `DeepSeek V4 Flash`
- **Default reasoning worker:** `DeepSeek V4 Pro`
- **Default premium synthesis:** `Claude Opus 4.7`
- **Default ideation / creative strategy:** `GPT-5.4`
- **Default long-horizon / architecture:** `GLM-5.1`

## Exact Model Assignments by Job Type

### 1) Daily reminders, alerts, status pings
Use when the job is mostly formatting, sending, or checking a single condition.

- **Primary:** `gpt-5.4-mini`
- **Fallback 1:** `DeepSeek V4 Flash`
- **Fallback 2:** `Qwen3.5 Plus`

Examples:
- daily reminders
- cron delivery wrappers
- threshold alerts
- healthcheck notifications
- simple summaries with no deep synthesis

### 2) Knowledge-base ingest, extraction, and cleanup
Use when the job is mostly parsing, deduping, or summarizing source material.

- **Primary:** `DeepSeek V4 Flash`
- **Fallback 1:** `Qwen3.6 Plus`
- **Fallback 2:** `DeepSeek V4 Pro`

Examples:
- URL ingest
- article capture
- transcript cleanup
- note normalization
- metadata extraction

### 3) User research / market research / article synthesis
Use when the job needs synthesis across messy inputs and judgment.

- **Primary:** `Claude Opus 4.7`
- **Fallback 1:** `Kimi K2.6`
- **Fallback 2:** `DeepSeek V4 Pro`

Examples:
- research briefings
- competitor analysis
- product discovery
- source comparison
- “what does this mean?” tasks

### 4) Brainstorming / ideation / naming / product strategy
Use when the job needs breadth, taste, and non-obvious alternatives.

- **Primary:** `GPT-5.4`
- **Fallback 1:** `Claude Opus 4.7`
- **Fallback 2:** `Qwen3.6 Plus`

Examples:
- feature ideation
- naming
- product positioning
- content hooks
- strategic options

### 5) Coding, implementation, and normal agentic workflows
Use when the job is a real build task but not yet a huge refactor.

- **Primary:** `DeepSeek V4 Pro`
- **Fallback 1:** `Kimi K2.6`
- **Fallback 2:** `DeepSeek V4 Flash`

Examples:
- feature implementation
- bug fixes across a few files
- terminal-heavy automation
- standard agent workflows
- code review and PR feedback

### 6) Large refactors, architecture, long autonomous runs
Use when the task spans many files, requires deep planning, or runs for hours.

- **Primary:** `Kimi K2.6`
- **Fallback 1:** `GLM-5.1`
- **Fallback 2:** `Claude Opus 4.7`

Examples:
- multi-file refactors
- architecture decisions
- long-horizon autonomous runs
- complex orchestration
- agent-swarm style coordination

### 7) Debugging and terminal-heavy automation
Use when the task is exploratory and changes direction a lot.

- **Primary:** `Qwen3.6 Plus`
- **Fallback 1:** `DeepSeek V4 Pro`
- **Fallback 2:** `DeepSeek V4 Flash`

Examples:
- runtime debugging
- script repairs
- log analysis
- environment troubleshooting
- cron job recovery

### 8) Multimodal / screenshot-to-code / visual reasoning
Use when the task involves images, UI screenshots, or mixed modalities.

- **Primary:** `MiMo-V2-Omni`
- **Fallback 1:** `GPT-5.4`
- **Fallback 2:** `Claude Opus 4.7`

Examples:
- screenshot-to-code
- UI audits
- visual debugging
- mixed image + text reasoning

## Hermes Job Routing Rules

### Internal ops
- Use `DeepSeek V4 Flash` by default.
- Escalate to `DeepSeek V4 Pro` only when the job spans multiple steps or depends on uncertain interpretation.

### User-facing synthesis
- Use `Claude Opus 4.7` or `GPT-5.4`.
- Do not use flash-tier models for polished external writing unless it is truly trivial.

### Research
- Start with `Claude Opus 4.7`.
- If the task is broad or long-horizon, escalate to `Kimi K2.6`.
- If the task is price-sensitive but still analytical, use `DeepSeek V4 Pro`.

### Code work
- Start with `DeepSeek V4 Pro`.
- Escalate to `Kimi K2.6` for refactors or complex orchestration.
- Use `DeepSeek V4 Flash` for tiny fixes and mechanical edits.

### Reminders and simple notifications
- Keep on `gpt-5.4-mini` or `DeepSeek V4 Flash`.
- Never spend premium reasoning on routine delivery.

## Suggested Hermes Defaults

If you want a simple universal setup:

- **Default Hermes model:** `gpt-5.4-mini`
- **Default worker for automation:** `DeepSeek V4 Flash`
- **Default research model:** `Claude Opus 4.7`
- **Default coding model:** `DeepSeek V4 Pro`
- **Default architecture model:** `Kimi K2.6`

## Fallback Discipline

Use this fallback order whenever a job fails because of rate limits or weak output:

1. Primary model
2. Fallback 1
3. Fallback 2
4. Degrade to a smaller model only if the task is mechanical

Do not keep retrying the same model after a 429, 503, or clearly bad-output loop.
