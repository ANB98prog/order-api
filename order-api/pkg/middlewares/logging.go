package middlewares

import (
	"log/slog"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func Logging(logger *slog.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
			defer func() {
				duration := time.Since(start)
				logger.Info("Request processed",
					slog.String("method", r.Method),
					slog.String("path", r.URL.Path),
					slog.String("remote_ip", r.RemoteAddr),
					slog.Int("status", rw.statusCode),
					slog.Int64("duration_ms", duration.Milliseconds()),
				)
			}()
			next.ServeHTTP(rw, r)
		})
	}
}
