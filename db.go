package promcommon

import (
	"github.com/prometheus/client_golang/prometheus"
)

type DBIncrementObserver interface {
	IncDatabaseSelects(db, operation, result string)
	IncDatabaseInserts(db, operation, result string)
	IncDatabaseUpdates(db, operation, result string)
	IncDatabaseDeletes(db, operation, result string)
	NewDatabaseSelectTimer(db, operation string) *prometheus.Timer
	NewDatabaseInsertTimer(db, operation string) *prometheus.Timer
	NewDatabaseUpdateTimer(db, operation string) *prometheus.Timer
	NewDatabaseDeleteTimer(db, operation string) *prometheus.Timer
}

type DBMetrics struct {
	databaseSelectsTotal   *prometheus.CounterVec
	databaseInsertsTotal   *prometheus.CounterVec
	databaseUpdatesTotal   *prometheus.CounterVec
	databaseDeletesTotal   *prometheus.CounterVec
	databaseSelectDuration *prometheus.HistogramVec
	databaseInsertDuration *prometheus.HistogramVec
	databaseUpdateDuration *prometheus.HistogramVec
	databaseDeleteDuration *prometheus.HistogramVec
}

var _ DBIncrementObserver = &DBMetrics{}

func NewDBMetrics() *DBMetrics {
	metrics := new(DBMetrics)

	metrics.databaseSelectsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "database_selects_total",
		Help: "Count of selects sent to database, attempts and failures separately",
	}, []string{"database", "operation", "result"})
	prometheus.MustRegister(metrics.databaseSelectsTotal)

	metrics.databaseInsertsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "database_inserts_total",
		Help: "Count of inserts sent to database, attempts and failures separately",
	}, []string{"database", "operation", "result"})
	prometheus.MustRegister(metrics.databaseInsertsTotal)

	metrics.databaseUpdatesTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "database_updates_total",
		Help: "Count of updates sent to database, attempts and failures separately",
	}, []string{"database", "operation", "result"})
	prometheus.MustRegister(metrics.databaseUpdatesTotal)

	metrics.databaseDeletesTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "database_deletes_total",
		Help: "Count of deletes sent to database, attempts and failures separately",
	}, []string{"database", "operation", "result"})
	prometheus.MustRegister(metrics.databaseDeletesTotal)

	metrics.databaseSelectDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "database_select_duration",
		Help:    "A Histogram of the database selects duration in seconds",
		Buckets: CustomBuckets,
	}, []string{"database", "operation"})
	prometheus.MustRegister(metrics.databaseSelectDuration)

	metrics.databaseInsertDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "database_insert_duration",
		Help:    "A Histogram of the database inserts duration in seconds",
		Buckets: CustomBuckets,
	}, []string{"database", "operation"})
	prometheus.MustRegister(metrics.databaseInsertDuration)

	metrics.databaseUpdateDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "database_update_duration",
		Help:    "A Histogram of the database updates duration in seconds",
		Buckets: CustomBuckets,
	}, []string{"database", "operation"})
	prometheus.MustRegister(metrics.databaseUpdateDuration)

	metrics.databaseDeleteDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "database_delete_duration",
		Help:    "A Histogram of the database deletes duration in seconds",
		Buckets: CustomBuckets,
	}, []string{"database", "operation"})
	prometheus.MustRegister(metrics.databaseDeleteDuration)

	return metrics
}

func (m *DBMetrics) IncDatabaseSelects(db, operation, result string) {
	m.databaseSelectsTotal.WithLabelValues(db, operation, result).Inc()
}

func (m *DBMetrics) IncDatabaseInserts(db, operation, result string) {
	m.databaseInsertsTotal.WithLabelValues(db, operation, result).Inc()
}

func (m *DBMetrics) IncDatabaseUpdates(db, operation, result string) {
	m.databaseUpdatesTotal.WithLabelValues(db, operation, result).Inc()
}

func (m *DBMetrics) IncDatabaseDeletes(db, operation, result string) {
	m.databaseDeletesTotal.WithLabelValues(db, operation, result).Inc()
}

func (m *DBMetrics) NewDatabaseSelectTimer(db, operation string) *prometheus.Timer {
	return prometheus.NewTimer(m.databaseSelectDuration.WithLabelValues(db, operation))
}

func (m *DBMetrics) NewDatabaseInsertTimer(db, operation string) *prometheus.Timer {
	return prometheus.NewTimer(m.databaseInsertDuration.WithLabelValues(db, operation))
}

func (m *DBMetrics) NewDatabaseUpdateTimer(db, operation string) *prometheus.Timer {
	return prometheus.NewTimer(m.databaseUpdateDuration.WithLabelValues(db, operation))
}

func (m *DBMetrics) NewDatabaseDeleteTimer(db, operation string) *prometheus.Timer {
	return prometheus.NewTimer(m.databaseDeleteDuration.WithLabelValues(db, operation))
}
