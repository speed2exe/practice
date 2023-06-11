package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type ListNodeWithIndex struct {
	listNode *ListNode
	index    int
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := make([]ListNodeWithIndex, 0, len(lists))
	for i, n := range lists {
		if n != nil {
			h = append(h, ListNodeWithIndex{
				listNode: n,
				index:    i,
			})
			lists[i] = n.Next
		}
	}
	if len(h) == 0 {
		return nil
	}
	heapify(h)

	var nilNode *ListNode = nil
	var head **ListNode = &nilNode
	var cur = head
	for len(h) > 0 {
		var smallest ListNodeWithIndex = h[0]
    smallest.listNode.Next = nil
    *cur = smallest.listNode
		cur = &(*cur).Next

		// insert next smallest
		next := lists[smallest.index]
		if next == nil {
			if len(h) == 1 {
				break
			} else {
				h[0] = h[len(h)-1]
				h = h[:len(h)-1]
				siftDown(h, 0)
			}
		} else {
			h[0] = ListNodeWithIndex{
				listNode: next,
				index:    smallest.index,
			}
			siftDown(h, 0)
			lists[smallest.index] = next.Next
		}
	}

	return *head
}

func heapify(h []ListNodeWithIndex) {
	for i := len(h) / 2; i >= 0; i-- {
		siftDown(h, i)
	}
}

// [0 1 2 3 4 5 6]
//
//	     0
//	  1     2
//	3  4   5   6
func siftDown(h []ListNodeWithIndex, i int) {
	for {
		val := h[i].listNode.Val
		toSift := i

		right := i*2 + 1
		left := right - 1
		if right < len(h) { // right child exists
			rightVal := h[right].listNode.Val
			if rightVal < val {
				toSift = right
				val = rightVal
			}

			// if right child exists, left child must exist! (no need to check for len again)
			leftVal := h[left].listNode.Val
			if leftVal < val {
				toSift = left
			}
		} else {
			if left < len(h) { // if left child exists
				leftVal := h[left].listNode.Val
				if leftVal < val {
					toSift = left
				}
			}
		}

		if toSift == i {
			return
		}

		h[i], h[toSift] = h[toSift], h[i]
		i = toSift
	}
}

func main() {
	tests := []struct {
		lists    []*ListNode
		expected *ListNode
	}{
		{
			lists: []*ListNode{
				{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
				{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
				{Val: 2, Next: &ListNode{Val: 6}},
			},
			expected: &ListNode{Val: 1,
				Next: &ListNode{Val: 1,
					Next: &ListNode{Val: 2,
						Next: &ListNode{Val: 2,
							Next: &ListNode{Val: 4,
								Next: &ListNode{Val: 5,
									Next: &ListNode{Val: 5,
										Next: &ListNode{Val: 6,
											Next: &ListNode{Val: 1}}}}}}}}},
		},
	}
	for _, test := range tests {
		actual := mergeKLists(test.lists)
		actual_str := sprintListNode(actual)
		fmt.Println(actual_str)
		// expected_str := sprintListNode(test.expected)
		// if actual_str != expected_str {
		// 	fmt.Println("[WRONG]", "test:", test, "actual:", actual, "expected:", test.expected)
		// }
	}
}

func sprintListNode(n *ListNode) string {
	var b strings.Builder
	b.WriteByte('[')
	for n != nil {
		b.WriteString(fmt.Sprint(n.Val))
		b.WriteByte(' ')
		n = n.Next
	}
	b.WriteByte(']')
	return b.String()
}
