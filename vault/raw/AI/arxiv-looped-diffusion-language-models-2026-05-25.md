---
source: arxiv
ingested_at: 2026-05-26T09:17:30Z
type: note
status: uncompiled
category: AI
subcategory: arxiv
title: Looped Diffusion Language Models
arxiv_id: 2605.26106
---

## Looped Diffusion Language Models

**Authors:** Sanghyun Lee, Chunsan Hong, Seungryong Kim, Jonghyun Lee et al.
**Published:** 2026-05-25
**Categories:** cs.LG
**URL:** https://arxiv.org/abs/2605.26106v1

### Abstract
Masked diffusion models (MDMs) have emerged as a promising alternative to autoregressive models for language modeling, yet the effective design of transformer architectures for MDMs remains underexplored. In this paper, we show that selectively looping the early-middle transformer layers significantly improves both training efficiency and model performance in MDMs. We call this approach LoopMDM(Looped Masked Diffusion Model), which brings two key benefits: looping layers at training-time yields a depth-scaling effect without adding parameters, while varying the number of loops at inference-time enables flexible compute scaling. Despite the simplicity, the results are striking: across multiple pre-training corpora, LoopMDM matches the performance of same-size MDMs with up to 3.3 fewer training FLOPs, while its final performance outperforms them on various reasoning benchmarks, including up to 8.5 points on GSM8K. It even surpasses deeper non-looped MDMs trained with comparable per-step compute, indicating that selective looping is more effective than naive depth scaling. Furthermore, LoopMDM can scale inference-time compute by increasing the number of loops. Adaptively adjusting the number of loops throughout the sampling process further yields additional gains in compute efficiency while maintaining performance. Lastly, with attention analysis, we provide evidence that looping is effective in MDMs by promoting interactions among masked positions. Our code and weights will be publicly released.
