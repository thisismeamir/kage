package graph_analysis

import "github.com/thisismeamir/kage/pkg/graph"

func DFSWithCycleDetection(gr []graph.GraphObject, nodeID int, visited map[int]bool, recStack map[int]bool) bool {
	// If the node is already in the recursion stack, a cycle is detected
	if recStack[nodeID] {
		return true
	}

	// If the node is already visited, skip it (don't reprocess it)
	if visited[nodeID] {
		return false
	}

	// Mark the current node as visited and add it to the recursion stack
	visited[nodeID] = true
	recStack[nodeID] = true

	// Recursively visit all outgoing neighbors
	for _, neighborID := range GetNeighbors(gr, nodeID) {
		if DFSWithCycleDetection(gr, neighborID, visited, recStack) {
			return true
		}
	}

	// After processing all neighbors, remove the node from the recursion stack
	recStack[nodeID] = false
	return false
}
