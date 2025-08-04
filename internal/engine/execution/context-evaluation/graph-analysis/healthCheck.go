package graph_analysis

import (
	"errors"
	"fmt"
	"github.com/thisismeamir/kage/internal/internal-pkg/registry"
	"github.com/thisismeamir/kage/pkg/graph"
)

func HealthCheck(graph *graph.Graph, reg registry.Registry) (bool, error) {
	if graph == nil {
		return false, nil
	}
	// Check if every nodes and maps exist in registry
	for _, identifier := range graph.Model.Structure {
		if identifier.Type != "node" && identifier.Type != "map" {
			return false, errors.New(fmt.Sprintf("unknown node type for graph %s, at execution identifier %s", graph.Name, identifier))
		}
		if path := reg.GetPath(identifier.ExecutionIdentifier); len(path) == 0 {
			return false, errors.New(fmt.Sprintf("identifier %s not found in registry for graph %s", identifier, graph.Name))
		}
	}

	return true, nil
}
