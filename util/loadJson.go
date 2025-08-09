package util

import (
	"encoding/json"
	"os"
)

func LoadJson(filename string) map[string]interface{} {
	var target map[string]interface{}
	data, err := os.ReadFile(filename)
	if err != nil {
		return target // Return an empty map if the file cannot be read
	}

	if err := json.Unmarshal(data, &target); err != nil {
		return target
	}

	return target
}
