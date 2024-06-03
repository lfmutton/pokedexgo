package main

import (
    "os"
    "bufio"
    "fmt"
    "strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func commandHelp(cfg *config, arg ...string) error{
    fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandError(cfg *config, arg ...string) error{
    os.Exit(0)
	return nil
}

func getCommands() map[string]cliCommand{
    return map[string]cliCommand{
        "help": {
            name:        "help",
            description: "Displays a help message",
            callback:    commandHelp,
        },
        "exit": {
            name:        "exit",
            description: "Exit the Pokedex",
            callback:    commandError,
        },
		"map": {
			name: 		 "map",
			description: "Show 20 locations",
			callback:    showMap,
		},
		"mapb": {
			name:		 "mapb",
			description: "Show the last 20 locations",
			callback:    showMapb,
		},
		"explore":{
			name:		 "explore <location_area>",
			description: "Show the pokemons encounters in the area",
			callback:	 exploreMap,
		},
		"catch":{
			name: 		 "catch <pokemon>",
			description: "Try to catch a pokemon",
			callback:    catchPokemon,
		},
		"inspect":{
			name: 		 "inspect <pokemon>",
			description: "Inspect the stats of the pokemon",
			callback:    inspectPokemon,
		},
		"pokedex":{
			name: 		 "pokedex",
			description: "Show your pokemons",
			callback:    showYourPokemons,
		},
    }
}

func Start(cfg *config) {
    scanner := bufio.NewScanner(os.Stdin)
    allCommand := make(map[string]cliCommand)
    allCommand = getCommands()
    for {
        fmt.Printf("Pokedex > ")
        scanner.Scan()
        text := Clean(scanner.Text())
		if len(text)==0{
			continue
		}
		command := text[0]
		comd, ok := allCommand[command]
		if !ok{
			fmt.Printf("Invalid Command '%v'\n", command)
			continue
		}
		var arg []string
		if len(text) > 1 {
			arg = text[1:]
		}
		err := comd.callback(cfg, arg...)
		if err != nil{
			fmt.Println(err)
		}
	}
} 

func Clean(text string) []string{
    output := strings.ToLower(text)
	word := strings.Fields(output)
    return word
}