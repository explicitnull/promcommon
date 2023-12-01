package promcommon

import (
	"github.com/prometheus/client_golang/prometheus"
)

type BrokerIncrementObserver interface {
	IncBrokerMessagesReads(queue, result string)
	IncBrokerMessagesWrites(queue, result string)
	NewBrokerWriteTimer(queue string) *prometheus.Timer
}

type BrokerMetrics struct {
	brokerMessagesReadsTotal  *prometheus.CounterVec
	brokerMessagesWritesTotal *prometheus.CounterVec
	brokerWriteDuration       *prometheus.HistogramVec
}

var _ BrokerIncrementObserver = &BrokerMetrics{}

func NewBrokerMetrics() *BrokerMetrics {
	metrics := new(BrokerMetrics)

	metrics.brokerMessagesReadsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "broker_messages_reads_total",
		Help: "Count of messages read from queues, attempts and failures separately",
	}, []string{"queue", "result"})
	prometheus.MustRegister(metrics.brokerMessagesReadsTotal)

	metrics.brokerMessagesWritesTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "broker_messages_writes_total",
		Help: "Count of messages written to queues, attempts and failures separately",
	}, []string{"queue", "result"})
	prometheus.MustRegister(metrics.brokerMessagesWritesTotal)

	metrics.brokerWriteDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "broker_write_duration",
		Help:    "A Histogram of the write operation duration in seconds",
		Buckets: CustomBuckets,
	}, []string{"queue"})
	prometheus.MustRegister(metrics.brokerWriteDuration)

	return metrics
}

func (m *BrokerMetrics) IncBrokerMessagesReads(queue, result string) {
	m.brokerMessagesReadsTotal.WithLabelValues(queue, result).Inc()
}

func (m *BrokerMetrics) IncBrokerMessagesWrites(queue, result string) {
	m.brokerMessagesWritesTotal.WithLabelValues(queue, result).Inc()
}

func (m *BrokerMetrics) NewBrokerWriteTimer(queue string) *prometheus.Timer {
	return prometheus.NewTimer(m.brokerWriteDuration.WithLabelValues(queue))
}
