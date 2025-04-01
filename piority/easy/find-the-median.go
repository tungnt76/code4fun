// mock test
package easy

import (
	"sort"
)

// length is odd
func FindTheMedian(arr []int32) int32 {
	// Write your code here
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	return arr[len(arr)/2+1]
}
