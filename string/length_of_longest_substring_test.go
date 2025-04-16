// https://leetcode.com/problems/longest-substring-without-repeating-characters/
package string

import "testing"

// Example 1:

// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
// Example 2:

// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Example 3:

// Input: s = "pwpkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

// p: 0
//

func lengthOfLongestSubstring(s string) int {
	seen := map[rune]int{}
	start, end := 0, 0
	max := 0
	for end < len(s) {
		c := rune(s[end])
		if idx, ok := seen[c]; ok && idx >= start {
			start = idx + 1
		}

		seen[c] = end
		end++

		if max < end-start {
			max = end - start
		}
	}

	return max
}

func TestLengthOfLongestSubString(t *testing.T) {
	s := "abcabcbb"
	expected := 3
	result := lengthOfLongestSubstring(s)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	s = "bbbbb"
	expected = 1
	result = lengthOfLongestSubstring(s)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	s = "pwpwkpew"
	expected = 4
	result = lengthOfLongestSubstring(s)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
	
	s = " "
	expected = 1
	result = lengthOfLongestSubstring(s)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
