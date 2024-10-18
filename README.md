# Go Performance Monitoring Toolkit

![Go Logo](https://golang.org/lib/godoc/images/go-logo-blue.svg)

A lightweight and customizable performance monitoring toolkit for Go applications. This toolkit provides real-time metrics on CPU usage, memory consumption, goroutines, and more. It integrates seamlessly with Prometheus for metric scraping and visualization, making it an essential tool for devs aiming to optimize and monitor their Go applications effectively.

## Table of Contents

- [Features](#features)
- [Demo](#demo)
- [Installation](#installation)
- [Usage](#usage)
  - [Running the Toolkit](#running-the-toolkit)
  - [Accessing Metrics](#accessing-metrics)
  - [Viewing the Dashboard](#viewing-the-dashboard)
- [Prometheus Integration](#prometheus-integration)
- [Custom Metrics](#custom-metrics)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Features

- **CPU Monitoring:** Track CPU utilization per core.
- **Memory Monitoring:** Monitor memory usage, including total, used, and free memory.
- **Goroutine Monitoring:** Keep an eye on the number of active goroutines.
- **Custom Metrics:** Define and register custom counters and gauges tailored to your application's needs.
- **Prometheus Integration:** Export metrics in a Prometheus-compatible format for scraping and visualization.
- **Web Dashboard:** View real-time metrics through a built-in web interface.
- **Logging:** Comprehensive logging for monitoring and debugging purposes.
- **Extensible Architecture:** Easily extend the toolkit with additional metrics and integrations.

## Demo

![Dashboard Screenshot](./web/templates/dashboard_screenshot.png)

*Example screenshot of the web dashboard displaying real-time metrics.*

## Installation

### Prerequisites

- **Go:** Ensure you have Go installed (version 1.20 or later). You can download it from [here](https://golang.org/dl/).
- **Git:** To clone the repository. Download from [here](https://git-scm.com/downloads).
- **Prometheus:** (Optional) For advanced metrics scraping and visualization. Install instructions can be found [here](https://prometheus.io/docs/prometheus/latest/installation/).

### Clone the Repository

```bash
git clone https://github.com/amahmood561/go-performance-monitor.git
cd go-performance-monitor
```

### Build the Application

```bash
go build -o monitor ./cmd/monitor
```

### Install Dependencies

The project uses Go modules for dependency management. Ensure all dependencies are installed:

```bash
go mod tidy
```

## Usage

### Running the Toolkit

After building the application, you can run it using the following command:

```bash
./monitor
```

Alternatively, you can use the provided startup script:

```bash
./scripts/start.sh
```

### Accessing Metrics

Once the application is running, metrics are exposed at:

- **Prometheus Metrics Endpoint:** `http://localhost:2112/metrics`

### Viewing the Dashboard

The toolkit includes a simple web dashboard to visualize metrics in real-time. Access it at:

- **Web Dashboard:** `http://localhost:8080/`

**Note:** The dashboard currently displays raw metrics data. Future updates will include enhanced visualization using charts and graphs.

## Prometheus Integration

Prometheus is a powerful metrics aggregation and alerting tool. Integrating this toolkit with Prometheus allows for advanced monitoring and visualization.

### Setting Up Prometheus

1. **Download Prometheus:**

   Visit the [Prometheus Downloads](https://prometheus.io/download/) page and download the latest version for your operating system.

2. **Configure Prometheus:**

   Add the following job to your `prometheus.yml` configuration file:

   ```yaml
   scrape_configs:
     - job_name: 'go-performance-monitor'
       static_configs:
         - targets: ['localhost:2112']
   ```

3. **Start Prometheus:**

   ```bash
   ./prometheus --config.file=prometheus.yml
   ```

4. **Access Prometheus Dashboard:**

   Navigate to `http://localhost:9090/` to access the Prometheus web interface.

### Querying Metrics

You can use Prometheus's query language (PromQL) to query and visualize metrics. For example:

- **CPU Usage:**

  ```promql
  cpu_usage_percent
  ```

- **Memory Usage:**

  ```promql
  memory_usage_bytes{type="used"}
  ```

- **Goroutines:**

  ```promql
  goroutines_total
  ```

## Custom Metrics

The toolkit allows you to define and register custom metrics to monitor specific aspects of your application.

### Creating Custom Counters

```go
import (
    "github.com/amahmood561/go-performance-monitor/internal/custommetrics"
)

func main() {
    customCounter := custommetrics.NewCounter("custom_requests_total", "Total number of custom requests")
    customCounter.Register()

    // Increment the counter
    customCounter.Increment()
}
```

### Creating Custom Gauges

```go
import (
    "github.com/amahmood561/go-performance-monitor/internal/custommetrics"
)

func main() {
    customGauge := custommetrics.NewGauge("custom_metric", "A custom gauge metric")
    customGauge.Register()

    // Set the gauge value
    customGauge.Set(42.0)
}
```

## Configuration

The toolkit can be customized by modifying configuration files or environment variables. Future updates may include a dedicated configuration file for more granular control.

### Environment Variables

- **PROMETHEUS_PORT:** Port for Prometheus metrics endpoint (default: `2112`).
- **DASHBOARD_PORT:** Port for the web dashboard (default: `8080`).

### Example

To change the Prometheus metrics port to `2200`:

```bash
export PROMETHEUS_PORT=2200
./monitor
```

## Contributing

Contributions are welcome! Whether you're fixing bugs, improving documentation, or adding new features, your help is appreciated.

### How to Contribute

1. **Fork the Repository**

   Click the "Fork" button at the top-right corner of the repository page to create your own fork.

2. **Clone Your Fork**

   ```bash
   git clone https://github.com/amahmood561/go-performance-monitor.git
   cd go-performance-monitor
   ```

3. **Create a Branch**

   ```bash
   git checkout -b feature/your-feature-name
   ```

4. **Make Your Changes**

   Implement your feature or fix. Ensure your code follows the project's coding standards.

5. **Run Tests**

   ```bash
   go test ./...
   ```

6. **Commit Your Changes**

   ```bash
   git add .
   git commit -m "Add your commit message"
   ```

7. **Push to Your Fork**

   ```bash
   git push origin feature/your-feature-name
   ```

8. **Create a Pull Request**

   Go to the original repository and click "Compare & pull request". Provide a clear description of your changes.

### Coding Standards

- Follow Go's [idiomatic conventions](https://golang.org/doc/effective_go).
- Write clear and concise commit messages.
- Ensure all tests pass before submitting.

### Reporting Issues

If you encounter any bugs or have suggestions for improvements, please open an [issue](https://github.com/amahmood561/go-performance-monitor/issues).

## Contact

For any questions, suggestions, or feedback, feel free to reach out:

- **Email:** amahmood561@gmail.com
- **GitHub:** [amahmood561](https://github.com/amahmood561)

---

