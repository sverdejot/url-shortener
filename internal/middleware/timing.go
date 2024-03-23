package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func Timer(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		slog.Info("started handling", "method", r.Method, "path", r.URL.Path, "start", start)
		f.ServeHTTP(w, r)
		slog.Info("ended handling", "method", r.Method, "path", r.URL.Path, "end", time.Now(), "took", time.Since(start))
	})
}
