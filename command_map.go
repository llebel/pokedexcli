package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func commandMapb(context *cliConfig) error {
	// When going back, call map with Previous url passed as Next
	if context.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	context.Next = context.Previous
	commandMap(context)

	return nil
}

func commandMap(context *cliConfig) error {
	// Check command context
	if context.Next == "" {
		context.Next = "https://pokeapi.co/api/v2/location-area/"
	}

	// Calling PokeAPI location-area endpoitn
	// https://pokeapi.co/docs/v2#location-areas
	res, err := http.Get(context.Next)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	// Unmarchalling into struct
	location := Location{}
	err = json.Unmarshal(body, &location)
	if err != nil {
		return fmt.Errorf("error Unmarshalling %s (%v)", body, err)
	}

	// Update command context
	context.Next = location.Next
	context.Previous = location.Previous

	// Display results for this batch
	for _, result := range location.Results {
		fmt.Println(result.Name)
	}
	return nil
}

type Location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
