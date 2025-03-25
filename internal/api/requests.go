package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jdjaxon/pokedexcli/internal/pokedex"
)

// GetLocations -
func (c *Client) GetLocations(reqURL *string) (LocationResponse, error) {
	url := baseURL + locationEndpoint
	if reqURL != nil {
		url = *reqURL
	}

	var locResp LocationResponse
	err := c.getJSON(url, &locResp)
	if err != nil {
		return LocationResponse{}, err
	}

	return locResp, nil
}

// ExploreLocation -
func (c *Client) ExploreLocation(location string) (ExploreResponse, error) {
	if location == "" {
		return ExploreResponse{}, ErrLocation
	}

	url := baseURL + locationEndpoint + location

	var expResp ExploreResponse
	err := c.getJSON(url, &expResp)
	if err != nil {
		return ExploreResponse{}, err
	}

	return expResp, nil
}

// CatchPokemon -
func (c *Client) CatchPokemon(name string) (pokedex.Pokemon, error) {
	if name == "" {
		return pokedex.Pokemon{}, ErrPokemon
	}

	url := baseURL + pokemonEndpoint + name

	var pokemon pokedex.Pokemon
	err := c.getJSON(url, &pokemon)
	if err != nil {
		return pokedex.Pokemon{}, err
	}

	return pokemon, nil
}

func (c *Client) getJSON(url string, target any) error {
	cachedResp, ok := c.reqCache.Get(url)
	if ok {
		return json.Unmarshal(cachedResp, target)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return fmt.Errorf("failed with status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	c.reqCache.Add(url, data)

	return json.Unmarshal(data, target)
}
