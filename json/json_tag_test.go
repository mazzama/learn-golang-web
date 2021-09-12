package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Employee struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Occupation string `json:"occupation"`
	TotalExperience int `json:"total_experience"`
}

func TestEncodeJsonWithTag(t *testing.T) {
	employee := Employee{
		FirstName:       "Muhammad",
		LastName:        "Azzam",
		Occupation:      "Programmer",
		TotalExperience: 4,
	}

	bytes, _ := json.Marshal(employee)
	fmt.Println(string(bytes))
}

func TestDecodeJsonWithTag(t *testing.T) {
	jsonString := "{\"first_name\":\"Muhammad\",\"last_name\":\"Azzam\",\"occupation\":\"Programmer\",\"total_experience\":4}"
	jsonBytes := []byte(jsonString)
	employee := &Employee{}

	json.Unmarshal(jsonBytes, employee)
	fmt.Println(employee)
}