package main

import (
	"encoding/json"
	"fmt"
)

// Petstore for Dogs and Cats
type Petstore struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Dogs     []*Pet `json:"dogs,omitempty"`
	Cats     []*Pet `json:"cats,omitempty"`
}

// Pet description
type Pet struct {
	Name  string `json:"name"`
	Breed string `json:"breed,omitempty"`
}

// PetStoreList of Petstore
type PetStoreList []*Petstore

func main() {
	petstorelist := PetStoreList{}
	petstore := &Petstore{
		Name:     "Fuzzy's",
		Location: "New York, 5th and Broadway",
	}
	petstore.Dogs = append(petstore.Dogs,
		&Pet{
			Name:  "Whiskers",
			Breed: "Pomeranian",
		},
	)
	petstore.Dogs = append(petstore.Dogs,
		&Pet{
			Name: "Trinity",
		},
	)
	petstorelist = append(petstorelist, petstore)

	jsonString, _ := json.MarshalIndent(petstorelist, "", "\t")
	fmt.Printf("%s", jsonString)
}
