package event

import (
	"encoding/json"
	i "github.com/thisismeamir/kage/internal/bootstrap"
	. "github.com/thisismeamir/kage/pkg/graph"
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

func AddToGraphRegistry(graphPath string, graph Graph) {
	data, err := os.ReadFile(i.GetGlobalConfig().BasePath + "/data/graph.registry.json")
	if err != nil {
		log.Fatalf("Error finding graphs registry JSON: %s", err)
	} else {
		var registry GraphRegistry
		if err := json.Unmarshal(data, &registry); err != nil {
			log.Fatalf("Error unmarshalling graph registry JSON: %s", err)
		}
		var exists bool = false
		for _, m := range registry.Graphs {
			if m.Identifier == graph.Identifier {
				exists = true
				log.Printf("[Warning] Map with identifier %s already exists in the registry, skipping addition.", graph.Identifier)
			}
		}
		if exists != true {
			registry.Graphs = append(registry.Graphs, GraphRegister{
				Identifier: graph.Identifier,
				Path:       graph.Path,
			})
		}
	}
}
