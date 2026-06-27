---
source: arxiv-cs-AI
ingested_at: 2026-05-21T11:09:39Z
type: ai_paper
subcategory: arxiv-AI
status: uncompiled
---

# arXiv cs.AI — Recent Papers

*Fetched: 2026-05-21T11:09:39Z*


### [Variance Reduction for Expectations with Diffusion Teachers](http://arxiv.org/abs/2605.21489v1)

- **Published:** 2026-05-20
- Pretrained diffusion models serve as frozen teachers feeding downstream pipelines such as text-to-3D, single-step distillation, and data attribution. The teacher gradients these pipelines consume are Monte Carlo (MC) expectations over noise levels and Gaussian noise samples; their estimator variance dominates compute cost because each draw requires expensive upstream work (rendering, simulation, encoding). We introduce CARV, a compute-aware variance-accounting framework that motivates a hierarch…

### [Quantifying Hyperparameter Transfer and the Importance of Embedding Layer Learning Rate](http://arxiv.org/abs/2605.21486v1)

- **Published:** 2026-05-20
- Hyperparameter transfer allows extrapolating optimal optimization hyperparameters from small to large scales, making it critical for training large language models (LLMs). This is done either by fitting a scaling law to the hyperparameters or by a judicious choice of parameterization, such as Maximal Update ($μ$P), that renders optimal hyperparameters approximately scale invariant. In this paper, we first develop a framework to quantify hyperparameter transfer through three metrics: (1) the qual…

### [DeepWeb-Bench: A Deep Research Benchmark Demanding Massive Cross-Source Evidence and Long-Horizon Derivation](http://arxiv.org/abs/2605.21482v1)

- **Published:** 2026-05-20
- Deep research, in which an agent searches the open web, collects evidence, and derives an answer through extended reasoning, is a prominent use case for frontier language models. Frontier deep research products score high on existing benchmarks, making it difficult to distinguish their capabilities from current evaluation data alone. We introduce DeepWeb-Bench, a deep research benchmark that is substantially harder than existing benchmarks for the current frontier. Difficulty comes from three pr…

### [AiraXiv: An AI-Driven Open-Access Platform for Human and AI Scientists](http://arxiv.org/abs/2605.21481v1)

- **Published:** 2026-05-20
- Recent advances in artificial intelligence (AI) have accelerated the growth of both human-authored and AI-generated research outputs, placing increasing strain on traditional academic publishing systems and challenging the scalability of conference- and journal-centered paradigms amid rising submission volumes, reviewer workload, and venue size. To address these challenges, we explore an AI-era publishing paradigm in which both human and AI scientists participate as authors and readers, and pape…

### [WikiVQABench: A Knowledge-Grounded Visual Question Answering Benchmark from Wikipedia and Wikidata](http://arxiv.org/abs/2605.21479v1)

- **Published:** 2026-05-20
- Visual Question Answering (VQA) benchmarks have largely emphasized perception-based tasks that can be solved from visual content alone. In contrast, many real-world scenarios require external knowledge that is not directly observable in the image to answer correctly. We introduce WikiVQABench, a human-curated knowledge-grounded VQA benchmark constructed by systematically combining Wikipedia images, their associated article captions, and structured knowledge from Wikidata. Our pipeline uses large…

### [Agent JIT Compilation for Latency-Optimizing Web Agent Planning and Scheduling](http://arxiv.org/abs/2605.21470v1)

- **Published:** 2026-05-20
- Computer-use agents (CUA) automate tasks specified with natural language such as "order the cheapest item from Taco Bell" by generating sequences of calls to tools such as click, type, and scroll on a browser. Current implementations follow a sequential fetch-screenshot-execute loop where each iteration requires an LLM call, resulting in high latency and frequent errors from incorrect tool use. We present agent just-in-time (JIT) compilation, an alternative that compiles task descriptions direct…

### [Mem-$π$: Adaptive Memory through Learning When and What to Generate](http://arxiv.org/abs/2605.21463v1)

- **Published:** 2026-05-20
- We present Mem-$π$, a framework for adaptive memory in large language model (LLM) agents, where useful guidance is generated on demand rather than retrieved from external memory stores. Existing memory-augmented agents typically rely on similarity-based retrieval from episodic memory banks or skill libraries, returning static entries that often misalign with the current context. In contrast, Mem-$π$ uses a dedicated language or vision-language model with its own parameters, separate from the dow…

