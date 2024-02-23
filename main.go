package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func randomBetweenInt64(min, max int64) int64 {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		panic(err) // Handle errors appropriately
	}

	return n.Int64() + min
}

type User struct {
	A int
	B map[int]int
}

func main() {
	arr := [][]int32{{112, 42, 83, 119}, {56, 125, 56, 49}, {15, 78, 101, 43}, {62, 98, 114, 119}}
	fmt.Println(flippingMatrix(arr))
}

func flippingMatrix(matrix [][]int32) int32 {
	// Write your code here
	n := len(matrix[0]) / 2
	length := len(matrix)
	for i := range matrix {
		start := []int32{}
		end := []int32{}
		for j := 0; j < n; j++ {
			start = append(start, matrix[j][i])
			end = append(end, matrix[j][length-i-1])
		}

		if sum(start) < sum(end) {
			for j := 0; j < n; j++ {
				temp := matrix[j][i]
				matrix[j][i] = matrix[j][length-i-1]
				matrix[j][length-i-1] = temp
			}
		}
	}

	sumUpperLeftQuadrant := int32(0)
	for i := 0; i < n; i++ {
		row := matrix[i]
		if sum(row[:n]) > sum(row[len(matrix[i])-n:]) {
			sumUpperLeftQuadrant += sum(row[:n])
		} else {
			sumUpperLeftQuadrant += sum(row[len(matrix[i])-n:])
		}
	}

	return sumUpperLeftQuadrant
}

func sum(arr []int32) int32 {
	sum := int32(0)
	for _, num := range arr {
		sum += num
	}
	return sum
}

func countContinuedValues(input []string) []int {
	counts := make([]int, 0)
	currentCount := 1
	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			currentCount++
		} else {
			counts = append(counts, currentCount)
			currentCount = 1
		}
	}
	counts = append(counts, currentCount)

	fmt.Println(counts)
	results := make([]int, 0)
	for _, count := range counts {
		results = append(results, count)
		if count > 1 {
			for i := 1; i < count; i++ {
				results = append(results, count)
			}
		}
	}

	return results
}
