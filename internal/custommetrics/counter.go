package custommetrics

import (
    "sync"

    "github.com/prometheus/client_golang/prometheus"
)

// Counter represents a custom counter metric
type Counter struct {
    metric *prometheus.CounterVec
    name   string
    help   string
    mu     sync.Mutex
}

// NewCounter creates a new Counter
func NewCounter(name, help string) *Counter {
    c := &Counter{
        metric: prometheus.NewCounterVec(
            prometheus.CounterOpts{
                Name: name,
                Help: help,
            },
            []string{},
        ),
        name: name,
        help: help,
    }
    return c
}

// Increment increases the counter by 1
func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.metric.WithLabelValues().Inc()
}

// Value returns the current value of the counter
func (c *Counter) Value() float64 {
    // Prometheus counters do not support retrieving current value directly
    // This method can be used to track value internally if needed
    return 0 // Placeholder
}

// Register registers the counter with Prometheus
func (c *Counter) Register() {
    prometheus.MustRegister(c.metric)
}
