package main

import(
	"fmt"
)

func showYourPokemons(cfg *config, arg ...string) error{
	for _ ,pokemon := range cfg.myPokemons{
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}