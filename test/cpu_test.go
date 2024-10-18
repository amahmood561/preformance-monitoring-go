package test

import (
    "log"
    "testing"

    "github.com/yourusername/go-performance-monitor/internal/monitoring"
)

func TestCPUMonitor(t *testing.T) {
    logger := log.New(nil, "", 0)
    cpuMonitor := monitoring.NewCPUMonitor(logger)

    // Test Collect method
    cpuMonitor.Collect()

    // Since Prometheus metrics are registered globally, check if they are registered
    // Note: In a real test, you might mock the CPU usage
}
