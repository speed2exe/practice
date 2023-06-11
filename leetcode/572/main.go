// https://leetcode.com/problems/subtree-of-another-tree/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var isSameTree func(p, q *TreeNode) bool
	isSameTree = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}

	var helper func(p, q *TreeNode) bool
	helper = func(p, q *TreeNode) bool {
		if p == nil {
			return false
		}
		if isSameTree(p, q) {
			return true
		}
		return helper(p.Left, q) || helper(p.Right, q)
	}

	return helper(root, subRoot)
}

func main() {
	tests := []struct {
		root     *TreeNode
		subRoot  *TreeNode
		expected bool
	}{
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
			subRoot: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			expected: true,
		},
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 0,
						},
					},
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
			subRoot: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			expected: false,
		},
	}
	for _, test := range tests {
		actual := isSubtree(test.root, test.subRoot)
		if actual != test.expected {
			fmt.Println("[WRONG]", "test:", test, ", actual:", actual)
		}
	}
}
