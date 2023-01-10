// https://leetcode.com/problems/linked-list-cycle/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head

	for {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		if fast == nil {
			return false
		}
		if slow == fast {
			return true
		}
	}
}
