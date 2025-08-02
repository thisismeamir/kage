package graph_analysis

import (
	"github.com/thisismeamir/kage/pkg/graph"
	"sort"
)

func FindIsland(gr []graph.GraphObject) []int {
	adj := BuildUndirectedAdjList(gr)
	components := FindComponents(adj)

	if len(components) <= 1 {
		return nil // no island
	}

	// Sort components by size
	sort.Slice(components, func(i, j int) bool {
		return len(components[i]) > len(components[j])
	})

	// Return all nodes that are **not in the largest component**
	nodeInMain := make(map[int]bool)
	for _, id := range components[0] {
		nodeInMain[id] = true
	}

	var island []int
	for _, comp := range components[1:] {
		for _, id := range comp {
			island = append(island, id)
		}
	}

	return island
}
