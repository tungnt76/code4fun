package linkedlist

import "testing"

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return root
	}

	if root.Val == q.Val || root.Val == p.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

func TestLowestCommonAncestor(t *testing.T) {
	// Example usage:
	// Construct a binary tree
	//       3
	//      / \
	//     5   1
	//    / \ / \
	//   6  2 0  8
	//     / \
	//    7   4

	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}

	p := root.Left       // Node with value 5
	q := root.Left.Right // Node with value 2

	lca := lowestCommonAncestor(root, p, q)
	if lca != nil {
		println("Lowest Common Ancestor:", lca.Val) // Should print "Lowest Common Ancestor: 5"
	} else {
		println("Lowest Common Ancestor not found")
	}
}
