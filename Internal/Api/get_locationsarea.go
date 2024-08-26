package api

import(
	"io"
	"net/http"
	"fmt"
	"encoding/json"
)

func (c *Client) GetLocationsArea(page *string) (Locations ,error) {
	endpoint := "/location-area"
	fullURL := baseUrl + endpoint
	if page != nil{
		fullURL = *page
	}

	data, ok := c.cache.Get(fullURL)
	if ok{

		LocationAreas := Locations{}
		err := json.Unmarshal(data, &LocationAreas)
		if err != nil{
			return Locations{}, err
		}
	
		return LocationAreas, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil{
		return Locations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil{
		return Locations{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399{
		return Locations{}, fmt.Errorf("Bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil{
		return Locations{}, err
	}

	LocationAreas := Locations{}
	err = json.Unmarshal(data, &LocationAreas)
	if err != nil{
		return Locations{}, err
	}

	c.cache.Add(fullURL, data)

	return LocationAreas, err
}

func (c *Client) GetLocationInfo(local string) (LocationInfo ,error) {
	endpoint := fmt.Sprintf("/location-area/" + local)
	fullURL := baseUrl + endpoint

	data, ok := c.cache.Get(fullURL)
	if ok{

		location := LocationInfo{}
		err := json.Unmarshal(data, &location)
		if err != nil{
			return LocationInfo{}, err
		}
	
		return location, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil{
		return LocationInfo{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil{
		return LocationInfo{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399{
		return LocationInfo{}, fmt.Errorf("Bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil{
		return LocationInfo{}, err
	}

	location := LocationInfo{}
	err = json.Unmarshal(data, &location)
	if err != nil{
		return LocationInfo{}, err
	}

	c.cache.Add(fullURL, data)

	return location, err
}