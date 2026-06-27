---
source: arxiv
ingested_at: 2026-05-10T09:15:18.492053Z
type: web
status: uncompiled
topic: AI
---
# Verifier-Backed Hard Problem Generation for Mathematical Reasoning

**Authors:** Yuhang Lai, Jiazhan Feng, Yee Whye Teh, Ning Miao

**Published:** 2026-05-07T17:58:32Z

**Summary:**
Large Language Models (LLMs) demonstrate strong capabilities for solving scientific and mathematical problems, yet they struggle to produce valid, challenging, and novel problems - an essential component for advancing LLM training and enabling autonomous scientific research. Existing problem generation approaches either depend on expensive human expert involvement or adopt naive self-play paradigms, which frequently yield invalid problems due to reward hacking. This work introduces VHG, a verifier-enhanced hard problem generation framework built upon three-party self-play. By integrating an independent verifier into the conventional setter-solver duality, our design constrains the setter's reward to be jointly determined by problem validity (evaluated by the verifier) and difficulty (assessed by the solver). We instantiate two verifier variants: a Hard symbolic verifier and a Soft LLM-based verifier, with evaluations conducted on indefinite integral tasks and general mathematical reasoning tasks. Experimental results show that VHG substantially outperforms all baseline methods by a clear margin.

**Link:** http://arxiv.org/abs/2605.06660v1
