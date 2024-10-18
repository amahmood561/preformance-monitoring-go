package custommetrics

import (
    "sync"

    "github.com/prometheus/client_golang/prometheus"
)

// Gauge represents a custom gauge metric
type Gauge struct {
    metric *prometheus.GaugeVec
    name   string
    help   string
    mu     sync.Mutex
}

// NewGauge creates a new Gauge
func NewGauge(name, help string) *Gauge {
    g := &Gauge{
        metric: prometheus.NewGaugeVec(
            prometheus.GaugeOpts{
                Name: name,
                Help: help,
            },
            []string{},
        ),
        name: name,
        help: help,
    }
    return g
}

// Set sets the gauge to a specific value
func (g *Gauge) Set(value float64) {
    g.mu.Lock()
    defer g.mu.Unlock()
    g.metric.WithLabelValues().Set(value)
}

// Register registers the gauge with Prometheus
func (g *Gauge) Register() {
    prometheus.MustRegister(g.metric)
}
