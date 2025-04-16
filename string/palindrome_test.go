// https://leetcode.com/problems/valid-palindrome/description/

package string

import (
	"testing"
	"unicode"
)

func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	for l < r {
		if !unicode.IsLetter(rune(s[l])) && !unicode.IsDigit(rune(s[l])) {
			l++
			continue
		}
		if !unicode.IsLetter(rune(s[r])) && !unicode.IsDigit(rune(s[r])) {
			r--
			continue
		}

		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}
		l++
		r--
	}

	return true
}

func TestIsPalindrome(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	expected := true
	result := isPalindrome(s)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	s = ""
	expected = true
	result = isPalindrome(s)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	s = "race a car"
	expected = false
	result = isPalindrome(s)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
