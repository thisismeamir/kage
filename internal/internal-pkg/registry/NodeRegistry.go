package registry

import (
	"encoding/json"
	"github.com/thisismeamir/kage/pkg/node"
	"log"
	"os"
)

type NodeRegistry struct {
	Nodes []NodeRegister `json:"nodes"`
}

type NodeRegister struct {
	Identifier   string                 `json:"identifier"`
	Path         string                 `json:"path"`
	InputSchema  map[string]interface{} `json:"input_schema"`
	OutputSchema map[string]interface{} `json:"output_schema"`
}

func LoadNodeRegistry(registryPath string) (NodeRegistry, error) {
	data, err := os.ReadFile(registryPath)
	if err != nil {
		log.Fatalf("[FATAL] Could not load node registry JSON: %s", err)
		return NodeRegistry{}, err
	} else {
		var nodeRegistry NodeRegistry
		if err := json.Unmarshal(data, &nodeRegistry); err != nil {
			log.Fatalf("[FATAL] Error unmarshalling node registry JSON: %s", err)
			return NodeRegistry{}, err
		}
		return nodeRegistry, nil
	}
}

func (registry NodeRegistry) SaveNodeRegistry(registryPath string) NodeRegistry {
	data, err := json.Marshal(registry)
	if err != nil {
		log.Fatalf("[FATAL] Error marshalling node registry JSON: %s", err)
	} else if err := os.WriteFile(registryPath, data, 0755); err != nil {
		log.Fatalf("[FATAL] Error saving node registry JSON: %s", err)
	}
	return registry
}

func (registry NodeRegistry) NodeExists(identifier string) bool {
	for _, node := range registry.Nodes {
		if node.Identifier == identifier {
			return true
		}
	}
	return false
}

func (registry NodeRegistry) AddNodeToRegistry(node node.Node) NodeRegistry {
	if !registry.NodeExists(node.Identifier) {
		registry.Nodes = append(registry.Nodes, NodeRegister{
			Identifier:   node.Identifier,
			Path:         node.Path,
			InputSchema:  node.Model.ExecutionModel.InputSchema,
			OutputSchema: node.Model.ExecutionModel.OutputSchema,
		})
	}
	return registry
}

func (registry NodeRegistry) RemoveNodeFromRegistry(identifier string) NodeRegistry {
	newRegistry := NodeRegistry{
		Nodes: make([]NodeRegister, 0),
	}
	for _, r := range registry.Nodes {
		if r.Identifier != identifier {
			newRegistry.Nodes = append(newRegistry.Nodes, r)
		} else {
			node.LoadNode(r.Path)
		}
	}
	return newRegistry
}
