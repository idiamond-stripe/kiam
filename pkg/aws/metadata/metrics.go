package metadata

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	handlerTimer = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "kiam",
			Subsystem: "metadata",
			Name:      "handlerTiming",
			Help:      "Bucketed histogram of handler timings",

			// 1ms to 5min
			Buckets: prometheus.ExponentialBuckets(.001, 2, 13),
		},
		[]string{"handler"},
	)

	findRoleError = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "kiam",
			Subsystem: "metadata",
			Name:      "findRoleErrorCount",
			Help:      "Number of errors finding the role for a pod",
		},
		[]string{"handler"},
	)

	emptyRole = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "kiam",
			Subsystem: "metadata",
			Name:      "emptyRoleCount",
			Help:      "Number of empty roles returned",
		},
		[]string{"handler"},
	)

	success = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "kiam",
			Subsystem: "metadata",
			Name:      "successCount",
			Help:      "Number of successful responses from a handler",
		},
		[]string{"handler"},
	)

	responses = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "kiam",
			Subsystem: "metadata",
			Name:      "responsesCount",
			Help:      "Responses from mocked out metadata handlers",
		},
		[]string{"handler", "code"},
	)
)

func init() {
	prometheus.MustRegister(handlerTimer)
	prometheus.MustRegister(findRoleError)
	prometheus.MustRegister(emptyRole)
	prometheus.MustRegister(success)
	prometheus.MustRegister(responses)
}
