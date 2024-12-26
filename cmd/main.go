package main

import (
	"hysteria_exporter/web"
	"log"
)

func main() {
	// Start the Prometheus exporter server on port 9090
	go web.StartExporterServer(":9090")
	// Log message to confirm the server is running
	log.Println("Prometheus exporter is running at http://localhost:9090/metrics")
	// Blocking main thread (keep server running)
	select {}
}
