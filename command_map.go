package main

import (
	"fmt"
)

var pokeURl string = ""

func commandMap(cfg *config) error {
	var locationUrl *string
	if cfg.nextLocationsUrl != nil {
		locationUrl = cfg.nextLocationsUrl
	}

	locationResp, err := cfg.pokeApiClient.ListLocations(locationUrl)
	if err != nil {
		fmt.Println(err)
		return err
	}

	cfg.nextLocationsUrl = &locationResp.Next
	cfg.previousLocationsUrl = &locationResp.Previous

	for _, results := range locationResp.Results {
		fmt.Println(results.Name)
	}

	return nil
}

func commandMapB(cfg *config) error {
	var locationUrl *string
	if cfg.previousLocationsUrl != nil {
		locationUrl = cfg.previousLocationsUrl
	}

	locationResp, err := cfg.pokeApiClient.ListLocations(locationUrl)
	if err != nil {
		fmt.Println(err)
		return err
	}

	cfg.nextLocationsUrl = &locationResp.Next
	cfg.previousLocationsUrl = &locationResp.Previous

	for _, results := range locationResp.Results {
		fmt.Println(results.Name)
	}

	return nil

}
