package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type locationResp struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []location `json:"results"`
}

func (c *Client) ListLocations(locationUrl *string) (locationResp, error) {
	url := baseURL + "/location-area"
	if locationUrl != nil {
		url = *locationUrl
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return locationResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationResp{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return locationResp{}, err
	}

	locationResp := locationResp{}

	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		fmt.Println(err)
		return locationResp, err
	}

	return locationResp, err
}
