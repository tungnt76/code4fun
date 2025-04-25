// https://leetcode.com/problems/linked-list-cycle/

package linkedlist

// use two pointer
// slow and fast
// slow will move 1 step
// fast will move 2 steps

func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}
