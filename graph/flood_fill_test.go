package graph

import (
	"fmt"
	"testing"
)

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	m := len(image)
	n := len(image[0])
	queue := []int{sr*n + sc}
	oldColor := image[sr][sc]

	visited := make(map[int]bool)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if visited[node] {
			continue
		}

		x := node / n
		y := node % n
		if image[x][y] != oldColor {
			continue
		}

		image[x][y] = color
		visited[node] = true

		directions := []struct{ dx, dy int }{
			{0, 1},  // right
			{0, -1}, // left
			{1, 0},  // down
			{-1, 0}, // up
		}

		for _, dir := range directions {
			newX := x + dir.dx
			newY := y + dir.dy

			if newX >= 0 && newX < m && newY >= 0 && newY < n && image[newX][newY] == oldColor && !visited[newX*n+newY] {
				queue = append(queue, newX*n+newY)
			}
		}
	}

	return image
}

func TestFloorFill(t *testing.T) {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	sr := 1
	sc := 1
	color := 2

	expected := [][]int{
		{2, 2, 2},
		{2, 2, 0},
		{2, 0, 1},
	}

	result := floodFill(image, sr, sc, color)

	fmt.Println("Result:", result)

	for i := range expected {
		for j := range expected[i] {
			if result[i][j] != expected[i][j] {
				t.Errorf("Expected %d at (%d,%d), got %d", expected[i][j], i, j, result[i][j])
			}
		}
	}
}
