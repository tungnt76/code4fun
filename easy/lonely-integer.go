// https://www.hackerrank.com/challenges/one-week-preparation-kit-lonely-integer/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=preparation-kits&playlist_slugs%5B%5D=one-week-preparation-kit&playlist_slugs%5B%5D=one-week-day-two
package easy

func LonelyInteger(arr []int32) int32 {
	// Write your code here
	lonelyMap := map[int32]bool{}
	for _, num := range arr {
		if _, ok := lonelyMap[num]; ok {
			lonelyMap[num] = false
		} else {
			lonelyMap[num] = true
		}
	}

	for num, isAlone := range lonelyMap {
		if isAlone {
			return num
		}
	}
	return 0
}
