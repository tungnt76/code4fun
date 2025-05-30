package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func newListNode(values []int) *ListNode {
	head := &ListNode{}
	pointer := head

	for _, value := range values {
		curNode := &ListNode{
			Val: value,
		}

		pointer.Next = curNode
		pointer = pointer.Next
	}

	return head.Next
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
