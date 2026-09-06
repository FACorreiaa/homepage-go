package pages

import (
	"myapp/internal/model"
	"myapp/ui/layouts"
)

func norviqFAQs() []model.DetailFeature {
	return []model.DetailFeature{
		{Title: "Is Norviq free?", Body: "Yes. The free tier is a complete tracker, not a trial. Pro adds sync, AI, scenarios, tax, and MCP."},
		{Title: "Can Norviq trade or move my money?", Body: "No. Broker and bank connections are read-only by design, and there is no order or transfer feature anywhere in the app."},
		{Title: "Where is my data stored?", Body: "On servers in the European Union. Bank tokens are encrypted at rest. You can export or delete everything from Settings."},
		{Title: "What is MCP access?", Body: "A remote Model Context Protocol server. You create a scoped personal access token and your own LLM client (Claude, Codex, Cursor) can read your portfolio and spending through it. You choose and pay the model."},
		{Title: "Which brokers and banks work?", Body: "Interactive Brokers through its reporting feed, plus CSV and spreadsheet import for any broker. Banks through Plaid (US) and GoCardless (EU)."},
		{Title: "Does it work on Android?", Body: "The web app works in any browser, including on Android. The native app is iPhone only."},
	}
}

func norviqExtraJSONLD(p model.ProjectItem) string {
	faqs := norviqFAQs()
	items := make([]layouts.FAQItem, 0, len(faqs))
	for _, q := range faqs {
		items = append(items, layouts.FAQItem{Question: q.Title, Answer: q.Body})
	}
	return layouts.ProjectJSONLD(p.Slug, p.Title, p.Description, p.LogoAsset) +
		layouts.FAQPageJSONLD("/projects/norviq", items)
}
