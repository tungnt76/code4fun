// https://www.hackerrank.com/challenges/breaking-best-and-worst-records/problem?isFullScreen=true
package easy

import (
	"math"
)

/*
 * Complete the 'breakingRecords' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY scores as parameter.
 */

func BreakingRecords(scores []int32) []int32 {
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
