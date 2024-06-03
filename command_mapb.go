package main

import (
	"fmt"
	"log"
)

func showMapb(cfg *config, arg ...string) error{
	if cfg.previousLocationURL == nil{
		return fmt.Errorf("Still in the first page")
	}
	locations, err := cfg.apiClient.GetLocationsArea(cfg.previousLocationURL)

	if err != nil{
		log.Fatal(err)
	}

	for _, loc := range locations.Results{
		fmt.Printf("-> %v\n", loc.Name)
	}
	cfg.nextLocationURL = locations.Next
	cfg.previousLocationURL = locations.Previous

	return nil
}