---
source: https://huggingface.co/blog/emo-pretraining-mixture-of-experts-for-emergent-m
ingested_at: 2026-05-11T09:26:18.190430Z
type: web
status: uncompiled
topic: AI
---
How do we get modularity to emerge?
In an MoE, a small network called the router decides which experts each token activates. We want the router to learn that tokens from similar domains should activate similar subsets of experts. Our key observation is that 
tokens from the same document usually come from the same domain
. We therefore use document boundaries as a weak supervisory signal: during training, all tokens in a document are restricted to choose their active experts from a shared expert pool.
As shown in the figure above, we first split the corpus into chunks of 2,048 sub-tokens. We then sample a subset of 256 experts for each chunk. During training, all tokens in a chunk are forced to choose their active experts from this subset. This encourages the router to learn to route all tokens from the same document — which usually come from the same domain — to similar subsets of experts. As a result, the router learns to group experts into coherent clusters that specialize in specific domains or capabilities.
EMO scaling behavior
EMO is a family of models. The flagship configuration uses 8 experts during inference and 128 total experts during training, with 1 trillion tokens. We also explore smaller and larger variants, including a 4-expert active / 64 total expert model trained on 500B tokens and a 16 expert active / 256 total expert model trained on 2T tokens. In the paper, we show that across model sizes, EMO retains its modular structure and selective expert use property.
Benchmarks
On standard English language modeling benchmarks (zero-shot, no chain-of-thought), EMO achieves: - 33% lower perplexity than a dense 7B model - 11% lower perplexity than a standard MoE with the same architecture - On BigBench-Hard, EMO's selective use capability allows it to match full-model performance while using only 12.5% of its experts, resulting in a 4.5x speedup during inference.
Take-aways and what's next