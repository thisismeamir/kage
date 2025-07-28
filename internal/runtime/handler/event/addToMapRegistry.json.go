package event

import (
	"encoding/json"
	i "github.com/thisismeamir/kage/internal/bootstrap"
	. "github.com/thisismeamir/kage/pkg/mapping"
	"log"
	"os"
)

type MapRegistry struct {
	Maps []MapRegister `json:"maps"`
}

type MapRegister struct {
	Identifier   string                 `json:"identifier"`
	Path         string                 `json:"path"`
	InputSchema  map[string]interface{} `json:"input_schema"`
	OutputSchema map[string]interface{} `json:"output_schema"`
}

func AddToMapRegister(mapping Map) {
	data, err := os.ReadFile(i.GetGlobalConfig().BasePath + "/data/map.registry.json")
	if err != nil {
		log.Fatalf("Error finding map registry JSON: %s", err)
	} else {
		var registry MapRegistry
		if err := json.Unmarshal(data, &registry); err != nil {
			log.Fatalf("Error unmarshalling map registry JSON: %s", err)
		}
		var exists bool = false
		for _, m := range registry.Maps {
			if m.Identifier == mapping.Identifier {
				exists = true
				log.Printf("[Warning] Map with identifier %s already exists in the registry, skipping addition.", mapping.Identifier)
			}
		}
		if exists != true {
			registry.Maps = append(registry.Maps, MapRegister{
				Identifier:   mapping.Identifier,
				Path:         mapping.Path,
				InputSchema:  mapping.Model.InputSchema,
				OutputSchema: mapping.Model.OutputSchema,
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
