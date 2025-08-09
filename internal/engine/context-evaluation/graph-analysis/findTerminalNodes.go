package graph_analysis

import "github.com/thisismeamir/kage/pkg/graph"

func FindTerminalNodes(gr []graph.GraphObject) []int {
	// Terminal nodes are those that have no outgoing edges
	terminalNodes := make([]int, 0)
	for _, obj := range gr {
		if len(obj.Outgoing) == 0 {
			// This node has no outgoing edges, so it's a terminal node
			terminalNodes = append(terminalNodes, obj.Id)
		}
	}
	return terminalNodes
}
