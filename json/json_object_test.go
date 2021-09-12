package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Person struct {
	FirstName string
	LastName string
	Age int
	Married bool
	Hobbies []string
}

func TestEncodeStruct(t *testing.T) {
	person := Person{
		FirstName: "Muhammad",
		LastName:  "Azzam",
		Age:       25,
		Married: false,
		Hobbies: []string{"Gaming", "Reading"},
	}

	bytes, _ := json.Marshal(person)
	fmt.Println(string(bytes))
}