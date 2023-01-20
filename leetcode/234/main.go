// https://leetcode.com/problems/palindrome-linked-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// A -> B -> C -> D
// count: 2
// middle: 1

// A -> B -> C -> D -> E
// count: 5
// middle: 1
// A -> B -> null
// C -> D -> E -> null
// if extra, discard middle
// reverse: D -> C -> null
// compare A -> B and D -> C
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// Find the middle of the list
	count := 0
	current := head
	for current != nil {
		count++
		current = current.Next
	}

	middle := count/2 - 1 // because we want to get the node just before the middle
	current = head
	for i := 0; i < middle; i++ {
		current = current.Next
	}

	next := current.Next
	current.Next = nil // cut off the link from first linkedlist to second linkedlist
	if count%2 == 1 {  // edge case handling for odd number of nodes(discard middle)
		next = next.Next
	}

	reversed := reverseListIterative(next)

	// compare the two linkedlists
	for {
		if head == nil && reversed == nil {
			return true
		}
		if head == nil || reversed == nil {
			return false
		}
		if head.Val != reversed.Val {
			return false
		}
		head = head.Next
		reversed = reversed.Next
	}
}

// func printLinkedList(head *ListNode) {
// 	fmt.Print("[")
// 	for head != nil {
// 		fmt.Print(" ", head)
// 		head = head.Next
// 	}
// 	fmt.Println("]")
// }

// solution imported from previous problem
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

		if nextNext == nil {
			return next
		}

		// update the pointers
		current = next
		next = nextNext
	}
}

// TODO: recursive solution
