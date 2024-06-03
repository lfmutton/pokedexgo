package main

import (
	"github.com/lfmutton/pokedex/Internal/Api"
	"time"
)

type config struct{
	apiClient api.Client
	nextLocationURL *string
	previousLocationURL *string
	myPokemons map[string]api.Pokemon
}

func main(){
	cfg := config{
		apiClient:	api.NewClient(time.Minute*5),
		myPokemons:	make(map[string]api.Pokemon),
	}

	Start(&cfg)
}