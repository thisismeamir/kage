package event

import (
	"encoding/json"
	i "github.com/thisismeamir/kage/internal/bootstrap"
	. "github.com/thisismeamir/kage/pkg/node"
	"log"
	"os"
)

type NodeRegistry struct {
	Nodes []NodeRegister `json:"nodes"`
}

type NodeRegister struct {
	Identifier   string                 `json:"identifier"`
	Path         string                 `json:"path"`
	InputSchema  map[string]interface{} `json:"input_schema"`
	OutputSchema map[string]interface{} `json:"output_schema,omitempty"`
}

func AddToNodeRegistry(node Node) {
	data, err := os.ReadFile(i.GetGlobalConfig().BasePath + "/data/node.registry.json")
	if err != nil {
		log.Fatalf("Error finding node registry JSON: %s", err)
	} else {
		var registry NodeRegistry
		if err := json.Unmarshal(data, &registry); err != nil {
			log.Fatalf("Error unmarshalling node registry JSON: %s", err)
		}
		var exists bool = false
		for _, n := range registry.Nodes {
			if n.Identifier == node.Identifier {
				exists = true
			}
		}
		if exists != true {
			registry.Nodes = append(registry.Nodes, NodeRegister{
				Identifier:   node.Identifier,
				Path:         node.Path,
				InputSchema:  node.Model.ExecutionModel.InputSchema,
				OutputSchema: node.Model.ExecutionModel.OutputSchema,
			})
			newData, err := json.MarshalIndent(registry, "", "  ")
			if err != nil {
				log.Printf("error marshalling updated map registry JSON: %w", err)
			}

			if err := os.WriteFile(i.GetGlobalConfig().BasePath+"/data/map.registry.json", newData, 0644); err != nil {
				log.Printf("error writing updated map registry JSON: %w", err)
			}
		}
	}
}
