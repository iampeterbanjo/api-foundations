package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/davecgh/go-spew/spew"
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

	jsonBlob, err := ioutil.ReadFile("example2.json")
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
	}

	err = json.Unmarshal(jsonBlob, &petstorelist)
	if err != nil {
		fmt.Printf("Error decoding json: %s\n", err)
	}

	spew.Dump(petstorelist)
}
