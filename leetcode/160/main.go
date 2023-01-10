// https://leetcode.com/problems/intersection-of-two-linked-lists/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	tailA := headA
	lenA := 1
	for tailA.Next != nil {
		tailA = tailA.Next
		lenA++
	}

	tailB := headB
	lenB := 1
	for tailB.Next != nil {
		tailB = tailB.Next
		lenB++
	}

	if tailA != tailB {
		return nil
	}

	curA := headA
	curB := headB
	if lenA > lenB {
		for i := 0; i < lenA-lenB; i++ {
			curA = curA.Next
		}
	} else {
		for i := 0; i < lenB-lenA; i++ {
			curB = curB.Next
		}
	}

	for curA != curB {
		curA = curA.Next
		curB = curB.Next
	}

	return curA
}
