package main

import (
	"encoding/json"
	"testing"
)

// how to use json struct tag
// 1. omitemtpy
// 2. "-" disable output
type BirdOmiteEmpty struct {
	Species string `json:"birdType"`
	// we can set the "omitempty" property as part of the JSON tag
	Description string `json:"what it does,omitempty"`
}

type BirdNoOutput struct {
	Species     string `json:"birdType"`
	Description string `json:"-"`
}

//Ignoring Empty Fields
func TestIgnoringEmptyFields(t *testing.T) {
	birdByte, _ := json.Marshal(&BirdOmiteEmpty{Species: "Pigeon"})
	t.Logf("%s", string(birdByte))
	// {"birdType":"Pigeon"}

	birdByte, _ = json.Marshal(&BirdNoOutput{Description: "Pigeon can fly"})
	t.Logf("%s", string(birdByte))
	// {"birdType":""}
}


