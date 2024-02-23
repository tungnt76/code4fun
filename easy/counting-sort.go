// https://www.hackerrank.com/challenges/one-week-preparation-kit-countingsort1/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=preparation-kits&playlist_slugs%5B%5D=one-week-preparation-kit&playlist_slugs%5B%5D=one-week-day-two
package easy

func countingSort(arr []int32) []int32 {
	// Write your code here
	countArr := make([]int32, 100)
	for _, num := range arr {
		countArr[num]++
	}
	
	return countArr
}
