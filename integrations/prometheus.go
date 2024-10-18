package integrations

import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/amahmood561/go-performance-monitor/internal/custommetrics"
    "github.com/amahmood561/go-performance-monitor/internal/monitoring"
)

// RegisterMetrics registers all provided metrics with Prometheus
func RegisterMetrics(metrics ...prometheus.Registerer) {
    for _, metric := range metrics {
        prometheus.MustRegister(metric)
    }
}

// Alternatively, if you have specific metric collectors:
func RegisterMetrics(cpu *monitoring.CPUMonitor, memory *monitoring.MemoryMonitor, goroutine *monitoring.GoroutineMonitor, counters ...*custommetrics.Counter) {
    cpu.Register()
    memory.Register()
    goroutine.Register()
    for _, counter := range counters {
        counter.Register()
    }
}
