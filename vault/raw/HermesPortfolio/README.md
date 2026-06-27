# 📊 Hermes Portfolio Management System

Your personal portfolio tracking + review workflow. Designed for the StockPlan investor who wants systematic, not emotional, decision-making.

---

## 📁 Files & Tools

### `portfolio_tracker.py`
**Purpose:** Maintain the canonical holdings list + live prices snapshot  
**Input:** `live_prices.csv` (manual entry or bulk import from your daily broker reports)  
**Output:**  
- `portfolio_snapshot.csv` — latest view of all positions with market values  
- `portfolio_history.csv` — append-only timestamped history for time-series tracking  
- `wishlist_conviction.json` — structured thesis tracking for planned additions

**Usage:**
```bash
# Interactive price update (prompts for each active ticker)
python ~/.hermes/scripts/portfolio_tracker.py

# Bulk update from CSV (ticker,price columns)
python ~/.hermes/scripts/portfolio_tracker.py --from-prices /path/to/prices.csv

# Write snapshot only (no prompts)
python ~/.hermes/scripts/portfolio_tracker.py --snapshot

# Quick refresh with cached prices
python ~/.hermes/scripts/portfolio_tracker.py --once
```

### `portfolio_summary.py`
**Purpose:** Terminal dashboard — one-glance portfolio health check  
**Output:** ASCII table with ticker, shares, last price, market value, % of portfolio

```bash
python ~/.hermes/scripts/portfolio_summary.py        # read latest snapshot
python ~/.hermes/scripts/portfolio_summary.py --refresh   # re-run tracker first
```

### `portfolio_monthly_review.py`
**Purpose:** Formal monthly review report — thesis validation, concentration check, action items  
**Output:** Markdown report at `~/.hermes/portfolio/reports/review-<YYYY-MM>.md`

```bash
python ~/.hermes/scripts/portfolio_monthly_review.py
```

---

## 🗂 Directory Structure

```
~/.hermes/
├── scripts/
│   ├── portfolio_tracker.py
│   ├── portfolio_summary.py
│   └── portfolio_monthly_review.py
└── portfolio/
    ├── portfolio_snapshot.csv     ← latest state
    ├── portfolio_history.csv      ← time-series log
    ├── live_prices.csv            ← your price entries (MANUAL)
    ├── wishlist_conviction.json   ← 1% candidates thesis/exit criteria
    └── reports/
        └── review-2026-04.md      ← monthly markdown reviews
```

---

## 🧑‍💻 Daily / Weekly Workflow

### Daily (5 min)
1. Open your broker app / daily report
2. Run: `python ~/.hermes/scripts/portfolio_tracker.py --once`
   - If prices are stale, run interactive: `python …/portfolio_tracker.py`
   - Enter last price for each active ticker from your broker
3. Quick peek: `python ~/.hermes/scripts/portfolio_summary.py`

### Monthly (30 min, calendar reminder)
1. Open the latest review: `cat ~/.hermes/portfolio/reports/review-*.md`
2. Update `wishlist_conviction.json`:
   - Set `"conviction": 1–10` for each planned 1% ticker
   - Refine `"thesis"`, `"catalyst"`, `"exit"` based on new information
3. Run the review generator: `python ~/.hermes/scripts/portfolio_monthly_review.py`
4. Commit the report to git / copy to Obsidian
5. Decide capital deployment for the month:
   - Top 1–2 conviction names → new 1% positions
   - Underperforming speculative → trim or exit

---

## ⚙️ Automation (cron)

Currently you have the **skill-based** cron infrastructure via Hermes. To automate reviews:

```bash
# Add a monthly cron job (first Monday of month at 9 AM)
crontab -e
0 9 1 * * /opt/data/home/.hermes/scripts/portfolio_monthly_review.py
```

*(Note: The script writes to files only; it does NOT deliver to Discord/Telegram by default.)*

If you want it delivered:
```bash
# Pipe to your Hermes report skill (adjust path)
0 9 1 * * /opt/data/home/.hermes/scripts/portfolio_monthly_review.py && python /path/to/kb-report-skill.py --from ~/.hermes/portfolio/reports/review-$(date +\%Y-\%m).md
```

---

## 🎯 Portfolio Design Principles

### Current risk profile
- **Active holdings:** 17 (very diversified)
- **Speculative / micro-cap:** ~43% of value (ADUR, KRKNF, A6I, ONDS, SIDU, etc.)
- **Large-cap core:** AMD, GOOGL, AMZN, NFLX, UBER (~45%)
- **Mid-cap growth:** ZETA, HIMS, OSCR (~30%)

### 1% position sizing rule
Your plan: every new idea (RDW, OUST, SMR, ELF) gets 1% of portfolio (~$333 at $33k).  
**Why 1%?** Limits any single speculative loss to 1% of total capital.

### Rebalancing checkpoints
- No speculative > 3% without explicit "core" conviction
- Keep at least 5–10% cash dry powder for dips in ZETA/HIMS/ABCL
- Quarterly: trim anything >2× its target 1% allocation back to 1%

---

## 📝 Price Update Tips

### Manual entry
Edit `~/.hermes/portfolio/live_prices.csv` directly:
```csv
ticker,price,updated_at
AMD,329.54,2026-04-29T07:00:00Z
ZETA,18.33,2026-04-29T07:00:00Z
```
The tracker reads this on every run.

### Bulk from broker export
Most brokers export CSV with columns like `Symbol`, `Last`, `Close`. Convert:
```python
# Convert your broker's CSV to Hermes prices format
import csv
with open("broker_export.csv") as src, open("live_prices.csv","w",newline="") as dst:
    w = csv.writer(dst)
    w.writerow(["ticker","price","updated_at"])
    for row in csv.DictReader(src):
        w.writerow([row["Symbol"], row["Last"], "2026-04-29"])
```

### From your daily Hermes stock news script
If you already have a cron job that outputs daily ticker prices, pipe to this:
```bash
# Suppose your daily script outputs TSV: TICKER<tab>PRICE
python ~/.hermes/scripts/daily_stock_prices.py | tee /tmp/todays_prices.tsv
# Then import:
python -c "
import csv, sys
with open('/tmp/todays_prices.tsv') as f, open('~/.hermes/portfolio/live_prices.csv','w',newline='') as out:
    w = csv.writer(out); w.writerow(['ticker','price','updated_at'])
    for line in f:
        t,p = line.strip().split('\\t')
        w.writerow([t, p, 'now'])
"
```

---

## 🔧 Wishlist: Setting Conviction

Edit `wishlist_conviction.json` before buying:
```json
{
  "RDW": {
    "conviction": 7,
    "max_allocation_pct": 1.0,
    "thesis": "Space infrastructure …",
    "catalyst": "NASA contract …",
    "exit": "Sell if: revenue flat 2 quarters; dilution >10%"
  },
  …
}
```

**Conviction scale:** 1–10  
- 1–3: speculative bet, small size only  
- 4–6: moderate confidence, monitor quarterly  
- 7–8: high confidence, core holding material  
- 9–10: buy aggressively on dips

---

## 🧠 Philosophy

- **Data before emotion:** Update prices from broker statements, not feelings.
- **Thesis-driven:** Every position (existing or planned) needs a clear catalyst + exit rule.
- **Mechanical sizing:** 1% for speculative, 3–5% max for core growth.
- **Monthly review cadence:** Prevents reactive day-trading; forces strategic thinking.

---

**Last updated:** 2026-04-29 (initial setup)  
**Portfolio value baseline:** $33,317 (screenshot); grew to $37,567 with recent gains

