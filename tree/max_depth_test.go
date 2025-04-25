// https://leetcode.com/problems/maximum-depth-of-binary-tree/
package tree

import (
	"fmt"
	"testing"
)

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func TestMaxDept(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	fmt.Println(maxDepth(root))
}
