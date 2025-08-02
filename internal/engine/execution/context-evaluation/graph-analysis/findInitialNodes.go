package graph_analysis

import "github.com/thisismeamir/kage/pkg/graph"

// FindInitialNodes finds all initial nodes in the graph that have no incoming edges.
func FindInitialNodes(gr []graph.GraphObject) []int {
	// First we get all the outgoing edges from each node
	outgoingEdges := make(map[int]bool)
	initialNodes := make([]int, 0)
	for _, obj := range gr {
		for _, outgoing := range obj.Outgoing {
			outgoingEdges[outgoing] = true
		}
	}
	// Now we find nodes that are not in the outgoing edges map
	for _, obj := range gr {
		if _, exists := outgoingEdges[obj.Id]; !exists {
			// This node has no incoming edges, so it's an initial node
			initialNodes = append(initialNodes, obj.Id)
		}
	}
	return initialNodes
}
