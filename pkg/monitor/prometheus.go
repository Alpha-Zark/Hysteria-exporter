package monitor

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Define the traffic usage metric
var (
	Hysteria2TrafficUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "hysteria2_traffic_usage",
			Help: "Tracks the traffic usage (Tx/Rx) for each user in hysteria2.",
		},
		[]string{"user", "direction"}, // Labels: user and direction (tx or rx)
	)

	Hysteria2OnlineUser = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "hysteria2_online_user",
			Help: "Tracks the number of online devices (instances) for each user in hysteria2.",
		},
		[]string{"user"}, // Label: user
	)
)

//func init() {
//	// Register the custom metrics with Prometheus
//	prometheus.MustRegister(Hysteria2TrafficUsage)
//	prometheus.MustRegister(Hysteria2OnlineUser)
//}

// Update the traffic usage for a user
func UpdateTrafficUsage(user string, tx, rx int) {
	// Set the tx and rx traffic values for the user
	Hysteria2TrafficUsage.WithLabelValues(user, "tx").Set(float64(tx))
	Hysteria2TrafficUsage.WithLabelValues(user, "rx").Set(float64(rx))
}

// Update the online user count (number of devices) for a user
func UpdateOnlineUser(user string, deviceCount int) {
	// Set the number of devices (connections) for the user
	Hysteria2OnlineUser.WithLabelValues(user).Set(float64(deviceCount))
}
