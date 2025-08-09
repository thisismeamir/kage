package graph_analysis

import "github.com/thisismeamir/kage/pkg/graph"

func GetNeighbors(gr []graph.GraphObject, nodeId int) []int {
	// Find the node with the given ID
	for _, obj := range gr {
		if obj.Id == nodeId {
			return obj.Outgoing // Return the outgoing edges of the node
		}
	}
	return nil // Return nil if the node is not found
}
