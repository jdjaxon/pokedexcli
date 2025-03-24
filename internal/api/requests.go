package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocations(reqURL *string) (LocationResponse, error) {
	url := baseURL + "location-area"
	if reqURL != nil {
		url = *reqURL
	}

	var locResp LocationResponse
	cachedResp, ok := c.reqCache.Get(url)
	if ok {
		err := json.Unmarshal(cachedResp, &locResp)
		if err != nil {
			return LocationResponse{}, err
		}
		return locResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locResp, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locResp, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		return locResp, fmt.Errorf("failed with status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return locResp, err
	}

	c.reqCache.Add(url, data)

	err = json.Unmarshal(data, &locResp)
	if err != nil {
		return LocationResponse{}, err
	}

	return locResp, nil
}
