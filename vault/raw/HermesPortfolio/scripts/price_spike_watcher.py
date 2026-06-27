#!/usr/bin/env python3
import csv, json, os, sys
from datetime import datetime, timedelta
from pathlib import Path
import subprocess

BASE = Path.home() / ".hermes" / "portfolio"
HISTORY_FILE = BASE / "portfolio_history.csv"
ALERT_FILE  = BASE / "price_spike_alerts.json"

WATCHLIST = [
    "AIDR", "OCO", "LEEF", "GOAT",
    "RDW", "OUST", "SMR", "ELF",
    "ZETA", "HIMS", "OSCR", "ASTS", "SOFI", "ABCL", "TE", "ONDS",
]

THRESHOLD_PCT = 20.0
BASELINE_DAYS = 7

def load_history():
    history = {}
    if not HISTORY_FILE.exists():
        return history
    with HISTORY_FILE.open(newline="") as f:
        for row in csv.DictReader(f):
            ticker = row["ticker"]
            if ticker.startswith("_"):
                continue
            try:
                ts = row["snapshot_dt"].replace("Z", "").split("+")[0].strip()
                dt = datetime.fromisoformat(ts)
                price = float(row["live_price"] or 0)
                if price <= 0:
                    continue
                history.setdefault(ticker, []).append((dt, price))
            except Exception:
                pass
    for t in history:
        history[t].sort(reverse=True)
    return history

def detect_spikes():
    history = load_history()
    alerts = []
    cutoff = datetime.utcnow() - timedelta(days=BASELINE_DAYS)
    for ticker in WATCHLIST:
        if ticker not in history:
            continue
        prices = history[ticker]
        if not prices:
            continue
        latest_dt, latest_price = prices[0]
        baseline_pts = [(dt, p) for dt, p in prices if dt >= cutoff]
        if len(baseline_pts) < 2:
            continue
        baseline_dt, baseline_price = baseline_pts[-1]
        pct = ((latest_price - baseline_price) / baseline_price) * 100
        if abs(pct) >= THRESHOLD_PCT:
            alerts.append({"ticker": ticker, "direction": "up" if pct > 0 else "down",
                "pct_change": round(pct, 2), "latest_price": latest_price,
                "baseline_price": baseline_price,
                "baseline_date": baseline_dt.isoformat(),
                "latest_date": latest_dt.isoformat()})
    return alerts

def format_md(alerts):
    if not alerts:
        return f"# 📈 Price Spike Watch — {datetime.now().strftime('%Y-%m-%d %H:%M')}\n\n✅ No tickers moved >{THRESHOLD_PCT}% in the last {BASELINE_DAYS} days."
    lines = [f"# 🚨 PRICE SPIKE ALERT — {datetime.now().strftime('%Y-%m-%d %H:%M')}", f"**Threshold:** ±{THRESHOLD_PCT}% over {BASELINE_DAYS} days", ""]
    for a in sorted(alerts, key=lambda x: abs(x["pct_change"]), reverse=True):
        arrow = "↗️" if a["direction"] == "up" else "↘️"
        sign = "+" if a["pct_change"] > 0 else ""
        lines.append(f"## {arrow} {a['ticker']}  {sign}{a['pct_change']}%")
        lines.append(f"- Latest: ${a['latest_price']:.2f}  ({a['latest_date'][:10]})")
        lines.append(f"- Baseline: ${a['baseline_price']:.2f}  ({a['baseline_date'][:10]})")
        lines.append("")
    lines.append("---")
    lines.append("**Action:** Review thesis. Catalyst or noise?")
    return "\n".join(lines)

def main():
    BASE.mkdir(parents=True, exist_ok=True)
    alerts = detect_spikes()
    md = format_md(alerts)
    with ALERT_FILE.open("w") as f:
        json.dump({"generated_at": datetime.now().isoformat(), "threshold_pct": THRESHOLD_PCT, "alerts": alerts}, f, indent=2)
    alert_md = BASE / "price_spike_report.md"
    alert_md.write_text(md)
    print(md)
    if alerts:
        hermes = Path("/opt/hermes/bin/hermes")
        if hermes.exists():
            subprocess.run([str(hermes), "skills", "run", "kb-report", "--args", f"--title \"🚨 Price Spike Alert — {len(alerts)} tickers\" --priority high"], input=md, text=True, capture_output=True)
    print(f"\n✅ Spike check → {alert_md}")
    if alerts:
        print(f"   ⚠️  {len(alerts)} ticker(s) exceeded threshold")

if __name__ == "__main__":
    main()
