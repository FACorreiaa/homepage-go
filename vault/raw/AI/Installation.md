---
title: "Installation"
source: "https://docs.claude-mem.ai/installation"
author:
published:
created: 2026-04-15
description: "Install Claude-Mem plugin for persistent memory across sessions"
tags:
  - "clippings"
---
## Installation Guide

## Quick Start

### Option 1: npx (Recommended)

Install and configure Claude-Mem with a single command:

```shellscript
npx claude-mem install
```

The interactive installer will:
- Detect your installed IDEs (Claude Code, Cursor, Gemini CLI, Windsurf, etc.)
- Copy plugin files to the correct locations
- Register the plugin with Claude Code
- Install all dependencies (including Bun and uv)
- Auto-start the worker service

### Option 2: Plugin Marketplace

Install Claude-Mem directly from the plugin marketplace inside Claude Code:

```shellscript
/plugin marketplace add thedotmack/claude-mem
/plugin install claude-mem
```

Both methods will automatically configure hooks and start the worker service. Start a new Claude Code session and you’ll see context from previous sessions automatically loaded.

> **Important:** Claude-Mem is published on npm, but running `npm install -g claude-mem` installs the **SDK/library only**. It does **not** register plugin hooks or start the worker service. Always install via `npx claude-mem install` or the `/plugin` commands above.

## System Requirements

- **Node.js**: 18.0.0 or higher
- **Claude Code**: Latest version with plugin support
- **Bun**: JavaScript runtime and process manager (auto-installed if missing)
- **SQLite 3**: For persistent storage (bundled)

## Advanced Installation

For development or testing, you can clone and build from source:

### Clone and Build

```shellscript
# Clone the repository
git clone https://github.com/thedotmack/claude-mem.git
cd claude-mem

# Install dependencies
npm install

# Build hooks and worker service
npm run build

# Worker service will auto-start on first Claude Code session
# Or manually start with:
npm run worker:start

# Verify worker is running
npm run worker:status
```

### Post-Installation Verification

#### 1\. Automatic Dependency Installation

Dependencies are installed automatically during plugin installation. The SessionStart hook also ensures dependencies are up-to-date on each session start (this is fast and idempotent). Works cross-platform on Windows, macOS, and Linux.

#### 2\. Verify Plugin Installation

Check that hooks are configured in Claude Code:

```shellscript
cat plugin/hooks/hooks.json
```

#### 3\. Data Directory Location

Data is stored in `~/.claude-mem/`:
- Database: `~/.claude-mem/claude-mem.db`
- PID file: `~/.claude-mem/.worker.pid`
- Port file: `~/.claude-mem/.worker.port`
- Logs: `~/.claude-mem/logs/worker-YYYY-MM-DD.log`
- Settings: `~/.claude-mem/settings.json`
Override with environment variable:

```shellscript
export CLAUDE_MEM_DATA_DIR=/custom/path
```

#### 4\. Check Worker Logs

```shellscript
npm run worker:logs
```

#### 5\. Test Context Retrieval

```shellscript
npm run test:context
```

## Upgrading

Upgrades are automatic when updating via the plugin marketplace. Key changes in recent versions: **v7.1.0**: PM2 replaced with native Bun process management. Migration is automatic on first hook trigger. **v7.0.0+**: 11 configuration settings, dual-tag privacy system. **v5.4.0+**: Skill-based search replaces MCP tools, saving ~2,250 tokens per session. See [CHANGELOG](https://github.com/thedotmack/claude-mem/blob/main/CHANGELOG.md) for complete version history.

## Next Steps

- [Getting Started Guide](https://docs.claude-mem.ai/usage/getting-started) - Learn how Claude-Mem works automatically
- [MCP Search Tools](https://docs.claude-mem.ai/usage/search-tools) - Query your project history
- [Configuration](https://docs.claude-mem.ai/configuration) - Customize Claude-Mem behavior