package app

import (
	"fmt"
	"log"
	"os"
	"time"

	"hysteria_exporter/internal"
)

func MetricJob(interval time.Duration) {

	// Read the API_URL and PORT environment variables
	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = "127.0.0.1"
		log.Printf("API_URL not set, defaulting to %s", apiURL)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "9980"
		log.Printf("PORT not set, defaulting to %s", port)
	}

	// Build the full API URL
	fullAPIURL := fmt.Sprintf("http://%s:%s", apiURL, port)

	// Read the SECRET environment variable
	secret := os.Getenv("SECRET")
	if secret == "" {
		log.Fatal("SECRET is not set, exiting.")

	}

	// Start the metrics update loop
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("Updating metrics...")
		err := internal.UpdateMetrics(fullAPIURL, secret)
		if err != nil {
			log.Printf("Failed to update metrics: %v", err)
		} else {
			log.Println("Metrics updated successfully")
		}
	}
}
