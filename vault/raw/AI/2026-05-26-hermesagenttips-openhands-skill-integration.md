---
author: HermesAgentTips
date: 2026-05-26
source: X (Twitter)
url: https://x.com/HermesAgentTips/status/2059047565878780060
metrics: likes=90, reposts=8, comments=7, views=9.0K
tags: [AI]
classification: llm
---

# Hermes can now orchestrate OpenHands as a built-in skill

Hermes Agent now supports OpenHands (formerly OpenDevin) as an installable skill, making it a model-agnostic autonomous coding agent backed by LiteLLM — works with ANY LLM provider (OpenAI, Anthropic, OpenRouter, DeepSeek, Qwen, Ollama, vLLM, Nous, 100+ models).

## Relevance

Directly relevant to the Hermes ecosystem I run. OpenHands sits alongside claude-code, codex, and opencode in the "delegation skills" family, but is the only one that's truly model-agnostic. Worth installing and testing — especially useful for running coding subagents without locking into a single provider.

## Key Details

- **Install:** `hermes skills install official/autonomous-ai-agents/openhands` + `uv tool install openhands --python 3.12`
- **Flags:** `--headless`, `--json`, `-t TEXT`, `--resume [ID] / --last`, `--override-with-envs`, `--exit-without-confirmation`
- **Pitfalls:** Use `OPENHANDS_SUPPRESS_BANNER=1` to avoid banner spam; `--override-with-envs` is mandatory in headless mode; resume IDs use dashed-UUID format.
- **Credit:** Built by Tim Koepsel (xzessmedia) + Hermes Agent. MIT licensed.