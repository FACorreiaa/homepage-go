---
source: arxiv
ingested_at: 2026-05-26T09:17:30Z
type: note
status: uncompiled
category: AI
subcategory: arxiv
title: Confidence and Calibration of Activation Oracles for Reliable Interpretation of Language Model Internals
arxiv_id: 2605.26045
---

## Confidence and Calibration of Activation Oracles for Reliable Interpretation of Language Model Internals

**Authors:** Federico Torrielli, Peter Schneider-Kamp, Lukas Galke Poech
**Published:** 2026-05-25
**Categories:** cs.CL, cs.AI
**URL:** https://arxiv.org/abs/2605.26045v1

### Abstract
Activation oracles aim to make the activations of other models legible to humans and yield promising results compared to white-box interpretability techniques. However, uncertainty quantification (UQ) for the natural-language outputs of such activation oracles is so far understudied. Here, we investigate 6 different methods for estimating the confidence of activation oracles and evaluate how well-calibrated their confidence scores are. Our experiments on 6,000 samples per oracle (varying verbalizer and context prompts) reveal that bootstrap mode frequency is the best-calibrated method among those tested (ECE 5.7% vs. 25.5% for the answer-word log-probability on Qwen3-8B; 10.3% vs. 13.1% on Qwen3.6-27B), and that the log-prob baseline can serve as a fast triage signal at a fraction of the cost.
  Code and the patched trainer are available at https://github.com/federicotorrielli/probabilistic_activation_oracles.
