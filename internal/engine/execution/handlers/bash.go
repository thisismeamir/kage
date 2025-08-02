package handlers

import (
	"fmt"
)

type BashHandler struct{}

func (b *BashHandler) Run(source string, input map[string]interface{}) (map[string]interface{}, error) {
	// command arguments should be --key=value recursively, meaning that {key1: {ke2: value2}} should be converted to --key1.key2=value2
	var arguments string
	for key, value := range input {
		if value == nil {
			continue
		}
		switch v := value.(type) {
		case string:
			arguments += "--" + key + "=" + v + " "
		case int, float64, bool:
			arguments += "--" + key + "=" + fmt.Sprintf("%v", v) + " "
		default:
			return nil, fmt.Errorf("unsupported type for key %s: %T", key, v)
		}
	}
	// Execute the bash command with the arguments
	println(arguments)
	return nil, nil // Placeholder return
}
