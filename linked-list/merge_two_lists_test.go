// https://leetcode.com/problems/merge-two-sorted-lists/
package linkedlist

import (
	"fmt"
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			p.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			p.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		p = p.Next
	}

	if list1 != nil {
		p.Next = list1
	} else {
		p.Next = list2
	}

	return dummy.Next
}

func TestMergeTwoLists(t *testing.T) {
	list1 := newListNode([]int{1, 2, 4})
	list2 := newListNode([]int{1, 3, 4})

	mergedList := mergeTwoLists(list1, list2)
	print(mergedList)
	fmt.Printf("%+v\n", &mergedList)
}
