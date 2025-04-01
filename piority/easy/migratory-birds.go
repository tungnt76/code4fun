// https://www.hackerrank.com/challenges/migratory-birds/problem
package easy

/* The function accepts INTEGER_ARRAY arr as parameter.
 */

func MigratoryBirds(arr []int32) int32 {
	// Write your code here
	m := map[int32]int{}
	for _, num := range arr {
		m[num] += 1
	}

	max := 0
	var result int32
	for k, v := range m {
		if v > max {
			max = v
			result = k
		}
	}

	return int32(result)
}
