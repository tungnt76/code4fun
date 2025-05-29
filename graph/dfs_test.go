package graph

import (
	"fmt"
	"testing"
)

// DFS performs a depth-first search starting from the `node`
func dfs(graph map[int][]int, node int, visited map[int]bool) {
	if visited[node] {
		return
	}

	fmt.Printf("Visited: %d\n", node)
	visited[node] = true

	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			dfs(graph, neighbor, visited)
		}
	}
}

func TestDfs(t *testing.T) {
	graph := map[int][]int{
		1: {2, 3, 4},
		3: {5, 6},
	}

	visited := make(map[int]bool)
	dfs(graph, 1, visited)
}
