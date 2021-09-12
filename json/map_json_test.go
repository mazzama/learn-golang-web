package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapJSON(t *testing.T) {
	jsonRequest := `{"id": 1, "name": "Azzam"}`
	jsonBytes := []byte(jsonRequest)

	var result map[string]interface{}
	_ = json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
}