### [HITL-D: Human In The Loop Diffusion Assisted Shared Control](http://arxiv.org/abs/2605.21460v1)

- **Published:** 2026-05-20
- Autonomous manipulation systems have achieved remarkable capabilities, yet the integration of human expertise with diffusion-based policies in shared control remains relatively unexplored. In this paper, we propose Human-In-The-Loop Diffusion (HITL-D), a shared control framework that enhances user performance in multi-step, insertion, and fine manipulation tasks. HITL-D leverages a novel combination of diffusion-based policies and human control to provide autonomous end effector orientation upda…

### [Mind the Sim-to-Real Gap & Think Like a Scientist](http://arxiv.org/abs/2605.21458v1)

- **Published:** 2026-05-20
- Suppose a planner has a pre-trained simulator of a sequential decision problem and the option to run real experiments in the field. The simulator is cheap to query but inherits confounding and drift from its calibration data. Experimentation is unbiased but consumes one real unit per trial. We study when, and how, the planner should supplement the simulator with experiments. We give three results. First, an extended simulation lemma decomposes the simulator's value error into a calibration--depl…

### [Quality and Security Signals in AI-Generated Python Refactoring Pull Requests](http://arxiv.org/abs/2605.21453v1)

- **Published:** 2026-05-20
- As AI agents increasingly contribute to code development and maintenance, there is still limited empirical evidence on the quality and risk characteristics of their changes in real-world projects, particularly for refactoring-oriented contributions. It remains unclear how agent-authored refactoring edits affect maintainability, code quality, and security once merged into GitHub repositories. To address this gap, we conduct an empirical study of Python refactoring pull requests (PRs) from the AID…

### [Approximation Theory for Neural Networks: Old and New](http://arxiv.org/abs/2605.21451v1)

- **Published:** 2026-05-20
- Universal approximation theorems provide a mathematical explanation for the expressive power of neural networks. They assert that, under mild conditions on the activation function, feedforward neural networks are dense in broad function classes, such as continuous functions on compact subsets of $\mathbb{R}^d$, $L^p$ spaces, or Sobolev spaces. Over the past four decades, these qualitative universality results have evolved into a rich quantitative theory addressing approximation rates, parameter …

### [Lost in Fog: Sensor Perturbations Expose Reasoning Fragility in Driving VLAs](http://arxiv.org/abs/2605.21446v1)

- **Published:** 2026-05-20
- Interpretable autonomous driving planners depend not only on generating explanations, but also on those explanations remaining reliable under real-world sensor degradation. In this paper we present a controlled perturbation study of Vision-Language-Action (VLA) robustness in autonomous driving, evaluating Alpamayo R1 (10B parameters) across 1,996 scenarios under eight sensor perturbations (Gaussian noise at four intensities, two lighting extremes, and two fog levels; ${\sim}18{,}000$ inference t…

### [TempGlitch: Evaluating Vision-Language Models for Temporal Glitch Detection in Gameplay Videos](http://arxiv.org/abs/2605.21443v1)

- **Published:** 2026-05-20
- Vision-language models (VLMs) are increasingly being explored for video game quality assurance, especially gameplay glitch detection. Most existing evaluations, however, treat glitches as static visual anomalies, asking models to detect failures from a single frame. We argue that this framing misses a key distinction: some glitches are spatial and visible in an isolated frame, whereas others are temporal and become evident only through changes across ordered frames. A preliminary study confirms …

### [torchtune: PyTorch native post-training library](http://arxiv.org/abs/2605.21442v1)

- **Published:** 2026-05-20
- Modern LLMs typically require multistage training pipelines to achieve strong downstream performance, with post-training serving as the main interface for adapting open-weight models. We introduce torchtune, a PyTorch-native library designed to streamline the post-training lifecycle of LLMs, enabling efficient fine-tuning, experimentation, and deployment-oriented workflows. Unlike many existing fine-tuning frameworks, which often optimize for ease of use, specialized recipes, or hardware efficie…

### [PALS: Power-Aware LLM Serving for Mixture-of-Experts Models](http://arxiv.org/abs/2605.21427v1)

- **Published:** 2026-05-20
- Large language model (LLM) inference has become a dominant workload in modern data centers, driving significant GPU utilization and energy consumption. While prior systems optimize throughput and latency by batching, scheduling, and parallelism, they largely treat GPU power as a static constraint rather than a controllable resource. In this paper, we present a power-aware runtime for LLM serving, PALS, that treats GPU power caps as a first-class control knob and jointly optimizes them with softw…
