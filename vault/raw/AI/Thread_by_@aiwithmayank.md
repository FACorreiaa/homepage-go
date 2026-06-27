---
title: "Thread by @aiwithmayank"
source: "https://x.com/aiwithmayank/status/2046914454353510893?s=48"
author:
  - "[[@aiwithmayank]]"
published: 2026-04-22
created: 2026-04-25
description: "Holy shit…Karpathy dropped autoresearch and the internet rebuilt it 40 different ways in weeks. Someone just cataloged every single fork, p"
tags:
  - "clippings"
---
**Mayank Vora** @aiwithmayank [2026-04-22](https://x.com/aiwithmayank/status/2046914454353510893)

Holy shit…Karpathy dropped autoresearch and the internet rebuilt it 40 different ways in weeks.

Someone just cataloged every single fork, port, and descendant in one place.

Here's what the community built on top of it:

→ A macOS fork for Apple Silicon that runs the full loop on M-series chips

→ A Windows RTX version for consumer NVIDIA GPUs with VRAM floor configs

→ A WebGPU port that runs the entire experiment loop in your browser

→ A multi-GPU version with crash recovery and adaptive search strategy

→ A Colab/Kaggle T4 port for people who want to run it for free with zero local setup

Then it got stranger.

People started applying the loop to completely different domains.

→ A trading agent optimizing prompts against rolling Sharpe ratio instead of model loss

→ A genealogy researcher that iteratively expands and verifies family history

→ A Spring Boot service that grew from 119 lines to 950 in 5 autonomous cycles

The original idea was: give an AI a metric and let it self-improve until it wins.

Turns out that idea works on almost anything.

1.1k stars. 100% Opensource.

Repo:

---

**Mayank Vora** @aiwithmayank [2026-04-22](https://x.com/aiwithmayank/status/2046914466416366028)

If you made it this far, you'll love The Shift.

Every day, there's a new AI breakthrough changing how people work and live.

We break it down in under 5 minutes a day, so you don't.

Plus, get 3,000+ AI tools and free AI courses when you join.

Subscribe👇

---

**Sathish Harry** @SathishAiHype [2026-04-22](https://x.com/SathishAiHype/status/2046923214132244973)

Everyone's calling autoresearch AI doing research while you sleep.

The more precise framing: it's a ratchet.

Propose change → measure → keep if better → revert if not → repeat.

Git is the memory. The metric is the fitness function. The agent is the mutation operator.

That's

---

**Danny Shmueli** @dannyshmueli [2026-04-23](https://x.com/dannyshmueli/status/2047195218357674096)

Love this. The weak spot in auto-research is usually LLM-as-judge. I built a version that uses real web analytics to judge automatic A/B tests on auto-research results instead: https://github.com/Agent-Analytics/autoresearch-growth… Curious if anyone else has tried this.

![Image](https://pbs.twimg.com/media/HGkVczyWgAApTYh?format=jpg&name=large)

---

**Sergio Pereira** @sergiopreira [2026-04-23](https://x.com/sergiopreira/status/2047290710978682886)

Open source moves at the speed of developers who actually need the thing. Karpathy ships, the internet fork-bombs it in every direction. That's not chaos. That's distributed product-market fit testing.