package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs details of each HTTP request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Capture the start time
		startTime := time.Now()

		// Wrap the ResponseWriter to capture the status code
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// Call the next handler
		next.ServeHTTP(lrw, r)

		// Log request details
		log.Printf("[%s] %s %s %d %s",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
			lrw.statusCode,
			time.Since(startTime))
	})
}

// loggingResponseWriter is a wrapper for http.ResponseWriter to capture the status code
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code before passing it to the underlying ResponseWriter
func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
