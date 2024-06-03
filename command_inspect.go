package main

import(
	"fmt"
)

func inspectPokemon(cfg *config, arg ...string) error{
	name := arg[0]
	pokeInfo, ok := cfg.myPokemons[name]
	if !ok{
		return fmt.Errorf("This pokemon does not exist in your party.")
	}
	fmt.Printf("Name: %v\n", pokeInfo.Name)
	fmt.Printf("height: %v\n", pokeInfo.Height)
	fmt.Printf("Weight: %v\n", pokeInfo.Weight)
	fmt.Printf("Stats:\n")
	for _, stats := range pokeInfo.Stats{
		fmt.Printf("  -%s: %d\n", stats.Stat.Name, stats.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, types := range pokeInfo.Types{
		fmt.Printf("  - %s\n", types.Type.Name)
	}
	return nil
}