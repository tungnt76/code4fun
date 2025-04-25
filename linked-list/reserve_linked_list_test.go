// https://leetcode.com/problems/reverse-linked-list/description/

package linkedlist

import "testing"

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curHead := head
	for curHead != nil {
		nextTemp := curHead.Next
		curHead.Next = prev
		prev = curHead
		curHead = nextTemp
	}
	
	return prev
}

func TestReserveLinkedList(t *testing.T) {
	head := newListNode([]int{1, 2, 3, 4, 5})
	print(head)
	reserved := reverseList(head)
	print(reserved)
}
