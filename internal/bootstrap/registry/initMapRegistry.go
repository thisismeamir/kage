package registry

import (
	"encoding/json"
	"github.com/thisismeamir/kage/internal/bootstrap/config"
	"github.com/thisismeamir/kage/internal/internal-pkg/registry"
	"github.com/thisismeamir/kage/pkg/mapping"
	"log"
	"os"
)

func InitMapRegistry() {
	registryPath := config.GetGlobalConfig().BasePath + "/data/map.registry.json"
	mapRegistry, err := registry.LoadMapRegistry(registryPath)
	if err != nil {
		log.Fatalf("[ERROR] initializing map registry: %v", err)
	} else {
		for _, register := range mapRegistry.Maps {
			if TestMapRegistry(register) {
				log.Printf("Node %s, test passed!", register.Identifier)
			} else {
				log.Printf("Node %s, couldn't be found. Removing From registry.")
				mapRegistry.RemoveMapFromRegistry(register.Identifier)
			}
		}
	}
}

// TestNodeValidity checks if the nodes in the registry are the same as the node definition iteself.
func TestMapRegistry(register registry.MapRegister) bool {
	data, err := os.ReadFile(register.Path)
	if err != nil {
		log.Fatalf("failed to read registry file: %v", err)
		return false
	} else {
		var m mapping.Map
		if err := json.Unmarshal(data, &m); err != nil {
			log.Fatalf("failed to unmarshal node JSON: %v", err)
			return false
		}

		if m.Identifier != register.Identifier {
			log.Printf("Node identifier mismatch: expected %s, got %s", register.Identifier, m.Identifier)
			return false
		}

		if m.Model.InputSchema != nil && len(m.Model.InputSchema) != len(register.InputSchema) {
			log.Printf("Input schema mismatch for node %s", register.Identifier)
			return false
		}

		if m.Model.OutputSchema != nil && len(m.Model.OutputSchema) != len(register.OutputSchema) {
			log.Printf("Output schema mismatch for node %s", register.Identifier)
			return false
		}

		return true
	}
}
