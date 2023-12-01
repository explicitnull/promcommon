package promcommon

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

func RegisterGoBuildInfo() {
	prometheus.MustRegister(collectors.NewBuildInfoCollector())
}
