package handler

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
	"sync"
)

var gzPool = sync.Pool{
	New: func() any {
		w, _ := gzip.NewWriterLevel(io.Discard, gzip.BestSpeed)
		return w
	},
}

type gzipResponseWriter struct {
	http.ResponseWriter
	gz          *gzip.Writer
	wroteHeader bool
}

func (w *gzipResponseWriter) WriteHeader(code int) {
	if !w.wroteHeader {
		w.Header().Del("Content-Length")
		w.wroteHeader = true
	}
	w.ResponseWriter.WriteHeader(code)
}

func (w *gzipResponseWriter) Write(b []byte) (int, error) {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}
	return w.gz.Write(b)
}

func compressible(path string) bool {
	for _, ext := range []string{".css", ".js", ".svg", ".json", ".xml", ".txt", ".html", ".woff"} {
		if strings.HasSuffix(path, ext) {
			return true
		}
	}
	// Extensionless paths are HTML routes.
	return !strings.Contains(path[strings.LastIndex(path, "/")+1:], ".")
}

// Gzip compresses compressible responses when the client accepts it. Behind
// Caddy this is mostly a no-op for clients (Caddy negotiates zstd first), but
// it keeps the binary fast when serving directly.
func Gzip(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") ||
			r.Header.Get("Range") != "" ||
			!compressible(r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}
		gz := gzPool.Get().(*gzip.Writer)
		defer gzPool.Put(gz)
		gz.Reset(w)
		defer gz.Close() //nolint:errcheck

		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Add("Vary", "Accept-Encoding")
		next.ServeHTTP(&gzipResponseWriter{ResponseWriter: w, gz: gz}, r)
	})
}
