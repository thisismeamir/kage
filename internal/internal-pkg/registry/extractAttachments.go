package registry

import (
	. "github.com/thisismeamir/kage/pkg/graph"
	"log"
)

func ExtractNodes(graph Graph, registry NodeRegistry, defaultSavePath string) {
	for _, node := range graph.Model.Attachments.GraphNodes {
		err := node.Save(defaultSavePath)
		if err != nil {
			log.Printf("Failed to save node %s. Error: %v", node.Name, err)
		} else {
			registry.AddNodeToRegistry(node)
		}

	}
}

func ExtractMaps(graph Graph, registry MapRegistry, defaultPath string) {
	for _, mapp := range graph.Model.Attachments.GraphMaps {
		err := mapp.Save(defaultPath)
		if err != nil {
			log.Printf("Failed to save map %s. Error: %v", mapp, err)
		} else {
			registry.AddMapToRegistry(mapp)
		}
	}
}

func ExtractAllAttachments(graph Graph, nodeRegistry NodeRegistry, mapRegistry MapRegistry, defaultSavePath string) {
	ExtractNodes(graph, nodeRegistry, defaultSavePath)
	ExtractMaps(graph, mapRegistry, defaultSavePath)
}
