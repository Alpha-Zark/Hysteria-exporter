package internal

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"hysteria_exporter/pkg/struct"
	"io"
	"io/ioutil"
	"net/http"
)

func GetTraffic(url, token string) (_struct.TrafficInfo, error) {
	// Create a new HTTP request with the Authorization header
	req, err := http.NewRequest("GET", url+"/traffic", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Add Authorization header
	req.Header.Add("Authorization", token)

	// Send the request and get the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logrus.Errorf("failed to close body: %v", err)
		}
	}(resp.Body)

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch traffic data, status code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Parse the JSON response into TrafficInfo
	var traffic _struct.TrafficInfo
	err = json.Unmarshal(body, &traffic)
	if err != nil {
		return nil, fmt.Errorf("failed to parse traffic data: %v", err)
	}

	return traffic, nil
}
