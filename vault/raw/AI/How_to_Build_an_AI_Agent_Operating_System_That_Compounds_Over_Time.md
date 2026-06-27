---
title: "How to Build an AI Agent Operating System That Compounds Over Time"
source: "https://maxmitcham.substack.com/p/how-to-build-an-ai-agent-operating"
author:
  - "[[Max Mitcham]]"
published: 2026-04-10
created: 2026-04-25
description: "Setting up Hermes the OpenClaw killer."
tags:
  - "clippings"
---
![](https://www.youtube.com/watch?v=e9RyhHYzrao)

If you want an AI setup that actually compounds over time, you need more than a chatbot. You need a system that stores raw information properly, turns it into structured knowledge, and produces useful outputs from that knowledge.

This guide breaks down the stack and the exact prompt used to create an Obsidian-based memory system that supports multiple agents and gets more useful the more you use it.

---

## What you are building

- Hermes as the main controller
- OpenClaw for heavy tool calling when needed
- Codex and Claude Code for execution tasks
- An Obsidian vault that acts as long-term memory

The point is not just automation. The point is compounding automation.

## Why this setup works

Most people use AI like a chat window. That works for one-off tasks, but it falls apart when you need memory, consistency, workflows, and multiple agents working together.

This setup works because it separates three things clearly:

1. raw = source material
2. wiki = structured knowledge compiled by agents
3. output = deliverables generated from that knowledge

That structure stops your system turning into a messy folder full of chats, drafts, and half-finished documents.

## Step 1: Install Hermes

Start by installing Hermes. Hermes is the control layer for the whole stack. It manages the rest of the tools and gives you a cleaner, more reliable way to run agent workflows.

- Less setup than OpenClaw
- More reliable day to day
- Strong memory system
- Good security and prompt injection protection
- Works well as the orchestration layer

## Step 2: Add the API keys you actually need

During setup, Hermes will ask for integrations and API keys. Useful ones include Firecrawl for scraping and browser control, Browserbase for automation, and CameoFox for more resilient browsing.

Do not add everything just because you can. Start with the tools you know you will use.

## Step 3: Connect Hermes to Discord or Slack

If you want to interact with your agents from Discord or Slack, connect Hermes to one of those platforms.

For Discord, the basic flow is:

1. Create a new application in the Discord Developer Portal
2. Create a bot inside that application
3. Enable the required intents, especially message content and presence
4. Give the bot the permissions it needs, such as reading message history and sending messages
5. Copy the bot token into the Hermes setup flow
6. Use the OAuth installation URL to add the bot to your server

It sounds more painful than it is. It is mostly just permission setup.

## Step 4: Use Hermes as the main controller

Once Hermes is installed, run it from the terminal and switch models depending on the task. A practical way to think about the stack is:

- Hermes = control tower
- Codex = strong coding and execution
- Claude Code = another coding agent option
- OpenClaw = useful when a task needs lots of tool calls

You do not need all of these on day one. Start with the smallest stack that gets real work done.

## Step 5: Only use OpenClaw where it actually helps

You do not need OpenClaw just because people talk about it online. Use it where it has a real advantage: tool-heavy tasks and complex coordination. For daily orchestration, Hermes is the simpler control layer.

## Step 6: Build the memory layer in Obsidian

This is where the setup becomes genuinely powerful. Your Obsidian vault should not be treated like a notes folder. It should be treated like a memory system.

> This is not a notes folder. It is a three-layer memory architecture.

1. raw = immutable source material
2. wiki = structured agent-compiled knowledge
3. output = generated deliverables like reports, briefs, and slides

Most AI systems fail because they dump raw inputs, summaries, and outputs into one place. Do not do that.

## Step 7: Use this prompt to create the Obsidian system

<aside> ⚠️ Important: your followers should edit this prompt based on their own goals, workflows, content categories, and agent setup. The architecture should stay the same. The categories should change to fit their business.

</aside>

```markup
Create an Obsidian-based knowledge system that compounds over time, supports multiple agents, and separates raw inputs from structured knowledge and generated outputs.

Core principle:
This is not a notes folder. It is a three-layer memory architecture:
1. raw = immutable source material
2. wiki = structured agent-compiled knowledge
3. output = generated deliverables like reports, briefs, and slides

Set it up exactly like this.

Folder structure:
<ROOT>/
  SCHEMA.md
  log.md
  raw/
    articles/
    papers/
    competitors/
    repos/
    tweets/
    misc/
  wiki/
    _index.md
    market/
    content/
    product/
    ops/
    concepts/
  output/
    reports/
    slides/
    charts/
  templates/

System rules:
1. Raw is append-only. Never edit or delete raw source files after ingestion.
2. Wiki is where agents synthesize knowledge from raw sources.
3. Output is derived from wiki, never treated as source truth.
4. Every wiki page must use YAML frontmatter.
5. Every wiki page must link to at least 2 related pages.
6. Every meaningful action must be appended to log.md.
7. A central SCHEMA.md must define folder conventions, tag taxonomy, file naming rules, and page types.
8. Agents must search for existing pages before creating new ones to avoid duplicates.
9. New durable knowledge must be filed into this system, not left in chat.

Create these files first:

1. SCHEMA.md
It must define:
- Purpose of the vault
- The 3-layer architecture: raw, wiki, output
- Allowed domains under wiki: market, content, product, ops, concepts
- File naming convention: lowercase, hyphenated, no spaces
- Required frontmatter schema for wiki files
- Confidence levels: high, medium, low
- Tag taxonomy section
- Rules for cross-linking and index maintenance

2. log.md
Append-only operational log.
```
```markup
Format:
## [YYYY-MM-DD HH:MM] action | title
- Source: path if relevant
- Notes: short explanation

3. wiki/_index.md
This must include:
- Purpose of the wiki
- Domain directories
- Recent updates table
- Basic stats block
- Navigation links to sub-indexes

4. templates/article.md
Use this exact structure:

---
title: "Page Title"
type: entity
domain: market
tags: [example-tag]
sources:
  - raw/articles/example-source.md
created: YYYY-MM-DD
updated: YYYY-MM-DD
confidence: medium
---

# Summary

One-paragraph summary.

## Key Points

- Point 1
- Point 2
- Point 3

## Evidence

Detailed synthesis from sources.

## Related

- [[related-page-1]]
- [[related-page-2]]

5. templates/raw-source.md
Use this structure:

---
title: "Source Title"
source_url: "<https://example.com>"
date: YYYY-MM-DD
type: article
author:
tags: []
---

# Full Content

Paste the complete source content here.

Agent workflow:
When new information arrives, the agent must:
1. Save it to raw/<category>/<slug>.md using the raw-source template
2. Append an ingest entry to log.md
3. Search wiki for existing related pages
4. Either update an existing page or create a new one
5. Add at least 2 wiki links
6. Update the relevant index pages
7. Append a compile/update entry to log.md

Page creation rules:
- Create a new wiki page when a concept/entity appears in 2+ sources or is central to 1 major source
- Do not create thin stub pages
- Update existing pages before creating new ones

Retrieval behavior:
When answering future questions, the agent must:
1. Search the wiki first
2. Prefer structured wiki pages over raw notes
3. Cite which internal pages informed the answer
4. If memory is weak, say so and recommend ingesting more sources

Maintenance behavior:
Set up recurring checks for:
- orphan pages
- broken li
```
```markup
nks
- stale articles
- missing frontmatter
- pages not listed in indexes
- raw sources not yet compiled into wiki

If automation is available, create recurring jobs for:
- daily or nightly digest
- weekly health report
- periodic indexing/embedding refresh if semantic search is enabled

Search/indexing:
If semantic search is available, index both:
- raw sources collection
- wiki articles collection

The agent should prefer hybrid search for retrieval and keyword search for exact lookups.

Non-negotiables:
- Do not let useful work die in chat
- Do not overwrite raw source truth
- Do not create duplicate wiki pages
- Do not leave indexes stale
- Do not present speculative information as high confidence

Deliverables:
Create the full folder structure, generate the initial files, and leave the system ready for ongoing agent ingestion.
```

## Step 8: Customize the prompt to fit your business

Do not blindly use the prompt as-is. The architecture should stay the same, but the categories, outputs, and workflows should be adapted to the business.

- SaaS company: keep folders like product, market, and ops
- Agency: add clients, delivery, playbooks, and case studies
- Creator business: add scripts, sponsors, newsletters, and audience insights

## Step 9: Set up crons so the system improves automatically

Crons are scheduled automations. They let the system keep working without waiting for you to prompt it every time.

### Two useful cron types

- Brain-improving crons: update the memory system, categorize conversations, maintain indexes, check for stale pages
- Task-executing crons: monitor trends, draft content, review performance, watch releases, update help docs

That is the difference between a toy setup and a real operating system.

## Recommended starter workflow

1. Install Hermes
2. Connect the model providers you need
3. Connect Discord or Slack if you want messaging access
4. Create the Obsidian vault using the prompt above
5. Set up one daily brain digest cron
6. Set up one work cron, like trend scouting or release monitoring
7. Review outputs weekly and improve prompts based on what you learn

## Practical hosting advice

Do not overcomplicate hosting. A Mac Mini, an old machine, or a VPS is enough. The real priority is reliability, isolation, and clean access to the tools your agents need.

## Final takeaway

Stop thinking of AI as a chatbot. Start thinking of it as an operating system with memory, workflows, and scheduled execution. Once you do that, the real advantage is not just that the system can do tasks. It is that the system gets better the more it is used.

---

## Quick summary

- Use Hermes as the main orchestration layer
- Use OpenClaw only where heavy tool calling actually helps
- Use Codex and Claude Code for execution work
- Store knowledge in Obsidian using a raw / wiki / output architecture
- Use crons to keep the system updating itself
- Customize the prompt to fit your business instead of copying it blindly