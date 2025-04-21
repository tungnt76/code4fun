// https://leetcode.com/problems/ransom-note/description/
package hashmap

import "testing"

func canConstruct(ransomNote string, magazine string) bool {
	seen := map[rune]int{}
	for _, c := range magazine {
		seen[c]++
	}

	for _, c := range ransomNote {
		seen[c]--
		if seen[c] < 0 {
			return false
		}
	}

	return true
}

func TestCanConstruct(t *testing.T) {
	scenarios := []struct {
		ransomNote string
		magazine   string
		expected   bool
	}{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
		{"aabb", "ab", false},
		{"aabb", "aab", false},
		{"aabb", "aabb", true},
		{"aabb", "aabbc", true},
	}

	for _, scenario := range scenarios {
		result := canConstruct(scenario.ransomNote, scenario.magazine)
		if result != scenario.expected {
			t.Errorf("Expected %v, got %v", scenario.expected, result)
		}
	}
}
