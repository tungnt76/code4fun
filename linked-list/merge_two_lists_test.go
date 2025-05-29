// https://leetcode.com/problems/merge-two-sorted-lists/
package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var head *ListNode

	for list1 == nil && list2 == nil {
		if list1.Val < list2.Val {
			if head == nil {
				head = list1
			} else {
				head.Next = list1
			}

			continue
		}
		
		if head == nil {
			
		}
	}

	if list1.Val <= list2.Val {
		return mergeTwoLists(list1.Next, list2)
	}

	return mergeTwoLists(list1, list2.Next)
}
