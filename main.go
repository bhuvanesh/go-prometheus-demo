// main.go
package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Define metrics
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path"},
	)

	httpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path"},
	)
)

func instrumentHandler(path string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler(w, r)
		duration := time.Since(start).Seconds()

		httpRequestsTotal.WithLabelValues(path).Inc()
		httpRequestDuration.WithLabelValues(path).Observe(duration)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the monitored app!"))
}

func doubleHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Double handler called")
	// Extract number from path
	path := strings.TrimPrefix(r.URL.Path, "/double/")
	num, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := num * 2
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%d doubled is %d", num, result)
	// num, err := strconv.Atoi(r.PathValue("num"))
	// path := strings.TrimPrefix(r.URL.Path, "/double/")
	// num, err := strconv.Atoi(path)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// result := num * 2
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte(fmt.Sprintf("%d doubled is %d", num, result)))
}

func main() {
	// Register metrics endpoint
	http.Handle("/metrics", promhttp.Handler())

	// Register application endpoints with instrumentation
	http.HandleFunc("/", instrumentHandler("/", homeHandler))
	http.HandleFunc("/double/{num}", instrumentHandler("/double/{num}", doubleHandler))

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
