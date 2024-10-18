package monitoring

import (
    "log"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/shirou/gopsutil/v3/mem"
)

// MemoryMonitor collects memory usage metrics
type MemoryMonitor struct {
    memoryUsage *prometheus.GaugeVec
    logger      *log.Logger
}

// NewMemoryMonitor creates a new MemoryMonitor
func NewMemoryMonitor(logger *log.Logger) *MemoryMonitor {
    return &MemoryMonitor{
        memoryUsage: prometheus.NewGaugeVec(
            prometheus.GaugeOpts{
                Name: "memory_usage_bytes",
                Help: "Memory usage in bytes",
            },
            []string{"type"},
        ),
        logger: logger,
    }
}

// Collect gathers memory usage data
func (m *MemoryMonitor) Collect() {
    v, err := mem.VirtualMemory()
    if err != nil {
        m.logger.Println("Error collecting memory usage:", err)
        return
    }

    m.memoryUsage.WithLabelValues("total").Set(float64(v.Total))
    m.memoryUsage.WithLabelValues("used").Set(v.Used)
    m.memoryUsage.WithLabelValues("free").Set(v.Free)
}

// Register registers the memory metrics with Prometheus
func (m *MemoryMonitor) Register() {
    prometheus.MustRegister(m.memoryUsage)
}
