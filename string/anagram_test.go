// https://leetcode.com/problems/valid-anagram/submissions/1604277103/
package string

import "testing"

func isAnagram(s string, t string) bool {
	seen := map[rune]int{}
	for _, char := range s {
		seen[char]++
	}

	for _, char := range t {
		seen[char]--
		if seen[char] < 0 {
			return false
		}
	}

	for _, count := range seen {
		if count != 0 {
			return false
		}
	}
	return true
}

func TestIsAnagram(ts *testing.T) {
	s := "anagram"
	t := "nagaram"
	expected := true
	result := isAnagram(s, t)
	if result != expected {
		ts.Errorf("Expected %v, got %v", expected, result)
	}
}
