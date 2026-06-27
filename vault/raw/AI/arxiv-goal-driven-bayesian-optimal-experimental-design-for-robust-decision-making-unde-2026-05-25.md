---
source: arxiv
ingested_at: 2026-05-26T09:17:30Z
type: note
status: uncompiled
category: AI
subcategory: arxiv
title: Goal-driven Bayesian Optimal Experimental Design for Robust Decision-Making Under Model Uncertainty
arxiv_id: 2605.26093
---

## Goal-driven Bayesian Optimal Experimental Design for Robust Decision-Making Under Model Uncertainty

**Authors:** Jinwoo Go, Xiaoning Qian, Byung-Jun Yoon
**Published:** 2026-05-25
**Categories:** cs.LG, stat.ML
**URL:** https://arxiv.org/abs/2605.26093v1

### Abstract
Bayesian optimal experimental design (BOED) selects experiments to maximize information gain about model parameters. However, in decision-critical settings, reducing parameter uncertainty does not necessarily improve downstream decisions, as only specific parameter directions relevant to the objective truly matter. We propose GoBOED, a goal-driven BOED framework that directly optimizes experimental designs for a specified decision-making objective. GoBOED combines an amortized variational posterior surrogate with a differentiable convex decision layer, enabling gradient-based design optimization that is fully decision-focused. We theoretically show that GoBOED gradients are insensitive to parameter directions irrelevant to the decision objective, providing a formal justification for why goal-driven design achieves equivalent decision quality over a wider set of experimental designs than information-gain maximization. Empirically, across source localization, epidemic management, and pharmacokinetic control, GoBOED identifies designs that better align with downstream decision objectives and reveals that near-optimal design windows are substantially wider than those predicted by goal-agnostic BOED approaches.
