---
source: arxiv-cs-LG
ingested_at: 2026-05-21T11:09:39Z
type: ai_paper
subcategory: arxiv-LG
status: uncompiled
---

# arXiv cs.LG — Recent Papers

*Fetched: 2026-05-21T11:09:39Z*


### [Variance Reduction for Expectations with Diffusion Teachers](http://arxiv.org/abs/2605.21489v1)

- **Published:** 2026-05-20
- Pretrained diffusion models serve as frozen teachers feeding downstream pipelines such as text-to-3D, single-step distillation, and data attribution. The teacher gradients these pipelines consume are Monte Carlo (MC) expectations over noise levels and Gaussian noise samples; their estimator variance dominates compute cost because each draw requires expensive upstream work (rendering, simulation, encoding). We introduce CARV, a compute-aware variance-accounting framework that motivates a hierarch…

### [Equilibrium Reasoners: Learning Attractors Enables Scalable Reasoning](http://arxiv.org/abs/2605.21488v1)

- **Published:** 2026-05-20
- Scaling test-time compute by iteratively updating a latent state has emerged as a powerful paradigm for reasoning. Yet the internal mechanisms that enable these iterative models to generalize beyond memorized patterns remain unclear. We hypothesize that generalizable reasoning arises from learning task-conditioned attractors: latent dynamical systems whose stable fixed points correspond to valid solutions.   We formalize this process through Equilibrium Reasoners (EqR), which enable test-time sc…

### [Quantifying Hyperparameter Transfer and the Importance of Embedding Layer Learning Rate](http://arxiv.org/abs/2605.21486v1)

- **Published:** 2026-05-20
- Hyperparameter transfer allows extrapolating optimal optimization hyperparameters from small to large scales, making it critical for training large language models (LLMs). This is done either by fitting a scaling law to the hyperparameters or by a judicious choice of parameterization, such as Maximal Update ($μ$P), that renders optimal hyperparameters approximately scale invariant. In this paper, we first develop a framework to quantify hyperparameter transfer through three metrics: (1) the qual…

### [EvoStruct: Bridging Evolutionary and Structural Priors for Antibody CDR Design via Protein Language Model Adaptation](http://arxiv.org/abs/2605.21485v1)

- **Published:** 2026-05-20
- Equivariant graph neural network (GNN) methods for antibody complementarity-determining region (CDR) design achieve the highest sequence recovery but suffer from severe vocabulary collapse. The current best GNN methods over-predict very few amino acids, such as tyrosine and glycine, while ignoring functionally important residues. We trace this failure to GNN encoders learning amino acid distributions de novo from limited structural data, discarding substitution patterns encoded in evolutionary d…

### [Velocityformer: Broken-Symmetry-Matched Equivariant Graph Transformers for Cosmological Velocity Reconstruction](http://arxiv.org/abs/2605.21483v1)

- **Published:** 2026-05-20
- Precise measurement of the kinematic Sunyaev-Zel'dovich (kSZ) effect - a probe of the large-scale distribution of baryonic matter, a key observable for cosmological inference - requires accurate reconstruction of galaxy velocities from spectroscopic surveys. The signal-to-noise ratio (SNR) of kSZ measurements scales directly with the correlation coefficient $r$ between reconstructed and true velocities. We introduce Velocityformer, an equivariant graph transformer architecture designed to match …

### [AiraXiv: An AI-Driven Open-Access Platform for Human and AI Scientists](http://arxiv.org/abs/2605.21481v1)

- **Published:** 2026-05-20
- Recent advances in artificial intelligence (AI) have accelerated the growth of both human-authored and AI-generated research outputs, placing increasing strain on traditional academic publishing systems and challenging the scalability of conference- and journal-centered paradigms amid rising submission volumes, reviewer workload, and venue size. To address these challenges, we explore an AI-era publishing paradigm in which both human and AI scientists participate as authors and readers, and pape…

### [Is Fixing Schema Graphs Necessary? Full-Resolution Graph Structure Learning for Relational Deep Learning](http://arxiv.org/abs/2605.21475v1)

- **Published:** 2026-05-20
- Relational prediction tasks are fundamental in many real-world applications, where data are naturally stored in relational databases (RDBs). Relational Deep Learning (RDL) addresses this problem by modeling RDBs as graphs and applying graph neural networks (GNNs) for end-to-end learning. However, the full-resolution property is commonly adopted as a design principle in graph construction for RDBs to preserve relational semantics, which leads most existing methods to rely on fixed graph structure…

