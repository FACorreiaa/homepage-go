---
title: "How to Build an AI Second Brain with Claude Code and Obsidian"
source: "https://www.mindstudio.ai/blog/build-ai-second-brain-claude-code-obsidian-2"
author:
  - "[[MindStudio Team]]"
published: 2026-04-02
created: 2026-04-14
description: "Learn how to build a personal AI second brain using Claude Code and Obsidian that learns from every session and automates your daily business tasks."
tags:
  - "clippings"
---
![How to Build an AI Second Brain with Claude Code and Obsidian](https://i.mscdn.ai/70cbb1ad-08d7-4fdc-ab31-e343780966a6/generated-images/dbdc29dc-b7da-44cd-8b59-ac6483d74a64.png?fm=auto&w=1200&fit=cover?fm=auto&w=1200&fit=cover)

## Why Your Notes Are Useless Without a System That Thinks

Most people have a notes problem. They save links, write meeting summaries, jot down ideas — and then never find any of it when they actually need it. The problem isn’t the volume. It’s that notes don’t connect, don’t resurface, and definitely don’t do anything.

An AI second brain changes that. By combining Claude Code with Obsidian, you can build a personal knowledge system that captures what matters, links ideas automatically, and runs real tasks on your behalf — not just during a session, but continuously, as the system learns more about how you work.

This guide walks through exactly how to do it: setting up Obsidian as your knowledge vault, connecting Claude Code to read and write your notes, and building workflows that get smarter over time.

---

## What an AI Second Brain Actually Means

The term “second brain” comes from Tiago Forte’s productivity framework, which treats an external system — notes, documents, references — as an extension of your memory. The idea is to offload cognitive load so your working memory is free for actual thinking.

Pairing that with AI takes it further. Instead of just storing information, your system can:

- Surface relevant notes when you start a new task
- Write summaries and extract key points automatically
- Connect ideas across different projects and time periods
- Trigger actions (send emails, update databases, draft documents) based on what it knows about you

Obsidian handles the storage. Claude Code handles the intelligence layer. Together, they form something that actually gets better the more you use it.

---

## Setting Up Obsidian as Your Knowledge Vault

Obsidian stores notes as plain Markdown files in a local folder — your “vault.” This matters a lot. Because files are plain text, Claude Code can read, write, and search them directly without needing any special API or integration.

### Structure Your Vault for Machine Readability

How you organize your vault determines how useful it is for AI. A few principles:

- **Use consistent YAML frontmatter.** Every note should have metadata at the top — date, tags, project, status. Claude can filter and query this.
- **Keep atomic notes.** One idea per note is easier to link and retrieve than long documents.
- **Use a clear folder structure.** Something like `/Projects`, `/People`, `/Daily Notes`, `/Resources`, `/Archive` gives Claude a mental map.
- **Use Obsidian’s `[[wikilinks]]`.** These backlinks create a graph Claude can traverse to understand how concepts relate.

A simple frontmatter block looks like this:

```yaml
yaml---
date: 2025-06-10
tags: [project, strategy, q3]
status: active
project: product-launch
---
```

### Install the Right Plugins

Obsidian has a rich plugin ecosystem. For AI second brain purposes, these matter most:

- **Dataview** — lets you query your vault like a database (e.g., “all notes tagged #meeting from the last 30 days”)
- **Templater** — enforces consistent note structure via templates
- **Daily Notes** — creates a dated note each day where you can log context, tasks, and decisions
- **Periodic Notes** — extends daily notes to weekly and monthly reviews

You don’t need to install every plugin. But Dataview in particular becomes very useful when Claude needs to find specific notes programmatically.

---

## Connecting Claude Code to Your Obsidian Vault

Claude Code is Anthropic’s agentic coding tool that runs in your terminal. It can read files, write code, execute commands, and take multi-step actions — which makes it perfect for interacting with a local Obsidian vault.

### Prerequisites

- Node.js installed
- Claude Code installed (`npm install -g @anthropic-ai/claude-code`)
- Your Obsidian vault in a known local path (e.g., `~/obsidian-vault`)
- Anthropic API access

### Give Claude Code a CLAUDE.md File

The key to making Claude Code useful for your second brain is a `CLAUDE.md` file in the root of your vault. This is a persistent context file Claude reads at the start of every session.

Think of `CLAUDE.md` as Claude’s orientation document. It tells Claude:

- What your vault structure is
- What naming conventions you use
- What recurring tasks it should be able to help with
- Any personal context it should keep in mind (your role, your current projects, your preferences)

Here’s a minimal example:

```markdown
markdown# Vault Overview

This is [Your Name]'s personal knowledge vault.

## Folder Structure
- /Daily Notes — one note per day, format: YYYY-MM-DD.md
- /Projects — active project folders, each with a README.md
- /People — notes on contacts and relationships
- /Resources — articles, books, references

## Naming Conventions
- Notes use kebab-case: project-kickoff-notes.md
- Tags are lowercase, hyphenated: #client-work

## Recurring Tasks
- When I say "morning brief," summarize today's Daily Note, open tasks, and relevant project context.
- When I say "capture [topic]," create a new note in /Resources with a summary and tags.
- When I say "weekly review," aggregate last 7 days of Daily Notes into a summary.

## Current Projects
- Product Launch (Q3) — folder: /Projects/product-launch
- Client: Acme Corp — folder: /Projects/acme-corp
```

The more specific this file is, the better Claude performs. Update it regularly.

### Test the Connection

From your terminal, navigate to your vault and run Claude Code:

```bash
bashcd ~/obsidian-vault
claude
```

Ask it something specific: “What notes have I tagged #strategy this month?” Claude will read your files and respond. If you’ve set up Dataview syntax, it can even construct queries for you.

---

## Building Workflows That Learn From Every Session

A second brain isn’t useful if it only works when you’re actively asking it questions. The real value comes from workflows that capture context automatically and make it available later.

### The Daily Note Workflow

The foundation of most second brain systems is a daily note — a running log for the day. Use Templater to auto-create one each morning with sections for:

- Top 3 priorities
- Meeting notes
- Decisions made
- Open questions
- Links to relevant projects

Then add a Claude Code routine you can trigger at the end of each day:

**Prompt:** “Review today’s Daily Note, extract any open tasks, link them to their relevant projects, and add a TL;DR summary at the top.”

Claude reads the note, identifies action items, updates the project files, and writes a summary. You don’t have to do any of this manually. Over time, your project files accumulate rich context from every day’s log.

### The Meeting Notes Workflow

Before a meeting, run:

**Prompt:** “I’m about to meet with \[Person\] about \[Topic\]. Pull relevant notes from /People/\[person\].md and any active projects linked to them.”

Claude surfaces what you need in seconds. After the meeting, paste your rough notes and run:

**Prompt:** “Format these meeting notes, extract action items with owners and deadlines, and append them to \[person\].md and the relevant project README.”

This is where the second brain starts to feel genuinely useful. You’re not just storing notes — you’re building structured, searchable records that Claude can reason over.

### The Weekly Review Workflow

A weekly review is where your second brain compounds. Set this as a recurring prompt:

**Prompt:** “Review all Daily Notes from this week. Summarize key decisions, flag any tasks that appear open across multiple days, identify which projects had the most activity, and write a weekly summary note in /Reviews/\[date\]-weekly.md.”

Claude stitches together the week’s context into a coherent picture. After a few months, you’ll have a detailed record of how your time was actually spent — useful for performance reviews, client reporting, or just understanding your own patterns.

---

## Automating Daily Business Tasks

Beyond note management, Claude Code can take real actions using the context it’s built up in your vault.

### Email Drafting With Full Context

Instead of starting from scratch every time you need to write an email, let Claude pull context first:

**Prompt:** “Draft a follow-up email to \[Person\]. Check their note for recent meeting history and any open action items from our last conversation.”

Claude reads the contact note, finds the relevant context, and writes a draft that actually reflects the relationship history. You review and send.

### Project Status Reports

If your project folders follow a consistent structure, Claude can write status reports automatically:

**Prompt:** “Generate a status report for the product-launch project. Include: current phase, completed milestones, open blockers, and next steps. Format it for sharing with stakeholders.”

Claude reads the README, meeting notes, and task logs across the project folder and produces a coherent report. This task alone can save an hour every week.

### Research Capture and Synthesis

When you’re researching a topic, Claude can help you build a knowledge base in real time:

**Prompt:** “I’m researching \[topic\]. I’ll paste three articles. Synthesize the key points, identify where they agree and disagree, and save the result as a new note in /Resources with appropriate tags.”

Over time, your vault becomes a curated library of synthesized research — not just a pile of bookmarks.

### Prompting Yourself at the Right Time

You can use Claude Code with simple shell scripts and cron jobs to trigger context-aware reminders. For example, a morning script that:

1. Reads your Daily Note template
2. Pulls open tasks from the last three days
3. Summarizes active project status
4. Outputs a morning brief in your terminal

This gives you a personalized, context-rich briefing every morning with no manual work.

---

## Taking It Further With MindStudio

The Claude Code and Obsidian setup is powerful, but it’s built for developers comfortable working in a terminal. If you want to extend your second brain to team workflows, automated pipelines, or web-based interfaces — or if you’re not a developer — MindStudio offers a natural next step.

MindStudio’s [Agent Skills Plugin](https://mindstudio.ai/) is an npm SDK that lets Claude Code and other AI agents call over 120 typed capabilities as simple method calls. That means the same Claude agent reasoning over your Obsidian vault can also call `agent.sendEmail()`, `agent.searchGoogle()`, `agent.runWorkflow()`, or push data to tools like HubSpot, Notion, and Google Workspace — without you building and maintaining that infrastructure yourself.

So instead of just drafting an email in your terminal, Claude can send it. Instead of summarizing a project status, Claude can post it to Slack. Instead of identifying a follow-up task, Claude can add it to your project management tool.

For teams, MindStudio’s no-code builder lets non-technical members create their own AI agents that connect to shared knowledge bases — no API keys, no terminal, just a visual interface. You can expose your Claude Code second brain logic as a workflow others can trigger through a form or chat interface.

[Try MindStudio free at mindstudio.ai](https://mindstudio.ai/) — building a basic agent takes under an hour.

---

## Common Mistakes to Avoid

Getting this setup right takes some iteration. Here are the problems people run into most often:

**Inconsistent note structure.** If some notes have YAML frontmatter and others don’t, Claude has to guess. Enforce templates from the start.

**Overly broad prompts.** “Summarize my work” is too vague. “Summarize all notes tagged #client-work from the last two weeks and identify open deliverables” gives Claude something to work with.

**Not updating CLAUDE.md.** As your vault evolves, your context file gets stale. Set a reminder to review it monthly.

**Trying to automate everything at once.** Start with one workflow — the daily note is the easiest. Get that working reliably before adding more.

**Treating Claude Code as stateless.** Unlike a chat session, Claude Code has persistent context through your vault. Use that. Reference past notes, ask Claude to compare current context with historical notes, build on what’s already there.

---

## Frequently Asked Questions

### What is an AI second brain?

An AI second brain is a personal knowledge system that not only stores your notes and ideas but actively uses AI to organize, link, and act on them. Unlike a standard note-taking app, an AI second brain can surface relevant context, draft content, extract tasks, and trigger actions based on accumulated knowledge. The combination of Claude Code and Obsidian is one of the most capable setups available because it uses plain-text files (easy for AI to read and write) alongside a powerful AI agent.

### Do I need to know how to code to build this?

Basic comfort with the terminal helps. Claude Code runs in the command line, and setting it up requires some technical steps. That said, once it’s configured, the day-to-day use is natural language — you type prompts, Claude takes actions. If you want a fully no-code version of this concept, platforms like [MindStudio](https://mindstudio.ai/) let you build AI agents with a visual interface and connect them to tools like Notion, Google Workspace, or Airtable without writing code.

### Is my data safe using Claude Code with a local vault?

Obsidian stores all your notes locally by default — they never leave your machine unless you choose to sync them. When you use Claude Code, your prompts and file contents are sent to Anthropic’s API. Review [Anthropic’s privacy policy](https://www.anthropic.com/privacy) to understand how your data is handled. For sensitive information, you can use Obsidian’s local-only mode and avoid including that content in Claude Code sessions.

### How does this differ from just using ChatGPT or Claude in a browser?

Browser-based AI tools are stateless — each conversation starts fresh. They don’t know what you worked on yesterday, who your clients are, or what projects are active. Claude Code with Obsidian is different because Claude has persistent access to your entire knowledge vault. It reads your history, builds on previous sessions, and maintains context over time. That’s what makes it a second brain rather than just a chatbot.

### Can I use this system for team collaboration, not just personal use?

The setup described here is primarily for individual use. For team collaboration, you’d need a shared knowledge base (a shared Obsidian vault synced via Obsidian Sync or a Git repo) and a way for multiple people to trigger Claude Code workflows. MindStudio handles this more elegantly — you can build shared AI agents that connect to team tools like Slack, Notion, and HubSpot, and expose them through interfaces anyone on the team can use, no terminal required. Learn more about [building AI workflows for teams](https://mindstudio.ai/).

### What’s the difference between Obsidian and Notion for this purpose?

Obsidian stores files as local Markdown, which makes them trivially easy for Claude Code to read and write. Notion uses a proprietary database format and requires API calls to access. Both can work, but Obsidian’s simplicity is a significant advantage for AI integration. Claude can open, read, and write Obsidian notes using basic file system operations — no extra setup required.

---

## Key Takeaways

- An AI second brain combines a structured note vault (Obsidian) with an AI agent (Claude Code) that can read, write, and act on your knowledge.
- The `CLAUDE.md` file is the most important piece — it gives Claude persistent context about who you are, how your vault is structured, and what tasks it can handle.
- Start with one workflow (daily notes), get it reliable, then expand to meeting notes, weekly reviews, and task automation.
- The system compounds over time — the more context Claude has, the more useful and accurate its outputs become.
- For teams or no-code setups, MindStudio extends this concept with 120+ built-in capabilities, no-code agent building, and integrations with the tools your team already uses.

If you’re ready to build something beyond a local setup — or want to create AI workflows your whole team can use — [MindStudio is free to start](https://mindstudio.ai/).