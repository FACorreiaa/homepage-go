---
source: The Defiant
ingested_at: 2026-06-01T10:04:57Z
type: web
status: uncompiled
category: Crypto
subcategory: The Defiant
title: "Alephium Bridge Loses $815K to Forged Guardian Messages, Not Stolen Keys"
url: "http://thedefiant.io/news/hacks/alephium-bridge-815k-forged-guardian-messages"
pubdate: ""
---

Title: Alephium Bridge Loses $815K to Forged Guardian Messages, Not Stolen Keys

URL Source: http://thedefiant.io/news/hacks/alephium-bridge-815k-forged-guardian-messages

Published Time: 2026-05-31T12:59:29Z

Markdown Content:
Alephium's Wormhole-fork TokenBridge was drained on Ethereum and BNB Chain in roughly seven minutes after an off-chain backend flaw let fraudulent messages slip past its four-guardian network, the team said in a public correction.

Alephium, a proof-of-work Layer 1 that runs a private fork of the Wormhole bridge, lost about $815,000 across Ethereum and BNB Chain on Friday after an attacker pushed forged messages through the bridge backend and out the other side as legitimate-looking transfers, [according to](https://x.com/alephium/status/2060768639100268782) the team. Alephium has shut the bridge down and said no new transactions can be initiated.

The attacker drained 200,967 USDT, 17,594 USDC, 5.18 WETH and 0.335 WBTC on Ethereum, plus 36,750 USDT and 24.386 WBNB on BNB Chain, [per](https://x.com/alephium/status/2060751960022663334) Alephium's accounting, and minted 13.76 million wrapped ALPH on Ethereum with no corresponding ALPH locked on the Alephium chain. The full sequence took about seven minutes, [according to](https://x.com/blockaid_/status/2060765689229484392) blockchain security firm Blockaid, which spotted the exploit first and brought in the [SEAL 911](https://x.com/SEAL_911) emergency-response unit.

The headline number is small by the standards of 2026's bridge year — PeckShield data put cumulative cross-chain bridge losses at roughly [$329 million](https://protos.com/bridge-hacks-back-in-vogue-as-verus-exploit-brings-2026-total-to-329m/) through mid-May — but the attack mechanism is the part to watch. Initial reports, including from Blockaid, blamed compromised guardian keys. Alephium and Blockaid have since revised that.

## What Actually Broke

In a [follow-up post](https://x.com/blockaid_/status/2060682200861831327), Blockaid said the exploit "does not appear to have involved a compromise of guardian private keys. Instead, it appears to have involved an exploit that allowed forged malicious events/messages to be observed and signed by guardians." Alephium's own [post-incident statement](https://x.com/alephium/status/2060768639100268782) traced the root cause to "an off-chain vulnerability in the bridge backend that could be triggered in specific edge cases" and ruled out both a smart-contract bug and a key compromise.

The distinction is the story. A stolen-key incident is an operational-security failure — a guardian custodian losing control of a signing device. A forged-event incident is a software-layer flaw between the bridge's on-chain contracts and the off-chain process that feeds the guardians what to sign. The guardians, in other words, signed valid signatures over invalid data. Six of those signed approvals were enough to drain the bridge.

## Four-Guardian Set

Alephium's fork uses four guardians with a quorum of three. [Wormhole's mainnet](https://wormhole.com/docs/protocol/infrastructure/guardians/), by comparison, runs 19 guardians and a 13-signature quorum. Under the larger Wormhole set, a backend bug that pushes a forged message to a single guardian still has to convince twelve more. Under a four-guardian set, the same bug only has to convince two. The math is unforgiving once the off-chain feed is wrong.

ALPH held inside the bridge itself was not drained and is recoverable, the team [said](https://x.com/alephium/status/2060751960022663334), and Alephium has [urged](https://x.com/alephium/status/2060751961855549948) ALPH holders on Ethereum and BNB Chain to pull liquidity from Uniswap and PancakeSwap pools immediately. Any swap activity now lets the attacker convert the 13.76 million unbacked wrapped ALPH still in the exploiter wallet into real value.

## A Pattern Beat Across 2026

Cross-chain bridges have been the busiest target in DeFi this year. The Verus-Ethereum bridge [lost $11.5 million](https://thedefiant.io/verus-ethereum-bridge-exploit-11-5-million-ri18bt) on May 18 to a backing-validation flaw — the same class of bug, fraudulent messages clearing a check they should have failed — and the Gravity Bridge was drained of [$5.4 million](https://beincrypto.com/gravity-bridge-hack-key-compromise-5m/) on May 30 in a suspected signing-key compromise. CrossCurve and Hyperbridge fell to fabricated-message and verifier bugs earlier in the year. The exposures vary; the architecture rarely does.

Alephium said the recovery process for users with ALPH locked in the bridge will be announced next week, alongside a full technical postmortem and details of a compensation plan. The bridge stays offline until then.

