package web

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"hysteria_exporter/pkg/monitor"
	"log"
	"net/http"
)

func StartExporterServer(address string) {
	// Register your custom metrics (from pkg/monitor)
	prometheus.MustRegister(monitor.Hysteria2TrafficUsage)
	prometheus.MustRegister(monitor.Hysteria2OnlineUser)

	// Expose metrics at /metrics endpoint
	http.Handle("/metrics", promhttp.Handler())

	// Log the server address and start the server
	fmt.Println("Prometheus exporter is running at", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
