package handler

import (
	"net/http"

	"myapp/internal/model"
	"myapp/ui/pages"
)

var trustItems = []model.TrustItem{
	{Icon: "GL", Value: "4+", Label: "Countries"},
	{Icon: "PR", Value: "15+", Label: "Projects Delivered"},
	{Icon: "YR", Value: "7+", Label: "Years Experience"},
	{Icon: "TK", Value: "10+", Label: "Technologies"},
}

// Proof slots (Metrics, Testimonial, ScreenshotAsset) are intentionally empty
// until real data exists — never fabricate numbers, quotes, or screenshots.
// Populate like:
//
//	Metrics: []model.ProjectMetric{{Value: "38ms", Label: "p95 API latency"}},
//
// Featured is the hosted-work gallery. Flip Featured to true when a product
// should sit in that first section (Norviq, LuminaVault, and later North / Loci).
var allProjects = []model.ProjectItem{
	{
		Slug:         "norviq",
		Title:        "Norviq",
		RoleTag:      "Independent",
		Description:  "A private financial command center for holdings, research, targets, and spending — available on the web and as a native iPhone app.",
		Outcome:      "I own and maintain the full stack in production: SwiftUI, Vapor, PostgreSQL, and Redis on Hetzner, secured with MFA, passkeys, and Face ID.",
		Tags:         []string{"Vapor (Swift)", "SwiftUI", "Web", "PostgreSQL", "Redis", "Hetzner VPS"},
		Category:     "Full Stack / Web & iOS",
		DisplayGroup: "Web & iOS products",
		GithubLink:   "Private",
		LiveLink:     "https://norviq.org",
		AppStoreLink: "https://apps.apple.com/pt/app/norviq/id6765849578?l=en-GB",
		HasLiveLink:  true,
		Featured:     true,
		Icon:         "NV",
		LogoAsset:    "/assets/static/projects/norviq-icon.webp",
		Status:       "Live · App Store",
	},
	{Slug: "luminavault", Title: "LuminaVault", RoleTag: "Independent", Description: "A private second brain for people who want to own their memory: organized Spaces plus Hermes, an AI agent that reasons only over what you've saved.", Outcome: "Live on TestFlight — capture becomes structured, searchable memory, and every AI answer is grounded in your own vault, self-hostable end to end.", Tags: []string{"SwiftUI", "Hermes Agent", "Spaces", "AI Memory", "TestFlight"}, Category: "iOS App / AI Memory", DisplayGroup: "iOS products", GithubLink: "Private", HasLiveLink: false, Featured: true, Icon: "LV", LogoAsset: "/assets/static/projects/luminavault-icon.webp", Status: "TestFlight"},
	{Slug: "hermesvault-backend", Title: "HermesVault Backend", RoleTag: "Independent", Description: "The self-hosted engine behind HermesVault: a Swift 6 / Hummingbird 2 API whose kb-compile pipeline turns raw Markdown into a queryable knowledge base with pgvector semantic search.", Outcome: "Runs in production as private infrastructure — per-tenant vaults, JWT auth, Docker deploys on a bare VPS.", Tags: []string{"Swift 6", "Hummingbird 2", "Postgres", "pgvector", "Docker"}, Category: "Backend API", DisplayGroup: "Backend systems", GithubLink: "Private", HasLiveLink: false, Featured: true, Icon: "HB", LogoAsset: "/assets/static/projects/hermesvault-icon.webp", Status: "Self-hosted"},
	{Slug: "hermes", Title: "Hermes", RoleTag: "Independent", Description: "A private gateway and automation hub hosted on a VPS, acting as a central coordinator for AI models, webhooks, and personal services.", Outcome: "Personal infra gateway for model routing, webhooks, and small automations.", Tags: []string{"Docker", "VPS", "Gateway", "Swift", "API"}, Category: "Backend API", DisplayGroup: "Private infrastructure", GithubLink: "Private", HasLiveLink: false, Icon: "HM", Status: "Private infra"},
	{Slug: "hermesvault-client", Title: "HermesVault Client", RoleTag: "Independent", Description: "Native iOS app with Vision-based OCR, native capture of photos, notes, and HealthKit data.", Outcome: "Capture client for photos, notes, OCR, and personal data ingestion.", Tags: []string{"SwiftUI", "SwiftData", "Vision OCR", "AVFoundation"}, Category: "iOS App", DisplayGroup: "iOS products", GithubLink: "Private", HasLiveLink: false, Icon: "HC", Status: "In development"},
	{Slug: "hermes-integrations", Title: "Hermes Integrations", RoleTag: "Independent", Description: "Messaging integrations for Hermes: Discord, Slack, and Telegram bots for remote commands and instant knowledge capture.", Outcome: "Chat-based capture and command surfaces for personal services.", Tags: []string{"Swift", "Discord API", "Slack API", "Telegram Bot API", "Webhooks"}, Category: "Integrations / Services", DisplayGroup: "Private infrastructure", GithubLink: "Private", HasLiveLink: false, Icon: "HI", Status: "Private infra", ScreenshotAsset: "/assets/static/projects/hermes-integrations-logs.webp", ScreenshotAlt: "Hermes agent log stream showing cron scheduler jobs delivering to Telegram and plugin discovery registering 54 providers"},
	{Slug: "fandemic", Title: "Fandemic", RoleTag: "Client Work", Description: "Client delivery for a real-time community platform: I built the Go backend — auth, membership, and live group chat over WebSockets — plus client-side product work.", Outcome: "Shipped with an OpenAPI spec so client teams integrated quickly; JWT-secured sessions and PostgreSQL persistence delivered on schedule.", Tags: []string{"Go", "REST API", "WebSocket", "JWT", "PostgreSQL"}, Category: "Backend / Client Work", DisplayGroup: "Client delivery", GithubLink: "Private", HasLiveLink: false, Featured: true, Icon: "FN", LogoAsset: "/assets/static/projects/fandemic-icon.webp", Status: "Delivered"},
	{Slug: "pamozi-crm", Title: "Pamozi CRM", RoleTag: "Client Work", Description: "Platform for managing data, users and stock for an NGO providing used glasses to communities in need.", Outcome: "Operational CRM for inventory, users, and field data.", Tags: []string{"Go", "Server Side Rendering", "Templ", "HTMX", "PostgreSQL"}, Category: "Backend API", DisplayGroup: "Client delivery", GithubLink: "https://github.com/FACorreiaa/glasses-management-platform", LiveLink: "https://pamozi.de", HasLiveLink: true, Icon: "PC", Status: "Live"},
	{Slug: "ios-ecommerce", Title: "iOS E-Commerce App", RoleTag: "Client Work", Description: "Native iOS e-commerce application with Apple Pay integration and smooth animations.", Outcome: "Native storefront flow with payments and polished mobile interactions.", Tags: []string{"Swift", "SwiftUI", "CoreData", "Combine"}, Category: "iOS App", DisplayGroup: "Client delivery", GithubLink: "https://github.com/FACorreiaa", HasLiveLink: false, Icon: "IE", Status: "Delivered"},
	{Slug: "android-fitness-tracker", Title: "Android Fitness Tracker", RoleTag: "Client Work", Description: "Native Android fitness tracking app integrating health APIs, built with Kotlin and Jetpack Compose.", Outcome: "Mobile fitness tracker with native health integrations.", Tags: []string{"Kotlin", "Jetpack Compose", "Android SDK"}, Category: "Android App", DisplayGroup: "Client delivery", GithubLink: "https://github.com/FACorreiaa", HasLiveLink: false, Icon: "AF", Status: "Delivered"},
	{Slug: "web-dashboard-mvp", Title: "Web Dashboard MVP", RoleTag: "Client Work", Description: "Responsive, high-performance web dashboard with real-time analytics for enterprise business needs.", Outcome: "Fast analytics MVP for internal business workflows.", Tags: []string{"React", "TypeScript", "Tailwind CSS"}, Category: "Web App", DisplayGroup: "Client delivery", GithubLink: "https://github.com/FACorreiaa", LiveLink: "https://facorreia.com", HasLiveLink: true, Icon: "WD", Status: "Live"},
	{Slug: "financial-services-platform", Title: "Financial Services Platform", RoleTag: "Client Work", Description: "Modular web services for finance sector with Go, Postgres, MSSQL, and GraphQL.", Outcome: "Service modules for financial data processing and delivery.", Tags: []string{"Go", "GraphQL", "Postgres", "MSSQL"}, Category: "Backend", DisplayGroup: "Client delivery", GithubLink: "https://github.com/FACorreiaa", HasLiveLink: false, Icon: "FS", Status: "Delivered"},
	{Slug: "network-tracking-app", Title: "Network Tracking App", RoleTag: "Client Work", Description: "Web version of mobile network tracking application using React and Rust with real-time telemetry.", Outcome: "Telemetry-heavy web interface connected to mobile network data.", Tags: []string{"React", "Rust", "gRPC", "OTLP"}, Category: "Full Stack", DisplayGroup: "Client delivery", GithubLink: "https://github.com/FACorreiaa", HasLiveLink: false, Icon: "NT", Status: "Delivered"},
	{Slug: "social-learning-platform", Title: "Social Learning Platform", RoleTag: "Client Work", Description: "Cross-platform mobile and web app with React Native and TypeScript, full backoffice configuration.", Outcome: "Cross-platform learning product with configurable backoffice workflows.", Tags: []string{"React Native", "TypeScript", "Go"}, Category: "Mobile", DisplayGroup: "Client delivery", GithubLink: "https://github.com/FACorreiaa", HasLiveLink: false, Icon: "SL", Status: "Delivered"},
}

