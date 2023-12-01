package promcommon

import (
	"github.com/prometheus/client_golang/prometheus"
)

type HTTPIncrementObserver interface {
	IncOutgoingHTTPRequests(destination, endpoint, result string)
	NewOutgoingHTTPRequestTimer(destination, endpoint string) *prometheus.Timer
}

type HTTPMetrics struct {
	outgoingHTTPRequestsTotal   *prometheus.CounterVec
	outgoingHTTPRequestDuration *prometheus.HistogramVec
}

var _ HTTPIncrementObserver = &HTTPMetrics{}

func NewHTTPMetrics() *HTTPMetrics {
	metrics := new(HTTPMetrics)

	metrics.outgoingHTTPRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "outgoing_http_requests_total",
		Help: "Count of requests sent to other services - failures or attempts",
	}, []string{"destination", "endpoint", "result"})
	prometheus.MustRegister(metrics.outgoingHTTPRequestsTotal)

	metrics.outgoingHTTPRequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "outgoing_http_request_duration",
		Help:    "A Histogram of requests sent to other services duration in seconds",
		Buckets: CustomBuckets,
	}, []string{"destination", "endpoint"})
	prometheus.MustRegister(metrics.outgoingHTTPRequestDuration)

	return metrics
}

func (m *HTTPMetrics) IncOutgoingHTTPRequests(destination, endpoint, result string) {
	m.outgoingHTTPRequestsTotal.WithLabelValues(destination, endpoint, result).Inc()
}

func (m *HTTPMetrics) NewOutgoingHTTPRequestTimer(destination, endpoint string) *prometheus.Timer {
	return prometheus.NewTimer(m.outgoingHTTPRequestDuration.WithLabelValues(destination, endpoint))
}
