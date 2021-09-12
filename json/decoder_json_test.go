package json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestDecoder(t *testing.T) {
	file, _ := os.Open("file.json")
	jsonDecoder := json.NewDecoder(file)

	person := &Person{}
	jsonDecoder.Decode(person)

	fmt.Println(person)
}

func TestEncoder(t *testing.T) {
	person := Person{
		FirstName: "Muhammad",
		LastName:  "Azzam",
		Age:       25,
		Married:   false,
	}

	file, _ := os.Create("output.json")
	jsonEncoder := json.NewEncoder(file)
	err := jsonEncoder.Encode(person)
	if err != nil {
		return
	}
}
