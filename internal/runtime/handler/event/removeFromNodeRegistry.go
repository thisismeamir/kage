package event

import (
	"encoding/json"
	i "github.com/thisismeamir/kage/internal/bootstrap"
	"log"
	"os"
)

func RemoveFromNodeRegistry(identifier string) {
	// Read the existing node registry
	data, err := os.ReadFile(i.GetGlobalConfig().BasePath + "/data/node.registry.json")
	if err != nil {
		log.Printf("error reading node registry JSON: %w", err)
		return
	}

	var registry NodeRegistry
	if err := json.Unmarshal(data, &registry); err != nil {
		log.Printf("error unmarshalling node registry JSON: %w", err)
		return
	}

	// Remove the node with the specified identifier
	for i, n := range registry.Nodes {
		if n.Identifier == identifier {
			registry.Nodes = append(registry.Nodes[:i], registry.Nodes[i+1:]...)
			break
		}
	}

	// Write the updated registry back to the file
	newData, err := json.MarshalIndent(registry, "", "  ")
	if err != nil {
		log.Printf("error marshalling updated node registry JSON: %w", err)
		return
	}

	if err := os.WriteFile(i.GetGlobalConfig().BasePath+"/data/node.registry.json", newData, 0644); err != nil {
		log.Printf("error writing updated node registry JSON: %w", err)
	}
}
