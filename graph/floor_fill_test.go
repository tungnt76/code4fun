// https://leetcode.com/problems/flood-fill/
package graph

func distinctNumber(arr []int) int {
	if len(arr) < 0 {
		return -1
	}

	m := map[int]int{}
	for _, c := range m {
		m[c]++
	}

	max := -1
	for k, v := range m {
		if v == 1 {
			if max < k {
				max = k
			}
		}
	}

	return max
}
