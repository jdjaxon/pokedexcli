package api

import (
	"encoding/json"
	"net/http"
)

func (c *Client) GetLocations(reqURL *string) (LocationResponse, error) {
	url := baseURL + "location-area"
	if reqURL != nil {
		url = *reqURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationResponse{}, err
	}
	defer resp.Body.Close()

	var areaResp LocationResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&areaResp)
	if err != nil {
		return LocationResponse{}, err
	}

	return areaResp, nil
}
