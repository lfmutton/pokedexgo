package main

import(
	"fmt"
	"math/rand"
)

func catchPokemon(cfg *config, arg ...string) error{
	if arg == nil{
		return fmt.Errorf("Arguments not found")
	}

	pokemonInfo, err := cfg.apiClient.GetPokemonInfo(arg[0])

	if err != nil{
		return err
	}

	pokeExp := pokemonInfo.BaseExperience
	threshold := 100
	for i := 0; pokeExp > i; i += 25{
		threshold -= 10
	}
	if threshold < 1{
		threshold = 1
	}
	number := rand.Intn(150)
	fmt.Println(pokeExp, number, threshold)
	if number <= threshold{
		fmt.Printf("%v was caugh!!\n", arg[0])
		fmt.Printf("You may now inspect it with the inspect command.\n")
		cfg.myPokemons[pokemonInfo.Name] = pokemonInfo
		return nil
	}
	return fmt.Errorf("The pokemon ran away!")
}