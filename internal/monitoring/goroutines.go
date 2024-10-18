package monitoring

import (
    "log"
    "runtime"

    "github.com/prometheus/client_golang/prometheus"
)

// GoroutineMonitor collects the number of active goroutines
type GoroutineMonitor struct {
    goroutines *prometheus.GaugeVec
    logger     *log.Logger
}

// NewGoroutineMonitor creates a new GoroutineMonitor
func NewGoroutineMonitor(logger *log.Logger) *GoroutineMonitor {
    return &GoroutineMonitor{
        goroutines: prometheus.NewGaugeVec(
            prometheus.GaugeOpts{
                Name: "goroutines_total",
                Help: "Number of active goroutines",
            },
            []string{},
        ),
        logger: logger,
    }
}

// Collect gathers goroutine data
func (g *GoroutineMonitor) Collect() {
    count := runtime.NumGoroutine()
    g.goroutines.WithLabelValues().Set(float64(count))
}

// Register registers the goroutine metrics with Prometheus
func (g *GoroutineMonitor) Register() {
    prometheus.MustRegister(g.goroutines)
}
