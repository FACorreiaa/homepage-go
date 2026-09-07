package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRobotsListsLLMsTxt(t *testing.T) {
	rec := httptest.NewRecorder()
	RobotsTxt(rec, httptest.NewRequestWithContext(context.Background(), http.MethodGet, "/robots.txt", nil))
	body := rec.Body.String()
	for _, want := range []string{"Sitemap: https://facorreia.com/sitemap.xml", "llms.txt"} {
		if !strings.Contains(body, want) {
			t.Fatalf("robots.txt missing %q:\n%s", want, body)
		}
	}
}

func TestLLMsTxt(t *testing.T) {
	rec := httptest.NewRecorder()
	LLMsTxt(rec, httptest.NewRequestWithContext(context.Background(), http.MethodGet, "/llms.txt", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("status %d", rec.Code)
	}
	if ct := rec.Header().Get("Content-Type"); !strings.HasPrefix(ct, "text/plain") {
		t.Fatalf("content type %q", ct)
	}
	body := rec.Body.String()
	for _, want := range []string{
		"# Fernando Correia Software Studio",
		"https://facorreia.com/",
		"linkedin.com/in/fernando-correia",
		"norviq.org",
		"kheprios.com",
		"lociai.fyi",
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("llms.txt missing %q:\n%s", want, body)
		}
	}
}
