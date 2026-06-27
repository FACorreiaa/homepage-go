---
title: "rohitg00/agentmemory: #1 Persistent memory for AI coding agents based on real-world benchmarks"
source: "https://github.com/rohitg00/agentmemory"
author:
published:
created: 2026-05-07
description: "#1 Persistent memory for AI coding agents based on real-world benchmarks - rohitg00/agentmemory"
tags:
  - "clippings"
---
[![agentmemory — Persistent memory for AI coding agents](https://github.com/rohitg00/agentmemory/raw/main/assets/banner.png)](https://github.com/rohitg00/agentmemory/blob/main/assets/banner.png)

**Your coding agent remembers everything. No more re-explaining.**  
Persistent memory for Claude Code, Cursor, Gemini CLI, Codex CLI, pi, OpenCode, and any MCP client.

**The gist extends Karpathy's LLM Wiki pattern with confidence scoring, lifecycle, knowledge graphs, and hybrid search.  
agentmemory is the implementation.**

![95.2% retrieval R@5](https://github.com/rohitg00/agentmemory/raw/main/assets/tags/stat-recall.svg) ![92% fewer tokens](https://github.com/rohitg00/agentmemory/raw/main/assets/tags/stat-tokens.svg) ![51 MCP tools](https://github.com/rohitg00/agentmemory/raw/main/assets/tags/stat-tools.svg) ![12 auto hooks](https://github.com/rohitg00/agentmemory/raw/main/assets/tags/stat-hooks.svg) ![0 external DBs](https://github.com/rohitg00/agentmemory/raw/main/assets/tags/stat-deps.svg) ![827 tests passing](https://github.com/rohitg00/agentmemory/raw/main/assets/tags/stat-tests.svg)

[![agentmemory demo](https://github.com/rohitg00/agentmemory/raw/main/assets/demo.gif)](https://github.com/rohitg00/agentmemory/blob/main/assets/demo.gif)

[Quick Start](#quick-start) • [Benchmarks](#benchmarks) • [vs Competitors](#vs-competitors) • [Agents](#works-with-every-agent) • [How It Works](#how-it-works) • [MCP](#mcp-server) • [Viewer](#real-time-viewer) • [iii Console](#iii-console--trace-level-engine-inspection) • [Config](#configuration) • [API](#api)

---

agentmemory works with any agent that supports hooks, MCP, or REST API. All agents share the same memory server.

| [![Claude Code](https://camo.githubusercontent.com/f5376c03e89d0240a856b2879a31b8aec2b92ac6414b3c22fe3e31bc5618d7f1/68747470733a2f2f6d61747468696173726f6465722e636f6d2f636f6e74656e742f696d616765732f323032362f30312f436c617564652e706e673f73697a653d313230)](https://claude.com/product/claude-code)   **Claude Code**   <sub>12 hooks + MCP + skills</sub> | [![OpenClaw](https://github.com/openclaw.png?size=120)](https://github.com/rohitg00/agentmemory/blob/main/integrations/openclaw)   **OpenClaw**   <sub>MCP + <a href="https://github.com/rohitg00/agentmemory/blob/main/integrations/openclaw">plugin</a></sub> | [![Hermes](https://github.com/NousResearch.png?size=120)](https://github.com/rohitg00/agentmemory/blob/main/integrations/hermes)   **Hermes**   <sub>MCP + <a href="https://github.com/rohitg00/agentmemory/blob/main/integrations/hermes">plugin</a></sub> | [![Cursor](https://camo.githubusercontent.com/559a53ed1552d8bd45143984c97ad060913c128edf6331207c67c8bb25ff2f45/68747470733a2f2f7777772e667265656c6f676f766563746f72732e6e65742f77702d636f6e74656e742f75706c6f6164732f323032352f30362f637572736f722d6c6f676f2d667265656c6f676f766563746f72732e6e65745f2e706e67)](https://cursor.com/)   **Cursor**   <sub>MCP server</sub> | [![Gemini CLI](https://github.com/google-gemini.png?size=120)](https://github.com/google-gemini/gemini-cli)   **Gemini CLI**   <sub>MCP server</sub> | [![OpenCode](https://github.com/opencode-ai.png?size=120)](https://github.com/opencode-ai/opencode)   **OpenCode**   <sub>MCP server</sub> | [![Codex CLI](https://github.com/openai.png?size=120)](https://github.com/openai/codex)   **Codex CLI**   <sub>MCP server</sub> | [![Cline](https://github.com/cline.png?size=120)](https://github.com/cline/cline)   **Cline**   <sub>MCP server</sub> |
| --- | --- | --- | --- | --- | --- | --- | --- |
| [![Goose](https://github.com/block.png?size=120)](https://github.com/block/goose)   **Goose**   <sub>MCP server</sub> | [![Kilo Code](https://github.com/Kilo-Org.png?size=120)](https://github.com/Kilo-Org/kilocode)   **Kilo Code**   <sub>MCP server</sub> | [![Aider](https://github.com/Aider-AI.png?size=120)](https://github.com/Aider-AI/aider)   **Aider**   <sub>REST API</sub> | [![Claude Desktop](https://github.com/anthropics.png?size=120)](https://claude.ai/download)   **Claude Desktop**   <sub>MCP server</sub> | [![Windsurf](https://camo.githubusercontent.com/14295bc149d2b65ee125eee3d4700438034c76a9c079ef595fde2b5ccb385acd/68747470733a2f2f65786166756e6374696f6e2e6769746875622e696f2f7075626c69632f6272616e642f77696e64737572662d626c61636b2d73796d626f6c2e7376673f73697a653d313230)](https://windsurf.com/)   **Windsurf**   <sub>MCP server</sub> | [![Roo Code](https://github.com/RooCodeInc.png?size=120)](https://github.com/RooCodeInc/Roo-Code)   **Roo Code**   <sub>MCP server</sub> | [![Claude SDK](https://github.com/anthropics.png?size=120)](https://github.com/anthropics/claude-agent-sdk-typescript)   **Claude SDK**   <sub>AgentSDKProvider</sub> | **Any agent**   <sub>REST API</sub> |

<sub>Works with <strong>any</strong> agent that speaks MCP or HTTP. One server, memories shared across all of them.</sub>

---

You explain the same architecture every session. You re-discover the same bugs. You re-teach the same preferences. Built-in memory (CLAUDE.md,.cursorrules) caps out at 200 lines and goes stale. agentmemory fixes this. It silently captures what your agent does, compresses it into searchable memory, and injects the right context when the next session starts. One command. Works across agents.

**What changes:** Session 1 you set up JWT auth. Session 2 you ask for rate limiting. The agent already knows your auth uses jose middleware in `src/middleware/auth.ts`, your tests cover token validation, and you chose jose over jsonwebtoken for Edge compatibility. No re-explaining. No copy-pasting. The agent just *knows*.

```
npx @agentmemory/agentmemory
```

> **New in v0.9.0** — Landing site at [agent-memory.dev](https://agent-memory.dev/), filesystem connector (`@agentmemory/fs-watcher`), standalone MCP now proxies to the running server so hooks and the viewer agree, audit policy codified across every delete path, health stops flagging `memory_critical` on tiny Node processes. Full notes in [CHANGELOG.md](https://github.com/rohitg00/agentmemory/blob/main/CHANGELOG.md#090--2026-04-18).

---

| ### Retrieval Accuracy  **LongMemEval-S** (ICLR 2025, 500 questions)  \| System \| R@5 \| R@10 \| MRR \| \| --- \| --- \| --- \| --- \| \| **agentmemory** \| **95.2%** \| **98.6%** \| **88.2%** \| \| BM25-only fallback \| 86.2% \| 94.6% \| 71.5% \| | ### Token Savings  \| Approach \| Tokens/yr \| Cost/yr \| \| --- \| --- \| --- \| \| Paste full context \| 19.5M+ \| Impossible (exceeds window) \| \| LLM-summarized \| ~650K \| ~$500 \| \| **agentmemory** \| **~170K** \| **~$10** \| \| agentmemory + local embeddings \| ~170K \| **$0** \| |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |

> Embedding model: `all-MiniLM-L6-v2` (local, free, no API key). Full reports: [`benchmark/LONGMEMEVAL.md`](https://github.com/rohitg00/agentmemory/blob/main/benchmark/LONGMEMEVAL.md), [`benchmark/QUALITY.md`](https://github.com/rohitg00/agentmemory/blob/main/benchmark/QUALITY.md), [`benchmark/SCALE.md`](https://github.com/rohitg00/agentmemory/blob/main/benchmark/SCALE.md). Competitor comparison: [`benchmark/COMPARISON.md`](https://github.com/rohitg00/agentmemory/blob/main/benchmark/COMPARISON.md) — agentmemory vs mem0, Letta, Khoj, claude-mem, Hippo.

---

|  | agentmemory | mem0 (53K ⭐) | Letta / MemGPT (22K ⭐) | Built-in (CLAUDE.md) |
| --- | --- | --- | --- | --- |
| **Type** | Memory engine + MCP server | Memory layer API | Full agent runtime | Static file |
| **Retrieval R@5** | **95.2%** | 68.5% (LoCoMo) | 83.2% (LoCoMo) | N/A (grep) |
| **Auto-capture** | 12 hooks (zero manual effort) | Manual `add()` calls | Agent self-edits | Manual editing |
| **Search** | BM25 + Vector + Graph (RRF fusion) | Vector + Graph | Vector (archival) | Loads everything into context |
| **Multi-agent** | MCP + REST + leases + signals | API (no coordination) | Within Letta runtime only | Per-agent files |
| **Framework lock-in** | None (any MCP client) | None | High (must use Letta) | Per-agent format |
| **External deps** | None (SQLite + iii-engine) | Qdrant / pgvector | Postgres + vector DB | None |
| **Memory lifecycle** | 4-tier consolidation + decay + auto-forget | Passive extraction | Agent-managed | Manual pruning |
| **Token efficiency** | ~1,900 tokens/session ($10/yr) | Varies by integration | Core memory in context | 22K+ tokens at 240 obs |
| **Real-time viewer** | Yes (port 3113) | Cloud dashboard | Cloud dashboard | No |
| **Self-hosted** | Yes (default) | Optional | Optional | Yes |

---

Compatibility: this release targets stable `iii-sdk` `^0.11.0` and iii-engine v0.11.x.

### Try it in 30 seconds

```
# Terminal 1: start the server
npx @agentmemory/agentmemory

# Terminal 2: seed sample data and see recall in action
npx @agentmemory/agentmemory demo
```

`demo` seeds 3 realistic sessions (JWT auth, N+1 query fix, rate limiting) and runs semantic searches against them. You'll see it find "N+1 query fix" when you search "database performance optimization" — keyword matching can't do that.

Open `http://localhost:3113` to watch the memory build live.

### Session Replay

Every session agentmemory records is replayable. Open the viewer, pick the **Replay** tab, and scrub through the timeline: prompts, tool calls, tool results, and responses render as discrete events with play/pause, speed control (0.5×–4×), and keyboard shortcuts (space to toggle, arrows to step).

Already have older Claude Code JSONL transcripts you want to bring in?

```
# Import everything under the default ~/.claude/projects
npx @agentmemory/agentmemory import-jsonl

# Or import a single file
npx @agentmemory/agentmemory import-jsonl ~/.claude/projects/-my-project/abc123.jsonl
```

Imported sessions show up in the Replay picker alongside native ones. Under the hood each entry routes through the `mem::replay::load`, `mem::replay::sessions`, and `mem::replay::import-jsonl` iii functions — no side-channel servers.

### Upgrade / Maintenance

Use the maintenance command when you intentionally want to update your local runtime:

```
npx @agentmemory/agentmemory upgrade
```

Warning: this command mutates the current workspace/runtime. It can update JavaScript dependencies, may run `cargo install iii-engine --force`, and may pull Docker images.

Implementation details live in `src/cli.ts` (see `runUpgrade` around the `src/cli.ts:544-595` region).

### Claude Code (one block, paste it)

```
Install agentmemory: run \`npx @agentmemory/agentmemory\` in a separate terminal to start the memory server. Then run \`/plugin marketplace add rohitg00/agentmemory\` and \`/plugin install agentmemory\` — the plugin registers all 12 hooks, 4 skills, AND auto-wires the \`@agentmemory/mcp\` stdio server via its \`.mcp.json\`, so you get 51 MCP tools (memory_smart_search, memory_save, memory_sessions, memory_governance_delete, etc.) without any extra config step. Verify with \`curl http://localhost:3111/agentmemory/health\`. The real-time viewer is at http://localhost:3113.
```

**OpenClaw (paste this prompt)**

```
Install agentmemory for OpenClaw. Run \`npx @agentmemory/agentmemory\` in a separate terminal to start the memory server on localhost:3111. Then add this to my OpenClaw MCP config so agentmemory is available with all 43 memory tools:

{
  "mcpServers": {
    "agentmemory": {
      "command": "npx",
      "args": ["-y", "@agentmemory/mcp"]
    }
  }
}

Restart OpenClaw. Verify with \`curl http://localhost:3111/agentmemory/health\`. Open http://localhost:3113 for the real-time viewer. For deeper memory-slot integration, copy \`integrations/openclaw\` to \`~/.openclaw/extensions/agentmemory\` and enable \`plugins.slots.memory = "agentmemory"\` in \`~/.openclaw/openclaw.json\`.
```

Full guide: [`integrations/openclaw/`](https://github.com/rohitg00/agentmemory/blob/main/integrations/openclaw)

**Hermes Agent (paste this prompt)**

```
Install agentmemory for Hermes. Run \`npx @agentmemory/agentmemory\` in a separate terminal to start the memory server on localhost:3111. Then add this to ~/.hermes/config.yaml so Hermes can use agentmemory as an MCP server with all 43 memory tools:

mcp_servers:
  agentmemory:
    command: npx
    args: ["-y", "@agentmemory/mcp"]

memory:
  provider: agentmemory

Verify with \`curl http://localhost:3111/agentmemory/health\`. Open http://localhost:3113 for the real-time viewer. For deeper 6-hook memory provider integration (pre-LLM context injection, turn capture, MEMORY.md mirroring, system prompt block), copy integrations/hermes from the agentmemory repo to ~/.hermes/plugins/agentmemory.
```

Full guide: [`integrations/hermes/`](https://github.com/rohitg00/agentmemory/blob/main/integrations/hermes)

### Other agents

Start the memory server: `npx @agentmemory/agentmemory`

Then add the MCP config for your agent:

| Agent | Setup |
| --- | --- |
| **Cursor** | Add to `~/.cursor/mcp.json`: `{"mcpServers": {"agentmemory": {"command": "npx", "args": ["-y", "@agentmemory/mcp"]}}}` |
| **OpenClaw** | Add to MCP config: `{"mcpServers": {"agentmemory": {"command": "npx", "args": ["-y", "@agentmemory/mcp"]}}}` or use the [memory plugin](https://github.com/rohitg00/agentmemory/blob/main/integrations/openclaw) |
| **Gemini CLI** | `gemini mcp add agentmemory npx -y @agentmemory/mcp --scope user` |
| **Codex CLI** | `codex mcp add agentmemory -- npx -y @agentmemory/mcp` or add `[mcp_servers.agentmemory]` to `.codex/config.toml` |
| **pi** | Copy [`integrations/pi`](https://github.com/rohitg00/agentmemory/blob/main/integrations/pi) to `~/.pi/agent/extensions/agentmemory` and restart pi |
| **OpenCode** | Add to `opencode.json`: `{"mcp": {"agentmemory": {"type": "local", "command": ["npx", "-y", "@agentmemory/mcp"], "enabled": true}}}` |
| **Hermes Agent** | Add to `~/.hermes/config.yaml` with `memory.provider: agentmemory` or use the [memory provider plugin](https://github.com/rohitg00/agentmemory/blob/main/integrations/hermes) |
| **Cline / Goose / Kilo Code** | Add MCP server in settings |
| **Claude Desktop** | Add to `claude_desktop_config.json`: `{"mcpServers": {"agentmemory": {"command": "npx", "args": ["-y", "@agentmemory/mcp"]}}}` |
| **Aider** | REST API: `curl -X POST http://localhost:3111/agentmemory/smart-search -d '{"query": "auth"}'` |
| **Any agent (32+)** | `npx skillkit install agentmemory` |

### From source

```
git clone https://github.com/rohitg00/agentmemory.git && cd agentmemory
npm install && npm run build && npm start
```

This starts agentmemory with a local `iii-engine` if `iii` is already installed, or falls back to Docker Compose if Docker is available. REST, streams, and the viewer bind to `127.0.0.1` by default.

Install `iii-engine` manually:

- **macOS / Linux:** `curl -fsSL https://install.iii.dev/iii/main/install.sh | sh`
- **Windows:** download `iii-x86_64-pc-windows-msvc.zip` from [iii-hq/iii releases](https://github.com/iii-hq/iii/releases/latest), extract `iii.exe`, add to PATH

Or use Docker (the bundled `docker-compose.yml` pulls `iiidev/iii:latest`). Full docs: [iii.dev/docs](https://iii.dev/docs).

### Windows

agentmemory runs on Windows 10/11, but the Node.js package alone isn't enough — you also need the `iii-engine` runtime (a separate native binary) as a background process. The official upstream installer is a `sh` script and there is no PowerShell installer or scoop/winget package today, so Windows users have two paths:

**Option A — Prebuilt Windows binary (recommended):**

```
# 1. Open https://github.com/iii-hq/iii/releases/latest in your browser
# 2. Download iii-x86_64-pc-windows-msvc.zip
#    (or iii-aarch64-pc-windows-msvc.zip if you're on an ARM machine)
# 3. Extract iii.exe somewhere on PATH, or place it at:
#    %USERPROFILE%\.local\bin\iii.exe
#    (agentmemory checks that location automatically)
# 4. Verify:
iii --version

# 5. Then run agentmemory as usual:
npx -y @agentmemory/agentmemory
```

**Option B — Docker Desktop:**

```
# 1. Install Docker Desktop for Windows
# 2. Start Docker Desktop and make sure the engine is running
# 3. Run agentmemory — it will auto-start the bundled compose file:
npx -y @agentmemory/agentmemory
```

**Option C — standalone MCP only (no engine):** if you only need the MCP tools for your agent and don't need the REST API, viewer, or cron jobs, skip the engine entirely:

```
npx -y @agentmemory/agentmemory mcp
# or via the shim package:
npx -y @agentmemory/mcp
```

**Diagnostics for Windows:** if `npx @agentmemory/agentmemory` fails, re-run with `--verbose` to see the actual engine stderr. Common failure modes:

| Symptom | Fix |
| --- | --- |
| `iii-engine process started` then `did not become ready within 15s` | Engine crashed on startup — re-run with `--verbose`, check stderr |
| `Could not start iii-engine` | Neither `iii.exe` nor Docker is installed. See Option A or B above |
| Port conflict | `netstat -ano \| findstr :3111` to see what's bound, then kill it or use `--port <N>` |
| Docker fallback skipped even though Docker is installed | Make sure Docker Desktop is actually running (system tray icon) |

> Note: there is no `cargo install iii-engine` — `iii` is not published to crates.io. The only supported install methods are the prebuilt binary above, the upstream `sh` install script (macOS/Linux only), and the Docker image.

---

Every coding agent forgets everything when the session ends. You waste the first 5 minutes of every session re-explaining your stack. agentmemory runs in the background and eliminates that entirely.

```
Session 1: "Add auth to the API"
  Agent writes code, runs tests, fixes bugs
  agentmemory silently captures every tool use
  Session ends -> observations compressed into structured memory

Session 2: "Now add rate limiting"
  Agent already knows:
    - Auth uses JWT middleware in src/middleware/auth.ts
    - Tests in test/auth.test.ts cover token validation
    - You chose jose over jsonwebtoken for Edge compatibility
  Zero re-explaining. Starts working immediately.
```

### vs built-in agent memory

Every AI coding agent ships with built-in memory — Claude Code has `MEMORY.md`, Cursor has notepads, Cline has memory bank. These work like sticky notes. agentmemory is the searchable database behind the sticky notes.

|  | Built-in (CLAUDE.md) | agentmemory |
| --- | --- | --- |
| Scale | 200-line cap | Unlimited |
| Search | Loads everything into context | BM25 + vector + graph (top-K only) |
| Token cost | 22K+ at 240 observations | ~1,900 tokens (92% less) |
| Cross-agent | Per-agent files | MCP + REST (any agent) |
| Coordination | None | Leases, signals, actions, routines |
| Observability | Read files manually | Real-time viewer on:3113 |

---

### Memory Pipeline

```
PostToolUse hook fires
  -> SHA-256 dedup (5min window)
  -> Privacy filter (strip secrets, API keys)
  -> Store raw observation
  -> LLM compress -> structured facts + concepts + narrative
  -> Vector embedding (6 providers + local)
  -> Index in BM25 + vector

Stop / SessionEnd hook fires
  -> Summarize session
  -> Knowledge graph extraction (if GRAPH_EXTRACTION_ENABLED=true)
  -> Slot reflection (if SLOT_REFLECT_ENABLED=true)

SessionStart hook fires
  -> Load project profile (top concepts, files, patterns)
  -> Hybrid search (BM25 + vector + graph)
  -> Token budget (default: 2000 tokens)
  -> Inject into conversation
```

### 4-Tier Memory Consolidation

Inspired by how human brains process memory — not unlike sleep consolidation.

| Tier | What | Analogy |
| --- | --- | --- |
| **Working** | Raw observations from tool use | Short-term memory |
| **Episodic** | Compressed session summaries | "What happened" |
| **Semantic** | Extracted facts and patterns | "What I know" |
| **Procedural** | Workflows and decision patterns | "How to do it" |

Memories decay over time (Ebbinghaus curve). Frequently accessed memories strengthen. Stale memories auto-evict. Contradictions are detected and resolved.

### What Gets Captured

| Hook | Captures |
| --- | --- |
| `SessionStart` | Project path, session ID |
| `UserPromptSubmit` | User prompts (privacy-filtered) |
| `PreToolUse` | File access patterns + enriched context |
| `PostToolUse` | Tool name, input, output |
| `PostToolUseFailure` | Error context |
| `PreCompact` | Re-injects memory before compaction |
| `SubagentStart/Stop` | Sub-agent lifecycle |
| `Stop` | End-of-session summary |
| `SessionEnd` | Session complete marker |

### Key Capabilities

| Capability | Description |
| --- | --- |
| **Automatic capture** | Every tool use recorded via hooks — zero manual effort |
| **Semantic search** | BM25 + vector + knowledge graph with RRF fusion |
| **Memory evolution** | Versioning, supersession, relationship graphs |
| **Auto-forgetting** | TTL expiry, contradiction detection, importance eviction |
| **Privacy first** | API keys, secrets, `<private>` tags stripped before storage |
| **Self-healing** | Circuit breaker, provider fallback chain, health monitoring |
| **Claude bridge** | Bi-directional sync with MEMORY.md |
| **Knowledge graph** | Entity extraction + BFS traversal |
| **Team memory** | Namespaced shared + private across team members |
| **Citation provenance** | Trace any memory back to source observations |
| **Git snapshots** | Version, rollback, and diff memory state |

---

Triple-stream retrieval combining three signals:

| Stream | What it does | When |
| --- | --- | --- |
| **BM25** | Stemmed keyword matching with synonym expansion | Always on |
| **Vector** | Cosine similarity over dense embeddings | Embedding provider configured |
| **Graph** | Knowledge graph traversal via entity matching | Entities detected in query |

Fused with Reciprocal Rank Fusion (RRF, k=60) and session-diversified (max 3 results per session).

### Embedding providers

agentmemory auto-detects your provider. For best results, install local embeddings (free):

```
npm install @xenova/transformers
```

| Provider | Model | Cost | Notes |
| --- | --- | --- | --- |
| **Local (recommended)** | `all-MiniLM-L6-v2` | Free | Offline, +8pp recall over BM25-only |
| Gemini | `text-embedding-004` | Free tier | 1500 RPM |
| OpenAI | `text-embedding-3-small` | $0.02/1M | Highest quality |
| Voyage AI | `voyage-code-3` | Paid | Optimized for code |
| Cohere | `embed-english-v3.0` | Free trial | General purpose |
| OpenRouter | Any model | Varies | Multi-model proxy |

---

51 tools, 6 resources, 3 prompts, and 4 skills — the most comprehensive MCP memory toolkit for any agent.

### 50 Tools

Core tools (always available)

| Tool | Description |
| --- | --- |
| `memory_recall` | Search past observations |
| `memory_compress_file` | Compress markdown files while preserving structure |
| `memory_save` | Save an insight, decision, or pattern |
| `memory_patterns` | Detect recurring patterns |
| `memory_smart_search` | Hybrid semantic + keyword search |
| `memory_file_history` | Past observations about specific files |
| `memory_sessions` | List recent sessions |
| `memory_timeline` | Chronological observations |
| `memory_profile` | Project profile (concepts, files, patterns) |
| `memory_export` | Export all memory data |
| `memory_relations` | Query relationship graph |

Extended tools (50 total — set AGENTMEMORY\_TOOLS=all)

| Tool | Description |
| --- | --- |
| `memory_patterns` | Detect recurring patterns |
| `memory_timeline` | Chronological observations |
| `memory_relations` | Query relationship graph |
| `memory_graph_query` | Knowledge graph traversal |
| `memory_consolidate` | Run 4-tier consolidation |
| `memory_claude_bridge_sync` | Sync with MEMORY.md |
| `memory_team_share` | Share with team members |
| `memory_team_feed` | Recent shared items |
| `memory_audit` | Audit trail of operations |
| `memory_governance_delete` | Delete with audit trail |
| `memory_snapshot_create` | Git-versioned snapshot |
| `memory_action_create` | Create work items with dependencies |
| `memory_action_update` | Update action status |
| `memory_frontier` | Unblocked actions ranked by priority |
| `memory_next` | Single most important next action |
| `memory_lease` | Exclusive action leases (multi-agent) |
| `memory_routine_run` | Instantiate workflow routines |
| `memory_signal_send` | Inter-agent messaging |
| `memory_signal_read` | Read messages with receipts |
| `memory_checkpoint` | External condition gates |
| `memory_mesh_sync` | P2P sync between instances |
| `memory_sentinel_create` | Event-driven watchers |
| `memory_sentinel_trigger` | Fire sentinels externally |
| `memory_sketch_create` | Ephemeral action graphs |
| `memory_sketch_promote` | Promote to permanent |
| `memory_crystallize` | Compact action chains |
| `memory_diagnose` | Health checks |
| `memory_heal` | Auto-fix stuck state |
| `memory_facet_tag` | Dimension:value tags |
| `memory_facet_query` | Query by facet tags |
| `memory_verify` | Trace provenance |

### 6 Resources · 3 Prompts · 4 Skills

| Type | Name | Description |
| --- | --- | --- |
| Resource | `agentmemory://status` | Health, session count, memory count |
| Resource | `agentmemory://project/{name}/profile` | Per-project intelligence |
| Resource | `agentmemory://memories/latest` | Latest 10 active memories |
| Resource | `agentmemory://graph/stats` | Knowledge graph statistics |
| Prompt | `recall_context` | Search + return context messages |
| Prompt | `session_handoff` | Handoff data between agents |
| Prompt | `detect_patterns` | Analyze recurring patterns |
| Skill | `/recall` | Search memory |
| Skill | `/remember` | Save to long-term memory |
| Skill | `/session-history` | Recent session summaries |
| Skill | `/forget` | Delete observations/sessions |

### Standalone MCP

Run without the full server — for any MCP client. Either of these works:

```
npx -y @agentmemory/agentmemory mcp   # canonical (always available)
npx -y @agentmemory/mcp                # shim package alias
```

Or add to your agent's MCP config:

Most agents (Cursor, Claude Desktop, Cline, etc.):

```
{
  "mcpServers": {
    "agentmemory": {
      "command": "npx",
      "args": ["-y", "@agentmemory/mcp"]
    }
  }
}
```

OpenCode (`opencode.json`):

```
{
  "mcp": {
    "agentmemory": {
      "type": "local",
      "command": ["npx", "-y", "@agentmemory/mcp"],
      "enabled": true
    }
  }
}
```

---

Auto-starts on port `3113`. Live observation stream, session explorer, memory browser, knowledge graph visualization, and health dashboard.

```
open http://localhost:3113
```

The viewer server binds to `127.0.0.1` by default. The REST-served `/agentmemory/viewer` endpoint follows the normal `AGENTMEMORY_SECRET` bearer-token rules. CSP headers use a per-response script nonce and disable inline handler attributes (`script-src-attr 'none'`).

### iii console — trace-level engine inspection

agentmemory runs on the [iii engine](https://iii.dev/), so the official [iii console](https://iii.dev/docs/console) gives you OpenTelemetry traces, the raw key/value state store, the stream monitor, and a direct function invoker for every piece of memory machinery. Use it to watch a `memory.search` call hit BM25 → embeddings → reranker in real time, replay a hook invocation, or poke individual functions without going through MCP.

[![iii console dashboard — system counters, application flow, registered triggers, live WebSocket status](https://github.com/rohitg00/agentmemory/raw/main/assets/iii-console/dashboard.png)](https://github.com/rohitg00/agentmemory/blob/main/assets/iii-console/dashboard.png)  
*Dashboard: functions, triggers, workers, streams, live flow graph. Screenshot from [iii.dev/docs/console](https://iii.dev/docs/console).*

**Install once:**

```
curl -fsSL https://install.iii.dev/console/main/install.sh | sh
```

**Launch alongside agentmemory:**

```
# The agentmemory viewer already holds port 3113, so run the console on 3114.
iii-console --port 3114 --engine-port 3111 --ws-port 3112
```

Then open `http://localhost:3114`.

**What you can do from the console:**

| Page | Use it to |
| --- | --- |
| **Functions** | Invoke any of agentmemory's ~33 functions directly with a JSON payload — handy for testing `memory.recall`, `memory.consolidate`, `graph.query` without wiring a client. |
| **Triggers** | Replay HTTP triggers (the agentmemory REST endpoints), fire the consolidation cron manually, or emit queue events. |
| **States** | Browse the KV store — sessions, memory slots, lifecycle timers, embeddings index — and edit values in place. |
| **Streams** | Watch live memory writes, hook events, and observation updates as they flow through iii's WebSocket stream. |
| **Traces** | OpenTelemetry waterfall / flame / service-breakdown views. Filter by `trace_id` to see exactly which functions, DB calls, and embedding requests a single `memory.search` produced. |
| **Logs** | Structured OTEL logs correlated to trace/span IDs. |

[![iii console trace waterfall view showing per-span duration](https://github.com/rohitg00/agentmemory/raw/main/assets/iii-console/traces-waterfall.png)](https://github.com/rohitg00/agentmemory/blob/main/assets/iii-console/traces-waterfall.png)  
*Traces: waterfall / flame / service breakdown for every memory operation.*

**Traces are already on:**

`iii-config.yaml` ships with the `iii-observability` worker enabled (`exporter: memory`, `sampling_ratio: 1.0`, metrics + logs). No extra config needed — the moment agentmemory starts, every memory operation emits a trace span and a structured log the console can read.

If you want to export to Jaeger/Honeycomb/Grafana Tempo instead, change `exporter: memory` to `exporter: otlp` and set the collector endpoint per iii's observability docs.

> **Heads-up:** no auth is enforced on the console itself — keep it bound to `127.0.0.1` (the default) and never expose it publicly.

---

### LLM Providers

agentmemory auto-detects from your environment. No API key needed if you have a Claude subscription.

| Provider | Config | Notes |
| --- | --- | --- |
| **No-op (default)** | No config needed | LLM-backed compress/summarize is DISABLED. Synthetic BM25 compression + recall still work. See `AGENTMEMORY_ALLOW_AGENT_SDK` below if you used to rely on the Claude-subscription fallback. |
| Anthropic API | `ANTHROPIC_API_KEY` | Per-token billing |
| MiniMax | `MINIMAX_API_KEY` | Anthropic-compatible |
| Gemini | `GEMINI_API_KEY` | Also enables embeddings |
| OpenRouter | `OPENROUTER_API_KEY` | Any model |
| Claude subscription fallback | `AGENTMEMORY_ALLOW_AGENT_SDK=true` | Opt-in only. Spawns `@anthropic-ai/claude-agent-sdk` sessions — used to cause unbounded Stop-hook recursion (#149 follow-up) so it is no longer the default. |

### Environment Variables

Create `~/.agentmemory/.env`:

```
# LLM provider (pick one — default is the no-op provider: no LLM calls)
# ANTHROPIC_API_KEY=sk-ant-...
# ANTHROPIC_BASE_URL=...              # Optional: Anthropic-compatible proxy / Azure
# GEMINI_API_KEY=...
# OPENROUTER_API_KEY=...
# MINIMAX_API_KEY=...
# Opt-in Claude-subscription fallback (spawns @anthropic-ai/claude-agent-sdk);
# leave OFF unless you understand the Stop-hook recursion risk (#149 follow-up):
# AGENTMEMORY_ALLOW_AGENT_SDK=true

# Embedding provider (auto-detected, or override)
# EMBEDDING_PROVIDER=local
# VOYAGE_API_KEY=...
# OPENAI_API_KEY=sk-...
# OPENAI_BASE_URL=https://api.openai.com   # Override for Azure / vLLM / LM Studio / proxies
# OPENAI_EMBEDDING_MODEL=text-embedding-3-small
# OPENAI_EMBEDDING_DIMENSIONS=1536        # Required when the model is not in the known-models table

# Search tuning
# BM25_WEIGHT=0.4
# VECTOR_WEIGHT=0.6
# TOKEN_BUDGET=2000

# Auth
# AGENTMEMORY_SECRET=your-secret

# Ports (defaults: 3111 API, 3113 viewer)
# III_REST_PORT=3111

# Features
# AGENTMEMORY_AUTO_COMPRESS=false  # OFF by default (#138). When on,
                                   # every PostToolUse hook calls your
                                   # LLM provider to compress the
                                   # observation — expect significant
                                   # token spend on active sessions.
# AGENTMEMORY_SLOTS=false          # OFF by default. Editable pinned
                                   # memory slots — persona,
                                   # user_preferences, tool_guidelines,
                                   # project_context, guidance,
                                   # pending_items, session_patterns,
                                   # self_notes. Size-limited; agent
                                   # edits via memory_slot_* tools.
                                   # Pinned slots addressable for
                                   # SessionStart injection.
# AGENTMEMORY_REFLECT=false        # OFF by default. Requires SLOTS=on.
                                   # Stop hook fires mem::slot-reflect:
                                   # scans recent observations, auto-
                                   # appends TODOs to pending_items,
                                   # counts patterns in
                                   # session_patterns, records touched
                                   # files in project_context. Fire-
                                   # and-forget; does not block.
# AGENTMEMORY_INJECT_CONTEXT=false # OFF by default (#143). When on:
                                   # - SessionStart may inject ~1-2K
                                   #   chars of project context into
                                   #   the first turn of each session
                                   #   (this is what actually reaches
                                   #   the model — Claude Code treats
                                   #   SessionStart stdout as context)
                                   # - PreToolUse fires /agentmemory/enrich
                                   #   on every file-touching tool call
                                   #   (resource cleanup, not a token
                                   #   fix — PreToolUse stdout is debug
                                   #   log only per Claude Code docs)
                                   # Observations are still captured via
                                   # PostToolUse regardless of this flag.
# GRAPH_EXTRACTION_ENABLED=false
# CONSOLIDATION_ENABLED=true
# LESSON_DECAY_ENABLED=true
# OBSIDIAN_AUTO_EXPORT=false
# AGENTMEMORY_EXPORT_ROOT=~/.agentmemory
# CLAUDE_MEMORY_BRIDGE=false
# SNAPSHOT_ENABLED=false

# Team
# TEAM_ID=
# USER_ID=
# TEAM_MODE=private

# Tool visibility: "core" (8 tools) or "all" (51 tools)
# AGENTMEMORY_TOOLS=core
```

---

107 endpoints on port `3111`. The REST API binds to `127.0.0.1` by default. Protected endpoints require `Authorization: Bearer <secret>` when `AGENTMEMORY_SECRET` is set, and mesh sync endpoints require `AGENTMEMORY_SECRET` on both peers.

Key endpoints

| Method | Path | Description |
| --- | --- | --- |
| `GET` | `/agentmemory/health` | Health check (always public) |
| `POST` | `/agentmemory/session/start` | Start session + get context |
| `POST` | `/agentmemory/session/end` | End session |
| `POST` | `/agentmemory/observe` | Capture observation |
| `POST` | `/agentmemory/smart-search` | Hybrid search |
| `POST` | `/agentmemory/context` | Generate context |
| `POST` | `/agentmemory/remember` | Save to long-term memory |
| `POST` | `/agentmemory/forget` | Delete observations |
| `POST` | `/agentmemory/enrich` | File context + memories + bugs |
| `GET` | `/agentmemory/profile` | Project profile |
| `GET` | `/agentmemory/export` | Export all data |
| `POST` | `/agentmemory/import` | Import from JSON |
| `POST` | `/agentmemory/graph/query` | Knowledge graph query |
| `POST` | `/agentmemory/team/share` | Share with team |
| `GET` | `/agentmemory/audit` | Audit trail |

Full endpoint list: [`src/triggers/api.ts`](https://github.com/rohitg00/agentmemory/blob/main/src/triggers/api.ts)

---

Built on [iii-engine](https://iii.dev/) 's three primitives — no Express, no Postgres, no Redis.

**118 source files · ~21,800 LOC · 800 tests · 123 functions · 34 KV scopes**

What iii-engine replaces

| Traditional stack | agentmemory uses |
| --- | --- |
| Express.js / Fastify | iii HTTP Triggers |
| SQLite / Postgres + pgvector | iii KV State + in-memory vector index |
| SSE / Socket.io | iii Streams (WebSocket) |
| pm2 / systemd | iii-engine worker management |
| Prometheus / Grafana | iii OTEL + health monitor |

```
npm run dev               # Hot reload
npm run build             # Production build
npm test                  # 800 tests (~1.7s)
npm run test:integration  # API tests (requires running services)
```

**Prerequisites:** Node.js >= 20, [iii-engine](https://iii.dev/docs) or Docker

[Apache-2.0](https://github.com/rohitg00/agentmemory/blob/main/LICENSE)