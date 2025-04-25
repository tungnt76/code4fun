package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func newListNode(values []int) *ListNode {
	var head *ListNode
	var pointer *ListNode

	for i := 0; i < len(values); i++ {
		curNode := &ListNode{
			Val: i,
		}

		if i == 0 {
			head = curNode
			pointer = head
			continue
		}

		pointer.Next = curNode
		pointer = pointer.Next
	}

	return head
}

func print(head *ListNode) {
	for head != nil {
		if head.Next != nil {
			fmt.Printf("%d -> ", head.Val)
		} else {
			fmt.Printf("%d\n", head.Val)
		}
		head = head.Next
	}
}
