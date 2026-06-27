---
title: "Remote Server Env & Log Management Tools"
source: "https://grok.com/c/5790fd74-6234-4d07-8d48-00627f633119?rid=c68f4ee5-94bc-4686-a52f-bbfe40fac94d"
author:
published:
created: 2026-04-25
description: "Grok is an AI assistant built by xAI. Chat, create images, write code, and get real-time answers from the web and X."
tags:
  - "clippings"
---
Is there any service to manage env variables on remote machines? Like I have my server on Hetzner, could I use something to modify my env variables or to check logs

**Yes, there are several excellent services and tools for managing environment variables (env vars/secrets) on remote machines like your Hetzner VPS, as well as checking/viewing logs remotely.** These avoid manual SSH + editing files every time and follow better security practices (no plain.env files sitting on the server).

Here are the most practical options, grouped by use case:

### 1\. Dedicated Secrets/Env Var Managers (Best for injecting secrets securely)

These let you **store and edit env vars centrally** (via web dashboard or API), then **pull/inject them as real environment variables** into your apps on the remote server at runtime. No more.env files on disk.

- **Infisical** (strongly recommended – open-source + cloud option) Centralized dashboard for secrets, API keys, etc. On your Hetzner server: Install their lightweight CLI (infisical), then run apps like infisical run -- your-app-command. Secrets are injected as env vars automatically (no code changes needed). Supports multiple environments (dev/staging/prod), audit logs, access control, and self-hosting (you can even run it on the same Hetzner VPS). Great for teams or solo devs who want to modify vars centrally and have them update on restart.
- **Doppler** (SaaS, very developer-friendly) Similar workflow: Web UI to manage vars → CLI command doppler run -- your-app-command injects them live. Excellent for any Linux/VPS setup. Supports syncing, rotations, and integrations. No self-hosting, but very polished CLI.
- **EnvKey** (open-source option) Another solid encrypted config/secrets manager with similar CLI-based injection.

**How this works on Hetzner**: SSH once to set up the CLI + authentication, then use it in your systemd service, Docker entrypoint, or cron jobs. Modify vars in the web UI → restart your app/service → done. Much safer than editing files remotely.

### 2\. Self-Hosted PaaS for Full App + Env + Logs Management (Easiest if you use Docker)

If you run apps via Docker/Docker Compose (very common on Hetzner), **Coolify** is perfect:

- Install it with one command on your Hetzner VPS → get a beautiful web dashboard (like a self-hosted Heroku/Netlify).
- **Env vars**: Edit them directly in the UI (normal form or raw.env-style editor), per app and per environment. Coolify writes them securely and loads them into containers.
- **Logs**: Built-in real-time container logs, monitoring, and deployment logs — view/search everything from the browser without SSH.
- Git-based deploys, automatic SSL, reverse proxy, etc.
- Many Hetzner-specific tutorials exist; it runs great on a cheap CX/CPX VPS.

This turns your raw server into a managed platform where you handle env vars and logs centrally via web UI.

### 3\. Dedicated Log Management Services (For checking logs remotely)

- **Papertrail** (SaaS, super simple and recommended for VPS) Install a tiny agent (remote\_syslog2) or configure rsyslog on your Hetzner server once. It forwards app logs, text files, systemd logs, etc., to their cloud dashboard. Then tail/search live logs in the browser (or via CLI/API) — no SSH required. Fast, has a free plan, and works with any app/language.
- Other good ones: Datadog, Loggly, or self-hosted Graylog/ELK + Grafana Loki (if you want everything on your own hardware).

### Quick Recommendations Based on Your Setup

- **Just want env vars + occasional log checks** → Start with **Infisical** (or Doppler) + **Papertrail**. Minimal overhead.
- **Run Docker apps and want a full dashboard** → **Coolify** on your Hetzner VPS (covers both env vars *and* logs in one place).
- **Already using CI/CD (GitHub Actions, etc.)** → Most of these integrate natively.

All of these work great with Hetzner (cloud or dedicated) because they’re either agent/CLI-based or self-hosted. None require special Hetzner features.

