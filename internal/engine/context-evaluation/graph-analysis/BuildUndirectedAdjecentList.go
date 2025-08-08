package graph_analysis

import "github.com/thisismeamir/kage/pkg/graph"

func BuildUndirectedAdjList(gr []graph.GraphObject) map[int][]int {
	adj := make(map[int][]int)

	for _, node := range gr {
		for _, neighbor := range node.Outgoing {
			adj[node.Id] = append(adj[node.Id], neighbor)
			adj[neighbor] = append(adj[neighbor], node.Id) // reverse edge added
		}
	}
	return adj
}
