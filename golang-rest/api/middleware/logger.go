package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

// responseWriter is a wrapper for http.ResponseWriter that
// captures the written HTTP status code and byte size.
type responseWriter struct {
	http.ResponseWriter
	wroteHeader bool
	statusCode  int
	bytesOut    int // Reading the Content-Length header doesn't work.
}

func (w *responseWriter) WriteHeader(statusCode int) {
	if w.wroteHeader {
		return
	}
	w.ResponseWriter.WriteHeader(statusCode)
	w.wroteHeader = true
	w.statusCode = statusCode
}

func (w *responseWriter) Write(b []byte) (int, error) {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}
	n, err := w.ResponseWriter.Write(b)
	w.bytesOut += n
	return n, err
}

// Logger logs HTTP responses.
// It should be used after RequestID middleware.
func Logger(logger *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Use wrapper to get HTTP status code.
			ww := &responseWriter{ResponseWriter: w}
			start := time.Now()
			next.ServeHTTP(ww, r)

			logger.Info(
				fmt.Sprintf("%s %s", r.Method, r.URL.EscapedPath()),
				slog.String("request_id", GetRequestID(r.Context())),
				slog.Int64("bytes_in", int64(r.ContentLength)),
				slog.Int("bytes_out", ww.bytesOut),
				slog.Int("status", ww.statusCode),
				slog.Int64("duration_ms", time.Since(start).Milliseconds()),
			)
		})
	}
}
