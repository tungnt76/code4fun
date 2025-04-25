// https://leetcode.com/problems/insert-interval/description

package interval

import (
	"fmt"
	"testing"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	insertedIndex := 0
	for i := len(intervals) - 1; i >= 0; i-- {
		if newInterval[0] >= intervals[i][0] {
			insertedIndex = i
			break
		}
	}

	// var newIntervals [][]int
	// copy(newIntervals, intervals)
	
	// newIntervals = append(newIntervals, intervals[insertedIndex+1:]...)

	// fmt.Println(newIntervals)
	intervals = append(intervals[:insertedIndex], append([][]int{newInterval}, intervals[insertedIndex:]...)...)
	fmt.Println(intervals)
	return merge(intervals)
}

func TestInsert(t *testing.T) {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
}
