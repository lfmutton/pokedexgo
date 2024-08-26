package api

import(
	"io"
	"net/http"
	"fmt"
	"encoding/json"
)

func (c *Client) GetPokemonInfo(name string) (Pokemon, error) {
	endpoint := fmt.Sprintf("/pokemon/" + name)
	fullURL := baseUrl + endpoint

	data, ok := c.cache.Get(fullURL)
	if ok{

		pokemon := Pokemon{}
		err := json.Unmarshal(data, &pokemon)
		if err != nil{
			return Pokemon{}, err
		}
	
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil{
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil{
		return Pokemon{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399{
		return Pokemon{}, fmt.Errorf("Bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil{
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil{
		return Pokemon{}, err
	}

	c.cache.Add(fullURL, data)

	return pokemon, err
}