package main

import (
	http_metric "hysteria_exporter/web/metric"
	"log"
)

func main() {
	// Start the Prometheus exporter server on port 9090
	go http_metric.StartExporterServer(":9090")
	// Log message to confirm the server is running
	log.Println("Prometheus exporter is running at http://localhost:9090/metrics")
	// Blocking main thread (keep server running)
	select {}
}
