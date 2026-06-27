---
source: http://thedefiant.io/news/blockchains/base-launches-azul-upgrade-takes-step-toward-stage-2-decentralization
ingested_at: 2026-05-31T10:03:53Z
title: Base Launches Azul Upgrade, Takes Step Toward Stage 2 Decentralization
type: web
category: Crypto
subcategory: The Defiant
status: uncompiled
---

Title: Base Launches Azul Upgrade, Takes Step Toward Stage 2 Decentralization

URL Source: http://thedefiant.io/news/blockchains/base-launches-azul-upgrade-takes-step-toward-stage-2-decentralization

Published Time: 2026-05-28T21:27:03.857Z

Markdown Content:
The Coinbase-incubated layer-2 network activated multiproofs on mainnet, combining TEE and ZK proof systems to move closer to full decentralization while cutting empty blocks by 99%.

Base, the Layer 2 blockchain incubated by Coinbase, activated its Azul network upgrade on mainnet this month, marking the chain's first independent protocol upgrade and a significant step toward Stage 2 decentralization.

The upgrade launched on Base Sepolia testnet on April 21 and was targeted for mainnet on May 13, [according to a blog post](https://blog.base.dev/introducing-base-azul) published by the Base engineering team that day. Multiproofs went live on mainnet by May 21, [a follow-up post confirmed](https://blog.base.dev/multiproofs-on-base).

Azul's most consequential change is the activation of multiproofs, a system that combines trusted execution environment (TEE) provers and zero-knowledge (ZK) provers into a single proof architecture.

Either proof type can independently finalize a block proposal, but when both agree, withdrawals from Base to Ethereum can settle in as little as one day. ZK proofs override TEE proofs if the two contradict, and conflicting proofs trigger an automatic soundness alert that disables the associated prover. The design is modeled on a finalization roadmap proposed by Ethereum co-founder Vitalik Buterin.

Base currently holds $4.38 billion in total value locked, supports $1.37 billion in daily DEX volume, and hosts $4.66 billion in stablecoin market cap on-chain, according to [DefiLlama data](https://defillama.com/chain/Base).

## Stage 2

Buterin's decentralization framework for Layer 2 networks defines Stage 2 as the point at which a chain can detect and handle proof system bugs onchain without relying on a privileged security council to intervene. Base reached Stage 1 in April 2025, when it established a decentralized Security Council to approve upgrades. Fault proofs, which allow any user to challenge an incorrect state claim, were activated in October 2024.

Multiproofs satisfy a core technical requirement for Stage 2 by enabling onchain detection of proof system bugs. The ZK component uses SP1, a prover developed by [Succinct Labs](https://succinct.xyz/), which Base cited for its performance and audit history. The system is designed so that a compromise of fast withdrawals would require an attacker to defeat multiple independent proof systems simultaneously.

Before the upgrade, Base ran an audit competition on [Immunefi](https://immunefi.com/) from April 21 to May 4 with a maximum reward pool of $250,000 for critical vulnerabilities reported against testnet code.

## Client Consolidation and Performance

Azul consolidates Base onto a single client stack. The upgrade drops support for all execution and consensus clients except base-reth-node, an execution client based on Reth and one of Ethereum's highest-performing clients, and base-consensus, a new consensus client built on the [Kona](https://github.com/op-rs/kona) framework. Base said the consolidation is aimed at clearing a path to 1 gigagas per second throughput.

Over the two months preceding Azul's launch, Base said it reduced empty blocks by approximately 99%, from roughly 200 per day to about two, and sustained multiple bursts of 5,000 transactions per second.

Node operators are required to migrate to the new client stack before network upgrade activation. Step-by-step instructions are available in Base's [node operator upgrade guide](https://docs.base.org/).

Azul also aligns Base with Ethereum's latest execution-layer specification, known as Osaka. The changes include EIP-7825, which introduces a per-transaction gas cap of roughly 17 million gas to support future validator performance.

## What Comes Next

Base outlined two additional upgrades planned for the coming months. A performance-focused upgrade targeting the end of June is expected to include an enshrined token standard, Flashblock access lists, Glamsterdam EIPs, a single combined client binary, and reduced withdrawal times. A second upgrade, targeting the end of August, would introduce native account abstraction.

Base also plans to launch Base Vibenet in mid-May, a public devnet for developers to experiment with upcoming features ahead of mainnet. Base described Vibenet as a permanent testing environment not tied to any specific hard fork.
