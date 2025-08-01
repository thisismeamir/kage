package registry

import (
	"encoding/json"
	"github.com/thisismeamir/kage/pkg/mapping"
	"github.com/thisismeamir/kage/pkg/node"
	"log"
	"os"
)

type Registry struct {
	GraphRegistry []GraphRegister `json:"graphs"`
	NodeRegistry  []NodeRegister  `json:"nodes"`
	MapRegistry   []MapRegister   `json:"maps"`
}

func LoadRegistry(path string) (*Registry, error) {
	var registry Registry
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Error reading registry file: %v", err)
		return nil, err
	} else {

		if err := json.Unmarshal(data, &registry); err != nil {
			log.Printf("Error parsing registry file: %v", err)
			return nil, err
		}
	}
	return &registry, nil
}

func (registry Registry) Save(path string) {
	data, err := json.MarshalIndent(registry, "", "  ")
	if err != nil {
		log.Fatalf("[FATAL] Could not save map registry JSON: %s", err)
	} else if err := os.WriteFile(path, data, 0644); err != nil {
		log.Fatalf("[FATAL] Could not save map registry JSON: %s", err)
	}
}

func (registry Registry) Contains(identifier string, path string) bool {
	for _, n := range registry.NodeRegistry {
		if n.Identifier == identifier && n.Path == path {
			return true
		}
	}
	for _, m := range registry.MapRegistry {
		if m.Identifier == identifier && m.Path == path {
			return true
		}
	}
	for _, g := range registry.GraphRegistry {
		if g.Identifier == identifier && g.Path == path {
			return true
		}
	}
	return false
}

func (registry *Registry) AddNode(n node.Node, identifier string, path string) {
	registry.NodeRegistry = append(registry.NodeRegistry, NodeRegister{
		Identifier:   identifier,
		Path:         path,
		InputSchema:  n.Model.ExecutionModel.InputSchema,
		OutputSchema: n.Model.ExecutionModel.OutputSchema,
	})
}

func (registry *Registry) AddMap(m mapping.Map, identifier string, path string) {
	registry.MapRegistry = append(registry.MapRegistry, MapRegister{
		Identifier:   identifier,
		Path:         path,
		InputSchema:  m.Model.InputSchema,
		OutputSchema: m.Model.OutputSchema,
	})
}

func (registry *Registry) AddGraph(identifier string, path string) {
	registry.GraphRegistry = append(registry.GraphRegistry, GraphRegister{
		Identifier: identifier,
		Path:       path,
	})
}
