---
title: "I Run 6 AI Agents - Only This One Makes the Other 5 Smarter"
author: gkisokay
date: 2026-05-04
ingested_at: 2026-05-13T16:19:05
source: https://x.com/gkisokay/status/2051275483996909982
topic: AI
tags: [AI-agents, research-agent, vault, evidence, multi-agent, Hermes, compounding-intelligence]
status: uncompiled
---

I run 6 AI agents. Only this one makes the other 5 smarter.

It's the agent most builders skip. It turns the outside world into compounding intelligence. In 3 months, it's logged 8,000+ pieces of evidence across 16 topics, and every other agent in my stack starts smarter as a result.

It's the research agent. The always-on research department inside Hermes or OpenClaw. It watches what is happening, reads my own workspace, tracks what is moving, and writes everything into a local vault that every other agent in my stack can reuse.

It is not a scraper. This isn't a daily summary bot. This part of the system ensures tomorrow's agents don't start with yesterday's knowledge.

The first version of a research agent is usually a feed reader. The better version is a librarian. The best version is an evidence operator.

Skip that distinction, and research bots become hallucination laundries: confident prose with no separation between what was observed, what was claimed, and what was actually verified.

This guide is the template. Point Hermes, OpenClaw, or any agent system at it and build your own.

## Why This Matters

A typical research agent scrapes, summarizes, and posts a digest. Then the next run starts almost from scratch, except for a note.

That looks productive because there is output everywhere. But output is not the same thing as compounding evidence.

A real research agent should be able to answer:
- What do we know now that we did not know before?
- Which claims are strong, and which are interesting but under-evidenced?
- Which sources and topics keep being useful?
- Which signals belong to which agent, and which old beliefs are now stale?

The research agent exists because the outside world is too noisy for the main agent to parse from scratch every time.

The agent system needs a research layer that can absorb noise, preserve evidence, and route only what matters to you.

That is the difference between agents that stay busy and agents that get sharper.

## What The Vault Actually Holds

To make this concrete, here is what the live research vault looks like:
- 2,631 claim records: candidate beliefs extracted, clustered, tracked over time
- 2,694 findings: individual observed signals from docs, GitHub, feeds, X, search
- 2,694 source evidence records: citation trail tying findings back to URLs, excerpts, timestamps
- 13 indexed dossiers: living topic files for AI agents, frontier AI, crypto rails, memory orchestration, robotics
- 18 registered source surfaces the research agent watches
- 8 operational modes in the loop skill
- 6 handoff lanes routing to other agents
- 36 vault tools plus wrapper scripts for validation, compilation, backup, search, health checks

The counts are not the point. The point is separation:
- A finding is not a claim
- A claim is not verified knowledge
- A source record is not a conclusion
- A dossier is not a daily summary
- A weak signal is not a task

## The Current Architecture

Separate roles downstream from the research agent:
- Research agent = evidence collector
- Main = conscious operator
- Subconscious = pattern-noticer
- Coder = builder
- QA = auditor
- Content = publishing operator

When research, judgment, building, and publishing collapse into a single agent, the system treats every interesting signal as an action item. The research agent gathers, weighs, remembers, and routes. Then the rest decides what to do.

## The Vault Structure

Minimum recommended shape: research-vault with context/, config/, dossiers/, knowledge/ (claims.jsonl, findings.jsonl, sources.jsonl), raw/, queue/, notes/, wiki/, indexes/, health/, ops/

Key distinctions most setups skip:
- decisions/ — records what was decided, by whom, on what evidence
- runs/ — each refresh leaves a receipt
- raw/ — unprocessed capture layer, kept separate from knowledge/

## The Research Loop

Packaged as a skill with modes: BOOTSTRAP, REFRESH, DAILY_SUMMARY, SUBCONSCIOUS_BRIEF, MIDDAY_FOCUS, BACKUP, RESTORE, RECOVER

Each run updates: interest profile, source plan, topic dossiers, question backlog, operator brief, daily summary, source balance, collector health, knowledge ledgers, verification queue, wiki pages, run receipt.

## The Source Plan

Prefer sources that change decisions. If the agent collects everything, it drowns. If it only collects what's easy, it overfits to the most convenient API.

The source plan is where the research agent learns taste.

## The Interest Profile

Rebuilt from shared state, durable notes, thesis files, posting behavior, shipped builds, repeated questions, and prior outputs. Not trying to be a neutral news service — designed to be useful to a single operator.

## Dossiers, Ledgers, and Verification

Recurring topics become dossiers. Machine-readable ledgers in knowledge/ (claims.jsonl, findings.jsonl, sources.jsonl). Verification queue for weak/under-sourced claims.

## The Wiki Layer

Pipeline: collectors -> raw/*.json -> knowledge/*.jsonl -> dossiers/*.md -> wiki/concepts/*.md -> indexes + health

## Operator Surfaces and Handoffs

Writes operator brief answering: What changed? What deserves attention? What is blocked by weak evidence? What should downstream agents do?

Handoff lanes: buildroom, content, monetize, subconscious, verify, watch.

## Quality Gates

Tracks source balance: did this rely too much on social media? Did primary sources show up? Health checks for broken links, missing frontmatter, gaps in source trails, orphan candidates, stale reviews.

## Runtime and Models

Intentionally modest stack. Most jobs on cheaper models (MiniMax M2.7). Use cheaper models for routine refresh, stronger models for synthesis and judgment.

## The Guardrails

Research agent explicitly not allowed to: make trading decisions, publish public posts, make purchases, commit to partnerships, touch secrets, turn weak signals into approved tasks, pretend stale data is fresh.

## Minimal Pseudocode

on schedule -> load SOUL -> read shared state -> read prior vault -> rebuild interest profile -> collect sources -> score quality -> extract findings -> update ledgers -> update dossiers -> rebuild backlog -> write operator brief -> compile wiki -> run health check -> route implications -> validate.

## The Real Lesson

A real research agent knows what you care about, what it has seen before, which sources are strong, which collectors are degraded, and which claims are not yet ready. Most importantly, it knows when to say "this is interesting, but not safe to build on yet."

## If You Want To Build One

1. Give it a vault and an identity
2. Then a source plan and knowledge ledgers
3. Then a verification queue and handoff lanes
4. Then health checks and a schedule
5. Then let other agents consume its outputs


## Relevance to Fernando's Projects

This article is highly relevant to Fernando's Hermes Agent setup and multi-agent architecture (FACorreia vault, cron jobs, content ingestion pipeline).

**Key takeaways for the Hermes stack:**
- **Research agent pattern**: The concept of a dedicated research agent that feeds evidence to other agents aligns with Fernando's existing content ingestion crons (X/Twitter ingestor, URL ingestor, YouTube extractors). Could be worth considering a research agent profile for Hermes.
- **Vault separation**: The distinction between raw captures, findings, claims, and verified knowledge could improve the FACorreia vault's classification system. Currently content goes to raw/<category>/ with status:uncompiled — adding a claims.knowledge layer would be a natural evolution.
- **Interest profile**: The idea of an interest-profile.json that the research agent reads to personalize collection could make the content ingestion more targeted instead of ingesting all URLs equally.
- **Handoff lanes**: Fernando already has multi-agent patterns (main agent, cron workers) — explicit handoff lanes between them could clarify responsibilities.
- **Source balance tracking**: Quality gates checking if collection is over-dependent on low-trust sources or degraded collectors would improve the ingestion pipeline.
- **Verification queue**: A place for "interesting but not ready to build on" content that doesn't get promoted to tasks immediately.

**Not action-critical right now** but a strong reference for future Hermes architecture decisions around the vault and multi-agent separation.
