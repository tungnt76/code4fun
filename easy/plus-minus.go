// https://www.hackerrank.com/challenges/one-week-preparation-kit-plus-minus/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=preparation-kits&playlist_slugs%5B%5D=one-week-preparation-kit&playlist_slugs%5B%5D=one-week-day-one
package easy

import "fmt"

func PlusMinus(arr []int32) {
	// Write your code here
	positiveCount, negativeCount, zeroCount, n := 0, 0, 0, len(arr)
	for _, num := range arr {
		if num > 0 {
			positiveCount++
		} else if num < 0 {
			negativeCount++
		} else {
			zeroCount++
		}
	}
	fmt.Printf("%.6f\n", float64(positiveCount)/float64(n))
	fmt.Printf("%.6f\n", float64(negativeCount)/float64(n))
	fmt.Printf("%.6f\n", float64(zeroCount)/float64(n))
}
