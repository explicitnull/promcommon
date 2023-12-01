package promcommon

const (
	Attempt = "attempt"
	Failure = "failure"

	PostgreSQL = "postgresql"
	MySQL      = "mysql"
	Clickhouse = "clickhouse"
	Mongo      = "mongo"
	Aerospike  = "aerospike"
)

// CustomBuckets extends Prometheus built-in DefBuckets for long operations
// like 30s queries
var CustomBuckets = []float64{.005, .01, .025, 0.05, .1, .25, .5, 1, 2, 5, 10,
	20, 30, 40, 50, 60, 90, 120, 300,
}
