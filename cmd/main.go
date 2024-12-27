package main

import (
	"hysteria_exporter/cmd/app"
	http_metric "hysteria_exporter/web/metric"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	// Start the Prometheus exporter server on port 2112
	go http_metric.StartExporterServer(":2112")
	// Log message to confirm the server is running
	log.Println("Prometheus exporter is running at http://localhost:9090/metrics")
	// Blocking main thread (keep server running)
	intervalStr := os.Getenv("METRICS_INTERVAL")
	interval := 15 // 默认值为15秒
	if intervalStr != "" {
		var err error
		interval, err = strconv.Atoi(intervalStr)
		if err != nil {
			log.Fatalf("Invalid METRICS_INTERVAL value: %v", err)
		}
	}

	go app.MetricJob(time.Duration(interval) * time.Second)

	select {}
}
