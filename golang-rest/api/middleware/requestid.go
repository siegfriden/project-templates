package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type requestIDContextKey struct{}

const RequestIDHeader = "X-Request-ID"

// RequestID assigns a unique request ID to each incoming HTTP request.
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get(RequestIDHeader)
		if reqID == "" {
			reqID = uuid.New().String()
			r.Header.Set(RequestIDHeader, reqID)
		}

		ctx := context.WithValue(r.Context(), requestIDContextKey{}, reqID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetRequestID retrieves the request ID assigned by RequestID middleware.
func GetRequestID(ctx context.Context) string {
	requestID, _ := ctx.Value(requestIDContextKey{}).(string)
	return requestID
}
