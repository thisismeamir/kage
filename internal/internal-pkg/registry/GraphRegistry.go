package registry

import (
	"encoding/json"
	"github.com/thisismeamir/kage/pkg/graph"
	"log"
	"os"
)

type GraphRegistry struct {
	Graphs []GraphRegister `json:"graphs"`
}

type GraphRegister struct {
	Identifier string `json:"identifier"`
	Path       string `json:"path"`
}

func LoadGraphRegistry(registryPath string) (GraphRegistry, error) {
	data, err := os.ReadFile(registryPath)
	if err != nil {
		log.Fatalf("[FATAL] Could not load graph registry JSON: %s", err)
		return GraphRegistry{}, err
	} else {
		var registry GraphRegistry
		if err := json.Unmarshal(data, &registry); err != nil {
			log.Fatalf("[FATAL] Error unmarshalling graph registry JSON: %s", err)
			return GraphRegistry{}, err
		}
		return registry, nil
	}
}

func (registry GraphRegistry) SaveGraphRegistry(registryPath string) GraphRegistry {
	data, err := json.Marshal(registry)
	if err != nil {
		log.Fatalf("[FATAL] Error marshalling graph registry JSON: %s", err)
	} else if err := os.WriteFile(registryPath, data, 0644); err != nil {
		log.Fatalf("[FATAL] Could not save graph registry JSON: %s", err)
	}
	return registry
}

func (registry GraphRegistry) AddGraphToRegistry(graph graph.Graph) GraphRegistry {
	registry.Graphs = append(registry.Graphs, GraphRegister{
		Identifier: graph.Identifier,
		Path:       graph.Path,
	})

	return registry
}

func (registry GraphRegistry) RemoveGraphFromRegistry(identifier string) GraphRegistry {
	newRegistry := GraphRegistry{
		Graphs: []GraphRegister{},
	}
	for _, graph := range registry.Graphs {
		if graph.Identifier != identifier {
			newRegistry.Graphs = append(newRegistry.Graphs, graph)
		}
	}
	return newRegistry
}
