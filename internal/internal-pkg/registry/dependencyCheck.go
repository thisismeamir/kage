package registry

import (
	"github.com/thisismeamir/kage/pkg/graph"
	"log"
)

type DependencyCheck struct {
	Identifier string `json:"identifier"`
	Exists     bool   `json:"exists"`
}

func CheckDependency(graph graph.Graph, mapRegistry MapRegistry, nodeRegistry NodeRegistry) []DependencyCheck {
	dependencyChecks := make([]DependencyCheck, 0)

	for _, dependency := range graph.Model.Dependencies {
		if dependency.Type == "map" {
			dependencyChecks = append(dependencyChecks, DependencyCheck{
				Identifier: dependency.Identifier,
				Exists:     mapRegistry.MapExists(dependency.Identifier),
			})
		} else if dependency.Type == "node" {
			dependencyChecks = append(dependencyChecks, DependencyCheck{
				Identifier: dependency.Identifier,
				Exists:     nodeRegistry.NodeExists(dependency.Identifier),
			})
		} else {
			log.Printf("Unsupported dependency type: %s", dependency.Type)
			dependencyChecks = append(dependencyChecks, DependencyCheck{
				Identifier: dependency.Identifier,
				Exists:     false,
			})

		}
	}
	return dependencyChecks
}
