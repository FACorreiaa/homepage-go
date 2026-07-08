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

	sess.Init()

	mux := http.NewServeMux()
	setupAssets(mux)

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok\n"))
	})

	// Static pages
	mux.HandleFunc("GET /", handler.ProjectsList)
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
	if err := http.ListenAndServe(":"+port, withRequestPath(mux)); err != nil {
		panic(err)
	}
}

// withRequestPath exposes the request path to templates for canonical/og:url tags.
func withRequestPath(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(layouts.WithRequestPath(r.Context(), r.URL.Path)))
	})
}

func getEnvOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func setupAssets(mux *http.ServeMux) {
	goEnv := strings.ToLower(strings.TrimSpace(os.Getenv("GO_ENV")))
	isProd := goEnv == "" || goEnv == "production"
	isDev := !isProd

	assetHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isDev {
			w.Header().Set("Cache-Control", "no-store")
		} else {
			switch r.URL.Path {
			case "/assets/css/output.css", "/assets/static/sw.js", "/assets/static/manifest.json":
				w.Header().Set("Cache-Control", "public, max-age=0, must-revalidate")
			default:
				w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
			}
		}
		var h http.Handler
		if isDev {
			h = http.FileServer(http.Dir("./assets"))
		} else {
			h = http.FileServer(http.FS(assets.Assets))
		}
		h.ServeHTTP(w, r)
	})
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", assetHandler))
	utils.SetupScriptRoutes(mux, isDev)
}
