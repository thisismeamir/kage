package graph_analysis

import "github.com/thisismeamir/kage/pkg/graph"

// DepthFirstSearch performs a depth-first search on the task graph starting from the given task ID.
func DepthFirstSearch(gr []graph.GraphObject, startingNodeId int, visited map[int]bool) {
	visited[startingNodeId] = true
	for _, neighbor := range GetNeighbors(gr, startingNodeId) {
		if !visited[neighbor] {
			visited[neighbor] = true
			// Recursively visit the neighbors
			DepthFirstSearch(gr, neighbor, visited)
		}
	}
}
