package handler

import (
	"net/http"

	"myapp/internal/model"
	"myapp/ui/pages"
)

var trustItems = []model.TrustItem{
	{"GL", "4+", "Countries"},
	{"PR", "15+", "Projects Delivered"},
	{"YR", "7+", "Years Experience"},
	{"TK", "10+", "Technologies"},
}

var allProjects = []model.ProjectItem{
	{Slug: "norviq", Title: "Norviq", RoleTag: "Independent", Description: "A focused investing workspace for portfolios, watchlists, targets, and market context.", Tags: []string{"Vapor (Swift)", "SwiftUI", "Docker", "PostgreSQL", "Redis", "Hetzner VPS"}, Category: "Full Stack / iOS App", GithubLink: "Private", LiveLink: "https://apps.apple.com/", HasLiveLink: true, Featured: true, Icon: "NV", IconURL: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/swift/swift-original.svg", LogoAsset: "/assets/static/projects/norviq-icon.png"},
	{Slug: "luminavault", Title: "LuminaVault", RoleTag: "Independent", Description: "An early-stage iOS second brain combining organized Spaces with Hermes, an AI agent designed to reason over saved links and notes.", Tags: []string{"SwiftUI", "Hermes Agent", "Spaces", "AI Memory", "TestFlight"}, Category: "iOS App / AI Memory", GithubLink: "Private", HasLiveLink: false, Featured: true, Icon: "LV", LogoAsset: "/assets/static/projects/luminavault-icon.jpg"},
	{Slug: "hermes", Title: "Hermes", RoleTag: "Independent", Description: "A private gateway and automation hub hosted on a VPS, acting as a central coordinator for AI models, webhooks, and personal services.", Tags: []string{"Docker", "VPS", "Gateway", "Swift", "API"}, Category: "Backend API", GithubLink: "Private", HasLiveLink: false, Featured: true, Icon: "HM"},
	{Slug: "hermesvault-backend", Title: "HermesVault Backend", RoleTag: "Independent", Description: "Self-hosted backend for HermesVault with a kb-compile engine transforming Markdown into a queryable knowledge base with pgvector semantic search.", Tags: []string{"Swift 6", "Hummingbird 2", "Postgres", "pgvector", "Docker"}, Category: "Backend API", GithubLink: "Private", HasLiveLink: false, Featured: true, Icon: "HB", IconURL: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/swift/swift-original.svg"},
	{Slug: "hermesvault-client", Title: "HermesVault Client", RoleTag: "Independent", Description: "Native iOS app with Vision-based OCR, native capture of photos, notes, and HealthKit data.", Tags: []string{"SwiftUI", "SwiftData", "Vision OCR", "AVFoundation"}, Category: "iOS App", GithubLink: "Private", HasLiveLink: false, Featured: true, Icon: "HC", IconURL: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/apple/apple-original.svg"},
	{Slug: "hermes-integrations", Title: "Hermes Integrations", RoleTag: "Independent", Description: "Messaging integrations for Hermes: Discord, Slack, and Telegram bots for remote commands and instant knowledge capture.", Tags: []string{"Swift", "Discord API", "Slack API", "Telegram Bot API", "Webhooks"}, Category: "Integrations / Services", GithubLink: "Private", HasLiveLink: false, Featured: true, Icon: "HI"},
	{Slug: "fanapi", Title: "FanAPI", RoleTag: "Client Work", Description: "RESTful API for Fandemic with JWT auth, real-time group chat via WebSockets, and OpenAPI documentation.", Tags: []string{"Go", "REST API", "WebSocket", "JWT", "PostgreSQL"}, Category: "Backend API", GithubLink: "Private", HasLiveLink: false, Featured: true, Icon: "FA", IconURL: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/go/go-original.svg"},
	{Slug: "pamozi-crm", Title: "Pamozi CRM", RoleTag: "Client Work", Description: "Platform for managing data, users and stock for an NGO providing used glasses to communities in need.", Tags: []string{"Go", "Server Side Rendering", "Templ", "HTMX", "PostgreSQL"}, Category: "Backend API", GithubLink: "https://github.com/FACorreiaa/glasses-management-platform", LiveLink: "https://pamozi.de", HasLiveLink: true, Featured: true, Icon: "PC", IconURL: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/go/go-original.svg"},
	{Slug: "ios-ecommerce", Title: "iOS E-Commerce App", RoleTag: "Client Work", Description: "Native iOS e-commerce application with Apple Pay integration and smooth animations.", Tags: []string{"Swift", "SwiftUI", "CoreData", "Combine"}, Category: "iOS App", GithubLink: "https://github.com/FACorreiaa", HasLiveLink: false, Featured: true, Icon: "IE", IconURL: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/apple/apple-original.svg"},
	{Slug: "android-fitness-tracker", Title: "Android Fitness Tracker", RoleTag: "Client Work", Description: "Native Android fitness tracking app integrating health APIs, built with Kotlin and Jetpack Compose.", Tags: []string{"Kotlin", "Jetpack Compose", "Android SDK"}, Category: "Android App", GithubLink: "https://github.com/FACorreiaa", HasLiveLink: false, Featured: true, Icon: "AF", IconURL: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/android/android-original.svg"},
	{Slug: "web-dashboard-mvp", Title: "Web Dashboard MVP", RoleTag: "Client Work", Description: "Responsive, high-performance web dashboard with real-time analytics for enterprise business needs.", Tags: []string{"React", "TypeScript", "Tailwind CSS"}, Category: "Web App", GithubLink: "https://github.com/FACorreiaa", LiveLink: "https://www.facorreia.com", HasLiveLink: true, Featured: true, Icon: "WD", IconURL: "https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/react/react-original.svg"},
	{Slug: "financial-services-platform", Title: "Financial Services Platform", RoleTag: "Client Work", Description: "Modular web services for finance sector with Go, Postgres, MSSQL, and GraphQL.", Tags: []string{"Go", "GraphQL", "Postgres", "MSSQL"}, Category: "Backend", GithubLink: "https://github.com/FACorreiaa", HasLiveLink: false, Icon: "FS"},
	{Slug: "network-tracking-app", Title: "Network Tracking App", RoleTag: "Client Work", Description: "Web version of mobile network tracking application using React and Rust with real-time telemetry.", Tags: []string{"React", "Rust", "gRPC", "OTLP"}, Category: "Full Stack", GithubLink: "https://github.com/FACorreiaa", HasLiveLink: false, Icon: "NT"},
	{Slug: "social-learning-platform", Title: "Social Learning Platform", RoleTag: "Client Work", Description: "Cross-platform mobile and web app with React Native and TypeScript, full backoffice configuration.", Tags: []string{"React Native", "TypeScript", "Go"}, Category: "Mobile", GithubLink: "https://github.com/FACorreiaa", HasLiveLink: false, Icon: "SL"},
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
			Project:     p,
			Tagline:     "A focused investing workspace for portfolios, watchlists, targets, and market context.",
			LongDesc:    []string{"Norviq is a personal investing workspace for active investors who want research, holdings, and decisions in one calm, fast app.", "The iOS client is built natively in SwiftUI with Swift Charts, secured by MFA, Face ID, and an app-lock layer."},
			Features:    []model.DetailFeature{{"Portfolio clarity", "Track holdings, cost basis, and allocation in real time."}, {"Research workspace", "Watchlists, stock insights, valuation editor, projections, and notes."}, {"Security first", "MFA, Face ID, app-lock with security code, encrypted token storage."}},
			TechStack:   []string{"SwiftUI", "Swift Charts", "Vapor", "PostgreSQL", "Redis", "Docker", "RevenueCat"},
			SocialLinks: []model.DetailLink{{"Instagram", "https://instagram.com/norviqplan"}, {"X (Twitter)", "https://x.com/NorviqPlanner"}, {"Discord", "https://discord.gg/3QVkas3rH"}},
			BackendNote: "Powered by api.norviqa.io",
		}
	case "luminavault":
		return model.ProjectDetailData{
			Project:     p,
			Tagline:     "Your second brain, self-hosted. Private, AI-powered memory layer you actually own.",
			LongDesc:    []string{"LuminaVault is a self-improving memory layer for researchers and analysts who want a living second brain they truly own.", "One tap with kb-compile turns raw inputs into a smart, searchable wiki."},
			Features:    []model.DetailFeature{{"Effortless capture", "Screenshots, photos, Apple Maps, HealthKit — all saved as clean Markdown."}, {"kb-compile engine", "Turns raw notes into a queryable knowledge base with pgvector semantic search."}, {"100% yours", "Self-hosted via Docker. Per-tenant vault, JWT auth, BYO LLM key."}},
			TechStack:   []string{"SwiftUI", "SwiftData", "Vision OCR", "Swift 6", "Hummingbird 2", "PostgreSQL", "pgvector", "Docker"},
			BackendNote: "Self-hosted on your own VPS.",
			BannerAsset: "/assets/static/projects/luminavault-banner.jpg",
		}
	default:
		return model.ProjectDetailData{Project: p, Tagline: p.Category, LongDesc: []string{p.Description}, TechStack: p.Tags}
	}
}

func ProjectsList(w http.ResponseWriter, r *http.Request) {
	pages.Projects(trustItems, allProjects).Render(r.Context(), w)
}

func ProjectDetail(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	p, ok := projectBySlug(slug)
	if !ok {
		http.NotFound(w, r)
		return
	}
	if p.Slug == "luminavault" {
		pages.ProjectLuminavault(p).Render(r.Context(), w)
		return
	}
	pages.ProjectDetail(detailFor(p)).Render(r.Context(), w)
}
