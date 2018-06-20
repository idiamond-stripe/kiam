package server

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	rpcTimer = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "kiam",
			Subsystem: "server",
			Name:      "rpc",
			Help:      "Bucketed histogram of rpc call timings",

			// 1ms to 5min
			Buckets: prometheus.ExponentialBuckets(.001, 2, 13),
		},
		[]string{"rpc", "type"},
	)
)

func init() {
	prometheus.MustRegister(rpcTimer)
}
