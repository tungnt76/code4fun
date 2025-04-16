// https://leetcode.com/problems/longest-repeating-character-replacement/
package string

import "testing"

func characterReplacement(s string, k int) int {
	counter := make(map[byte]int)
	maxCount := 0
	left := 0
	maxLength := 0
	for right := 0; right < len(s); right++ {
		counter[s[right]]++
		maxCount = max(maxCount, counter[s[right]])

		windowSize := right - left + 1
		if windowSize-maxCount > k {
			counter[s[left]]--
			if counter[s[left]] == 0 {
				delete(counter, s[left])
			}
			left++
		}

		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}

func TestCharacterReplacement(t *testing.T) {
	s := "AABABBA"
	k := 1
	expected := 4
	result := characterReplacement(s, k)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	s = "ABAB"
	k = 2
	expected = 4
	result = characterReplacement(s, k)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	s = "AABABBA"
	k = 1
	expected = 4
	result = characterReplacement(s, k)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
