package registry

import (
	i "github.com/thisismeamir/kage/internal/bootstrap/config"
	"github.com/thisismeamir/kage/internal/internal-pkg/registry"
	. "github.com/thisismeamir/kage/pkg/graph"
	"log"
)

func InitGraphRegistry() {
	registryPath := i.GetGlobalConfig().BasePath + "/data/graph.registry.json"
	graphRegistry, err := registry.LoadGraphRegistry(registryPath)
	if err != nil {
		log.Fatalf("[FATAL] Could not load graph registry: %s, Error: %s", registryPath, err)
	} else {
		for _, graph := range graphRegistry.Graphs {
			loadedGraph, err := LoadGraph(graph.Path)
			if err != nil {
				log.Printf("[FATAL] Could not load graph: %s, Error: %s", graph.Identifier, err)
			} else {
				if TestGraphValidity(loadedGraph) {

				} else {
					log.Printf("[ERROR] Graph: %s, is not valid.", graph.Identifier)
				}
			}
		}
	}
}

func TestGraphValidity(graph Graph) bool {
	mapRegistryPath := i.GetGlobalConfig().BasePath + "/data/map.registry.json"
	nodeRegistryPath := i.GetGlobalConfig().BasePath + "/data/node.registry.json"
	mapRegistry, err := registry.LoadMapRegistry(mapRegistryPath)
	nodeRegistry, err := registry.LoadNodeRegistry(nodeRegistryPath)
	if err != nil {
		log.Fatalf("[FATAL] Could not load map registry: %s, Error: %s", mapRegistryPath, err)
	}
	depsValidity := registry.CheckDependency(graph, mapRegistry, nodeRegistry)
	for _, validity := range depsValidity {
		if validity.Exists {
			log.Printf("[INFO] Graph: %s, depends on %s is valid", graph.Identifier, validity.Identifier)
		} else {
			log.Fatalf("[FATAL] Graph contains node %s, which is invalid.", validity.Identifier)
		}
	}
	return true
}
