package _struct

import "encoding/json"

type FlowInfo struct {
	Tx int `json:"tx"` // 发送流量
	Rx int `json:"rx"` // 接收流量
}
type TrafficInfo map[string]FlowInfo

func ParseTrafficData(data string) (TrafficInfo, error) {
	var traffic TrafficInfo
	err := json.Unmarshal([]byte(data), &traffic)
	if err != nil {
		return nil, err
	}
	return traffic, nil
}
func SerializeTrafficData(traffic TrafficInfo) (string, error) {
	trafficJSON, err := json.MarshalIndent(traffic, "", "  ")
	if err != nil {
		return "", err
	}
	return string(trafficJSON), nil
}

type OnlineInfo map[string]int

func ParseOnlineData(data string) (OnlineInfo, error) {
	var online OnlineInfo
	err := json.Unmarshal([]byte(data), &online)
	if err != nil {
		return nil, err
	}
	return online, nil
}

func SerializeOnlineData(online OnlineInfo) (string, error) {
	onlineJSON, err := json.MarshalIndent(online, "", "  ")
	if err != nil {
		return "", err
	}
	return string(onlineJSON), nil
}
