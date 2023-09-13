package main

import (
	"math"
	"sort"
)

type User struct {
	Name    string
	Address string
}

func main() {
	// func main() {
	// fmt.Println(getTotalX([]int32{2, 4}, []int32{16, 32, 96}))
	// // 10 5 20 20 4 5 2 25 1
	// fmt.Println(breakingRecords([]int32{10, 5, 20, 20, 4, 5, 2, 25, 1}))

	// year := 1918
	// // dayOfYear := 256

	// // Construct the date using time.Date function
	// date := time.Date(year, time.February, 14, 0, 0, 0, 0, time.UTC).AddDate(0, 0, 256-32)

	// fmt.Printf("The 256th day of %d is on %s\n", year, date.Format("2006-01-02"))

	// user := User{}
	// user.Name = "nttung"

	// fmt.Println(user)
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
