---
title: The Ultimate Hermes Guide from One Agent to a 4-Profile Team
author: nyk_builderz
date: 2026-04-15
ingested_at: 2026-05-13T16:20:56.342003Z
source: https://x.com/nyk_builderz/status/2044472463279710344
url: https://x.com/nyk_builderz/status/2044472463279710344
topic: Hermes
tags: [hermes, multi-agent, profiles, operator, orchestration, agent-architecture]
status: uncompiled
---

# The Ultimate Hermes Guide - from one agent to a 4-profile team that still feels coherent on day 30

**By Nyk (@nyk_builderz)**
*Posted: Apr 15, 2026 | 158.8K views | 2.1K bookmarks*

The author ran one Hermes agent as researcher, writer, coder, and orchestrator for 14 days on a single claude-sonnet-4.6 profile before everything blurred into the same voice. Most operators blame the prompt when this happens, but it is not prompting and not the model — it is one agent carrying five roles on shared memory. The Hermes primitive that actually fixes it: **isolated profiles**.

Credit to @neoaiforecast for the team build: orchestrator + alan + mira + turing, four roles, clean handoffs, 1,317 bookmarks in a day. This guide picks up at day two.

---

## The mental model — roles, not personas

The wrong mental model: I need one genius AI that does everything.
The better mental model: I need a small team with distinct roles, clear handoffs, and less context pollution.

Hermes profiles isolate seven pieces of state: configuration, sessions, memory, skills, personality, cron state, gateway state. Specialization becomes durable only when the state remains separated.

## The 4-role team

* **Hermes** — orchestrator: plans, decomposes, routes, synthesizes. The traffic controller, not the bottleneck.
* **Alan** — research specialist: source-first, skeptical, uncertainty-aware, protects from hallucinated confidence.
* **Mira** — narrative architect: clarity, structure, audience awareness, turns validated material into communication.
* **Turing** — builder and debugger: implementation, logs, diffs, reproducibility, cares about tests, not narrative polish.

## The 7-step build

**Step 1** — Start from a working Hermes. Make sure base installation is healthy before cloning.

**Step 2** — Create specialist profiles:
```bash
hermes profile create alan --clone
hermes profile create mira --clone
hermes profile create turing --clone
```
--clone copies config.yaml, .env, and SOUL.md from working base.

**Step 3** — Verify on disk:
```bash
hermes profile list
ls ~/.hermes/profiles/
```

**Step 4** — Write a real SOUL.md per role:
* Hermes (orchestrator): planning, delegation, synthesis, final QA. Structured and decisive.
* Alan (research): evidence, verification, depth, uncertainty. Source-first and skeptical.
* Mira (writer): clarity, structure, audience. Clear and audience-aware.
* Turing (engineer): implementation, debugging, tests, reproducibility. Precise and test-oriented.

**Step 5** — Keep project context in AGENTS.md, not SOUL.md. Identity stays stable; project context rotates.

**Step 6** — Add a team reference file documenting roster and handoffs.

**Step 7** — Run profiles separately:
```bash
hermes -p alan
hermes -p mira
hermes -p turing
```

## The operator layer

### Handoff contracts between profiles

Store at `~/.hermes/team/handoffs/<from>-to-<to>.md` with four fields:
* **Input shape:** what the receiving profile expects
* **Output shape:** what the receiving profile will return
* **Failure action:** what happens if input is malformed
* **Verification gate:** one assertion that must be true before handoff completes

### Memory-KPI per profile

Run weekly memory audit per profile. Watch stale notes — once it crosses 15% of total notes, schedule a brain-resolve pass.

### Policy gates per role

* **Alan (research):** risk class safe. Read web/repo, write to research/ only. No shell commands.
* **Mira (writer):** risk class safe. Read research outputs, write to drafts/ only. No code execution.
* **Turing (engineer):** risk class review. Read repo, run sandboxed tests, write to feature branch. Main requires orchestrator approval.
* **Hermes (orchestrator):** risk class critical. Only profile that can approve merges, widen permissions, or trigger paid API calls above budget ceiling.

### Gateway messaging for remote supervision

Wire each profile to its own messaging channel:
* Alan → #research
* Mira → #writing
* Turing → #engineering
* Hermes → #ops with human approval on critical actions

## Day 30 failure modes

**Failure 1 — Profile drift:** SOUL.md edits accumulate, roles blur into each other. *Patch:* diff each SOUL.md weekly against day-one version.

**Failure 2 — Handoff rot:** Contract exists but nobody enforces it. Boundaries dissolve. *Patch:* Wire handoff files into harness — if input doesn't match declared shape, fail and require human review.

**Failure 3 — SOUL.md bloat:** Each role grows edge cases. Within a month each SOUL.md is 2kb of special cases. *Patch:* Cap SOUL.md at 400 words. Anything beyond goes into AGENTS.md or per-domain reference.

**Failure 4 — Cron collision:** Profiles run cron jobs and collide. *Patch:* one shared `~/hermes/team/cron.md` listing every scheduled task across profiles with exact time, duration, and dependency.

## Team reference template

```markdown
# ~/hermes/team-agents.md

## roster
* hermes (orchestrator): plans, routes, approves, synthesizes
* alan (research): source-first, skeptical, uncertainty-tagged
* mira (writer): clarity, structure, audience-aware
* turing (engineer): implementation, tests, reproducibility

## when to use which profile
- starting a new project -> hermes (scopes and decomposes)
- validating a claim -> alan (source check, uncertainty tag)
- drafting anything external-facing -> mira (audience-first)
- writing or debugging code -> turing (test-first)

## handoff rules
- alan -> mira: ranked claims with source urls, no raw transcripts
- mira -> hermes: drafted section + change log, not a finished article
- turing -> hermes: feature branch + passing tests + diff summary, not a PR
- hermes -> any: scoped task with acceptance criteria and failure action

## good output per profile
- alan: every claim has a source url and a confidence tag
- mira: every section has a named audience and a clear thesis
- turing: every change has a passing test and a reproducible diff
- hermes: every synthesis names the contributors and the open questions

## policy ceilings
- alan: read-only outside research/
- mira: read research/, write drafts/
- turing: read repo, write feature branch, run sandboxed tests
- hermes: only profile allowed to approve merges, widen permissions, or access a budget ceiling

## cron schedule
(edit weekly; stagger to avoid 3am collisions)
- mon 6am - alan: weekly research digest
- tue 6am - mira: draft refresh from alan's digest
- wed 6am - turing: test sweep + flaky test report
- thu 6am - hermes: weekly synthesis + handoff audit
```

## The agent extraction layer

* **Objective:** run a 4-profile Hermes team that stays coherent past day 30
* **Inputs:** working Hermes base, profile CLI, SOUL.md + AGENTS.md split, handoff contracts, per-role policy, gateway messaging
* **Process:** build 4 profiles with --clone, write distinct SOUL.md per role, keep project context in AGENTS.md, encode handoff contracts, run weekly memory-kpi, diff SOUL.md against day-one, stagger cron, enforce team-agents.md via commits
* **Guardrails:** No SOUL.md edit without logged reason; no handoff accepted without declared input shape; no role widened without orchestrator approval; no cron added without checking shared schedule

**Key quote:** "Profiles are the feature. Boundaries are the moat."
