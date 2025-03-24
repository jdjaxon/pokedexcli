package client

import (
	"encoding/json"
	"net/http"
)

const baseUrl = "https://pokeapi.co/api/v2/"

func GetLocationAreas() ([]string, error) {
	url := baseUrl + "location-area"
	client := &http.Client{}
	resp, err := client.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return []string{}, err
	}

	var areas LocationResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&areas)
	if err != nil {
		return []string{}, err
	}

	var locationNames []string
	for area := range areas.results {
		areas = append(areas, area.name)
	}

	return areas, nil
}
