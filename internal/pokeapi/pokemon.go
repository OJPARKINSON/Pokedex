package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemon string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemon

	if data, ok := c.cache.Get(url); ok {
		pokemonResp := PokemonResp{}

		err := json.Unmarshal(data, &pokemonResp)
		if err != nil {
			fmt.Println(err)
			return Pokemon{}, err
		}

		return Pokemon{
			Id:             pokemonResp.Id,
			Name:           pokemonResp.Name,
			BaseExperience: pokemonResp.BaseExperience,
			Stats:          pokemonResp.Stats,
			Height:         pokemonResp.Height,
			Weight:         pokemonResp.Weight,
			Types:          pokemonResp.Types,
		}, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return Pokemon{}, err
	}

	pokemonResp := PokemonResp{}

	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		fmt.Println(err)
		return Pokemon{}, err
	}

	c.cache.Add(url, data)

	return Pokemon{
		Id:             pokemonResp.Id,
		Name:           pokemonResp.Name,
		BaseExperience: pokemonResp.BaseExperience,
	}, err

}
