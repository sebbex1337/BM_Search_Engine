package metrics

import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
)

// HTTP Metrics

var (
	HTTPRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP Requests",
		},
		[]string{"path","method","status"},
	)
	HTTPRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path","method","status"},
	)
)

// System Metrics

var(
	CPUUsage = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_usage",
		Help: "CPU usage in percentage",
	})
	MemoryUsage = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "memory_usage",
		Help: "Memory usage in percentage",
	})
	DiskUsage = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "disk_usage",
		Help: "Disk usage in percentage",
	})
)