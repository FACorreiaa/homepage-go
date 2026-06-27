# Hermes Portfolio System — Quick Reference

## 📍 Files
- **Scripts:** `~/.hermes/scripts/`
- **Data:** `~/.hermes/portfolio/`
- **Reports:** `~/.hermes/portfolio/reports/`

## 🚀 Common Commands

```bash
# Daily
python ~/.hermes/scripts/portfolio_tracker.py          # update prices (interactive)
python ~/.hermes/scripts/portfolio_tracker.py --once   # use cached prices
python ~/.hermes/scripts/portfolio_summary.py          # view snapshot

# Monthly
python ~/.hermes/scripts/portfolio_monthly_review.py   # generate markdown report
cat ~/.hermes/portfolio/reports/review-*.md            # read latest
# Edit: ~/.hermes/portfolio/wishlist_conviction.json    # set 1–10 scores before buying

# Analysis
python ~/.hermes/scripts/portfolio_performance.py      # show trend
python ~/.hermes/scripts/portfolio_daily_check.py      # health summary
```

## 📊 Portfolio Data
- **Baseline:** $33,317 (Apr 29 screenshot)
- **Latest:** $37,567.54 (with current prices)
- **1% position:** ~$333 at baseline; scales with AUM
- **Active holdings:** 17
- **Wishlist (planned 1%):** RDW, OUST, SMR, ELF
- **Price source:** Manual CSV (live_prices.csv) — update daily/weekly

## ⚙️ Key Files to Edit

| File | Purpose |
|------|---------|
| `live_prices.csv` | Last price for every ticker (update from broker) |
| `wishlist_conviction.json` | Set conviction 1–10, thesis, catalyst, exit before buying |
| `scripts/portfolio_tracker.py` | Add/remove tickers in `PORTFOLIO` list |

## 🗓 Cron Jobs (optional)

```bash
# Monthly review (first Monday 9 AM)
0 9 1 * * python /opt/data/home/.hermes/scripts/portfolio_monthly_report_cron.py

# Daily health after your stock news cron
0 8 * * * python /opt/data/home/.hermes/scripts/portfolio_daily_check.py >> ~/cron/daily_portfolio.log 2>&1
```

## 🎯 Decision Rules

- New position size: 1% per ticker (or as conviction dictates)
- Speculative limit: No single micro-cap >3% of portfolio
- Rebalance quarterly: trim anything >2× target allocation
- Update wishlist conviction score BEFORE buying (scale 1–10)

## 🔗 Integration

- **Obsidian vault:** `/opt/obsidian-vault/` — symlink or read from `~/.hermes/portfolio/`
- **Hermes skill:** `skill_view('portfolio-tracker')` loads docs
- **Deliverables:** monthly review MD → Discord via `kb-report` skill (manual or automated)

---

Last updated: 2026-04-29
Skill: `portfolio-tracker`
User: 𝓓𝓻𝓪𝓬𝓪𝓻𝔂𝓼 𝓣𝓲𝓰𝓮𝓻
