package promcommon

import (
	"github.com/prometheus/client_golang/prometheus"
)

type CacheIncrementer interface {
	IncHits(operation string)
	IncMisses(operation string)
}

type CacheMetrics struct {
	cacheHitsTotal   *prometheus.CounterVec
	cacheMissesTotal *prometheus.CounterVec
}

var _ CacheIncrementer = &CacheMetrics{}

func NewCacheMetrics() *CacheMetrics {
	metrics := new(CacheMetrics)

	metrics.cacheHitsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cache_hits_total",
		Help: "Count of records found in cache",
	}, []string{"operation"})
	prometheus.MustRegister(metrics.cacheHitsTotal)

	metrics.cacheMissesTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cache_misses_total",
		Help: "Count of records not found in cache",
	}, []string{"operation"})
	prometheus.MustRegister(metrics.cacheMissesTotal)

	return metrics
}

func (m *CacheMetrics) IncHits(operation string) {
	m.cacheHitsTotal.WithLabelValues(operation).Inc()
}

func (m *CacheMetrics) IncMisses(operation string) {
	m.cacheMissesTotal.WithLabelValues(operation).Inc()
}
