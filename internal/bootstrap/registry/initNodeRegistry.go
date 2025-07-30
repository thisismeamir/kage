package registry

import (
	"encoding/json"
	"github.com/thisismeamir/kage/internal/bootstrap/config"
	"github.com/thisismeamir/kage/internal/internal-pkg/registry"
	. "github.com/thisismeamir/kage/internal/internal-pkg/registry"
	"github.com/thisismeamir/kage/pkg/node"
	"log"
	"os"
)

func InitNodeRegistry() {
	registryPath := config.GetGlobalConfig().BasePath + "/data/node.registry.json"
	nodeRegistry, err := LoadNodeRegistry(registryPath)
	if err != nil {
		log.Fatalf("[ERROR] initializing node registry: %v", err)
	} else {
		for _, register := range nodeRegistry.Nodes {
			if TestNodeValidity(register) {
				log.Printf("Node %s, test passed!", register.Identifier)
			} else {
				log.Printf("Node %s, couldn't be found. Removing From registry.")
				nodeRegistry.RemoveNodeFromRegistry(register.Identifier)

			}
		}
	}
}

// TestNodeValidity checks if the nodes in the registry are the same as the node definition iteself.
func TestNodeValidity(register registry.NodeRegister) bool {
	data, err := os.ReadFile(register.Path)
	if err != nil {
		log.Fatalf("failed to read registry file: %v", err)
		return false
	} else {
		var n node.Node
		if err := json.Unmarshal(data, &n); err != nil {
			log.Fatalf("failed to unmarshal node JSON: %v", err)
			return false
		}

		if n.Identifier != register.Identifier {
			log.Printf("Node identifier mismatch: expected %s, got %s", register.Identifier, n.Identifier)
			return false
		}

		if n.Model.ExecutionModel.InputSchema != nil && len(n.Model.ExecutionModel.InputSchema) != len(register.InputSchema) {
			log.Printf("Input schema mismatch for node %s", register.Identifier)
			return false
		}

		if n.Model.ExecutionModel.OutputSchema != nil && len(n.Model.ExecutionModel.OutputSchema) != len(register.OutputSchema) {
			log.Printf("Output schema mismatch for node %s", register.Identifier)
			return false
		}

		return true
	}
}
