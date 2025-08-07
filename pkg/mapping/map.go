package mapping

import (
	"encoding/json"
	"fmt"
	"github.com/thisismeamir/kage/pkg/form"
	"log"
	"os"
	"strconv"
	"strings"
)

type Map struct {
	form.Form
	Model    MapModel      `json:"model"`
	Metadata form.Metadata `json:"metadata"`
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

// MapInputsToOutput maps multiple input JSONs to a single output JSON based on the map configuration
func MapInputsToOutput(mapConfigJSON string, inputJSONs []string) (string, error) {
	// Parse the map configuration
	var config Map
	if err := json.Unmarshal([]byte(mapConfigJSON), &config); err != nil {
		return "", fmt.Errorf("failed to parse map config: %v", err)
	}

	// Parse all input JSONs
	var inputs []map[string]interface{}
	for i, inputJSON := range inputJSONs {
		var input map[string]interface{}
		if err := json.Unmarshal([]byte(inputJSON), &input); err != nil {
			return "", fmt.Errorf("failed to parse input JSON %d: %v", i, err)
		}
		inputs = append(inputs, input)
	}

	// Validate input JSONs against their schemas
	if err := ValidateInputs(inputs, config.Model.InputSchemas); err != nil {
		return "", fmt.Errorf("input validation failed: %v", err)
	}

	// Create the output map
	output := make(map[string]interface{})

	// Get the output schema properties
	outputProperties, ok := config.Model.OutputSchema["properties"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("invalid output schema format")
	}

	// Map each output property
	for outputKey, mappingRule := range outputProperties {
		mappingStr, ok := mappingRule.(string)
		if !ok {
			return "", fmt.Errorf("invalid mapping rule for key %s", outputKey)
		}

		// Extract the value based on the mapping rule
		value, err := ExtractValue(inputs, mappingStr)
		if err != nil {
			return "", fmt.Errorf("failed to extract value for %s: %v", outputKey, err)
		}

		output[outputKey] = value
	}

	// Validate output against output schema
	if err := validateOutput(output, config.Model.OutputSchema); err != nil {
		return "", fmt.Errorf("output validation failed: %v", err)
	}

	// Convert output to JSON
	outputJSON, err := json.Marshal(output)
	if err != nil {
		return "", fmt.Errorf("failed to marshal output: %v", err)
	}

	return string(outputJSON), nil
}

// ExtractValue extracts a value from the inputs based on the mapping rule
func ExtractValue(inputs []map[string]interface{}, mappingRule string) (interface{}, error) {
	// Handle nested property access (e.g., "user.name", "user.age")
	parts := strings.Split(mappingRule, ".")

	// Search through all inputs for the first part of the path
	for _, input := range inputs {
		if value, exists := input[parts[0]]; exists {
			// If it's a simple property (no dots), return it directly
			if len(parts) == 1 {
				return value, nil
			}

			// Navigate through nested properties
			current := value
			for i := 1; i < len(parts); i++ {
				switch v := current.(type) {
				case map[string]interface{}:
					if nextValue, exists := v[parts[i]]; exists {
						current = nextValue
					} else {
						return nil, fmt.Errorf("property %s not found in nested object", parts[i])
					}
				default:
					return nil, fmt.Errorf("cannot access property %s on non-object type", parts[i])
				}
			}
			return current, nil
		}
	}

	return nil, fmt.Errorf("property %s not found in any input", parts[0])
}

// ValidateInputs validates all input JSONs against their corresponding schemas
func ValidateInputs(inputs []map[string]interface{}, schemas []map[string]interface{}) error {
	for i, input := range inputs {
		// Find matching schema for this input
		var matchingSchema map[string]interface{}
		var found bool

		for _, schema := range schemas {
			if SchemaMatches(input, schema) {
				matchingSchema = schema
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("no matching schema found for input %d", i)
		}

		if err := ValidateAgainstSchema(input, matchingSchema); err != nil {
			return fmt.Errorf("input %d validation failed: %v", i, err)
		}
	}
	return nil
}

// validateOutput validates the output against the output schema
func validateOutput(output map[string]interface{}, schema map[string]interface{}) error {
	return ValidateAgainstSchema(output, schema)
}

// SchemaMatches checks if an input matches a schema structure
func SchemaMatches(input map[string]interface{}, schema map[string]interface{}) bool {
	properties, ok := schema["properties"].(map[string]interface{})
	if !ok {
		return false
	}

	// Check if input has at least one property that matches the schema
	for key := range properties {
		if _, exists := input[key]; exists {
			return true
		}
	}
	return false
}

// ValidateAgainstSchema validates a JSON object against a schema
func ValidateAgainstSchema(data map[string]interface{}, schema map[string]interface{}) error {
	properties, ok := schema["properties"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid schema: missing properties")
	}

	for key, expectedType := range properties {
		value, exists := data[key]
		if !exists {
			continue // Optional fields
		}

		switch expectedTypeValue := expectedType.(type) {
		case string:
			if err := ValidateType(value, expectedTypeValue); err != nil {
				return fmt.Errorf("field %s: %v", key, err)
			}
		case map[string]interface{}:
			// Handle nested objects
			nestedData, ok := value.(map[string]interface{})
			if !ok {
				return fmt.Errorf("field %s: expected object, got %T", key, value)
			}

			// Create a schema-like structure for nested validation
			nestedSchema := map[string]interface{}{
				"properties": expectedTypeValue,
			}

			if err := ValidateAgainstSchema(nestedData, nestedSchema); err != nil {
				return fmt.Errorf("field %s: %v", key, err)
			}
		}
	}

	return nil
}

// ValidateType validates a value against an expected type string
func ValidateType(value interface{}, expectedType string) error {
	switch expectedType {
	case "string":
		if _, ok := value.(string); !ok {
			return fmt.Errorf("expected string, got %T", value)
		}
	case "int", "integer":
		switch value.(type) {
		case int, int32, int64, float64:
			// JSON numbers are typically float64, but we accept them as int if they're whole numbers
			if v, ok := value.(float64); ok && v != float64(int(v)) {
				return fmt.Errorf("expected integer, got float %v", v)
			}
		default:
			return fmt.Errorf("expected integer, got %T", value)
		}
	case "float", "number":
		switch value.(type) {
		case int, int32, int64, float32, float64:
			// Accept any numeric type
		default:
			return fmt.Errorf("expected number, got %T", value)
		}
	case "bool", "boolean":
		if _, ok := value.(bool); !ok {
			return fmt.Errorf("expected boolean, got %T", value)
		}
	case "object":
		if _, ok := value.(map[string]interface{}); !ok {
			return fmt.Errorf("expected object, got %T", value)
		}
	case "array":
		if _, ok := value.([]interface{}); !ok {
			return fmt.Errorf("expected array, got %T", value)
		}
	default:
		// If we don't recognize the type, we'll allow it (could be a custom type)
		return nil
	}

	return nil
}

// InterfaceToString Helper function to convert interface{} to string for display
func InterfaceToString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case int:
		return strconv.Itoa(val)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(val)
	default:
		return fmt.Sprintf("%v", val)
	}
}
