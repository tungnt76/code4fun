// https://www.hackerrank.com/challenges/between-two-sets/problem?isFullScreen=true
package easy

import (
	"sort"
)

/*
 * Complete the 'getTotalX' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func GetTotalX(a []int32, b []int32) int32 {
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
