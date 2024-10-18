package monitoring

import (
	"fmt"
    "log"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/shirou/gopsutil/v3/cpu"
)

// CPUMonitor collects CPU usage metrics
type CPUMonitor struct {
    cpuUsage *prometheus.GaugeVec
    logger   *log.Logger
}

// NewCPUMonitor creates a new CPUMonitor
func NewCPUMonitor(logger *log.Logger) *CPUMonitor {
    return &CPUMonitor{
        cpuUsage: prometheus.NewGaugeVec(
            prometheus.GaugeOpts{
                Name: "cpu_usage_percent",
                Help: "CPU usage percentage per core",
            },
            []string{"core"},
        ),
        logger: logger,
    }
}

// Collect gathers CPU usage data
func (c *CPUMonitor) Collect() {
    percentages, err := cpu.Percent(0, true)
    if err != nil {
        c.logger.Println("Error collecting CPU usage:", err)
        return
    }

    for i, perc := range percentages {
        coreLabel := fmt.Sprintf("core_%d", i)
        c.cpuUsage.WithLabelValues(coreLabel).Set(perc)
    }
}

// Register registers the CPU metrics with Prometheus
func (c *CPUMonitor) Register() {
    prometheus.MustRegister(c.cpuUsage)
}
