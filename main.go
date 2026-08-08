package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"myapp/assets"
	"myapp/content"
	"myapp/internal/db"
	"myapp/internal/handler"
	"myapp/internal/service"
	sess "myapp/internal/session"
	"myapp/ui/layouts"
	"myapp/ui/pages"

	"github.com/joho/godotenv"
	"github.com/templui/templui/utils"
)

func main() {
	godotenv.Load() //nolint:errcheck

	// DB
	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath == "" {
		dbPath = "app.db"
	}
	conn, err := db.Open(dbPath)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()

	q := db.New(conn)
	ctx := context.Background()

	if err := db.SeedAdminUser(ctx, q); err != nil {
		fmt.Println("seed admin:", err)
	}
	if err := service.SeedBlogPosts(ctx, q, content.Blog); err != nil {
		fmt.Println("seed blog:", err)
	}

	// Vault index. In production, mount a vault directory and set VAULT_PATH.
	// The directory should contain raw/ at its root.
	vaultFS := os.DirFS(getEnvOrDefault("VAULT_PATH", "./vault"))
	bookmarkStore := service.NewBookmarkStore(vaultFS)
	blogTracker := service.NewBlogTracker()

	visitTracker := service.NewVisitTracker(q)
	visitTracker.Start(ctx)

	sess.Init()

	mux := http.NewServeMux()
	setupAssets(mux)

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok\n"))
	})

	// Static pages. "/{$}" matches the root exactly — a bare "GET /" is a
	// subtree pattern and would swallow every unmatched path, leaving the 404
	// handler below unreachable.
	home := &handler.HomeHandler{Tracker: visitTracker}
	mux.HandleFunc("GET /{$}", home.Show)
	mux.HandleFunc("GET /projects", handler.ProjectsList)
	mux.HandleFunc("GET /projects/{slug}", handler.ProjectDetail)
	mux.HandleFunc("GET /about", handler.About)
	mux.HandleFunc("GET /curriculum", handler.Curriculum)
	mux.HandleFunc("GET /stack", handler.Stack)
	mux.HandleFunc("GET /book-call", handler.BookCall)
	mux.HandleFunc("GET /play", handler.Play)

	// Blog
	blog := &handler.BlogHandler{Q: q, Tracker: blogTracker}
	mux.HandleFunc("GET /blog", blog.List)
	mux.HandleFunc("GET /blog/{slug}", blog.Post)

	// Proposal
	proposal := &handler.ProposalHandler{Q: q}
	mux.HandleFunc("GET /proposal", proposal.Form)
	mux.HandleFunc("POST /proposal", proposal.Submit)

	// Bookmarks
	bk := &handler.BookmarksHandler{Store: bookmarkStore, VaultFS: vaultFS}
	mux.HandleFunc("GET /bookmarks", bk.List)
	mux.HandleFunc("GET /bookmarks/{slug}", bk.Show)
	mux.HandleFunc("GET /api/graph", bk.GraphData)

	// Stats
	stats := &handler.StatsHandler{Tracker: visitTracker}
	mux.HandleFunc("GET /stats", stats.Page)
	mux.HandleFunc("GET /api/stats", stats.JSON)
	mux.HandleFunc("GET /api/stats/stream", stats.Stream)

	// SEO
	sm := &handler.SitemapHandler{Q: q, Store: bookmarkStore}
	mux.HandleFunc("GET /robots.txt", handler.RobotsTxt)
	mux.HandleFunc("GET /sitemap.xml", sm.Serve)

	// Admin
	admin := &handler.AdminHandler{Q: q}
	mux.HandleFunc("GET /admin/login", admin.LoginPage)
	mux.HandleFunc("POST /admin/login", admin.LoginSubmit)

	protected := http.NewServeMux()
	protected.HandleFunc("GET /admin/dashboard", admin.Dashboard)
	protected.HandleFunc("POST /admin/logout", admin.Logout)
	protected.HandleFunc("GET /admin/blog/new", admin.BlogNewPage)
	protected.HandleFunc("POST /admin/blog", admin.BlogCreate)
	protected.HandleFunc("GET /admin/blog/{slug}/edit", admin.BlogEditPage)
	protected.HandleFunc("POST /admin/blog/{slug}", admin.BlogUpdate)

	protectedAdmin := sess.RequireAuth(protected)
	mux.Handle("GET /admin/dashboard", protectedAdmin)
	mux.Handle("POST /admin/logout", protectedAdmin)
	mux.Handle("GET /admin/blog/new", protectedAdmin)
	mux.Handle("POST /admin/blog", protectedAdmin)
	mux.Handle("GET /admin/blog/{slug}/edit", protectedAdmin)
	mux.Handle("POST /admin/blog/{slug}", protectedAdmin)

	// 404
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		if err := pages.ErrorPage(404, "Page Not Found").Render(r.Context(), w); err != nil {
			http.Error(w, "Failed to render page", http.StatusInternalServerError)
		}
	})

	port := getEnvOrDefault("PORT", "8090")
	fmt.Println("Server is running on http://localhost:" + port)
	if err := http.ListenAndServe(":"+port, withVisitTracking(visitTracker, withRequestPath(mux))); err != nil {
		panic(err)
	}
}

