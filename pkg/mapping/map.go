package mapping

import (
	"encoding/json"
	"github.com/thisismeamir/kage/pkg/form"
	"log"
	"os"
	"strings"
)

type Map struct {
	form.Form
	Model    map[string]interface{} `json:"model"`
	Metadata form.Metadata          `json:"metadata"`
}

func LoadMap(mapPath string) (*Map, error) {
	mapping := &Map{}
	data, err := os.ReadFile(mapPath)
	if err != nil {
		log.Fatalf("[ERROR] LoadMap failed to read file: %s", err)
		return nil, err
	} else {
		if err := json.Unmarshal(data, mapping); err != nil {
			log.Fatalf("[ERROR] LoadMap failed to unmarshal: %s", err)
			return nil, err
		}
		return mapping, nil
	}

}

func (mapp Map) Save(path string) error {
	data, err := json.MarshalIndent(mapp, "", "  ")
	if err != nil {
		log.Printf("[ERROR] Save failed to marshal JSON: %s", err)
		return err
	} else {
		if err := os.WriteFile(path, data, 0644); err != nil {
			log.Printf("[ERROR] Save failed to write file: %s", err)
			return err
		} else {
			log.Printf("[INFO] Map: %s saved successfully at %s", mapp.Name, path)
			return nil
		}
	}

}

func (mapp Map) MapList(inputs []map[string]interface{}) map[string]interface{} {
	output := make(map[string]interface{})

	// Iterate through each mapping rule in the model
	for newKey, sourcePath := range mapp.Model {
		sourcePathStr, ok := sourcePath.(string)
		if !ok {
			log.Printf("[WARNING] MapList: source path for key '%s' is not a string", newKey)
			continue
		}

		// Try to find the value in any of the input maps
		value := findValueInInputs(inputs, sourcePathStr)
		if value != nil {
			output[newKey] = value
		} else {
			log.Printf("[WARNING] MapList: could not find value for path '%s' in any input", sourcePathStr)
		}
	}

	return output
}

// Helper function to search for a value across all input maps
func findValueInInputs(inputs []map[string]interface{}, path string) interface{} {
	for _, input := range inputs {
		if value := getNestedValue(input, path); value != nil {
			return value
		}
	}
	return nil
}

// Helper function to extract nested values from a map using dot notation
func getNestedValue(data map[string]interface{}, path string) interface{} {
	keys := strings.Split(path, ".")
	current := data

	for i, key := range keys {
		value, exists := current[key]
		if !exists {
			return nil
		}

		// If this is the last key, return the value
		if i == len(keys)-1 {
			return value
		}

		// Otherwise, ensure the value is a map for the next iteration
		nextMap, ok := value.(map[string]interface{})
		if !ok {
			return nil
		}
		current = nextMap
	}

	return nil
}
