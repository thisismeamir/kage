package file

import (
	"encoding/json"
	"fmt"
	. "github.com/thisismeamir/kage/pkg/graph"
	. "github.com/thisismeamir/kage/pkg/mapping"
	. "github.com/thisismeamir/kage/pkg/node"
	"os"
)

type TypeHint struct {
	Type string `json:"type"` // "node", "map", or "graph"
}

func LoadForm(path string) (interface{}, error) {
	var hint TypeHint
	data, _ := os.ReadFile(path)
	if err := json.Unmarshal(data, &hint); err != nil {
		return nil, fmt.Errorf("failed to decode type hint: %w", err)
	}

	switch hint.Type {
	case "node":
		var node Node
		if err := json.Unmarshal(data, &node); err != nil {
			return nil, fmt.Errorf("failed to unmarshal node: %w", err)
		}
		return node, nil

	case "map":
		var m Map
		if err := json.Unmarshal(data, &m); err != nil {
			return nil, fmt.Errorf("failed to unmarshal map: %w", err)
		}
		return m, nil

	case "graph":
		var g Graph
		if err := json.Unmarshal(data, &g); err != nil {
			return nil, fmt.Errorf("failed to unmarshal graph: %w", err)
		}
		return g, nil

	default:
		return nil, fmt.Errorf("unknown type: %s", hint.Type)
	}
}