// withRequestPath exposes the request path to templates for canonical/og:url tags.
func withRequestPath(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(layouts.WithRequestPath(r.Context(), r.URL.Path)))
	})
}

// withVisitTracking feeds real page views to the /stats globe. Recording is
// fire-and-forget so it never sits on the response path, and it happens after
// the handler so that 404s and errors — including anything probing for URLs
// this site does not have — stay out of the public numbers.
//
// Requests that are not tracked pass through unwrapped, which keeps the
// ResponseWriter of the SSE endpoint a real http.Flusher.
func withVisitTracking(tracker *service.VisitTracker, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !service.ShouldTrack(r) {
			next.ServeHTTP(w, r)
			return
		}

		// Liveness is marked first so the page being rendered counts its own
		// reader; the persisted dot waits for a successful status.
		tracker.Touch(r)

		recorder := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(recorder, r)

		if recorder.status < http.StatusBadRequest {
			tracker.Record(r)
		}
	})
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

// Unwrap lets http.ResponseController reach the underlying writer.
func (r *statusRecorder) Unwrap() http.ResponseWriter { return r.ResponseWriter }

func getEnvOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// assetCacheControl picks the caching policy for an asset from its full request
// path — the path as the browser asked for it, before StripPrefix rewrites it.
// Matching after the strip was the original bug here: r.URL.Path is
// "css/output.css" by then, so the allowlist never fired.
//
// The policy is now allow-list-by-immutability rather than by filename. Only
// paths that genuinely cannot change — vendored libraries, which carry their
// version in the path, and fonts — are pinned. Everything this site builds
// itself is rebuilt at a fixed path on every deploy, and pinning any of it for
// a year is how a returning visitor ends up running new HTML against old CSS.
// One conditional request per asset is nothing next to that failure.
func assetCacheControl(requestPath string, isDev bool) string {
	if isDev {
		// Read-through so the Tailwind watcher's output shows up on reload.
		return "no-store"
	}
	if assets.IsImmutable(requestPath) {
		return "public, max-age=31536000, immutable"
	}
	return "public, max-age=0, must-revalidate"
}

func setupAssets(mux *http.ServeMux) {
	goEnv := strings.ToLower(strings.TrimSpace(os.Getenv("GO_ENV")))
	isProd := goEnv == "" || goEnv == "production"
	isDev := !isProd

	var files http.Handler
	if isDev {
		// Read from disk so the Tailwind watcher's output is picked up live.
		files = http.FileServer(http.Dir("./assets"))
	} else {
		files = http.FileServer(http.FS(assets.Assets))
	}

	// The header is chosen before StripPrefix runs, so the allowlist can be
	// written as the URLs a browser actually requests.
	assetHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", assetCacheControl(r.URL.Path, isDev))
		http.StripPrefix("/assets/", files).ServeHTTP(w, r)
	})
	mux.Handle("GET /assets/", assetHandler)
	utils.SetupScriptRoutes(mux, isDev)
}
