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

	var locResp LocationResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&locResp)
	if err != nil {
		return LocationResponse{}, err
	}

	return locResp, nil
}
