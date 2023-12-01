package promcommon

import (
	"github.com/prometheus/client_golang/prometheus"
)

type ProcessingIncrementObserver interface {
	IncEntitiesProcessed(entityType, result string)
	NewEntityProcessingTimer(entityType string) *prometheus.Timer
	IncInFlight(operationName string)
	DecInFlight(operationName string)
}

type ProcessingMetrics struct {
	entitiesProcessedTotal   *prometheus.CounterVec
	entityProcessingDuration *prometheus.HistogramVec
	entitiesInFlight         *prometheus.GaugeVec
}

var _ ProcessingIncrementObserver = &ProcessingMetrics{}

func NewProcessingMetrics() *ProcessingMetrics {
	metrics := new(ProcessingMetrics)

	metrics.entitiesProcessedTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "entities_processed_total",
		Help: "Count of entities processed, both sum and successes",
	}, []string{"entity_type", "result"})
	prometheus.MustRegister(metrics.entitiesProcessedTotal)

	metrics.entityProcessingDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "entity_processing_duration",
		Help:    "A Histogram of the entities processing duration in seconds",
		Buckets: CustomBuckets,
	}, []string{"entity_type"})
	prometheus.MustRegister(metrics.entityProcessingDuration)

	metrics.entitiesInFlight = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "entities_in_flight",
		Help: "Number of requests or messages in processing now",
	}, []string{"operation"})
	prometheus.MustRegister(metrics.entitiesInFlight)

	return metrics
}

func (m *ProcessingMetrics) IncEntitiesProcessed(entityType, result string) {
	m.entitiesProcessedTotal.WithLabelValues(entityType, result).Inc()
}

func (m *ProcessingMetrics) NewEntityProcessingTimer(entityType string) *prometheus.Timer {
	return prometheus.NewTimer(m.entityProcessingDuration.WithLabelValues(entityType))
}

func (m *ProcessingMetrics) IncInFlight(operationName string) {
	m.entitiesInFlight.WithLabelValues(operationName).Inc()
}

func (m *ProcessingMetrics) DecInFlight(operationName string) {
	m.entitiesInFlight.WithLabelValues(operationName).Dec()
}