If you tell me more details (e.g., are you using Docker? Node.js/Python? Single app or many? Team size?), I can give exact setup commands or the best starting point.

Docker, Docker Compose, Vapor monolith (Swift), dev and prod env variables for test flight and production, single dev

**For your setup (Docker + Docker Compose, Vapor Swift monolith, single dev, separate dev/prod envs including TestFlight/production), I recommend starting with Coolify as the primary tool.** It runs directly on your Hetzner VPS, gives you a nice web UI for managing everything, and handles Docker Compose natively while covering logs perfectly.

### Why Coolify Fits Perfectly

- **Docker Compose first-class support**: Point it at your docker-compose.yml (or docker-compose.production.yml). It treats the compose file as the source of truth but pulls in env vars from its UI.
- **Env vars management**: Define variables in your compose file like DATABASE\_URL=${DATABASE\_URL:?} (the:? makes them required). Coolify detects them automatically and shows a clean UI form (or raw editor) to set different values for dev vs prod deployments. You get separate environments/projects easily.
- **Dev vs Prod**: Create separate resources (or branches) in Coolify — one for dev (pointing to a dev branch), one for prod. Or use multiple compose files and switch via Coolify.
- **Logs**: Built-in real-time logs for all containers/services, searchable, with deployment history. No more constant SSH tailing.
- **Vapor/Swift compatibility**: Your Vapor app builds fine with a standard Dockerfile + Compose. Coolify handles the build/deploy cycle via Git (push → deploy).
- **Single dev friendly**: One-command install on Hetzner, free/open-source core, automatic HTTPS (Traefik), domains, backups, monitoring, etc. Runs great on cheap Hetzner CX/CPX instances.

**Installation on Hetzner (one command)**:

```markdown
curl -fsSL https://cdn.coollabs.io/coolify/install.sh | bash
```

Then access http://your-server-ip:8000, create admin user, and add your server (it auto-installs Docker if needed).

Connect your GitHub repo → choose **Docker Compose** build pack → specify path to your compose file.

### Handling Env Vars Securely in Coolify + Compose

In your docker-compose.yml (or production variant):

```markdown
services:
  vapor-app:
    build: .
    environment:
      - APP_ENV=${APP_ENV:?}           # Required, e.g. production or development
      - DATABASE_URL=${DATABASE_URL:?}
      - JWT_SECRET=${JWT_SECRET:?}
      - APP_KEY=${APP_KEY:?}           # Vapor-specific
      # ... other vars
    # ports, volumes, etc.
```

Coolify shows these in the UI for editing per deployment. Use shared variables across services if needed. For extra secrets security, combine with Docker secrets or external managers.

### Alternative/Complementary: Infisical (for Centralized Secrets)

If you want a dedicated secrets dashboard (highly recommended for production/TestFlight keys):

- Self-host Infisical on the same Hetzner VPS (easy Docker Compose setup).
- Use its CLI or machine identities to inject secrets at runtime into your containers.
- Works alongside Coolify: Manage vars centrally in Infisical → pull them into Coolify deployments.

**Doppler** is a great SaaS alternative if you prefer not to self-host the secrets manager.

### Logs Recommendation

- **Coolify built-in** → sufficient for most solo use (container logs, deployment logs).
- For advanced tailing/search/alerts: Add **Papertrail** or self-hosted **Loki + Grafana** (both integrate easily with Docker).

### Quick Workflow for You

1. Install Coolify on Hetzner.
2. Push your Vapor + Docker Compose repo to GitHub.
3. Deploy as Docker Compose resource in Coolify.
4. Set dev/prod env vars in the UI.
5. For TestFlight: Keep iOS-side config separate (e.g., via Xcode or CI), but backend API keys can live in Coolify/Infisical.
6. Deploy → view logs in browser.

This setup is lightweight, secure, and scales with you. No more manual SSH editing of.env files.

If you share your current docker-compose.yml structure (redacted) or Dockerfile, I can give exact config snippets for Coolify/Vapor. Or let me know if you want step-by-step install commands!

