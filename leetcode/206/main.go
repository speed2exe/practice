// https://leetcode.com/problems/reverse-linked-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// return reverseListRecursive(head)
	return reverseListIterative(head)
}

func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	return reverseListRecursiveHelper(head)
}

func reverseListRecursiveHelper(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	next := head.Next
	newHead := reverseListRecursiveHelper(next)
	next.Next = head
	head.Next = nil
	return newHead
}

func reverseListIterative(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	current := head
	next := head.Next
	head.Next = nil
	for {
		nextNext := next.Next
		next.Next = current
		// 1 <- 2 -> 3 -> 4 -> null
		// c    n   nn

		if nextNext == nil {
			return next
		}

		// update the pointers
		current = next
		next = nextNext
	}
}

// 1 -> 2 -> 3 -> 4 -> null
// c    n   nn

// null <- 1 <- 2 <- 3 <- 4
