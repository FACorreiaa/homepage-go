package model

// LegalDocument is one public legal page (privacy policy, terms of use) for a
// product this studio operates. Documents are served at /privacy/{app} and
// /terms/{app} — the App Store and Play listings link to them, so a slug must
// never change once published.
type LegalDocument struct {
	App         string // product slug, e.g. "norviq"
	AppName     string // display name, e.g. "Norviq"
	Kind        string // "privacy" | "terms"
	Title       string
	LastUpdated string
	Intro       string
	Sections    []LegalSection
}

type LegalSection struct {
	Heading string
	Body    []string
}

// LegalApp groups the documents and support details for one product.
type LegalApp struct {
	Slug         string
	Name         string
	Operator     string
	SupportEmail string
	PrivacyEmail string
	WebsiteURL   string
	AppStoreURL  string
	Privacy      LegalDocument
	Terms        LegalDocument
}

// LegalApps is the registry the legal routes consult. Order does not matter.
func LegalApps() []LegalApp {
	return []LegalApp{norviqLegal()}
}

func norviqLegal() LegalApp {
	const (
		operator = "Fernando Correia (FC Software Studio)"
		support  = "support@norviq.org"
		privacy  = "privacy@norviq.com"
	)
	app := LegalApp{
		Slug:         "norviq",
		Name:         "Norviq",
		Operator:     operator,
		SupportEmail: support,
		PrivacyEmail: privacy,
		WebsiteURL:   "https://norviq.org",
		AppStoreURL:  "https://apps.apple.com/app/money-manager-norviq/id6765849578",
	}
	app.Terms = LegalDocument{
		App: "norviq", AppName: "Norviq", Kind: "terms",
		Title:       "Norviq Terms of Use",
		LastUpdated: "2 September 2026",
		Intro:       "These Terms govern your use of Norviq (the \"App\") on iPhone and at norviq.org, operated by " + operator + ". By creating an account, downloading, or using the App you agree to them.",
		Sections: []LegalSection{
			{Heading: "Licence", Body: []string{
				"We grant you a personal, non-transferable, revocable licence to use the App on Apple devices you own or control, in accordance with the Apple Media Services Terms and these Terms, and to use the web app through your account.",
			}},
			{Heading: "Read-only financial data, not advice", Body: []string{
				"Norviq aggregates and displays financial data you enter, import, or connect for informational purposes only. It does not execute trades, transfers, or payments. Brokerage and bank connections are read-only.",
				"Nothing in the App is investment, tax, legal, or financial advice. Prices, sentiment, projections, scenario outputs, and AI-generated text can be wrong, delayed, or incomplete. Verify figures with the original source before acting on them.",
			}},
			{Heading: "Subscriptions and purchases", Body: []string{
				"Norviq Pro is an auto-renewing subscription sold through Apple In-App Purchase or web billing. Payment is charged to your Apple Account or payment method at confirmation. It renews automatically unless cancelled at least 24 hours before the end of the current period; manage or cancel it in your Apple Account settings or from Settings in the web app.",
				"Any free-trial portion is forfeited when you purchase a subscription. Refunds for App Store purchases are handled by Apple under its policies.",
			}},
			{Heading: "Your responsibilities", Body: []string{
				"You are responsible for the accuracy of the data you enter and for keeping your credentials, passkeys, and personal access tokens private. Tokens you create for MCP clients give that client access to your account; revoke them from Settings when no longer needed.",
				"You must be at least 18 years old to use the App.",
			}},
			{Heading: "Acceptable use", Body: []string{
				"You agree to use Norviq only for lawful purposes and not to attempt to disrupt, reverse-engineer, scrape, or gain unauthorised access to the service or to other users' data.",
			}},
			{Heading: "Availability and third parties", Body: []string{
				"Market data, macroeconomic statistics, brokerage and bank connections, AI models, and billing are provided by third parties and may change, be delayed, or become unavailable. We do not guarantee uninterrupted service.",
				"Economy pages display public statistics from Eurostat, the European Central Bank Data Portal, the Federal Reserve Bank of St. Louis FRED® API, IBGE and Banco Central do Brasil. Figures are reproduced as published by those institutions, which do not endorse Norviq. FRED® data is used under the FRED® API Terms of Use (https://fred.stlouisfed.org/docs/api/terms_of_use.html); by using features that display FRED® data you agree to those terms. FRED® is a registered trademark of the Federal Reserve Bank of St. Louis.",
			}},
			{Heading: "Disclaimer and liability", Body: []string{
				"The App is provided \"as is\" without warranties of any kind. To the maximum extent permitted by law, " + operator + " is not liable for indirect or consequential damages, including investment losses, and our total liability is limited to the amount you paid for the App in the 12 months before the claim. Nothing limits liability that cannot be excluded by law.",
			}},
			{Heading: "Changes and termination", Body: []string{
				"We may update these Terms; the date above changes when we do, and continued use after a change means you accept it. You can stop using the App and delete your account from Settings at any time. We may suspend accounts that breach these Terms.",
			}},
			{Heading: "Governing law", Body: []string{
				"These Terms are governed by the laws of Portugal, without regard to conflict-of-laws rules. Consumers in the EU keep the protections of their country of residence.",
			}},
			{Heading: "Contact", Body: []string{"Questions about these Terms: " + support + "."}},
		},
	}
	app.Privacy = LegalDocument{
		App: "norviq", AppName: "Norviq", Kind: "privacy",
		Title:       "Norviq Privacy Policy",
		LastUpdated: "2 September 2026",
		Intro:       "Norviq is a private financial planning tool operated by " + operator + ". This policy explains what data is processed, why, who helps us process it, and the controls you have.",
		Sections: []LegalSection{
			{Heading: "Data we collect", Body: []string{
				"Account information (email, username, authentication identifiers, passkeys), financial data you enter or import (portfolios, holdings, watchlists, expenses, budgets, goals, tax preferences), and usage and diagnostics data used to operate and improve the App.",
				"When subscriptions are enabled, we process billing state from RevenueCat and App Store or web billing events (subscription status, entitlement level, product and event identifiers). Raw billing-provider payloads are operational records and are not included in standard exports.",
			}},
			{Heading: "Connected financial accounts", Body: []string{
				"When you connect a brokerage or bank account, Norviq accesses that data on a read-only basis. Norviq cannot place trades, move money, or modify data at your financial institution.",
				"Interactive Brokers (IBKR): when enabled, portfolio and statement data may be ingested through the IBKR Web Service reporting feed using credentials you provide. IBKR acts as a sub-processor for that reporting data.",
				"Bank connections use Plaid (United States) and GoCardless (European Union) as data processors. Access tokens are stored encrypted at rest. You can revoke any connection at any time from Integrations.",
			}},
			{Heading: "AI features and MCP", Body: []string{
				"The in-app assistant, insight cards, and proactive tips may send limited portfolio and spending context to our AI model providers (for example OpenAI or OpenRouter) to generate a response. Those calls are paid by Norviq, subject to usage limits, and are not used to train models.",
				"Model Context Protocol (MCP) access is user-directed: you connect your own LLM client (such as Claude, Cursor, or ChatGPT) with a Norviq personal access token. Norviq provides tools and your account data to that client; the LLM provider is chosen and billed by you, not by Norviq.",
			}},
			{Heading: "Receipt scanning", Body: []string{
				"Receipt images captured for expense entry are processed transiently to extract amount, merchant, and date, and are not retained as a long-term image archive after processing completes.",
			}},
			{Heading: "Analytics and diagnostics", Body: []string{
				"We use PostHog (EU hosting) for product analytics and Sentry for crash and error reporting. Analytics events describe how features are used; they do not contain your holdings, balances, or bank data. IP addresses are anonymised where the provider supports it.",
			}},
			{Heading: "Sub-processors (summary)", Body: []string{
				"Infrastructure and hosting in the EU (compute, database, Redis, object storage), email and push delivery, market-data providers, billing (RevenueCat / App Store / web payments), analytics and observability (PostHog, Sentry), the IBKR reporting feed (when connected), Plaid, GoCardless, and AI model providers for in-app features.",
				"Contact " + privacy + " for the current sub-processor list or a data processing agreement request.",
			}},
			{Heading: "Retention and your controls", Body: []string{
				"Your data is kept while your account exists. You can disconnect financial accounts, delete individual records, export your data, or delete your entire account and associated product data from within the App. Deleted accounts are removed from backups within 30 days.",
				"Under the GDPR you may request access, rectification, erasure, restriction, portability, or object to processing by writing to " + privacy + ". You can also complain to your national data-protection authority (in Portugal, the CNPD).",
			}},
			{Heading: "Contact", Body: []string{
				"Privacy, export, and deletion requests: " + privacy + ". General support: " + support + ".",
			}},
		},
	}
	return app
}
