package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	baseUrl = "https://pokeapi.co/api/v2/location-area/"
)

type LocationArea struct {
	Id                   int                    `json:"id"`
	Name                 string                 `json:"name"`
	GameIndex            int                    `json:"game_index"`
	EncounterMethodRates []EncounterMethodRates `json:"encounter_method_rates"`
	Location             Location               `json:"location"`
	Names                []Names                `json:"names"`
	PokemonEncounters    []PokemonEncounters    `json:"pokemon_encounters"`
}

type EncounterMethod struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Version struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type VersionDetails struct {
	Rate    int     `json:"rate"`
	Version Version `json:"version"`
}

type EncounterMethodRates struct {
	EncounterMethod EncounterMethod  `json:"encounter_method"`
	VersionDetails  []VersionDetails `json:"version_details"`
}

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Language struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Names struct {
	Name     string   `json:"name"`
	Language Language `json:"language"`
}

type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonEncounters struct {
	Pokemon        Pokemon          `json:"pokemon"`
	VersionDetails []VersionDetails `json:"version_details"`
}

func (c *Client) Explore(locationArea string) ([]PokemonEncounters, error) {
	url := fmt.Sprintf("%s/location-area/%s", baseURL, locationArea)

	if data, ok := c.cache.Get(url); ok {
		locationPoke := []PokemonEncounters{}

		err := json.Unmarshal(data, &locationPoke)
		if err != nil {
			fmt.Println(err)
			return []PokemonEncounters{}, err
		}

		return locationPoke, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return []PokemonEncounters{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return []PokemonEncounters{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return []PokemonEncounters{}, err
	}

	locationData := LocationArea{}

	err = json.Unmarshal(data, &locationData)
	if err != nil {
		fmt.Println(err)
		return []PokemonEncounters{}, err
	}

	c.cache.Add(url, data)

	return locationData.PokemonEncounters, err

}
