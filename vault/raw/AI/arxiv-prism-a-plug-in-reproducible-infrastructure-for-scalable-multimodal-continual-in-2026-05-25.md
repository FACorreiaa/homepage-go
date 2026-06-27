---
source: arxiv
ingested_at: 2026-05-26T09:17:30Z
type: note
status: uncompiled
category: AI
subcategory: arxiv
title: Prism: A Plug-in Reproducible Infrastructure for Scalable Multimodal Continual Instruction Tuning
arxiv_id: 2605.26110
---

## Prism: A Plug-in Reproducible Infrastructure for Scalable Multimodal Continual Instruction Tuning

**Authors:** Jun-Tao Tang, Yu-Cheng Shi, Zhen-Hao Xie, Da-Wei Zhou
**Published:** 2026-05-25
**Categories:** cs.LG, cs.CL, cs.CV
**URL:** https://arxiv.org/abs/2605.26110v1

### Abstract
Multimodal Large Language Models (MLLMs) achieve versatility by reformulating diverse tasks into a unified instruction-following framework via instruction tuning. However, real-world deployment requires continuous adaptation to emerging tasks, motivating Multimodal Continual Instruction Tuning (MCIT). Despite its growing importance, current MCIT research is hindered by severe engineering bottlenecks. Existing methods are typically implemented by directly modifying the base MLLM codebase, which imposes substantial implementation overhead and yields method-specific architectures that severely limit code reuse and fair comparison. To address this, we introduce Prism, a plug-in reproducible codebase specifically designed for scalable MCIT research. It separates algorithmic development from the backbone implementation via a lightweight plugin registration mechanism, enabling new strategies to be integrated as independent plugins without modifying the underlying MLLM codebase, thereby eliminating structural fragmentation and accelerating method development. Prism natively supports widely used large-scale training pipeline, thereby enabling reproducible and scalable MCIT experimentation. Code is available at https://github.com/LAMDA-CL/Prism.
