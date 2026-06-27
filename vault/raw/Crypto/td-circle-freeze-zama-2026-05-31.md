---
source: http://thedefiant.io/news/defi/circle-freeze-on-zamas-confidential-usdc-locks-12-6m-of-user-funds-in-defi-crossfire
ingested_at: 2026-05-31T10:03:53Z
title: Circle Freeze on Zama's Confidential USDC Locks $12.6M of User Funds in DeFi 'Crossfire'
type: web
category: Crypto
subcategory: The Defiant
status: uncompiled
---

Title: Circle Freeze on Zama's Confidential USDC Locks $12.6M of User Funds in DeFi 'Crossfire'

URL Source: http://thedefiant.io/news/defi/circle-freeze-on-zamas-confidential-usdc-locks-12-6m-of-user-funds-in-defi-crossfire

Published Time: 2026-05-30T16:42:00.000Z

Markdown Content:
A court-ordered restraining order pushed Circle to blacklist Zama's cUSDC wrapper on Ethereum. The privacy protocol is not a defendant but every depositor in the pooled contract is locked out alongside the targeted funds.

Circle blacklisted the smart-contract address behind Zama's confidential USDC token on Friday, locking roughly $12.6 million of stablecoins inside a privacy protocol that is not itself the target of any litigation.

The freeze hit Zama's cUSDC wrapper, an ERC-1967 proxy that pools USDC on behalf of holders who use the protocol's fully homomorphic encryption to mask balances on Ethereum. On-chain investigator [ZachXBT pegged the trapped balance](https://cryip.co/zachxbt-reports-12-6m-usd-zama-cusdc-frozen-after-circle-blacklists-ethereum-address/) at 12,606,386.10 USDC and said the [block landed at 01:08 UTC](https://www.cryptotimes.io/2026/05/30/circle-blocks-zama-confidential-usdc-contract-freezing-12-6m/) on May 30. Circle has not publicly explained the action.

The action is the largest-dollar instance to date of a familiar DeFi problem: a stablecoin issuer's centralized blacklist reaching through a composed contract and locking innocent users alongside one targeted depositor. Zama, which raised a Series B last year for its confidential-blockchain stack, is the first named DeFi protocol caught in that crossfire on this scale.

## How the Freeze Reached an Innocent Contract

The trigger was a civil dispute over [Overnight Finance](https://docs.overnight.fi/core-concept/ovn-token), an OVN-token DeFi protocol whose holders have alleged the team was preparing to drain the treasury and recently [voted on Snapshot](https://snapshot.org/#/s:ovncommunity.eth/proposal/0x253dea808600c3cd198dfd8d73c188f6623ed520208058be13577ce23b98c061) to distribute funds. Plaintiffs in the underlying litigation are seeking damages against Overnight Finance founder Maxim Ermilov; ZachXBT has [named](https://x.com/zachxbt/status/2060725888262799504) Delaware-based Patagon Management, a firm with a history of suing DAO projects, as one of the plaintiffs.

A court issued a restraining order on wallets tied to the dispute the night before Circle moved. The address allegedly linked to the disputed funds had [deposited about $12.4 million USDC](https://www.cryptotimes.io/2026/05/30/circle-blocks-zama-confidential-usdc-contract-freezing-12-6m/) into Zama's cUSDC wrapper on May 11, accounting for more than 99% of the wrapper's balance. The order asked Circle to freeze the wrapper itself rather than the depositor; Circle complied. The blacklist sweeps every USDC in the pool, not just the disputed share.

## What Zama Says It Can Do

Zama CEO Rand Hindi [confirmed the investigation](https://x.com/randhindi/status/2060645486416085337) at 10:37 UTC. "We are investigating the cUSDC contract freeze. I will update here as the situation progresses," he wrote. Hindi said the incident is not related to Zama's privacy technology and that the depositing address was not flagged or sanctioned by Know-Your-Transaction monitoring tools at the time of the May 11 deposit. As a precaution, Zama paused its cUSDC, cUSDT and cWETH contracts pending review.

ZachXBT, who has [defended Zama](https://x.com/zachxbt/status/2060725888262799504) as an innocent third party, alleged the plaintiff misrepresented the relationship between the frozen address and the Zama wrapper during proceedings, and said the Zama team appeared to have received no advance notice before Circle executed the blacklist.

## The Composability Problem

Circle's blacklist function is built into the USDC ERC-20 contract: authorized accounts can mark an address; once marked, the address cannot send or receive USDC and the existing balance is blocked. The mechanic works cleanly for an externally owned wallet. It is indiscriminate against a pooled smart contract, where users' assets are commingled with the targeted depositor's.

That tension is not new. Circle [blacklisted Tornado Cash's sanctioned addresses](https://www.theblock.co/post/162172/circle-freezes-usdc-funds-in-tornado-cashs-us-treasury-sanctioned-wallets) in August 2022 after the U.S. Treasury added the mixer to the SDN list, freezing roughly 75,000 USDC. The Zama freeze runs the same mechanism on a far larger dollar scale — and without an OFAC designation underneath it. A court order in private civil litigation was sufficient.

## Where Holders Stand

Circle CEO Jeremy Allaire said in [April that the company would not freeze USDC absent a court order](https://www.coindesk.com/business/2026/04/13/circle-ceo-says-he-won-t-freeze-usdc-without-a-court-order-even-as-hackers-walk-away-with-millions), even when stolen funds were at stake. The Zama freeze meets that bar but cuts the other way: a court order from a private plaintiff has now locked an entire DeFi privacy contract.

cUSDC holders have no clear path to recovery while the freeze stands. Hindi has promised a post-mortem and updated compliance procedures after the review. Circle could reverse the action, narrow it, or publish a rationale; until it does, the contract — and the users inside it — wait.
