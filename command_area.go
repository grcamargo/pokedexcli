package main

import (
	"fmt"
	"log"
)

func callBackLocationArea() error {
	pokeapiClient := NewClient()
	resp, err := pokeapiClient.listLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	return nil
}
