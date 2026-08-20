---
title: "Latest changes"
summary: "What I have been shipping lately across North, Norviq, and the platform layer."
category: engineering
date: "2026-08-20"
---

The past few weeks have been mostly infra and plumbing. Nothing flashy, but the kind of work that makes everything else possible.

North is finally getting deployed. The app itself has been feature-complete for a while — it has a full Telegram bot, an MCP server, BYOK for AI keys, per-tier model chains. The missing piece was the deployment side: no ArgoCD app definition, no Helm values, no sealed secrets. That is what I am sorting out now. Once the sealed secret for the bot token is committed and ArgoCD picks it up, north will be live. The webhook registration with BotFather is the last manual step.

The bigger thing I set up is a base AI fallback chain that applies across all apps. The problem it solves: I keep opening dev environments and finding AI features broken because there is no key configured. The fix is a base OpenRouter API key baked into the platform secrets. When no BYOK key is set and no user-level key exists, the app falls back to OpenRouter with nvidia/nemotron-3-ultra-550b-a55b:free. It is slower than the paid models but it is always available and costs nothing. The chain goes: xAI Grok as primary, then OpenRouter with a capable paid model, then the nvidia free model as the floor. Free-tier users always land on the nvidia model. The key point is that AI should just work when you open the app, not fail silently.

Norviq is getting Telegram next. The architecture is different from north since the backend is Swift/Vapor rather than Go, but the idea is the same: a lifecycle handler that registers a webhook on boot, a route that validates the secret header and dispatches messages into the AI assistant pipeline. It is greenfield code in the backend, probably a few days of work.

All of this sits on the same foundation: k3s, ArgoCD, SealedSecrets, the same Helm chart pattern. Keeping the infra consistent across projects means each new thing is mostly config, not a new system to maintain.
