package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	jsonRequest := "{\"FirstName\":\"Muhammad\",\"LastName\":\"Azzam\",\"Age\":25,\"Married\":false,\"Hobbies\":[\"Gaming\",\"Reading\"]}"
	jsonBytes := []byte(jsonRequest)

	person := &Person{}
	json.Unmarshal(jsonBytes, person)
	fmt.Println(person)
}