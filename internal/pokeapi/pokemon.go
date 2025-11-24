package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (RespPokemon, error) {
	url := fmt.Sprintf("%s/pokemon/%s", baseURL, name)

	if cachedVal, ok := c.cache.Get(url); ok {
		pokemon := RespPokemon{}
		if err := json.Unmarshal(cachedVal, &pokemon); err != nil {
			return RespPokemon{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	c.cache.Add(url, dat)

	pokemon := RespPokemon{}
	if err := json.Unmarshal(dat, &pokemon); err != nil {
		return RespPokemon{}, err
	}

	return pokemon, nil
}
