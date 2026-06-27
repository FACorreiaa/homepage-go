---
source: https://medium.com/@jatinkrmalik/opencode-go-oh-my-openagent-the-complete-guide-to-sota-model-routing-without-hitting-limits-49fdc8cb3417
ingested_at: 2026-05-08T04:53:33Z
type: web
status: uncompiled
title: OpenCode Go + oh-my-openagent: The Complete Guide to SOTA Model Routing Without Hitting Limits!
---
Title: OpenCode Go + oh-my-openagent: The Complete Guide to SOTA Model Routing Without Hitting Limits!

URL Source: http://medium.com/@jatinkrmalik/opencode-go-oh-my-openagent-the-complete-guide-to-sota-model-routing-without-hitting-limits-49fdc8cb3417

Published Time: 2026-05-04T07:34:30Z

Markdown Content:
[![Image 1: Jatin K Malik](https://miro.medium.com/v2/resize:fill:32:32/2*9w8uaQp8uqA5wtRu-DltrQ.png)](https://medium.com/@jatinkrmalik?source=post_page---byline--49fdc8cb3417---------------------------------------)

9 min read

3 days ago

> **OpenCode Go** gives you access to 14 state-of-the-art open-source models like Kimi K2.6, DeepSeek V4 Pro, GLM-5.1, Qwen3.6 Plus, and more, for $10/month.
> 
> 
> Most subscribers route everything through a single model and burn through their limits in hours. Here’s how to use them correctly.

## What You Get with OpenCode Go

**For $10/month, you get access to:**

Press enter or click to view image in full size

![Image 2](https://miro.medium.com/v2/resize:fit:700/1*k9r8Os32l41UdpZn4BwySQ.png)

**Reference frontier models (for context):**

*   **Claude Opus 4.7**: 87.6% SWE-Verified, 64.3% SWE-Pro
*   **GPT-5.4**: 57.7% SWE-Pro
*   **Gemini 3.1 Pro**: 54.2% SWE-Pro

On SWE-Bench Pro (the hardest real-world engineering benchmark), Ki**mi K2.6** at 58.6% and **GLM-5.1** at 58.4% are within 6 points of **Opus 4.7** (64.3%). **DeepSeek V4 Pro** at 55.4% beats **Gemini 3.1 Pro** (54.2%).

> **T****h****e trade-off**: You’re getting ~80–90% of frontier quality at ~10–20x lower cost. For most production work, that’s the right trade-off.

## The Problem: Most People Use One Model for Everything

Press enter or click to view image in full size

![Image 3](https://miro.medium.com/v2/resize:fit:700/1*QwjyojMoJBb3rXduSytj1Q.png)

[https://opencode.ai/go](https://opencode.ai/go)

FYI, OpenCode Go limits are dollar-based, not token-based:

*   5-hour window: $12 of usage
*   Weekly: $30
*   Monthly: $60

A single agentic coding session **burns 50–200 requests** (one per tool call, file edit, shell command). If you route your main orchestrator through Kimi K2.6 for everything, you’ll hit the 1,150 request/5hr limit in 2–3 heavy sessions.

**The fix:** tiered routing with `fallback_models` chains. Use the right model for the right complexity.

## Step 1: Install OpenCode + Oh My Open Agent

## Prerequisites

*   OpenCode 1.0.133+ (older versions have a config bug)
*   Git, Node.js or Bun
*   macOS, Linux, or WSL

## Install OpenCode

> [https://opencode.ai/](https://opencode.ai/)

curl -fsSL https://opencode.ai/install | bash

opencode --version 
## Install Oh My Open Agent

> [https://github.com/code-yeongyu/oh-my-openagent](https://github.com/code-yeongyu/oh-my-openagent)

bunx oh-my-opencode install
For non-interactive setup (LLM agents, CI):

bunx oh-my-opencode install --no-tui \

 --opencode-go=yes \

 --opencode-zen=yes \

 --claude=no \

 --openai=no \

 --gemini=no \

 --copilot=no
## Verify Installation

opencode --version 

bunx oh-my-opencode doctor 

opencode models 

opencode auth login 

## Step 2: Understand the Agents

Oh My Open Agent provides 11 built-in agents, each designed for a specific task. The key insight: each agent has a different optimal model based on its job.

## Orchestrators & Planners

Press enter or click to view image in full size

![Image 4](https://miro.medium.com/v2/resize:fit:1000/1*RXvow1AGKJ1E6LOVb7eaTA.png)

## Deep Workers

Press enter or click to view image in full size

![Image 5](https://miro.medium.com/v2/resize:fit:1000/1*AT0iQCsjWPSCUTaaxAO6fg.png)

## Utility & Search Agents

Press enter or click to view image in full size

![Image 6](https://miro.medium.com/v2/resize:fit:1000/1*-F_RmA9RUb9-nE-ClKoZrQ.png)

## Step 3: The Tiered Model Architecture

> **The core principle:**match model capability to task complexity.

Press enter or click to view image in full size

![Image 7](https://miro.medium.com/v2/resize:fit:700/1*3XAmpyAs899I31wMwrokPg.png)

**Don’t burn your best model on trivial tasks!**

## Tier 1 // Volume Workhorse (Never Rate-Limited)

**_Models: DeepSeek V4 Flash, Qwen3.5 Plus, MiniMax M2.5_**

Use for:

*   Code completion and autocomplete
*   Simple bug fixes (1–2 files)
*   Code review and PR feedback
*   Explore / Librarian searches
*   Quick tasks and Sisyphus-Junior executions

Why: DeepSeek V4 Flash gives you 31,650 requests per 5 hours, you’ll never exhaust it. At 79% SWE-Verified, it’s good enough for 80% of coding tasks.

## Tier 2 // Standard Engineering (Balanced)

**_Models: DeepSeek V4 Pro, Qwen3.6 Plus, MiniMax M2.7_**

Use for:

*   Feature implementation (3–5 files)
*   Terminal-heavy automation
*   Multi-step debugging
*   Standard agentic workflows

Why: 3,300–3,450 requests per 5 hours. V4 Pro leads LiveCodeBench at 93.5%. Qwen3.6 Plus leads Terminal-Bench at 61.6%.

## Tier 3 // Complex Agentic (Elite)

**_Models: Kimi K2.6, GLM-5.1, MiMo-V2.5-Pro_**

Use for:

*   Multi-file refactoring (10+ files)
*   Long-horizon autonomous runs (4+ hours)
*   Architecture decisions
*   300-agent swarm coordination

Why: These score 57–59% on SWE-Bench Pro, the hardest real-world benchmark. But limits are tight (880–1,290 req/5hr). Use deliberately.

## Tier 4 // Specialized Capabilities

**_Models: MiMo-V2-Omni (multimodal), GLM-5.1 (long-horizon reasoning)_**

## Get Jatin K Malik’s stories in your inbox

Join Medium for free to get updates from this writer.

Remember me for faster sign in

Use for:

*   Screenshot-to-code (MiMo-V2-Omni)
*   8-hour autonomous engineering runs (GLM-5.1)
*   Spec-writing and architectural planning (GLM-5.1)

## Step 4: Configure Rate Limit Resilience

Press enter or click to view image in full size

![Image 8](https://miro.medium.com/v2/resize:fit:700/1*ueVBy4aSTSoAogvURT5xog.png)

Enable `model_fallback` and set up fallback chains so rate limits don't break your workflow:

{

 "model_fallback": true,

 "runtime_fallback": {

 "enabled": true,

 "retry_on_errors": [400, 429, 503, 529],

 "max_fallback_attempts": 3,

 "cooldown_seconds": 60,

 "timeout_seconds": 30,

 "notify_on_fallback": true

 }

}
When a provider hits rate limits, it’s globally blacklisted for `cooldown_seconds`. All new sessions skip blacklisted providers until cooldown expires.

## Step 5: The Complete Configuration

Save this to `~/.config/opencode/oh-my-openagent.json`:

{

 "$schema": "https://raw.githubusercontent.com/code-yeongyu/oh-my-openagent/dev/assets/oh-my-opencode.schema.json",

 "model_fallback": true,

 "runtime_fallback": {

 "enabled": true,

 "retry_on_errors": [400, 429, 503, 529],

 "max_fallback_attempts": 3,

 "cooldown_seconds": 60,

 "timeout_seconds": 30,

 "notify_on_fallback": true

 },

 "agents": {

 "sisyphus": {

 "model": "opencode-go/kimi-k2.6",

 "fallback_models": [

 "opencode-go/deepseek-v4-pro",

 "opencode-go/qwen3.6-plus"

 ],

 "ultrawork": {

 "model": "opencode-go/kimi-k2.6",

 "variant": "max"

 }

 },

 "hephaestus": {

 "model": "opencode-go/deepseek-v4-pro",

 "fallback_models": [

 "opencode-go/deepseek-v4-flash",

 "opencode-go/kimi-k2.6"

 ],

 "prompt_append": "Explore thoroughly, then implement. Prefer small, testable changes."

 },

 "oracle": {

 "model": "opencode-go/glm-5.1",

 "fallback_models": [

 "opencode-go/kimi-k2.6",

 "opencode-go/deepseek-v4-pro"

 ]

 },

 "librarian": {

 "model": "opencode-go/deepseek-v4-flash",

 "fallback_models": "opencode-go/qwen3.5-plus"

 },

 "explore": {

 "model": "opencode-go/deepseek-v4-flash"

 },

 "multimodal-looker": {

 "model": "opencode-go/mimo-v2-omni",

 "fallback_models": "opencode-go/qwen3.6-plus"

 },

 "prometheus": {

 "model": "opencode-go/glm-5.1",

 "fallback_models": [

 "opencode-go/qwen3.6-plus",

 "opencode-go/deepseek-v4-pro"

 ],

 "prompt_append": "Always interview first. Validate scope before planning."

 },

 "metis": {

 "model": "opencode-go/qwen3.6-plus",

 "fallback_models": "opencode-go/deepseek-v4-pro"

 },

 "momus": {

 "model": "opencode-go/qwen3.6-plus",

 "fallback_models": "opencode-go/kimi-k2.6"

 },

 "atlas": {

 "model": "opencode-go/deepseek-v4-pro",

 "fallback_models": "opencode-go/deepseek-v4-flash"

 },

 "code-reviewer": {

 "model": "opencode-go/kimi-k2.6",

 "fallback_models": "opencode-go/deepseek-v4-pro"

 },

 "sisyphus-junior": {

 "model": "opencode-go/deepseek-v4-flash"

 }

 },

 "categories": {

 "visual-engineering": {

 "model": "opencode-go/mimo-v2-omni",

 "fallback_models": "opencode-go/qwen3.6-plus"

 },

 "ultrabrain": {

 "model": "opencode-go/glm-5.1",

 "fallback_models": "opencode-go/kimi-k2.6"

 },

 "deep": {

 "model": "opencode-go/kimi-k2.6",

 "fallback_models": "opencode-go/deepseek-v4-pro"

 },

 "artistry": {

 "model": "opencode-go/glm-5.1",

 "fallback_models": "opencode-go/qwen3.6-plus"

 },

 "quick": {

 "model": "opencode-go/deepseek-v4-flash"

 },

 "unspecified-low": {

 "model": "opencode-go/deepseek-v4-flash"

 },

 "unspecified-high": {

 "model": "opencode-go/deepseek-v4-pro",

 "fallback_models": "opencode-go/kimi-k2.6"

 },

 "writing": {

 "model": "opencode-go/qwen3.6-plus"

 }

 },

 "background_task": {

 "defaultConcurrency": 5,

 "staleTimeoutMs": 180000,

 "providerConcurrency": {

 "opencode-go": 10

 },

 "modelConcurrency": {

 "opencode-go/kimi-k2.6": 2,

 "opencode-go/deepseek-v4-pro": 3,

 "opencode-go/deepseek-v4-flash": 20,

 "opencode-go/glm-5.1": 1,

 "opencode-go/qwen3.6-plus": 5

 }

 },

 "git_master": {

 "include_co_authored_by": false

 }

}
## Configuration Breakdown

### Sisyphus (Main Orchestrator)

*   Primary: `opencode-go/kimi-k2.6` — best agentic coder (58.6% SWE-Pro)
*   Fallback 1: `opencode-go/deepseek-v4-pro` — when K2.6 hits limits
*   Fallback 2: `opencode-go/qwen3.6-plus` — generous limits, good reasoning
*   Ultrawork mode also uses K2.6 for background parallel execution

### Hephaestus (Autonomous Deep Worker)

*   Primary: `opencode-go/deepseek-v4-pro` — principle-driven, GPT-like
*   Fallback 1: `opencode-go/deepseek-v4-flash` — cheap, fast
*   Fallback 2: `opencode-go/kimi-k2.6` — if both DeepSeek exhausted

### Oracle (Architecture Consultant)

*   Primary: `opencode-go/glm-5.1` — best reasoning among Go models, 8-hour autonomous runs
*   Fallback: `opencode-go/kimi-k2.6` — second-best agentic performance
*   Note: Without Opus/GPT models, Oracle is your weakest link. GLM-5.1 at 58.4% SWE-Pro is solid but trails Opus 4.7 (64.3%).

### Explore / Librarian (Search Agents)

*   Both use `opencode-go/deepseek-v4-flash` — 31,650 req/5hr, never rate-limited
*   These are speed agents. Don’t waste expensive models here!

### Prometheus (Planner)

*   Primary: `opencode-go/glm-5.1` — spec-writing, architectural planning
*   Fallback: `opencode-go/qwen3.6-plus` — good reasoning, generous limits
*   Planning consumes few requests, so tight GLM-5.1 limits (880/5hr) are acceptable

### Metis / Momus (Review Agents)

*   Both use `opencode-go/qwen3.6-plus` — 3,300 req/5hr, analytical
*   Momus falls back to Kimi K2.6 for critical reviews

### Code-reviewer (Quality Review)

*   Primary: `opencode-go/kimi-k2.6` — highest SWE-Pro among Go models (58.6%)
*   Fallback: `opencode-go/deepseek-v4-pro`

### Categories

*   `deep`: K2.6 → V4 Pro (complex research + execution)
*   `quick`: V4 Flash (trivial tasks)
*   `unspecified-low`: V4 Flash (moderate tasks)
*   `unspecified-high`: V4 Pro → K2.6 (complex work)
*   `writing`: Qwen3.6 Plus (instruction following)
*   `visual-enginee`[https://github.com/code-yeongyu/oh-my-openagent](https://github.com/code-yeongyu/oh-my-openagent)`ring`: MiMo-V2-Omni → Qwen3.6 Plus (multimodal → text fallback)
*   `ultrabrain`: GLM-5.1 → K2.6 (maximum reasoning available on Go)

## Step 6: Rate Limit Budgeting

## Per 5-Hour Window

Press enter or click to view image in full size

![Image 9](https://miro.medium.com/v2/resize:fit:700/1*I05D2hkOh_6KY9RrvwaKag.png)

## Session Cost Estimates

Press enter or click to view image in full size

![Image 10](https://miro.medium.com/v2/resize:fit:700/1*JmlQZ2795VGMkykd98_5pQ.png)

> **Rule of thumb:** If a task will take more than 100 requests, route it through V4 Flash first. Escalate to K2.6 or V4 Pro only if V4 Flash gets stuck.

## Step 7: The Quality Trade-Offs

Without Copilot or other frontier providers, you’re working with a pure OpenCode Go stack.

### Here’s where you win and where you compromise:

## Where You Win ✔

*   **Cost**: Save $50–200/month for frontier API access
*   **LiveCodeBench**: DeepSeek V4 Pro at 93.5% beats every frontier model on competitive programming
*   **Terminal-Bench:** Qwen3.6 Plus at 61.6% beats Claude 4.5 (59.3%) on agentic terminal work
*   **Token efficiency**: MiMo-V2.5-Pro uses 40–60% fewer tokens than Claude at comparable capability
*   **Rate limits**: DeepSeek V4 Flash at 31K req/5hr is effectively unlimited

## Where You Compromise ⁉

*   **SWE-Bench Verified**: Your best is 80.6% (V4 Pro) vs Opus 4.7 at 87.6%. That’s a 7-point gap on real-world bug fixes.
*   **Oracle quality**: GLM-5.1 is excellent (58.4% SWE-Pro, 8-hour runs) but not Opus 4.7. For architecture decisions on unfamiliar systems, expect ~15–20% more iterations.
*   **Multimodal**: MiMo-V2-Omni is capable but trails Gemini 3.1 Pro on vision-heavy tasks.

> V**_erdict_**: **_For 80% of coding tasks, the gap is invisible._**
> 
> 
> **For the remaining 20%**(complex architecture, cutting-edge security, maximum accuracy), you’ll need 1–2 extra iterations. The cost savings ($100–200/month) more than compensate.

## Step 8: Validate Your Setup

bunx oh-my-opencode doctor --verbose
bunx oh-my-opencode refresh-model-capabilities

cd your-project

opencode

Press enter or click to view image in full size

![Image 11](https://miro.medium.com/v2/resize:fit:700/1*WiQYWGiqzIuZHkBHREdH1w.png)

opencode + oh-my-openagent = ❤️

## Key Principles

1.   **Don’t route everything through one model.** K2.6 is great but 1,150 req/5hr is tight. Use V4 Flash for volume, V4 Pro for standard work, K2.6/GLM-5.1 for the hardest tasks.
2.   **Don’t “upgrade” utility agents.** Explore and Librarian need speed, not intelligence. V4 Flash at 31K req/5hr is perfect.
3.   **Use**`fallback_models`**everywhere**. Rate limits are not failures — they're signals to switch models.
4.   **Keep concurrency conservative.** Start at 2–3 parallel tasks, run cleanly, then increase. The `background_task.modelConcurrency` settings above are tuned for your limits.
5.   **When K2.6 (or some other top model) has a 3x promo (as it does this week), abuse it.** Route Sisyphus, Hephaestus, and Deep category through K2.6. When the promo ends, revert Hephaestus to V4 Pro primary.
6.   **GLM-5.1 is your Oracle substitute.**It has the tightest limits (880/5hr) but the best reasoning among Go models. Reserve it for planning, architecture, and long-horizon tasks.
7.   **Qwen3.6 Plus is your Swiss Army knife.** 3,300 req/5hr, 61.6% Terminal-Bench, good at everything. Use it for Metis, Momus, Writing, and as a universal fallback.

## Resources for you to read:

*   OpenCode Go Docs: [https://opencode.ai/docs/go/](https://opencode.ai/docs/go/)
*   Oh My Open Agent GitHub: [https://github.com/code-yeongyu/oh-my-openagent](https://github.com/code-yeongyu/oh-my-openagent)
*   Oh My Open Agent Docs: [https://ohmyopenagent.com/en/docs](https://ohmyopenagent.com/en/docs)
*   Config Schema: [https://raw.githubusercontent.com/code-yeongyu/oh-my-openagent/dev/assets/oh-my-opencode.schema.json](https://raw.githubusercontent.com/code-yeongyu/oh-my-openagent/dev/assets/oh-my-opencode.schema.json)
*   Agent-Model Matching: [https://github.com/code-yeongyu/oh-my-openagent/blob/dev/docs/guide/agent-model-matching.md](https://github.com/code-yeongyu/oh-my-openagent/blob/dev/docs/guide/agent-model-matching.md)
*   Configuration Reference: [https://github.com/code-yeongyu/oh-my-openagent/blob/dev/docs/reference/configuration.md](https://github.com/code-yeongyu/oh-my-openagent/blob/dev/docs/reference/configuration.md)
