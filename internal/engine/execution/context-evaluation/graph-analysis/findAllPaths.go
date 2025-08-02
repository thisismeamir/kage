package graph_analysis

func FindAllPaths(
	adj map[int][]int,
	current int,
	terminals map[int]bool,
	path []int,
	visited map[int]bool,
	results *[][]int,
) {
	if visited[current] {
		return // avoid cycles
	}

	path = append(path, current)

	if terminals[current] {
		// Found a complete path
		cp := make([]int, len(path))
		copy(cp, path)
		*results = append(*results, cp)
		return
	}

	visited[current] = true
	for _, neighbor := range adj[current] {
		FindAllPaths(adj, neighbor, terminals, path, visited, results)
	}
	visited[current] = false // backtrack
}
