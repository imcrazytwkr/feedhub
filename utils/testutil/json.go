package testutil

import (
	"encoding/json"
	"os"
)

func ReadJson(filename string, v any) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}
