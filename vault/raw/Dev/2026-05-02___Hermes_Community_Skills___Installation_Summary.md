---
date: 2026-05-02
tags: [Dev, Hermes, Skills, Setup]
category: Hermes Community Skills
---

# Hermes Community Skills — Installation Summary

**Date:** 2026-05-02
**Installed by:** Hermes Agent (automated ingestion from GitTrend article)

---

## 📦 Skills Installed to `~/.hermes/skills/`

### 1. `hermes-novel-generator` (Full SKILL.md)
- **Source:** rwcrosk-arch/hermes-novel-generator
- **Type:** Multi-Agent novel generation pipeline
- **Dependencies:** `pyyaml` (✅ already in venv)
- **Status:** ⚠️ Requires LLM calls via `hermes_tools.delegate_task` — needs Hermes agent running
- **Use case:** Creative writing, EPUB output, story generation

### 2. `music-skill` (Full SKILL.md)
- **Source:** KaleLjl/music-skill
- **Type:** MIDI analysis, instrument replacement, MIDI→WAV rendering
- **Dependencies:** `mido`, `basic-pitch`, `pyfluidsynth` + system `fluidsynth`
- **Status:** ⚠️ Dependencies missing (pip blocked in venv)
- **Use case:** Music composition, audio rendering, MIDI manipulation

### 3. `community-backup` (Wrapper — NEW)
- **Wraps:** art-solutions/hermes-agent-backup-skill
- **Type:** Nightly automatic backup to private GitHub
- **Dependencies:** `git` (system), SSH keys or GitHub token
- **Status:** ✅ Ready for setup (scripts provided)
- **Location:** `~/.hermes/skills/community-backup/`
- **Setup script:** `~/.hermes/scripts/setup-hermes-backup.sh`

### 4. `community-telepath` (Wrapper — NEW)
- **Wraps:** eren23/telepath
- **Type:** Memory-aware visualization via Kimi K2
- **Dependencies:** `MOONSHOT_API_KEY`, `openai` Python package
- **Status:** ✅ Ready for setup (scripts provided)
- **Location:** `~/.hermes/skills/community-telepath/`
- **Setup script:** `~/.hermes/scripts/setup-telepath.sh`

### 5. `community-music-lite` (Wrapper — NEW)
- **Wraps:** KaleLjl/music-skill (full version)
- **Type:** Lightweight placeholder pointing to full install
- **Dependencies:** None (doc-only)
- **Status:** ✅ Ready — documentation skill
- **Location:** `~/.hermes/skills/community-music-lite/`

---

## 🚫 Could Not Install

| Repo | SKILL.md | Blocked Reason |
|---|---|---|
| `psy-core` (jethros-project) | ❌ | Private repo — authentication required |
| `hermes-agent-backup-skill` (raw) | ❌ | No SKILL.md — wrapped instead |
| `telepath` (raw) | ❌ | No SKILL.md — wrapped instead |

---

## 🛠️ Manual Setup Required (on Hermes server 78.46.192.73)

### Priority 1 — Backup Skill (Critical for StockPlan)

```bash
# On Hermes server:
chmod +x ~/.hermes/scripts/setup-hermes-backup.sh
~/.hermes/scripts/setup-hermes-backup.sh

# Then configure:
hermes skill configure community-backup
# Set BACKUP_DIR=~/hermes-backup
```

### Priority 2 — Telepath (Visualization)

```bash
# Get Moonshot API key from https://platform.moonshot.cn/
export MOONSHOT_API_KEY="sk-..."

# Run setup:
chmod +x ~/.hermes/scripts/setup-telepath.sh
~/.hermes/scripts/setup-telepath.sh
```

### Priority 3 — Full Music Skill (Optional)

```bash
# Only if you need MIDI/WAV rendering:
chmod +x ~/.hermes/scripts/setup-music-skill-full.sh
~/.hermes/scripts/setup-music-skill-full.sh
```

---

## 📂 Vault Locations

All source repos saved here:
```
Raw/Dev/2026-05-02 — GitHub - hermes-agent-backup-skill - README.md
Raw/Dev/2026-05-02 — GitHub - telepath - README.md
Raw/Dev/2026-05-02 — GitHub - hermes-novel-generator - README.md
Raw/Dev/2026-05-02 — GitHub - psy-core - placeholder.md  (blocked)
Raw/Dev/2026-05-02 — GitHub - music-skill - README.md
```

Compiled wiki: `wiki/` (auto-cross-linked)

---

## 🎯 Quick Reference

| Skill | Trigger in Hermes | What it does |
|---|---|---|
| `community-backup` | "backup now", "restore from backup" | Git-backed persistence |
| `community-telepath` | "visualize my time", "create diagram" | Memory → Mermaid charts |
| `community-music-lite` | "list midi files", "midi help" | Docs-only (install full for rendering) |
| `hermes-novel-generator` | "write a novel" | Multi-agent story generation |
| `music-skill` | (disabled — missing deps) | Full MIDI/WAV toolkit |

---

## 🔗 Related

- GitTrend source: `Raw/Tech/2026-05-02 — X - gittrend0x - 20505721.md`
- Atlas project index: `wiki/@kevin-simback-on-x.md`
- Hermes voice support incoming: `wiki/@teknium-on-x.md`

---

*Installed automatically from GitTrend X article — 2026-05-02*
