---
source: governance.aave.com
ingested_at: 2026-06-03T10:00:00Z
type: market-intelligence
category: Crypto
subcategory: defi-governance
status: uncompiled
---
# Aave June 2026 Funding Update

Title: [Direct-to-AIP] May/June 2026 - Funding Update - Finance - Aave

URL Source: http://governance.aave.com/t/direct-to-aip-may-june-2026-funding-update/25000

Published Time: 2026-06-01T05:56:24+00:00

Markdown Content:
[![Image 1](https://dub1.discourse-cdn.com/flex013/user_avatar/governance.aave.com/tokenlogic/48/8248_2.png)](http://governance.aave.com/t/direct-to-aip-may-june-2026-funding-update/25000)

* * *

title: [Direct-to-AIP] May/June 2026 - Funding Update

 author: [@TokenLogic](https://governance.aave.com/u/tokenlogic)

 created: 2026-06-01

* * *

## [](http://governance.aave.com/t/direct-to-aip-may-june-2026-funding-update/25000#p-64233-summary-1)Summary

This publication presents the June Funding Update, consisting of the following key activities:

*   Acquire GHO to support the runway; and,
*   Create Allowances to support Operations.

## [](http://governance.aave.com/t/direct-to-aip-may-june-2026-funding-update/25000#p-64233-motivation-2)Motivation

This publication addresses near-term operational requirements, consolidates asset holdings, and refreshes the MainnetSwapSteward’s Allowances to support operations and future growth opportunities.

The MainnetSwapSteward and Aave Finance Committee (AFC) will continue executing a rolling GHO acquisition strategy to maintain adequate runway and preserve sufficient budget to fund ongoing growth initiatives.

### [](http://governance.aave.com/t/direct-to-aip-may-june-2026-funding-update/25000#p-64233-reimburse-audit-costs-3)Reimburse Audit Costs

Reimburse TokenLogic for costs incurred in facilitating the audit of one PR intended to streamline the UX for moving from USDC or USDT to GHO or sGHO. The [Gho Router](https://governance.aave.com/t/arfc-sgho-launch-configuration/24346) contract allows users to swap routes, use GSMs (Gho Stability Modules), and enter/exit sGHO. The final instalment of ChainSecurities’s audit was 11,655 GHO.

Reimburse Aave Labs for costs incurred in facilitating the audit of Aave Pro, Aave Kit, Aave App and general protocol services. The total amount is 392,746.66 GHO, spanning nine separate audit scopes and engaging five different audit teams.

### [](http://governance.aave.com/t/direct-to-aip-may-june-2026-funding-update/25000#p-64233-bug-bounty-payouts-4)**Bug Bounty Payouts**

This month’s funding update includes one 7,500 USD bug bounty program payment to a security researcher and the associated 750 USF fee payment to ImmuneFi. A separate publication from Aave Labs will provide insights into the findings, while this proposal will address payments only.

### [](http://governance.aave.com/t/direct-to-aip-may-june-2026-funding-update/25000#p-64233-incentive-campaign-5)Incentive Campaign

To support the continued success of the friendly Tydro deployment, a 5M GHO budget has been discussed with the Tydro team. This is in addition to the AAVE Allowance created during the April funding update and provides sufficient operational flexibility to support programs such as the Kraken Defi Borrow campaign, expected to commence shortly. How rewards are distributed will be determined in collaboration with the Ink team and Aave Service Providers.

## [](http://governance.aave.com/t/direct-to-aip-may-june-2026-funding-update/25000#p-64233-specification-6)Specification

### [](http://governance.aave.com/t/direct-to-aip-may-june-2026-funding-update/25000#p-64233-runway-7)Runway

Deposit idle ETH on the Collector into the Aave v3 Core instance.

In addition to the 3.7M GHO acquired during May 2026, use the [MainnetSwapSteward](https://governance.aave.com/t/arfc-steward-deployment-mainnetswapsteward-and-rewardssteward/23070) to acquire 4M of GHO to be deposited into the Prime instance during June 2026.

### [](http://governance.aave.com/t/direct-to-aip-may-june-2026-funding-update/25000#p-64233-ethereum-8)Ethereum

To support the continued success of the friendly Tydro deployment, a GHO Allowance from the Collector Contract to the Ahab SAFE fund will support an upcoming incentive campaign.

Asset: aEthLidoGHO `0x18eFE565A5373f430e2F809b97De30335B3ad96A`

 Amount: 5M

 Spender: Ahab `0xAA2461f0f0A3dE5fEAF3273eAe16DEF861cf594e`

Method: approve() aEthLidoGHO on the Aave Collector contract to the Merit Ahab address on Ethereum.

Following the rebalancing of the Ahab and Aave Finance Committee (AFC) SAFEs to facilitate the DeFi United donation, the AFC SAFE now holds AAVE, USDC, and GHO debt. The GHO debt is to be repaid from a combination of revenue derived from the Plasma instance and the USDC held within the AFC SAFE.

### [](http://governance.aave.com/t/direct-to-aip-may-june-2026-funding-update/25000#p-64233-refresh-mainnetswapsteward-allowances-9)Refresh mainnetSwapSteward Allowances

To support the acquisition of GHO, replenish Allowances on the `MainnetSwapSteward`.

| Token | Budget |
| --- | --- |
| ETH | 5k |
| USDC | 10M |
| USDT | 10M |
| USDe | 10M |
| USDS | 10M |
| DAI | 5M |

Reference: [MainnetSwapSteward](https://governance.aave.com/t/arfc-steward-deployment-mainnetswapsteward-and-rewardssteward/23070) forum post.

### [](http://governance.aave.com/t/direct-to-aip-may-june-2026-funding-update/25000#p-64233-reimbursements-10)Reimbursements

Reimburse 11,655 aEthLidoGHO to TokenLogic for GhoRouter audit expenses incurred.

Asset: aEthLidoGHO `0x18eFE565A5373f430e2F809b97De30335B3ad96A`

 Amount: 11,655

 Spender: TokenLogic `0xAA088dfF3dcF619664094945028d44E779F19894`

Reimburse Aave Labs for costs incurred in facilitating the audit of Aave Pro, Aave Kit, Aave App and general protocol services.

Asset: aEthLidoGHO `0x18eFE565A5373f430e2F809b97De30335B3ad96A`

 Amount: 392,746.66

 Spender: Aave Labs `0x1c037b3C22240048807cC9d7111be5d455F640bd`

### [](http://governance.aave.com/t/direct-to-aip-may-june-2026-funding-update/25000#p-64233-bug-bounty-payments-11)Bug Bounty Payments

The following bug bounty payments are to be processed to the researcher and ImmuneFi.

Asset: USDC `0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48` Amount: 7,500 Spender: Researcher `0xcC7383b24631d8BfC8571dbF9c81d6D094688628`

Asset: USDC `0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48`

 Amount: 750

 Spender: ImmuneFi `0x7119f398b6C06095c6E8964C1f58e7C1BAa79E18`

### [](http://governance.aave.com/t/direct-to-aip-may-june-2026-funding-update/25000#p-64233-plasma-12)Plasma

Refresh the existing aPlaUSDT0 Allowance down to 3M.

Asset: aPlaUSDT0 `0x5d72a9d9a9510cd8cbdba12ac62593a58930a948`

 Amount: 3M

 Spender: AFC `0x22740deBa78d5a0c24C58C740e3715ec29de1bFa`

### [](http://governance.aave.com/t/direct-to-aip-may-june-2026-funding-update/25000#p-64233-disclosure-13)Disclosure

TokenLogic is an active service provider to the Aave DAO, the beneficiary of stream 100072 and the KPI as outlined in this [publication](https://governance.aave.com/t/arfc-service-provider-compensation-reform-for-v4-alignment/23246). The scope of this engagement is available via this forum [proposal](https://governance.aave.com/t/arfc-tokenlogic-phase-ii/23223).

TokenLogic supports and maintains an independent [delegate voting platform](https://governance.aave.com/t/tokenlogic-delegate-platform/12516) within the Aave community.

TokenLogic and associated entities have no undisclosed material conflicts of interest at the time of submission.

## [](http://governance.aave.com/t/direct-to-aip-may-june-2026-funding-update/25000#p-64233-next-steps-14)Next Steps

1.   Using the Direct-to-AIP process, submit AIP(s) for vote.

## [](http://governance.aave.com/t/direct-to-aip-may-june-2026-funding-update/25000#p-64233-copyright-15)Copyright

Copyright and related rights waived via [CC0](https://creativecommons.org/publicdomain/zero/1.0/).

