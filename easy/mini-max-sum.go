// https://www.hackerrank.com/challenges/one-week-preparation-kit-mini-max-sum/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=preparation-kits&playlist_slugs%5B%5D=one-week-preparation-kit&playlist_slugs%5B%5D=one-week-day-one
package easy

import "fmt"

func MiniMaxSum(arr []int32) {
	// Write your code here
	sum, min, max := int64(0), int32(arr[0]), int32(arr[0])
	for _, num := range arr {
		sum += int64(num)
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	fmt.Printf("%d %d", sum-int64(max), sum-int64(min))
}
