package graph_analysis

import "github.com/thisismeamir/kage/pkg/graph"

func FindAllCycles(gr []graph.GraphObject) []map[int]bool {
	allCycles := make([]map[int]bool, 0)

	for _, node := range gr {
		visited := make(map[int]bool)
		restack := make(map[int]bool)
		if DFSWithCycleDetection(gr, node.Id, visited, restack) {
			allCycles = append(allCycles, restack)
		}
	}
	return allCycles
}
