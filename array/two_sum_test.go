package array

import "testing"

func twoSum(nums []int, target int) []int {
	indexes := map[int]int{}
	for i, num := range nums {
		if index, ok := indexes[target-num]; ok {
			return []int{index, i}
		}
		indexes[num] = i
	}

	return nil
}

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}
	result := twoSum(nums, target)
	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}
}
