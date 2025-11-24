package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(area string) (RespLocationArea, error) {
	url := fmt.Sprintf("%s/location-area/%s", baseURL, area)

	if cachedVal, ok := c.cache.Get(url); ok {
		locationArea := RespLocationArea{}
		if err := json.Unmarshal(cachedVal, &locationArea); err != nil {
			return RespLocationArea{}, err
		}
		return locationArea, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationArea{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationArea{}, err
	}

	c.cache.Add(url, dat)

	locationArea := RespLocationArea{}
	if err := json.Unmarshal(dat, &locationArea); err != nil {
		return RespLocationArea{}, err
	}

	return locationArea, nil
}
