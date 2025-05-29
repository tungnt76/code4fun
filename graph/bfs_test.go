package graph

import (
	"fmt"
	"testing"
)

// BFS performs a breadth-first search starting from the `start` node
func bfs(graph map[int][]int, start int) {
	visited := make(map[int]bool)
	queue := []int{start}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if visited[node] {
			continue
		}

		fmt.Printf("Visited: %d\n", node)
		visited[node] = true

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}
}

func TestBfs(t *testing.T) {
	graph := map[int][]int{
		1: {2, 3, 4},
		3: {5, 6},
	}

	bfs(graph, 1)
}
