// https://www.hackerrank.com/challenges/one-week-preparation-kit-diagonal-difference/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=preparation-kits&playlist_slugs%5B%5D=one-week-preparation-kit&playlist_slugs%5B%5D=one-week-day-two
package easy

import "math"

func DiagonalDifference(arr [][]int32) int32 {
	// Write your code here
	l2r, r2l := int32(0), int32(0)
	for i := range arr {
		l2r += arr[i][i]
		r2l += arr[i][len(arr)-i-1]
	}

	return int32(math.Abs(float64(l2r - r2l)))
}