### [Agent JIT Compilation for Latency-Optimizing Web Agent Planning and Scheduling](http://arxiv.org/abs/2605.21470v1)

- **Published:** 2026-05-20
- Computer-use agents (CUA) automate tasks specified with natural language such as "order the cheapest item from Taco Bell" by generating sequences of calls to tools such as click, type, and scroll on a browser. Current implementations follow a sequential fetch-screenshot-execute loop where each iteration requires an LLM call, resulting in high latency and frequent errors from incorrect tool use. We present agent just-in-time (JIT) compilation, an alternative that compiles task descriptions direct…

### [You Only Need Minimal RLVR Training: Extrapolating LLMs via Rank-1 Trajectories](http://arxiv.org/abs/2605.21468v1)

- **Published:** 2026-05-20
- Reinforcement learning with verifiable rewards (RLVR) has become a dominant paradigm for improving reasoning in large language models (LLMs), yet the underlying geometry of the resulting parameter trajectories remains underexplored. In this work, we demonstrate that RLVR weight trajectories are extremely low-rank and highly predictable. Specifically, we find that the majority of downstream performance gains are captured by a rank-1 approximation of the parameter deltas, where the magnitude of th…

### [DelTA: Discriminative Token Credit Assignment for Reinforcement Learning from Verifiable Rewards](http://arxiv.org/abs/2605.21467v1)

- **Published:** 2026-05-20
- Reinforcement learning from verifiable rewards (RLVR) has emerged as a central technique for improving the reasoning capabilities of large language models. Despite its effectiveness, how response-level rewards translate into token-level probability changes remains poorly understood. We introduce a discriminator view of RLVR updates, showing that the policy-gradient update direction implicitly acts as a linear discriminator over token-gradient vectors and thereby determines which token probabilit…

### [A Machine Learning Framework for Weighted Least Squares GNSS Positioning based on Activation Functions](http://arxiv.org/abs/2605.21461v1)

- **Published:** 2026-05-20
- Global Navigation Satellite Systems (GNSS) are widely used to provide position, velocity, and timing (PVT) information for various applications, including transportation, location-based communication services, and intelligent agriculture. In urban canyons, high-rise buildings and narrow streets can cause signal obstruction, non-line-of-sight (NLOS) reception, and multipath effects that introduce errors in GNSS pseudorange measurements. Although multi-constellations GNSS effectively increase the …

### [Mind the Sim-to-Real Gap & Think Like a Scientist](http://arxiv.org/abs/2605.21458v1)

- **Published:** 2026-05-20
- Suppose a planner has a pre-trained simulator of a sequential decision problem and the option to run real experiments in the field. The simulator is cheap to query but inherits confounding and drift from its calibration data. Experimentation is unbiased but consumes one real unit per trial. We study when, and how, the planner should supplement the simulator with experiments. We give three results. First, an extended simulation lemma decomposes the simulator's value error into a calibration--depl…

### [Mitigating Label Bias with Interpretable Rubric Embeddings](http://arxiv.org/abs/2605.21455v1)

- **Published:** 2026-05-20
- Statistical decision algorithms are increasingly deployed in domains where ground-truth labels are hard to obtain, such as hiring, university admissions, and content moderation. In these settings, models are typically trained on historical human evaluations -- for example, using past hiring decisions as a proxy for true applicant quality. However, if past evaluations unjustly favor certain groups, models trained on these labels may inherit those biases. To address this problem, we propose basing…

### [Approximation Theory for Neural Networks: Old and New](http://arxiv.org/abs/2605.21451v1)

- **Published:** 2026-05-20
- Universal approximation theorems provide a mathematical explanation for the expressive power of neural networks. They assert that, under mild conditions on the activation function, feedforward neural networks are dense in broad function classes, such as continuous functions on compact subsets of $\mathbb{R}^d$, $L^p$ spaces, or Sobolev spaces. Over the past four decades, these qualitative universality results have evolved into a rich quantitative theory addressing approximation rates, parameter …

### [torchtune: PyTorch native post-training library](http://arxiv.org/abs/2605.21442v1)

- **Published:** 2026-05-20
- Modern LLMs typically require multistage training pipelines to achieve strong downstream performance, with post-training serving as the main interface for adapting open-weight models. We introduce torchtune, a PyTorch-native library designed to streamline the post-training lifecycle of LLMs, enabling efficient fine-tuning, experimentation, and deployment-oriented workflows. Unlike many existing fine-tuning frameworks, which often optimize for ease of use, specialized recipes, or hardware efficie…
