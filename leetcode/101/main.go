// https://leetcode.com/problems/symmetric-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	// return isSymmetricRecursive(root.Left, root.Right)
	return isSymmetricIterative(root.Left, root.Right)
}

func isSymmetricRecursive(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left != nil || right != nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}

	return isSymmetricRecursive(left.Left, right.Right) && isSymmetricRecursive(right.Left, left.Right)
}

func isSymmetricIterative(left, right *TreeNode) bool {
	leftStack := []*TreeNode{left}
	rightStack := []*TreeNode{right}

	for {
		if len(leftStack) == 0 && len(rightStack) == 0 {
			return true
		}
		leftTop := leftStack[len(leftStack)-1]
		rightTop := rightStack[len(rightStack)-1]

		leftStack = leftStack[:len(leftStack)-1]
		rightStack = rightStack[:len(rightStack)-1]

		if leftTop == nil && rightTop == nil {
			continue
		}
		if leftTop == nil || rightTop == nil {
			return false
		}
		if leftTop.Val != rightTop.Val {
			return false
		}

		leftStack = append(leftStack, leftTop.Left)
		rightStack = append(rightStack, rightTop.Right)

		leftStack = append(leftStack, leftTop.Right)
		rightStack = append(rightStack, rightTop.Left)
	}
}
