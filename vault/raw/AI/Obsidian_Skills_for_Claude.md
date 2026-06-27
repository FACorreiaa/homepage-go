---
title: "Obsidian Skills for Claude"
source: "https://grok.com/c/b5ae9ce3-c4ad-418c-a573-18674b3b2ca5?rid=e846fa2f-ebe6-4522-a2f3-a54f8a139ab7"
author:
published:
created: 2026-04-24
description: "Grok is an AI-powered assistant built by xAI. Chat via text or voice, generate images and videos, search the web in real-time, write code, and reason through complex problems."
tags:
  - "clippings"
---
What is the name of the Claude skills with obsidian convo?

**The main one is "obsidian-skills" (by Obsidian's creator, kepano).**

This is a popular, official-style collection of **Claude skills** (Agent Skills) designed specifically for use with Obsidian vaults alongside Claude Code (or similar agents like Codex CLI). You drop them into your.claude/skills folder (or install via marketplace/npx), and they let Claude expertly handle Obsidian-specific tasks like:

- Editing Obsidian Flavored Markdown (.md files with wikilinks, callouts, embeds, properties, etc.).
- Working with Obsidian Bases (.base files).
- Creating/editing JSON Canvas (.canvas files).
- Using the Obsidian CLI for vault management, plugins, etc.
- Tools like Defuddle for clean web-to-markdown extraction.

These skills turn your Obsidian vault into a living "second brain" that Claude can read, edit, and maintain seamlessly during conversations/sessions. Many users combine them with custom skills for context loading, project management, note ingestion from convos, and more.

### Other notable related skills

- **Deep Recon** (/deep-recon): An experimental multi-agent skill for broader research across your vault + the web.
- Various custom/user-created ones for conversation ingestion, task management, thinking tools, etc., often built on top of or alongside obsidian-skills.

If you're referring to a specific custom skill, conversation-handling setup ("convo"), or something else (e.g., a particular YouTube tutorial's skills), feel free to share more details for a more targeted answer! You can find the main repo here: [github.com/kepano/obsidian-skills](https://github.com/kepano/obsidian-skills).

What about an iOS app where the users can mimic the same? Create notes, screenshots, etc all text saved with md format and then connecting the app with an LLM (Claude or OpenAI api) and have a diary and memories handled by ai that learns the more the user adds things to his memo. Could this be a business idea? 14 day trial and subscription?

**Yes, this is a solid business idea with real market potential—especially as a polished, iOS-native “Obsidian + Claude skills” experience focused on quick capture (notes + screenshots → Markdown) and an AI-powered diary/memories system that gets smarter over time.**

It directly addresses what many Obsidian users on mobile want: seamless Markdown vaults they own/control, frictionless input (including screenshots via Share Sheet with OCR + AI structuring), and built-in LLM integration (Claude or OpenAI API) for a living “second brain” that auto-connects ideas, summarizes days, surfaces patterns, and acts as an intelligent diary companion. The more the user adds (memos, screenshots, reflections), the richer the AI’s context becomes via retrieval-augmented generation (RAG) over their local MD files.

### Why it fits the market in 2026

The AI note-taking / second-brain / journaling space is crowded but still growing fast—lots of “best of 2026” roundups highlight demand for tools that turn raw input into intelligent, connected memories. Your concept stands out by being:

- **Markdown-first and portable** (local.md files in iCloud or a folder, just like Obsidian—not proprietary blobs).
- **Capture-first for mobile life** (one-tap screenshots → OCR text → clean Markdown note + AI tagging/summary added to the memory layer).
- **Diary + memories core** (AI learns from everything: auto-daily summaries, mood trends, connection graphs, queryable “what did I learn about X last month?”).

Closest existing apps (none perfectly match your full vision):

- **Mem.ai** → Excellent AI thought partner that auto-organizes notes, chats with your knowledge base, and has a strong iOS app. Supports voice/ideas/meetings and now Claude integration. But notes aren’t pure local Markdown files.
- **Reflect Notes** → Very close: minimalist iOS app with daily notes (diary feel), Markdown, backlinks, AI chat for insights, voice memos. Strong privacy focus. Pricing is ~$10/month (billed annually) with a 14-day trial—exactly your model.
- **TakeNote** → Apple-native Markdown app with local AI summaries, folders/tags, very private. Lightweight and fast, but lighter on the “AI memories that learn” side.

Your edge: True Obsidian-style vault + native iOS screenshot-to-structured-MD workflow + user-controlled LLM keys (they plug in their own Claude/OpenAI key so you don’t eat API costs and they keep full privacy/control). Apple Intelligence (on-device in 2026) could be a hybrid layer for offline OCR/summaries to reduce cloud dependency.

### Business model: 14-day trial + subscription = proven and smart

This is exactly how the successful players do it (Mem Pro ~$12–15/month, Reflect ~$10/month annual). It works beautifully for this category because:

- **Trial hooks users** — 14 days is long enough to build a habit (add 20–30 notes/screenshots, see the AI start making connections) and feel the “it’s learning me” magic.
- **Subscription justifies itself** — Tier it like this:
	- Free: Basic Markdown notes, screenshots → text, local storage, limited AI queries.
		- Pro ($4.99–$9.99/month or $49–$89/year): Unlimited AI (with user API key), full memory graph, auto-insights, advanced search, templates, iCloud sync extras, export/backup.
- Revenue potential is realistic — Sticky daily-use apps (diary + capture) have good retention. Many users already pay for Reflect/Mem or Obsidian Sync + separate Claude usage.

### Realistic challenges (and how to win)

- **Competition** → High, but most aren’t pure Markdown + screenshot-first + user-key LLM. Market it as “Obsidian skills, but native on your iPhone with zero desktop required.”
- **Tech** → Feasible today: SwiftUI + SwiftData for local MD, Apple Vision for OCR on screenshots, embed a vector DB (e.g., via CoreML or simple local store) for memories, let users input their own API keys. Start simple (RAG over notes) and add agent-like “skills” later (auto journal prompts, weekly reviews).
- **Privacy & costs** — Huge selling point: never store their notes or force your own API billing. Users love this (see TakeNote/Reflect emphasis on privacy).
- **App Store & retention** — Focus on delightful mobile UX (widgets for quick capture, Live Activities for streak reminders, Shortcuts integration). Diary apps thrive on streaks and insights.

**Bottom line: Yes—build it.** The timing is perfect (AI memory tools are hot, iOS users want native polish, Obsidian mobile power users are underserved on the AI side). Start with a focused MVP: Markdown vault + screenshot → AI-enhanced note + basic memory chat. Launch with the 14-day trial and watch conversions. If you nail the “it just gets me” feeling as the user adds more, it becomes sticky and monetizable fast.

If you want help refining features, pricing tiers, or even a quick competitive matrix, just say the word! This could be a winner.

Was thinking of doing the iOS app with SwiftUI and the backend with Vapor

**Yes, SwiftUI + Vapor is an excellent, cohesive choice for this app—especially if you're already comfortable with Swift.** It plays to your strengths as a solo or small-team developer and aligns well with the privacy-focused, Markdown-first vision.

### Why this stack works great for your idea

- **Unified language & models**: You can share Swift types, logic, and even some business rules (e.g., Markdown parsing, note models, memory/RAG structures) between the iOS app and backend. Many full-stack Swift developers highlight this as a huge productivity win.
- **SwiftUI strengths**: Perfect for a beautiful, native iOS capture experience—quick note entry, screenshot handling (via Share Sheet + Vision framework for OCR), Markdown editing/preview, widgets, Live Activities for streaks, and smooth animations. In 2026, SwiftUI is fully production-ready for complex apps.
- **Vapor strengths**: Mature, high-performance server-side Swift framework (still actively updated, with Vapor 4 stable and Vapor 5 on the horizon emphasizing async/concurrency). Great for building a clean REST/GraphQL API for:
	- User accounts & subscriptions (RevenueCat or Stripe integration).
		- Optional cloud sync (if users want multi-device or backups beyond iCloud).
		- Secure proxying or orchestration for LLM calls (if you ever want to add optional server-side AI features without exposing keys).
		- Analytics, backups, or advanced memory graph computation.

**Core architecture recommendation** (local-first + backend where needed):

- **Local layer** (iOS): SwiftData (or FileManager for raw.md files in iCloud/Documents) for the Markdown vault. This keeps notes fully user-owned and offline-capable. Use Apple’s on-device tools (Vision for screenshot OCR → structured MD, CoreML/Apple Intelligence for lightweight summaries) to reduce API calls.
- **Backend (Vapor)**: Handle auth, subscription management, optional encrypted sync, and any heavy RAG/vector search that’s better centralized. Deploy on a cheap VPS (Hetzner, DigitalOcean) with PostgreSQL + Docker—many devs do exactly this.
- **LLM integration**: Let users input their own Claude/OpenAI keys in the app. Call APIs directly from the device (or route through Vapor for logging/anonymization if desired). This keeps your costs near-zero and privacy high.

### Potential challenges & how to mitigate

- **Ops overhead**: Vapor means you manage a server (deployment, scaling, backups, security). Start simple—use Docker + a managed Postgres. It’s very doable for a side project or indie app, but expect more maintenance than Firebase/Supabase.
- **Sync complexity**: Pure local Markdown with iCloud works for many users. Only add Vapor sync if you want seamless multi-device or web access later.
- **Time to MVP**: Focus first on local capture + Markdown + basic on-device AI diary features. Add Vapor for accounts/subscriptions in phase 2.
- **Alternatives if you want less backend work**:
	- Supabase or Firebase for quicker auth/DB/sync.
		- Pure local + RevenueCat (no custom backend needed for subs).

Many indie devs successfully ship with SwiftUI + Vapor, and the community is supportive (Discord, etc.).

This stack gives you full control, a delightful native feel, and a strong “it just works with my existing Markdown/Obsidian habits” story. If you nail the screenshot-to-smart-note flow and the “AI that learns my life” diary insights, the 14-day trial → subscription conversion should be solid.

Want help with:

- High-level architecture diagram or folder structure?
- Key libraries (e.g., for Markdown, vector search, subscriptions)?
- MVP feature prioritization?
- Deployment tips for Vapor?

Just let me know—happy to dive deeper! This has real legs. 🚀

Instead of Vapor I might just use Humingbird. A full local experience won’t make the “vault” get too large?

**Hummingbird is a great alternative to Vapor—especially for a lean, privacy-first app like yours.** It's lighter, faster, and more modular, which fits a solo or small-team project perfectly.

### Hummingbird vs Vapor Quick Comparison

- **Hummingbird**: Lightweight, minimal dependencies, excellent performance (often 2x requests/sec in benchmarks), built on modern structured concurrency. Ideal for smaller-to-medium backends, Lambdas, or when you want to avoid bloat. Modular—you only add what you need (e.g., auth, WebSockets, Fluent ORM if required). Growing community with good examples and extensions.
- **Vapor**: More batteries-included, larger ecosystem, easier for complex apps out of the box. Heavier footprint.

**Recommendation**: Go with **Hummingbird** if your backend needs are focused (user accounts, subscriptions via RevenueCat/Stripe, optional encrypted sync, or proxying LLM calls). You'll have a smaller binary, faster compiles, and lower hosting costs (easy on cheap VPS or serverless). Many Swift server devs are shifting toward it for new projects.

You can still share a lot of code/models with your SwiftUI app. Start simple and scale as needed—Hummingbird handles production use cases well.

### Full Local Experience: Will the Vault Get Too Large?

**Short answer: No major issues for realistic personal use.** A pure local Markdown vault on iOS is very scalable.

- **No hard limits** — iOS apps are limited only by the device's available storage (e.g., 64GB–1TB+ phones). Markdown files are tiny (a dense daily note might be 5–50KB). Even with thousands of notes + screenshots (processed to text + embeddings), you'll stay under a few GB for years of heavy use.
- **Real-world Obsidian-like vaults** — Users routinely run 10,000–100,000+ notes locally with no problems. Performance depends more on search/indexing than raw size. Your app can use SwiftData or direct FileManager for.md files in the app's Documents/iCloud folder.
- **Screenshots & media** — Handle these smartly:
	- OCR → clean Markdown text (store the text, not always the full image).
		- Optional: Compress/resize images or let users choose attachments.
		- This keeps the vault lean while preserving the diary/memories value.

**Vector embeddings for "AI that learns" (RAG/memory layer)**: These add the most size. A typical 1536-dim embedding (OpenAI/Claude style) is ~6KB per chunk + metadata. For 10,000 notes/chunks that's ~60MB—perfectly fine. Use a local vector store (e.g., via Swift + SQLite + vector extension, or a lightweight lib). On-device Apple Intelligence can help with lighter summaries to reduce embedding volume.

**Potential gotchas & mitigations**:

- **Performance slowdown** — Very large vaults (tens of thousands of files) can slow full-text search or graph views. Solution: Smart indexing (only index recent/active notes), background processing, or optional cloud sync for heavy users.
- **iCloud sync** — Works great for local-first but has its own quotas. Offer it as optional; many users prefer pure local + manual backups.
- **App size & backups** — Keep the app binary small. Users can export their entire vault as a zip/folder anytime (huge selling point vs proprietary apps).
- **Edge cases** — Power users adding daily screenshots + long voice transcripts could grow faster. Add usage tips or a "vault health" view that suggests pruning/archiving old notes.

**Overall**: A full local experience is one of your strongest features—users own their data, zero vendor lock-in, works offline, and aligns with Obsidian lovers. The vault won't "get too large" for normal diary/memo use. Combine it with optional Hummingbird backend for sync/accounts/subs, and you have the best of both worlds.

This setup keeps development enjoyable and the product delightful. If you want specifics—like recommended Hummingbird packages for auth/sync, local vector DB options, or how to structure the Markdown + embeddings layer—let me know! You're on a strong path. 🚀