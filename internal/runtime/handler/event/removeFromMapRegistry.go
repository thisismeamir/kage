package event

import (
	"encoding/json"
	i "github.com/thisismeamir/kage/internal/bootstrap"
	"log"
	"os"
)

func RemoveFromMapRegistry(mapPath string, identifier string) {
	// Read the existing map registry
	data, err := os.ReadFile(i.GetGlobalConfig().BasePath + "/data/map.registry.json")
	if err != nil {
		log.Printf("error reading map registry JSON: %w", err)
	}

	var registry MapRegistry
	if err := json.Unmarshal(data, &registry); err != nil {
		log.Printf("error unmarshalling map registry JSON: %w", err)
	}

	// Remove the map with the specified identifier
	for i, m := range registry.Maps {
		if m.Identifier == identifier {
			registry.Maps = append(registry.Maps[:i], registry.Maps[i+1:]...)
			break
		}
	}

	// Write the updated registry back to the file
	newData, err := json.MarshalIndent(registry, "", "  ")
	if err != nil {
		log.Printf("error marshalling updated map registry JSON: %w", err)
	}

	if err := os.WriteFile(i.GetGlobalConfig().BasePath+"/data/map.registry.json", newData, 0644); err != nil {
		log.Printf("error writing updated map registry JSON: %w", err)
	}

}
