package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	sessionsGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "pgweb_sessions_count",
		Help: "Total number of database sessions",
	})

	queriesCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "pgweb_queries_count",
		Help: "Total number of custom queries executed",
	})

	healtyGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "pgweb_healty",
		Help: "Server health status",
	})
)

func IncrementQueriesCount() {
	queriesCounter.Inc()
}

func SetSessionsCount(val int) {
	sessionsGauge.Set(float64(val))
}

func SetHealty(val bool) {
	healthy := 0.0
	if val {
		healthy = 1.0
	}
	healtyGauge.Set(float64(healthy))
}
