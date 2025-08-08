package graph_analysis

import "github.com/thisismeamir/kage/pkg/graph"

func FindIngoingVertices(gr []graph.GraphObject, nodeId int) []int {
	// Find all nodes that have an outgoing edge to the given nodeId
	ingressNodes := make([]int, 0)
	for _, obj := range gr {
		for _, outgoing := range obj.Outgoing {
			if outgoing == nodeId {
				ingressNodes = append(ingressNodes, obj.Id)
			}
		}
	}
	return ingressNodes
}
