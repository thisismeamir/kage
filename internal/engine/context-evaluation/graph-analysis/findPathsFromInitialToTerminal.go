package graph_analysis

import (
	"github.com/thisismeamir/kage/pkg/graph"
)

func FindPathsFromInitialToTerminal(gr []graph.GraphObject) [][]int {
	initialNodes := FindInitialNodes(gr)
	terminalList := FindTerminalNodes(gr)

	// Fast lookup for terminals
	terminalMap := make(map[int]bool)
	for _, id := range terminalList {
		terminalMap[id] = true
	}

	adj := BuildAdjList(gr)

	results := [][]int{}
	for _, start := range initialNodes {
		FindAllPaths(adj, start, terminalMap, []int{}, make(map[int]bool), &results)
	}

	return results
}
