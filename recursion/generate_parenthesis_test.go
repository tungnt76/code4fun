// https://leetcode.com/problems/generate-parentheses/description/
package recursion

import (
	"fmt"
	"testing"
)

// [()]

func generateParenthesis(n int) []string {
	stack := ""
	results := []string{}

	backtrack(n, 0, 0, stack, &results)

	return results
}

func backtrack(n int, open int, close int, stack string, results *[]string) {
	if len(stack) == n*2 {
		*results = append(*results, stack)
	}

	if open < n {
		stack += "("
		backtrack(n, open+1, close, stack, results)
		stack = stack[:len(stack)-1]
	}

	if close < open {
		stack += ")"
		backtrack(n, open, close+1, stack, results)
		stack = stack[:len(stack)-1]
	}
}

func TestGenerateParenthesis(t *testing.T) {
	n := 3
	results := generateParenthesis(n)
	fmt.Println(results)
}
