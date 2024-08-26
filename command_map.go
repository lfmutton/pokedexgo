package main

import(
	"fmt"
	"log"
)

func showMap(cfg *config, arg ...string) error{
	locations, err := cfg.apiClient.GetLocationsArea(cfg.nextLocationURL)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("Location areas:\n")
	for _, loc := range locations.Results{
		fmt.Printf("-> %s\n", loc.Name)
	}
	cfg.nextLocationURL = locations.Next
	cfg.previousLocationURL = locations.Previous

	return nil
}

