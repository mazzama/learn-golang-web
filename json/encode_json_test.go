package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJSON(data interface{}) {
	result, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
}

func TestEncode(t *testing.T) {
	logJSON("Muhammad")
	logJSON([]string{"Muhammad", "Azzam", "Abduljabbar"})
	logJSON(1996)
}
