---
title: Portfolio System Build & Strategy Discussion
date: 2026-04-29T07:15:05.643866
tags: [hermes, portfolio, stockplan, automation]
---

# Hermes Conversation — April 29, 2026 07:15

**User:** [𝓓𝓻𝓪𝓬𝓪𝓻𝔂𝓼 𝓣𝓲𝓰𝓮𝓻]
**Context:** Building portfolio management system for StockPlan ($33k → $37k portfolio)

---

## Questions Asked

1. **"Is $ADUR dead money? What's the best strategy to follow on a portfolio like mine?"**
   - Provided portfolio analysis from April 29 screenshot
   - Discussed ADUR (micro-cap CAD, no price movement, high risk)
   - Outlined portfolio concentration risks, 1% position sizing, speculative allocation caps
   - Offered conviction-based framework for new positions

2. **"I'm thinking about Stock 1% position size RDW 333 OUST 333 SMR 333 ELF 333"**
   - Analyzed each 1% candidate (RDW space, OUST lidar, SMR nuclear, ELF cosmetics)
   - Compared risk profiles (ELF quality vs others speculative)
   - Warned against over-diversification across high-risk names
   - Asked about conviction, catalysts, opportunity cost vs. existing holdings

3. **"Create all of the above. Also for the portfolio list I have you earlier"**
   - Built complete portfolio tracking system:
     - `portfolio_tracker.py` — snapshot/history engine (manual CSV pricing)
     - `portfolio_summary.py` — terminal dashboard
     - `portfolio_monthly_review.py` — markdown report generator
     - `portfolio_daily_check.py`, `portfolio_performance.py`, import helpers
   - Initialized data files with Apr 29 prices
   - Set up wishlist conviction JSON for RDW/OUST/SMR/ELF
   - Created monthly review cron wrapper
   - Packaged as `portfolio-tracker` skill

4. **"Add this to our raw folder in the vault"**
   - Copied entire HermesPortfolio to `/opt/data/obsidian-vault/FACorreia/Raw/HermesPortfolio/`
   - Created conversation logger scripts
   - Set up sync mechanism

---

## Deliverables Created

| File/Folder | Purpose |
|-------------|---------|
| `~/.hermes/scripts/portfolio_*.py` (7 scripts) | Tracking, reporting, analysis |
| `~/.hermes/portfolio/` | Data: CSV snapshots, history, wishlist JSON |
| `~/.hermes/portfolio/reports/` | Monthly markdown reviews |
| `Raw/HermesPortfolio/` | Vault mirror of entire system |
| `sync_portfolio_vault.py` | Two-way sync between vault and home |
| `conversation_logger.py` | Auto-save Q&A to vault |

---

## Portfolio Snapshot (Apr 29 prices loaded)

| Ticker | Shares | Price | Mkt Val | Notes |
|--------|--------|-------|---------|-------|
| AMD | 20.29 | $329.54 | $6,687 | large-cap tech |
| GOOGL | 11.64 | $350.96 | $4,086 | large-cap tech |
| ZETA | 200.38 | $18.33 | $3,673 | mid-cap MarTech |
| OSCR | 178.92 | $17.97 | $3,215 | mid-cap health-tech |
| HIMS | 108.70 | $28.13 | $3,058 | mid-cap telehealth |
| AMZN | 11.42 | $260.77 | $2,977 | large-cap |
| ASTS | 30.00 | $72.73 | $2,182 | small-cap space |
| SOFI | 112.66 | $18.51 | $2,085 | fintech |
| ADUR | 130.05 | C11.97 | $1,557 | micro-cap (CAD) |
| TE | 259.74 | $5.00 | $1,299 | small-cap telecom |
| NFLX | 12.00 | $92.08 | $1,105 | large-cap streaming |
| ONDS | 104.52 | $10.53 | $1,101 | micro-cap |
| UBER | 14.49 | $74.35 | $1,077 | large-cap mobility |
| NVO | 25.00 | $41.20 | $1,030 | large-cap pharma |
| KRKNF | 153.32 | $5.67 | $869 | pink-sheet |
| A6I | 220.00 | $3.29 | $723 | frankfurt-listed |
| ABCL | 200.05 | $4.22 | $844 | biotech |
| **Wishlist (0-share)** | | | | |
| RDW | 0 | — | — | space infra |
| OUST | 0 | — | — | lidar |
| SMR | 0 | — | — | SMR nuclear |
| ELF | 0 | — | — | beauty |

**Total tracked:** $37,567.54

---

## User Preferences Captured

- **Portfolio baseline:** $33,317 (Mar/Apr snapshot)
- **Monthly contribution:** $500–600
- **Position sizing:** 1% per new ticker (~$333)
- **Wishlist conviction scale:** 1–10, must be set before buying
- **Automation preference:** Manual price entry (CSV-based), no external APIs
- **Vault storage:** All conversations + portfolio files to `Raw/HermesPortfolio/`

---

## Next Steps for User

1. Update `wishlist_conviction.json` with 1–10 scores for RDW, OUST, SMR, ELF
2. Run `python ~/.hermes/scripts/portfolio_tracker.py` to refresh prices from your broker
3. View `python ~/.hermes/scripts/portfolio_summary.py`
4. Set up monthly cron if desired (instructions in README)
5. Edit files directly in Obsidian vault (Raw/HermesPortfolio/) — sync back to home with `sync_portfolio_vault.py pull`

---

**Archived:** 2026-04-29T07:15:05.643883
**System version:** 1.0 — portfolio-tracker skill installed
