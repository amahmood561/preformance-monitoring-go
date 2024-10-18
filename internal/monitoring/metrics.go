package monitoring

// MetricCollector defines the interface for metric collectors
type MetricCollector interface {
    Collect()
    Register()
}
