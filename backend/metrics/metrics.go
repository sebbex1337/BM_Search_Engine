package metrics

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

var (
	// HTTP Requests Total
	RequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of HTTP requests",
	},
		[]string{"path"},
	)

	// CPU Usage
	CPUUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "system_cpu_usage_percent",
		Help: "Current system CPU usage percentage",
	})

	// Memory Usage
	MemoryUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "system_memory_usage_percent",
		Help: "Current system Memory usage percentage",
	})

	// Disk Usage
	DiskUsage = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "system_disk_usage_percent",
		Help: "Current system Disk usage percentage",
	},
		[]string{"mountpoint"},
	)

	// Request Duration
	RequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Duration of HTTP requests in seconds",
		Buckets: prometheus.DefBuckets,
	},
		[]string{"path"},
	)
)

// Init registers all the metrics
func Init() {
	prometheus.MustRegister(RequestsTotal)
	prometheus.MustRegister(CPUUsage)
	prometheus.MustRegister(MemoryUsage)
	prometheus.MustRegister(DiskUsage)
	prometheus.MustRegister(RequestDuration)
}

// Middleware captures HTTP requests
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		path := r.URL.Path

		// Normalize path by removing trailing slash
		if len(path) > 1 && path[len(path)-1] == '/' {
			path = path[:len(path)-1]
		}
		RequestsTotal.WithLabelValues(path).Inc()

		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rec, r)

		duration := time.Since(start).Seconds()
		RequestDuration.WithLabelValues(path).Observe(duration)
	})
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

// Handler exposes metrics endpoint
func Handler() http.Handler {
	return promhttp.Handler()
}

// CollectSystemMetrics periodically collects system metrics
func CollectSystemMetrics(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		collectCPUUsage()
		collectMemoryUsage()
		collectDiskUsage()
		<-ticker.C
	}
}

func collectCPUUsage() {
	percent, err := cpu.Percent(0, false)
	if err != nil || len(percent) == 0 {
		CPUUsage.Set(0)
		return
	}
	CPUUsage.Set(percent[0])
}

func collectMemoryUsage() {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		MemoryUsage.Set(0)
		return
	}
	MemoryUsage.Set(vmStat.UsedPercent)
}

func collectDiskUsage() {
	partitions, err := disk.Partitions(false)
	if err != nil {
		return
	}

	for _, p := range partitions {
		usage, err := disk.Usage(p.Mountpoint)
		if err != nil {
			continue
		}
		DiskUsage.WithLabelValues(p.Mountpoint).Set(usage.UsedPercent)
	}
}
