package main

import(
	"fmt"
)

func exploreMap(cfg *config, arg ...string) error{
	if arg == nil{
		return fmt.Errorf("Arguments not found")
	}
	explore, err := cfg.apiClient.GetLocationInfo(arg[0])
	if err != nil{
		return err
	}
	fmt.Printf("Exploring %v\n", arg[0])
	fmt.Println("Found Pokemon:")
	for _, result := range explore.PokemonEncounters{
		fmt.Printf("-> %v\n", result.Pokemon.Name)
	}
	return nil
}