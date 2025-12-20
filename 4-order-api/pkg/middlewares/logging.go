package middlewares

import (
	"github.com/sirupsen/logrus"
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

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		defer func() {
			duration := time.Since(start)
			logrus.WithFields(logrus.Fields{
				"method":      r.Method,
				"path":        r.URL.Path,
				"remote_ip":   r.RemoteAddr,
				"status":      rw.statusCode,
				"duration_ms": duration.Milliseconds(),
			}).Info("Request processed")
		}()
		next.ServeHTTP(rw, r)
	})
}
