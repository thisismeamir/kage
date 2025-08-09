package graph_analysis

import "github.com/thisismeamir/kage/pkg/graph"

func BuildAdjList(gr []graph.GraphObject) map[int][]int {
	adj := make(map[int][]int)
	for _, node := range gr {
		adj[node.Id] = node.Outgoing
	}
	return adj
}