I already have my server running with two environments for the same API. Dev and Prod. root@ubuntu-4gb-fsn1-1:~# cd /opt/stockplan root@ubuntu-4gb-fsn1-1:/opt/stockplan# ls AGENTS.md Makefile Public Tests docker-compose.production.yml monitoring Dockerfile Package.resolved README.md docker-compose.dev.yml docker-compose.yml scripts Dockerfile.dev Package.swift Sources docker-compose.observability.yml docs root@ubuntu-4gb-fsn1-1:/opt/stockplan# nginx, certbot, etc, fully working. Give me the detailed steps I need to use coolify and Infisical on my Hetzner machine

**Detailed step-by-step guide to add Coolify + Infisical to your existing Hetzner setup (/opt/stockplan with docker-compose.dev.yml + docker-compose.production.yml, nginx + certbot).**

Your current setup is already solid, but Coolify will take over **deployment management, env vars UI, logs, and rebuilds** via Git + Docker Compose. Infisical will become the **central secure source for all dev/prod secrets** (no more env vars in compose files or.env on disk).

**Important warnings before you start:**

- Coolify installs its own Traefik reverse proxy (ports 80/443 by default) → **this will conflict with your existing nginx**. We will migrate the proxy to Traefik (recommended — it handles SSL automatically with Let’s Encrypt, just like certbot). Nginx can be stopped/re-purposed or removed later.
- Downtime of ~10–30 minutes during migration is expected.
- **Backup everything first** (see step 1).
- Coolify works best with a **Git repository** (your compose files become the source of truth). Local paths only work with some limitations, so we’ll use Git.

### Step 1: Backup (do this NOW)

```markdown
cd /opt/stockplan
tar -czf ~/stockplan-backup-$(date +%Y%m%d).tar.gz .
cp -r /etc/nginx /root/nginx-backup
cp -r /etc/letsencrypt /root/certbot-backup   # if using certbot
docker compose -f docker-compose.production.yml ps   # note running containers
```

### Step 2: Put your code in Git (required for smooth Coolify Docker Compose deploys)

If you don’t have a repo yet:

```markdown
cd /opt/stockplan
git init
git add .
git commit -m "Initial commit - stockplan with dev/prod compose"
# Create a private GitHub/GitLab repo and push
git remote add origin https://github.com/yourusername/stockplan.git
git push -u origin main
```

Keep docker-compose.dev.yml, docker-compose.production.yml, docker-compose.yml, Dockerfile, Dockerfile.dev, etc. in the repo.

(You can keep the local /opt/stockplan folder as a fallback.)

### Step 3: Install Coolify on your Hetzner server

Run as root:

```markdown
curl -fsSL https://cdn.coollabs.io/coolify/install.sh | sudo bash
```

- This installs Docker (if missing), creates /data/coolify/, sets up Traefik + Coolify.
- It will show a URL like http://YOUR-SERVER-IP:8000.

**Immediately** open that URL in your browser and create the first admin account (do this before anyone else can!).

### Step 4: First-time Coolify setup

1. Log in at http://YOUR-IP:8000.
2. Go to **Servers** → your server should already be added (localhost). If not, add it manually via SSH (it uses the root user).
3. (Optional but recommended) Go to **Settings → General** and set your domain if you want Coolify dashboard on HTTPS later.

**Handle nginx conflict now:**

```markdown
systemctl stop nginx
systemctl disable nginx   # optional — you can re-enable later if needed
```

Traefik (Coolify) will now take ports 80/443. Your existing certbot certs are not needed anymore — Coolify handles SSL.

### Step 5: Deploy your Stockplan API with Coolify (Docker Compose build pack)

We’ll create **two separate resources** (one for dev, one for prod) so you can manage them independently.

1. In Coolify → **Projects** → Create new project called Stockplan.
2. Inside the project → **Create New Resource** → **Application** (or Service if you prefer).
3. Choose **Git Repository** → paste your repo URL + select main branch (or develop for dev).
4. **Build Pack** → change from Nixpacks to **Docker Compose**.
5. Configure:
	- **Base Directory**: / (or wherever your compose files live)
		- **Docker Compose Location**: docker-compose.production.yml (for prod resource)
		- Branch: main for prod, or create a second resource with develop branch + docker-compose.dev.yml
