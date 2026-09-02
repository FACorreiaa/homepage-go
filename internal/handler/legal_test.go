package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"myapp/internal/handler"

	"github.com/stretchr/testify/assert"
)

func legalRequest(t *testing.T, pattern, path string, fn http.HandlerFunc) *httptest.ResponseRecorder {
	t.Helper()
	mux := http.NewServeMux()
	mux.HandleFunc("GET "+pattern, fn)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, httptest.NewRequest(http.MethodGet, path, nil))
	return w
}

func TestNorviqTermsRenders(t *testing.T) {
	w := legalRequest(t, "/terms/{app}", "/terms/norviq", handler.LegalTerms)
	assert.Equal(t, http.StatusOK, w.Code)
	body := w.Body.String()
	assert.Contains(t, body, "Norviq Terms of Use")
	assert.Contains(t, body, "Governing law")
	assert.Contains(t, body, "laws of Portugal")
	assert.Contains(t, body, "Last updated")
}

func TestNorviqPrivacyRenders(t *testing.T) {
	w := legalRequest(t, "/privacy/{app}", "/privacy/norviq", handler.LegalPrivacy)
	assert.Equal(t, http.StatusOK, w.Code)
	body := w.Body.String()
	assert.Contains(t, body, "Norviq Privacy Policy")
	assert.Contains(t, body, "privacy@norviq.com")
	assert.Contains(t, body, "read-only basis")
}

func TestNorviqSupportRenders(t *testing.T) {
	w := legalRequest(t, "/support/{app}", "/support/norviq", handler.LegalSupport)
	assert.Equal(t, http.StatusOK, w.Code)
	body := w.Body.String()
	assert.Contains(t, body, "support@norviq.org")
	assert.Contains(t, body, `href="/privacy/norviq"`)
	assert.Contains(t, body, `href="/terms/norviq"`)
}

func TestUnknownLegalAppIs404(t *testing.T) {
	for _, tc := range []struct {
		pattern, path string
		fn            http.HandlerFunc
	}{
		{"/terms/{app}", "/terms/nope", handler.LegalTerms},
		{"/privacy/{app}", "/privacy/nope", handler.LegalPrivacy},
		{"/support/{app}", "/support/nope", handler.LegalSupport},
	} {
		w := legalRequest(t, tc.pattern, tc.path, tc.fn)
		assert.Equal(t, http.StatusNotFound, w.Code, tc.path)
	}
}
