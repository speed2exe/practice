// https://leetcode.com/problems/maximum-depth-of-binary-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	stack := []*TreeNode{root}
	var depth int
	for {
		if len(stack) == 0 {
			return depth
		}
		depth++
		newStack := []*TreeNode{}
		for _, node := range stack {
			if node.Left != nil {
				newStack = append(newStack, node.Left)
			}
			if node.Right != nil {
				newStack = append(newStack, node.Right)
			}
		}
		stack = newStack
	}
}
