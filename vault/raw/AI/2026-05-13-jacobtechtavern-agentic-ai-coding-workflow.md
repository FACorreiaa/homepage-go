---
title: "Agentic AI Coding Workflow — Multi-Agent with Stacked Diffs"
url: https://x.com/jacobtechtavern/status/2054639824980635985
author: Jacob Bartlett (@jacobtechtavern)
date: "2026-05-13"
ingested_at: "2026-05-13T19:14:29.228847"
tags: [agentic-ai, codex, workflow, multi-agent, code-review, stacked-diffs, ios]
topic: AI
status: uncompiled
---

# Agentic AI Coding Workflow — Multi-Agent Pattern

**Author:** Jacob Bartlett (@jacobtechtavern)  
**Source:** X/Twitter  
**Date:** 2026-05-13

## Original Content

The developer workflow:

1. Plan is in place with stacks and per-stack verification steps
2. Codex codes for hours while developer works on other tasks (multiple separate repo checkouts)
3. Result is a testable, mostly-there feature in stacked diffs
4. Code review in chosen client with inline comments across the stack
5. Tell the agent: "Go through the PR comments, one stacked diff at a time, addressing the issues flagged. Make changes on the appropriate point in the stack, and reply to each PR comment to confirm."

Key insight: The agent correctly oriented itself across a complex multi-hour stack of diffs and made changes at the appropriate points. Codex is cited as the only agent trusted for this.

Multi-agent breakthrough pattern: 1 Writer + N Readers on the same checkout
- Reader #1: Fetches PR comments, gathers context, proposes fixes
- Reader #2: Reviews code in each stack independently, outputs improvements/simplifications  
- Writer: Reads the markdown outputs and implements changes
- Optional idle agent at top of PR stack for manual testing (logs, test data, output review)

Practical limits: 1 Writer + 2 Readers is the mental processing limit.

Full article: https://blog.jacobstechtavern.com/p/ios-at-an-ai-unicorn

## Summary

Detailed agentic AI workflow that combines:
- Stacked diffs for clean, reviewable PR structure
- Multi-agent parallelization (1 Writer + 2-3 Readers)
- Asynchronous work (developer does other tasks while agent codes)
- PR comment-driven iteration (agent addresses review comments automatically)
- Separate repo checkouts to avoid agent conflicts

## Metrics

- Views: 15
- Full Article: https://blog.jacobstechtavern.com/p/ios-at-an-ai-unicorn

## Relevance to FACorreia Projects

**HIGH RELEVANCE** — directly applicable to Norviq, Loci, and ObsidianBrain iOS development workflow. The multi-agent pattern (Writer + Readers) overlaps with the Hermes multi-Agent skill architecture (CTO, PM, Dev, QA, Ops). Key takeaways:
- Stacked diffs enable granular code review
- Multi-agent parallelization requires careful coordination to avoid conflicts
- 1 Writer + 2 Readers is the practical limit for effective coordination
- Codex is trusted for complex, multi-hour multi-file changes; other agents less reliable

This is an iOS developer working at scale with AI agents — the exact workflow pattern you should consider for your iOS projects.
