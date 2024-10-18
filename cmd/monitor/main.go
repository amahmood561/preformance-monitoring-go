package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/amahmood561/go-performance-monitor/internal/custommetrics"
    "github.com/amahmood561/go-performance-monitor/internal/monitoring"
    "github.com/amahmood561/go-performance-monitor/integrations"
    "github.com/amahmood561/go-performance-monitor/pkg/logger"
    "github.com/amahmood561/go-performance-monitor/web"
)

func main() {
    // Initialize logger
    log := logger.NewLogger()

    // Initialize monitoring modules
    cpuMonitor := monitoring.NewCPUMonitor(log)
    memoryMonitor := monitoring.NewMemoryMonitor(log)
    goroutineMonitor := monitoring.NewGoroutineMonitor(log)

    // Initialize custom metrics
    customCounter := custommetrics.NewCounter("custom_requests_total", "Total number of custom requests")
    customGauge := custommetrics.NewGauge("custom_metric", "A custom gauge metric")

    // Register metrics with Prometheus
    integrations.RegisterMetrics(cpuMonitor, memoryMonitor, goroutineMonitor, customCounter, customGauge)

    // Start monitoring in separate goroutines
    go func() {
        for {
            cpuMonitor.Collect()
            time.Sleep(1 * time.Second)
        }
    }()

    go func() {
        for {
            memoryMonitor.Collect()
            time.Sleep(1 * time.Second)
        }
    }()

    go func() {
        for {
            goroutineMonitor.Collect()
            time.Sleep(1 * time.Second)
        }
    }()

    // Start custom metrics simulation
    go func() {
        for {
            customCounter.Increment()
            customGauge.Set(float64(customCounter.Value()))
            time.Sleep(2 * time.Second)
        }
    }()

    // Start Prometheus metrics server
    go func() {
        fmt.Println("Prometheus metrics available at :2112/metrics")
        log.Fatal(http.ListenAndServe(":2112", nil))
    }()

    // Start web dashboard
    fmt.Println("Web dashboard available at :8080")
    web.StartWebServer()

    // Prevent the main goroutine from exiting
    select {}
}
