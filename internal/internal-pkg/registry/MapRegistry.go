package registry

import (
	"encoding/json"
	i "github.com/thisismeamir/kage/internal/bootstrap/config"
	"github.com/thisismeamir/kage/pkg/mapping"
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

func LoadMapRegistry(registryPath string) (MapRegistry, error) {
	data, err := os.ReadFile(registryPath)
	if err != nil {
		log.Fatalf("[FATAL] Could not load map registry JSON: %s", err)
		return MapRegistry{}, err
	} else {
		var registry MapRegistry
		if err := json.Unmarshal(data, &registry); err != nil {
			log.Fatalf("[FATAL] Error unmarshalling map registry JSON: %s", err)
			return MapRegistry{}, err
		}
		return registry, nil
	}
}

func (registry MapRegistry) SaveMapRegistry() MapRegistry {
	registryPath := i.GetGlobalConfig().BasePath + "/data/map.registry.json"
	data, err := json.MarshalIndent(registry, "", "  ")
	if err != nil {
		log.Fatalf("[FATAL] Could not save map registry JSON: %s", err)
	} else if err := os.WriteFile(registryPath, data, 0644); err != nil {
		log.Fatalf("[FATAL] Could not save map registry JSON: %s", err)
	}
	return registry
}

func (registry MapRegistry) MapExists(identifier string) bool {
	for _, m := range registry.Maps {
		if m.Identifier == identifier {
			return true
		}
	}
	return false
}

func (registry MapRegistry) AddMapToRegistry(mapper mapping.Map) MapRegistry {
	if !registry.MapExists(mapper.Identifier) {
		registry.Maps = append(registry.Maps, MapRegister{
			Identifier:   mapper.Identifier,
			Path:         mapper.Path,
			InputSchema:  mapper.Model.InputSchema,
			OutputSchema: mapper.Model.OutputSchema,
		})
	}
	return registry
}

func (registry MapRegistry) RemoveMapFromRegistry(identifier string) MapRegistry {
	newRegistry := MapRegistry{
		Maps: make([]MapRegister, 0),
	}
	for _, r := range registry.Maps {
		if r.Identifier != identifier {
			newRegistry.Maps = append(newRegistry.Maps, r)
		}
	}
	return newRegistry
}
