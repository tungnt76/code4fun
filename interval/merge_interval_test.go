// https://leetcode.com/problems/merge-intervals/description/
package interval

import (
	"fmt"
	"sort"
	"testing"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	results := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		last := results[len(results)-1]

		if last[1] >= cur[0] {
			results[len(results)-1] = []int{last[0], max(last[1], cur[1])}
		} else {
			results = append(results, intervals[i])
		}
	}

	return results
}

func TestMerge(t *testing.T) {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
}