6. Click **Continue** → Coolify will parse your compose file.

**Env vars setup (critical for your dev/prod):**

- In your docker-compose.production.yml (and dev variant), change env vars to this format:
	```markdown
	services:
	  vapor-app:   # or whatever your service name is
	    environment:
	      - APP_ENV=${APP_ENV:?}
	      - DATABASE_URL=${DATABASE_URL:?}
	      - JWT_SECRET=${JWT_SECRET:?}
	      - APP_KEY=${APP_KEY:?}
	      # ... all your Vapor vars
	```
	(The:? makes them required in Coolify UI.)
- Coolify will auto-detect them and show a nice form. Fill them in the UI for now (we’ll move them to Infisical in step 7).
1. Assign a domain to your Vapor service (e.g. api.stockplan.com or dev.api.stockplan.com).
	- Coolify/Traefik will auto-issue SSL.
2. Click **Deploy**. Coolify will:
	- Clone the repo
		- Run docker compose -f docker-compose.production.yml up -d --build
		- Show real-time logs

Repeat for the dev environment (second resource).

**Logs**: In Coolify, click your resource → **Logs** tab (real-time, searchable, no SSH needed).

### Step 6: Deploy Infisical inside Coolify (easiest way)

Coolify has a one-click template for Infisical.

1. In the same Stockplan project → **Create New Resource** → **Service** → search for **Infisical**.
2. Choose the official Infisical template (it sets up PostgreSQL + Redis + Infisical backend).
3. Assign a domain, e.g. secrets.stockplan.com.
4. Deploy it.

Alternatively (if no one-click shows), deploy manually as a Docker Compose service using the official file:

```markdown
# In Coolify you can also paste this as a raw compose service, but one-click is simpler.
```

After deployment, open https://secrets.stockplan.com and create your first admin account.

### Step 7: Move all secrets to Infisical (dev + prod)

1. In Infisical dashboard:
	- Create two projects: stockplan-dev and stockplan-prod.
		- Add all your env vars (DATABASE\_URL, JWT\_SECRET, etc.) separately for each environment.
2. (Recommended) Create **Machine Identities** (under Project Settings → Identities) for secure, tokenless access from your app.
3. Update your Vapor Docker Compose / Dockerfile to inject secrets at runtime:
	- Install Infisical CLI in your Dockerfile (or Dockerfile.dev):
		```markdown
		# Add near the end, before final CMD
		RUN curl -sSL https://infisical.com/install.sh | bash
		```
		- Change your entrypoint / command in compose to:
		```markdown
		command: >
		  sh -c "infisical run \
		    --project-id YOUR_PROJECT_ID \
		    --env production \   # or development
		    --token YOUR_MACHINE_TOKEN \
		    -- /app/.build/release/Run serve --env production --hostname 0.0.0.0"
		```
		(Adjust the Vapor run command to match your current one.)
	Or use infisical run wrapper in the service definition.
4. In Coolify, remove the actual secret **values** from the UI (keep only the keys with ${VAR:?}) — they will now come from Infisical at deploy time.
5. Redeploy both dev and prod resources.

### Step 8: Final cleanup & workflow

- Stop your old manual containers if still running:
	```markdown
	cd /opt/stockplan
	docker compose -f docker-compose.production.yml down
	```
- Update any monitoring/observability compose if you want Coolify to manage it too.
- New workflow:
	1. Edit secrets → Infisical UI (or CLI)
		2. Push code to GitHub → Coolify auto-deploys (or click Deploy)
		3. View logs, health, envs, backups all in Coolify browser UI
		4. For TestFlight → your prod backend is now at the domain you assigned, with secrets pulled securely.

You now have:

- Coolify managing deployments, logs, scaling, updates
- Infisical managing all secrets centrally (dev/prod separated)
- No more manual SSH editing of env files

If you hit any error (e.g. port conflict, compose parsing, or Vapor-specific command), paste the exact error here and I’ll give the fix. Also share your (redacted) docker-compose.production.yml if you want me to show the exact modified version with Infisical integration.

You’re all set — this is a huge upgrade for a single-dev Vapor setup! 🚀