package middleware

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	ReturnCodeCounter  *prometheus.CounterVec
	RequestLatencyHist *prometheus.HistogramVec
)

func init() {
	ReturnCodeCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_request_counter",
		Help: "http request counter",
	}, []string{"path", "method", "code"})

	RequestLatencyHist = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "his_request_latency_sec",
		Help:    "history request duration distribution",
		Buckets: []float64{0.05, 0.2, 0.5, 1, 5, 10, 30},
	}, []string{"path", "method"})
}

func Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path
		method := request.Method
		start := time.Now()
		ReturnCodeCounter.WithLabelValues(path, method).Inc()

		next.ServeHTTP(writer, request)

		duration := time.Since(start)
		RequestLatencyHist.WithLabelValues(path, method).Observe(duration.Seconds())
	})
}
