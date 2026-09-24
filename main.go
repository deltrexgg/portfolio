package main

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/deltrexgg/profolio/functions"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "myapp_http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "path", "code"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "myapp_http_request_duration_seconds",
			Help:    "HTTP request duration in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)

	httpResponseBytes = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "myapp_http_response_bytes_total",
			Help: "Total HTTP response bytes sent.",
		},
		[]string{"method", "path"},
	)

	httpRequestBytes = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "myapp_http_request_bytes_total",
			Help: "Total HTTP request bytes received.",
		},
		[]string{"method", "path"},
	)

	resumeViews = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "resume_views_total",
			Help: "Total resume views by company.",
		},
		[]string{"company"},
	)
)

func init() {
	prometheus.MustRegister(
		httpRequestsTotal,
		httpRequestDuration,
		httpResponseBytes,
		httpRequestBytes,
		resumeViews,
	)
}

type responseWriter struct {
	http.ResponseWriter
	status int
	bytes  int
}

func (w *responseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *responseWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = http.StatusOK
	}

	n, err := w.ResponseWriter.Write(b)
	w.bytes += n
	return n, err
}

func metricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Don't count Prometheus scraping itself.
		if r.URL.Path == "/metrics" {
			next.ServeHTTP(w, r)
			return
		}

		start := time.Now()

		rw := &responseWriter{
			ResponseWriter: w,
		}

		next.ServeHTTP(rw, r)

		duration := time.Since(start).Seconds()

		path := r.URL.Path
		method := r.Method
		code := strconv.Itoa(rw.status)

		if path == "/" {
			company := strings.ToLower(
				strings.TrimSpace(
					r.URL.Query().Get("companyName"),
				),
			)

			if company != "" {
				resumeViews.WithLabelValues(company).Inc()
			}
		}

		httpRequestsTotal.WithLabelValues(
			method,
			path,
			code,
		).Inc()

		httpRequestDuration.WithLabelValues(
			method,
			path,
		).Observe(duration)

		httpResponseBytes.WithLabelValues(
			method,
			path,
		).Add(float64(rw.bytes))

		if r.ContentLength > 0 {
			httpRequestBytes.WithLabelValues(
				method,
				path,
			).Add(float64(r.ContentLength))
		}
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

	// Prometheus endpoint
	mux.Handle("/metrics", promhttp.Handler())

	println("Server running :", port)

	handler := metricsMiddleware(mux)

	http.ListenAndServe(":"+port, handler)
}