// https://leetcode.com/problems/add-two-numbers/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	ans := addTwoNumbersWithCarry(l1, l2, false)
	if ans == nil {
		return &ListNode{Val: 0}
	}
	return ans
}

func addTwoNumbersWithCarry(l1, l2 *ListNode, carry bool) *ListNode {
	switch {
	case l1 == nil && l2 == nil:
		if carry {
			return &ListNode{Val: 1}
		}
		return nil
	case l1 == nil && l2 != nil:
		return addTwoNumbersWithCarry(l2, l1, carry)
	case l1 != nil && l2 == nil:
		if carry {
			l1.Val++
		}
		var carryNext bool
		if l1.Val >= 10 {
			l1.Val -= 10
			carryNext = true
		}
		l1.Next = addTwoNumbersWithCarry(l1.Next, nil, carryNext)
		return l1
	default:
		var carryNext bool
		if carry {
			l1.Val++
		}
		l1.Val += l2.Val
		if l1.Val >= 10 {
			l1.Val -= 10
			carryNext = true
		}
		l1.Next = addTwoNumbersWithCarry(l1.Next, l2.Next, carryNext)
		return l1
	}
}
