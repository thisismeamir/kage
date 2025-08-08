package toposort

import (
	"fmt"
	"github.com/thisismeamir/kage/pkg/graph"
)

type TopoLevel struct {
	Level int   `json:"level"`
	Nodes []int `json:"nodes"`
}

type TopoSchedule struct {
	Order  []int       `json:"order"`
	Levels []TopoLevel `json:"levels"`
}

func TopoSort(gr []graph.GraphObject) (TopoSchedule, error) {
	// Map: node ID -> GraphObject
	nodeMap := make(map[int]graph.GraphObject)
	// Map: node ID -> number of incoming edges
	inDegree := make(map[int]int)
	// Adjacency list
	adj := make(map[int][]int)

	// Initialize graph structures
	for _, node := range gr {
		nodeMap[node.Id] = node
		if _, exists := inDegree[node.Id]; !exists {
			inDegree[node.Id] = 0
		}
		for _, target := range node.Outgoing {
			inDegree[target]++
			adj[node.Id] = append(adj[node.Id], target)
		}
	}

	// Queue for zero in-degree nodes (current level)
	var queue []int
	for id, deg := range inDegree {
		if deg == 0 {
			queue = append(queue, id)
		}
	}

	var (
		order  []int
		levels []TopoLevel
		level  = 0
	)

	for len(queue) > 0 {
		var nextQueue []int
		var levelNodes []int

		for _, nodeID := range queue {
			order = append(order, nodeID)
			levelNodes = append(levelNodes, nodeID)

			for _, neighbor := range adj[nodeID] {
				inDegree[neighbor]--
				if inDegree[neighbor] == 0 {
					nextQueue = append(nextQueue, neighbor)
				}
			}
		}

		levels = append(levels, TopoLevel{
			Level: level,
			Nodes: levelNodes,
		})
		level++
		queue = nextQueue
	}

	// If not all nodes are sorted, there's a cycle
	if len(order) != len(gr) {
		return TopoSchedule{}, fmt.Errorf("cycle detected: topological sort not possible")
	}

	return TopoSchedule{
		Order:  order,
		Levels: levels,
	}, nil
}
