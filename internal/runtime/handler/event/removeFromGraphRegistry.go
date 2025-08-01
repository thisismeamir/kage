package event

import (
	"encoding/json"
	i "github.com/thisismeamir/kage/internal/bootstrap"
	"log"
	"os"
)

func RemoveFromGraphRegistry(identifier string) {
	// Read the existing graph registry
	data, err := os.ReadFile(i.GetGlobalConfig().BasePath + "/data/graph.registry.json")
	if err != nil {
		log.Printf("error reading graph registry JSON: %w", err)
		return
	}

	var registry GraphRegistry
	if err := json.Unmarshal(data, &registry); err != nil {
		log.Printf("error unmarshalling graph registry JSON: %w", err)
		return
	}

	// Remove the graph with the specified identifier
	for i, g := range registry.Graphs {
		if g.Identifier == identifier {
			registry.Graphs = append(registry.Graphs[:i], registry.Graphs[i+1:]...)
			break
		}
	}

	// Write the updated registry back to the file
	newData, err := json.MarshalIndent(registry, "", "  ")
	if err != nil {
		log.Printf("error marshalling updated graph registry JSON: %w", err)
		return
	}

	if err := os.WriteFile(i.GetGlobalConfig().BasePath+"/data/graph.registry.json", newData, 0644); err != nil {
		log.Printf("error writing updated graph registry JSON: %w", err)
	}
}
