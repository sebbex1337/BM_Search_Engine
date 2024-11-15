package metrics

import (
    "net/http"
    "strconv"
    "time"
)


// ResponseWriter is a wrapper around http.ResponseWriter that records the response status code.
type ResponseWriter struct {
    http.ResponseWriter
    statusCode int
}

func (w *ResponseWriter) WriteHeader(statusCode int) {
    w.statusCode = statusCode
    w.ResponseWriter.WriteHeader(statusCode)
}

//PrometheusMiddleware instruments HTTP handlers with Prometheus metrics.

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	
	
	start := time.Now()

	//Wrap the ResponseWriter to record the status code
	rw := &ResponseWriter{w, http.StatusOK}
	next.ServeHTTP(rw, r)

	duration := time.Since(start).Seconds()
	path := r.URL.Path
	method := r.Method
	statusCode := strconv.Itoa(rw.statusCode)

	HTTPRequestsTotal.WithLabelValues(path, method, statusCode).Inc()
	HTTPRequestDuration.WithLabelValues(path, method, statusCode).Observe(duration)
})
}
