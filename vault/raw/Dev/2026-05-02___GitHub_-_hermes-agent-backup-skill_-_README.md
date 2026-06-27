---
source: GitHub
url: https://github.com/art-solutions/hermes-agent-backup-skill
date: 2026-05-02
tags: [Dev, Hermes, Skill]
project: hermes-agent-backup-skill
description: Automatically backs up Hermes Agent memory, skills, and configs to a private GitHub repo every night
---

# hermes-agent-backup-skill

# hermes-agent-core-backup

**Community skill for the Hermes Agent ecosystem**

Automated daily backup for your Hermes agent's entire brain — memories, skills, config, scripts, and secrets — pushed to a private git repository with zero manual effort.

---

## Why This Skill Matters

Your Hermes agent is more than software. Over time it accumulates:

- **Memories** — learned preferences, long-term context, behavioral rules
- **Skills** — custom workflows you built and tuned
- **State** — the SQLite database tracking everything the agent has done
- **Config & secrets** — API keys, integrations, environment variables

None of this lives in the Hermes installer. If your server crashes, your drive fails, or you move to a new machine — **everything is gone** unless you have a backup.

This skill solves that with a single shell script and one cron line. Once installed, backups run silently every night. Restore takes three commands.

### What you lose without it

| Event | Without backup | With this skill |
|-------|---------------|-----------------|
| Server crash | All memories, skills, state lost permanently | Restore from last night's backup |
| New machine / migration | Rebuild from scratch | Clone repo, unzip, done |
| Bad config change | No rollback possible | Restore yesterday's config |
| Accidental skill deletion | Gone | Pull from archive |

---

## Files in This Repository

| File | Purpose |
|------|---------|
| `hermes-agent-core-backup.md` | The skill file — install this into your Hermes agent |
| `README.md` | This document |

---

## How to Install the Skill

Copy `hermes-agent-core-backup.md` into your Hermes skills directory:

```bash
cp hermes-agent-core-backup.md ~/.hermes/skills/hermes-agent-core-backup.md
```

That's it. Hermes will discover the skill automatically on the next session start.

---

## How to Use It — Prompt Your Hermes Agent

Once the skill file is in place, give your agent one of these prompts depending on what you need:

---

### First-time setup (automated daily backup)

```
Use the hermes-agent-core-backup skill to set up automated daily backups for
my Hermes installation.

My details:
- Agent name: YOUR_NAME
- Hermes directory: ~/.hermes  (or your custom path)
- Local backup staging: ~/hermes_backups
- Remote backup repo: git@github.com:YOUR_USER/hermes_backups.git  (PRIVATE)
- Data directory: ~/.hermes/data  (or remove if not needed)

Create the backup script, make it executable, and add the cron job to run at
01:00 every night. Then run the script once so I can verify it works.
```

---

### Manual backup before a big change

```
Use the hermes-agent-core-backup skill to run a manual backup right now.
I'm about to make major changes and want a snapshot first.
```

---

### Restore after failure or migration

```
Use the hermes-agent-core-backup skill to restore my agent from backup.

My backup repo: git@github.com:YOUR_USER/hermes_backups.git
Target backup date: YYYY_MM_DD  (or "latest")

Stop any running Hermes processes before restoring, then walk me through
the full restore procedure.
```

---

### Verify backups are running

```
Use the hermes-agent-core-backup skill to check that my daily backups are
working correctly. Check the cron job, the last backup log, and confirm
the latest ZIP exists in the remote repo.
```

---

## Security Note

The backup archive includes `.hermes/.env` (your API keys). This is intentional — keys must survive a restore. For this reason:

- The backup repository **must be private**
- Never share the repo URL or ZIP files publicly
- If using GitHub, enable "Require authentication" on the repo

---

## Requirements

- Linux / macOS server running Hermes
- `git`, `zip`, `unzip` available on the system
- SSH key authorized on your backup repository host
- A **private** GitHub, GitLab, or Gitea repository

---

## Community

This skill was created for the Hermes community so every user can protect their agent with a reliable, standardized backup system.

If you improve the script or add support for S3 / rclone / other remotes, please share it back.
  # Limit README size
