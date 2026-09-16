package main

import (
	"net/http"
	"time"

	"github.com/deltrexgg/profolio/functions"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var httpRequestsTotal = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "portfolio",
		Help: "Total number of HTTP requests processed.",
	},
	[]string{"code", "method", "path"},
)

var httpRequestDuration = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "myapp_http_request_duration_seconds",
		Help:    "HTTP request duration in seconds.",
		Buckets: prometheus.DefBuckets,
	},
	[]string{"method", "path"},
)

type responseWriter struct {
	http.ResponseWriter
	statusCode  int
	bytesWritten int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	if rw.statusCode == 0 {
		rw.statusCode = http.StatusOK
	}

	n, err := rw.ResponseWriter.Write(b)
	rw.bytesWritten += n

	return n, err
}

func metricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(rw, r)

		duration := time.Since(start).Seconds()

		httpRequestsTotal.WithLabelValues(
			http.StatusText(rw.statusCode),
			r.Method,
			r.URL.Path,
		).Inc()

		httpRequestDuration.WithLabelValues(
			r.Method,
			r.URL.Path,
		).Observe(duration)
	})
}

func main() {
	port := "8090"

	mux := http.NewServeMux()

	mux.HandleFunc("/", functions.HomePage)

	mux.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	// Prometheus metrics endpoint
	mux.Handle("/metrics", promhttp.Handler())

	handler := metricsMiddleware(mux)

	println("Server running :", port)

	http.ListenAndServe(":"+port, handler)
}