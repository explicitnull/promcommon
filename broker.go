package promcommon

import (
	"github.com/prometheus/client_golang/prometheus"
)

type BrokerIncrementObserver interface {
	IncBrokerMessagesReads(result string)
	IncBrokerMessagesWrites(result string)
	NewBrokerWriteTimer() *prometheus.Timer
}

type BrokerMetrics struct {
	brokerMessagesReadsTotal  *prometheus.CounterVec
	brokerMessagesWritesTotal *prometheus.CounterVec
	brokerWriteDuration       prometheus.Histogram
}

var _ BrokerIncrementObserver = &BrokerMetrics{}

func NewBrokerMetrics() *BrokerMetrics {
	metrics := new(BrokerMetrics)

	metrics.brokerMessagesReadsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "broker_messages_reads_total",
		Help: "Count of messages read from queues, attempts and failures separately",
	}, []string{"result"})
	prometheus.MustRegister(metrics.brokerMessagesReadsTotal)

	metrics.brokerMessagesWritesTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "broker_messages_writes_total",
		Help: "Count of messages written to queues, attempts and failures separately",
	}, []string{"result"})
	prometheus.MustRegister(metrics.brokerMessagesWritesTotal)

	metrics.brokerWriteDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "broker_write_duration",
		Help:    "A Histogram of the write operation duration in seconds",
		Buckets: CustomBuckets,
	})
	prometheus.MustRegister(metrics.brokerWriteDuration)

	return metrics
}

func (m *BrokerMetrics) IncBrokerMessagesReads(result string) {
	m.brokerMessagesReadsTotal.WithLabelValues(result).Inc()
}

func (m *BrokerMetrics) IncBrokerMessagesWrites(result string) {
	m.brokerMessagesWritesTotal.WithLabelValues(result).Inc()
}

func (m *BrokerMetrics) NewBrokerWriteTimer() *prometheus.Timer {
	return prometheus.NewTimer(m.brokerWriteDuration)
}
