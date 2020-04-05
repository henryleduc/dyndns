package ip

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type response struct {
	IPAddr string `json:"ip"`
}

func GetIPv4() (string, error) {
	httpClient := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.ipify.org/?format=json", nil)
	if err != nil {
		log.Fatal("failed to create request")
		return "", err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatalf("failed to retrieve IPv4 request (%s): %v", resp.Status, err)
		return "", err
	}

	jsonBody := &response{}
	err = json.NewDecoder(resp.Body).Decode(jsonBody)
	if err != nil {
		log.Fatalf("failed to unmarshal json response")
		return "", err
	}

	if jsonBody.IPAddr == "" {
		return "", errors.New("failed to retrieve ip from api.ipify.com")
	}

	return jsonBody.IPAddr, nil
}
