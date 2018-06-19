package sts

import "github.com/prometheus/client_golang/prometheus"

var (
	cacheHit = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "kiam",
			Subsystem: "sts",
			Name:      "cacheHit",
			Help:      "Number of cache hits to the metadata cache",
		},
	)

	cacheMiss = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "kiam",
			Subsystem: "sts",
			Name:      "cacheMiss",
			Help:      "Number of cache misses to the metadata cache",
		},
	)

	errorIssuing = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "kiam",
			Subsystem: "sts",
			Name:      "errorIssuing",
			Help:      "Error issuing credentials",
		},
	)

	assumeRole = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "kiam",
			Subsystem: "sts",
			Name:      "assumeRole",
			Help:      "Bucketed histogram of assumeRole timings",

			// 1ms to 5min
			Buckets: prometheus.ExponentialBuckets(1, 2, 13),
		},
	)

	assumeRoleExecuting = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "kiam",
			Subsystem: "sts",
			Name:      "assumeRoleExecuting",
			Help:      "Number of assume role calls currently executing",
		},
	)
)

func init() {
	prometheus.MustRegister(cacheHit)
	prometheus.MustRegister(cacheMiss)
	prometheus.MustRegister(errorIssuing)
	prometheus.MustRegister(assumeRole)
	prometheus.MustRegister(assumeRoleExecuting)
}
