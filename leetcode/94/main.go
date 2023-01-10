// https://leetcode.com/problems/binary-tree-inorder-traversal/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)

	stack := []*TreeNode{root}
	for {
		if len(stack) == 0 {
			return result
		}

		top := stack[len(stack)-1]
		if top.Left != nil {
			stack = append(stack, top.Left)
			top.Left = nil
			continue
		}

		result = append(result, top.Val)
		stack = stack[:len(stack)-1]

		if top.Right != nil {
			stack = append(stack, top.Right)
		}
	}
}
