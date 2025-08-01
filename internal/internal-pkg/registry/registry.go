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

func (registry *Registry) GetAllPaths() []string {
	allPaths := make([]string, 0)
	for _, g := range registry.GraphRegistry {
		allPaths = append(allPaths, g.Path)
	}
	for _, m := range registry.MapRegistry {
		allPaths = append(allPaths, m.Path)
	}
	for _, n := range registry.NodeRegistry {
		allPaths = append(allPaths, n.Identifier)
	}
	return allPaths
}

func (registry *Registry) GetAllIdentifiers() []string {
	allIdentifiers := make([]string, 0)
	for _, g := range registry.GraphRegistry {
		allIdentifiers = append(allIdentifiers, g.Identifier)
	}
	for _, m := range registry.MapRegistry {
		allIdentifiers = append(allIdentifiers, m.Identifier)
	}
	for _, n := range registry.NodeRegistry {
		allIdentifiers = append(allIdentifiers, n.Identifier)
	}
	return allIdentifiers
}

// RemoveNode removes a node by identifier and path.
func (registry *Registry) RemoveNode(identifier, path string) {
	filtered := make([]NodeRegister, 0, len(registry.NodeRegistry))
	for _, n := range registry.NodeRegistry {
		if !(n.Identifier == identifier && n.Path == path) {
			filtered = append(filtered, n)
		}
	}
	registry.NodeRegistry = filtered
}

// RemoveMap removes a map by identifier and path.
func (registry *Registry) RemoveMap(identifier, path string) {
	filtered := make([]MapRegister, 0, len(registry.MapRegistry))
	for _, m := range registry.MapRegistry {
		if !(m.Identifier == identifier && m.Path == path) {
			filtered = append(filtered, m)
		}
	}
	registry.MapRegistry = filtered
}

// RemoveGraph removes a graph by identifier and path.
func (registry *Registry) RemoveGraph(identifier, path string) {
	filtered := make([]GraphRegister, 0, len(registry.GraphRegistry))
	for _, g := range registry.GraphRegistry {
		if !(g.Identifier == identifier && g.Path == path) {
			filtered = append(filtered, g)
		}
	}
	registry.GraphRegistry = filtered
}

// CleanMissingFiles removes registry entries whose referenced files do not exist.
func (registry *Registry) CleanMissingFiles() {
	// Clean nodes
	for _, n := range registry.NodeRegistry {
		if _, err := os.Stat(n.Path); os.IsNotExist(err) {
			log.Printf("[INFO] Removing missing file %s", n.Path)
			registry.RemoveNode(n.Identifier, n.Path)
		}
	}
	// Clean maps
	for _, m := range registry.MapRegistry {
		if _, err := os.Stat(m.Path); os.IsNotExist(err) {
			log.Printf("[INFO] Removing missing file %s", m.Path)
			registry.RemoveMap(m.Identifier, m.Path)
		}
	}
	// Clean graphs
	for _, g := range registry.GraphRegistry {
		if _, err := os.Stat(g.Path); os.IsNotExist(err) {
			log.Printf("[INFO] Removing missing file %s", g.Path)
			registry.RemoveGraph(g.Identifier, g.Path)
		}
	}
}
