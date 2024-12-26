package internal

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"hysteria_exporter/pkg/monitor"
)

// UpdateMetrics fetches traffic and online data, then updates the Prometheus metrics
func UpdateMetrics(apiURL, token string) error {
	// 1. 获取流量数据
	trafficData, err := GetTraffic(apiURL, token)
	if err != nil {
		return fmt.Errorf("failed to get traffic data: %v", err)
	}

	// 2. 获取在线用户数据
	onlineData, err := GetOnline(apiURL, token)
	if err != nil {
		return fmt.Errorf("failed to get online data: %v", err)
	}

	// 3. 更新 Prometheus 指标

	// 更新流量数据
	for user, flow := range trafficData {
		// 更新发送流量和接收流量指标
		monitor.UpdateTrafficUsage(user, flow.Tx, flow.Rx)
	}

	// 记录当前在线用户列表
	onlineUsers := make(map[string]bool)
	for user, connections := range onlineData {
		// 更新在线用户数指标
		monitor.UpdateOnlineUser(user, connections)
		onlineUsers[user] = true
	}

	// 处理不再在线的用户，设置其在线数为 0
	// 这样会在 Prometheus 中保留这些标签，值为 0
	for user := range onlineUsers {
		if !onlineUsers[user] {
			// 用户不再在线，将其在线数设置为 0
			monitor.UpdateOnlineUser(user, 0)
		}
	}
	logrus.Info("Successfully updated Prometheus metrics")
	return nil
}
