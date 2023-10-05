package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/iamtony/golang/hackerrank/easy"
)

// type User struct {
// 	Name    string
// 	Address string
// }

func main() {
	// fmt.Println(math.Ceil(3.2), math.Round(3.7))
	fmt.Println(easy.Birthday([]int32{1, 2, 1, 3, 2}, 3, 2))
	fmt.Println(easy.DivisibleSumPairs(6, 3, []int32{1, 3, 2, 6, 1, 2}))
}

func breakingRecords(scores []int32) []int32 {
	// Write your code here
	var min int32 = math.MaxInt32
	var max int32 = math.MinInt32

	var maxCount, minCount int32

	for _, score := range scores {
		if score < min {
			if min != math.MaxInt32 {
				minCount += 1
			}

			min = score
		}

		if score > max {
			if max != math.MinInt32 {
				maxCount += 1
			}
			max = score
		}
	}

	return []int32{maxCount, minCount}
}

func getTotalX(a []int32, b []int32) int32 {
	// Write your code here
	sort.SliceStable(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	sort.SliceStable(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	start := a[len(a)-1]
	end := b[0]

	rs := []int32{}

LOOP:
	for i := start; i <= end; i++ {
		for _, v := range a {
			if i%v != 0 {
				continue LOOP
			}
		}

		for _, v := range b {
			if v%i != 0 {
				continue LOOP
			}
		}

		rs = append(rs, i)
	}

	return int32(len(rs))
}
