# Norviq Feature Planning - Final Polish & Launch

## Project Context
- **App Name:** Norviq
- **Status:** Code complete, awaiting App Store review
- **Goal:** Polished launch with key differentiators highlighted
- **Target User:** Investors who want to track expenses and investments together

## Core Features (Already Implemented)
- Stock planning (DCA, thesis, advanced metrics)
- Earnings tracking
- News integration
- Expense tracking (manual entry)

## Differentiator
**Expense tracking is the killer feature** - simple, personal, integrated with investment dashboard

## Priority Features for Final Polish (Before Launch)

### P0 - Critical for Launch
1. **Expense Dashboard Integration**
   - Unified dashboard showing:
     - Monthly expenses trend
     - Investment portfolio performance
     - Correlation between expenses and investment capacity
   - Visual indicators when expenses exceed budget thresholds

2. **Expense Reporting**
   - Monthly expense summary report
   - Category breakdown (food, transport, entertainment, etc.)
   - Comparison vs previous months
   - Export to CSV/PDF

3. **Pro Feature Definition**
   - Identify 2-3 premium features for paid tier
   - Example: Advanced metrics, custom reports, multi-currency support
   - Set up in-app purchase infrastructure

### P1 - High Value Post-Launch
4. **QR Code Expense Entry (MVP)**
   - Scan receipt QR codes to auto-fill expense details
   - Basic OCR for extracting amount/date from receipt images

5. **IBKR Integration**
   - Connect to Interactive Brokers for automatic trade/import
   - Sync portfolio positions and transactions

6. **Crypto Tracking**
   - Support for major cryptocurrencies (BTC, ETH, SOL)
   - Price tracking and portfolio integration

### P2 - Future Enhancements
7. **Bank Feed Integration**
   - Plaid or Yodlee integration for automatic transaction import
   - Categorization rules engine

8. **Advanced Analytics**
   - Predictive cash flow based on expense trends
   - Investment scenario modeling with expense variables

## Technical Considerations

### Data Model
```
User {
  id: UUID
  email: string
  subscription_tier: free | pro | trial
}

Expense {
  id: UUID
  user_id: UUID
  amount: decimal
  category: string (food, transport, entertainment, etc.)
  date: datetime
  description: string
  receipt_image_url: string (optional)
}

StockPosition {
  id: UUID
  user_id: UUID
  symbol: string
  quantity: decimal
  avg_cost: decimal
  current_price: decimal
}

PortfolioSnapshot {
  id: UUID
  user_id: UUID
  date: datetime
  total_value: decimal
  expense_total: decimal
  investment_returns: decimal
}
```

### Architecture
- **Client:** SwiftUI iOS app
- **Backend:** Vapor (Swift on Server) with PostgreSQL
- **Storage:** Local CoreData for offline, sync with server
- **Images:** Store receipt images in S3 or similar

## UI/UX Priorities

### Dashboard (P0)
- Clean, minimalist design with clear metrics
- Expense vs Investment dual-axis chart
- Color coding for budget warnings
- Quick add expense button from dashboard

### Expense Entry (P0)
- Simple form: amount, category, date, description
- Recent categories autocomplete
- Receipt photo attachment
- Save and add another option

### Reports (P1)
- Monthly summary card
- Category pie chart
- Trend line chart (last 6 months)
- Export button

## Marketing Preparation

### App Store Assets
- Screenshots highlighting:
  - Dashboard with expense/investment integration
  - Expense entry simplicity
  - Reports and insights
- App icon (professional, clean design)
- App preview video (show key workflows)

### Description
- Lead with the unique value proposition: "Track expenses and investments in one place"
- Emphasize simplicity and personal use case
- Mention pro features and trial

## Launch Checklist
- [ ] Final QA on expense tracking
- [ ] Dashboard integration complete
- [ ] Reports functionality working
- [ ] App Store screenshots prepared
- [ ] App preview video recorded
- [ ] Description and keywords optimized
- [ ] In-app purchase setup configured
- [ ] TestFlight beta testing completed

## Success Metrics
- User activation rate (complete first expense entry)
- Dashboard usage frequency
- Conversion to pro tier
- App Store rating (target 4.5+)

---
*Document created: 2026-05-07*
*Next: HermiesVault MVP specification*