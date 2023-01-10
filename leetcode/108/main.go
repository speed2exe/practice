// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	midIdx := len(nums) / 2
	midElem := nums[midIdx]
	left := nums[:midIdx]
	right := nums[midIdx:]
	return &TreeNode{
		Val:   midElem,
		Left:  sortedArrayToBST(left),
		Right: sortedArrayToBST(right),
	}
}
