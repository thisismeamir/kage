package graph_analysis

func FindComponents(adj map[int][]int) [][]int {
	visited := make(map[int]bool)
	var components [][]int

	for node := range adj {
		if !visited[node] {
			component := []int{}
			queue := []int{node}
			visited[node] = true

			for len(queue) > 0 {
				current := queue[0]
				queue = queue[1:]
				component = append(component, current)

				for _, neighbor := range adj[current] {
					if !visited[neighbor] {
						visited[neighbor] = true
						queue = append(queue, neighbor)
					}
				}
			}

			components = append(components, component)
		}
	}

	return components
}
