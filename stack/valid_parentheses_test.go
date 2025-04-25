// https://leetcode.com/problems/valid-parentheses/description/
package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isValid(s string) bool {
	closeM := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := []rune{}
	for _, r := range s {
		if r == '(' || r == '{' || r == '[' {
			stack = append(stack, r)
		} else if r == ')' || r == '}' || r == ']' {
			if len(stack) == 0 {
				return false
			}

			if closeM[r] != stack[len(stack)-1] {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func TestIsValidParentheses(t *testing.T) {
	scenarios := []struct {
		Input    string
		Expected bool
	}{
		{
			Input:    "",
			Expected: true,
		},
		{
			Input:    "[]",
			Expected: true,
		},
	}

	for _, sc := range scenarios {
		assert.Equal(t, sc.Expected, isValid(sc.Input))
	}
}