func projectBySlug(slug string) (model.ProjectItem, bool) {
	for _, p := range allProjects {
		if p.Slug == slug {
			return p, true
		}
	}
	return model.ProjectItem{}, false
}

func detailFor(p model.ProjectItem) model.ProjectDetailData {
	switch p.Slug {
	case "norviq":
		return model.ProjectDetailData{
			Project:  p,
			Tagline:  "A private financial command center for portfolios, research, targets, and spending — synced across web and iPhone.",
			LongDesc: []string{"Norviq gives active investors one calm workspace for holdings, watchlists, due diligence, target scenarios, expenses, and reports.", "It is available in the browser and as a native SwiftUI app, with secure account access, passkeys, MFA, Face ID, exports, and account deletion built in."},
			Features: []model.DetailFeature{
				{Title: "Portfolio clarity", Body: "Track holdings, watchlists, performance, and allocation with current pricing."},
				{Title: "Research and targets", Body: "Capture thesis, risks, and catalysts, then model base, bear, and bull scenarios."},
				{Title: "Private by design", Body: "Passkeys, MFA, Face ID, private notes, data export, and in-app account deletion."},
			},
			TechStack: []string{"SwiftUI", "Swift Charts", "Vapor", "PostgreSQL", "Redis", "Docker", "RevenueCat"},
			SocialLinks: []model.DetailLink{
				{Label: "Instagram", URL: "https://instagram.com/norviqplan"},
				{Label: "X (Twitter)", URL: "https://x.com/NorviqPlanner"},
				{Label: "Discord", URL: "https://discord.gg/3QVkas3rH"},
			},
			BackendNote: "Powered by api.norviqa.io",
		}
	case "luminavault":
		return model.ProjectDetailData{
			Project:  p,
			Tagline:  "Your second brain, self-hosted. Private, AI-powered memory layer you actually own.",
			LongDesc: []string{"LuminaVault is a self-improving memory layer for researchers and analysts who want a living second brain they truly own.", "One tap with kb-compile turns raw inputs into a smart, searchable wiki."},
			Features: []model.DetailFeature{
				{Title: "Effortless capture", Body: "Screenshots, photos, Apple Maps, and HealthKit saved as clean Markdown."},
				{Title: "kb-compile engine", Body: "Turns raw notes into a queryable knowledge base with pgvector semantic search."},
				{Title: "100% yours", Body: "Self-hosted via Docker. Per-tenant vault, JWT auth, BYO LLM key."},
			},
			TechStack:   []string{"SwiftUI", "SwiftData", "Vision OCR", "Swift 6", "Hummingbird 2", "PostgreSQL", "pgvector", "Docker"},
			BackendNote: "Self-hosted on your own VPS.",
			BannerAsset: "/assets/static/projects/luminavault-banner.webp",
		}
	case "fandemic":
		return model.ProjectDetailData{
			Project:  p,
			Tagline:  "Backend and client work for Fandemic, a real-time community platform.",
			LongDesc: []string{"I built and delivered the backend for Fandemic: a Go REST API handling authentication, membership, and real-time group chat, plus client-side work across the product.", "The service is documented with OpenAPI and backed by PostgreSQL, with JWT-secured sessions and WebSocket channels for live messaging."},
			Features: []model.DetailFeature{
				{Title: "Real-time chat", Body: "Group messaging over WebSockets with presence and delivery handling."},
				{Title: "Secure by default", Body: "JWT authentication, scoped permissions, and hardened API endpoints."},
				{Title: "Documented API", Body: "OpenAPI spec so client teams could integrate quickly and reliably."},
			},
			TechStack:   []string{"Go", "REST API", "WebSocket", "JWT", "PostgreSQL", "OpenAPI"},
			BackendNote: "Backend delivered as client work.",
		}
	case "hermes-integrations":
		return model.ProjectDetailData{
			Project:  p,
			Tagline:  "Chat-native command surfaces for a private AI agent: Discord, Slack, and Telegram wired straight into Hermes.",
			LongDesc: []string{"Hermes Integrations turns messaging apps into remote controls for the Hermes agent platform: scheduled jobs report into Telegram, commands run from chat, and captured knowledge lands in the vault without opening a laptop.", "The runtime is observable end to end. A live log stream tracks the cron scheduler, delivery adapters, polling health with automatic conflict retries, and a plugin registry that discovers over 50 providers across web search, browsing, and media generation."},
			Features: []model.DetailFeature{
				{Title: "Cron to chat delivery", Body: "Scheduled agent jobs deliver results to Telegram through a live adapter, with silent runs skipped automatically."},
				{Title: "Self-healing polling", Body: "Telegram polling conflicts are detected, retried with backoff, and restarted with health checks before resuming delivery."},
				{Title: "Plugin provider registry", Body: "Startup discovery registers 54 plugins spanning web search, browser automation, and image and video generation providers."},
			},
			TechStack:   []string{"Swift", "Telegram Bot API", "Discord API", "Slack API", "Webhooks", "Cron", "Docker"},
			BackendNote: "Runs as private infrastructure on a VPS.",
			BannerAsset: "/assets/static/projects/hermes-integrations-logs.webp",
		}
	default:
		return model.ProjectDetailData{Project: p, Tagline: p.Category, LongDesc: []string{p.Description}, TechStack: p.Tags}
	}
}

func ProjectsList(w http.ResponseWriter, r *http.Request) {
	if err := pages.Projects(trustItems, allProjects).Render(r.Context(), w); err != nil {
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
	}
}

func ProjectDetail(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	p, ok := projectBySlug(slug)
	if !ok {
		http.NotFound(w, r)
		return
	}
	if p.Slug == "luminavault" {
		if err := pages.ProjectLuminavault(p).Render(r.Context(), w); err != nil {
			http.Error(w, "Failed to render page", http.StatusInternalServerError)
		}
		return
	}
	if err := pages.ProjectDetail(detailFor(p)).Render(r.Context(), w); err != nil {
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
	}
}
