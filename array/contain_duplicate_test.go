package array

import "testing"

func containsDuplicate(nums []int) bool {
	seen := map[int]bool{}
	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}

	return false
}

func TestContainsDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	expected := true
	result := containsDuplicate(nums)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	nums = []int{1, 2, 3, 4}
	expected = false
	result = containsDuplicate(nums)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
