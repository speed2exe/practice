// https://leetcode.com/problems/kth-smallest-element-in-a-bst/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var ans int

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)

		// action
		if k == 1 {
			ans = root.Val
		}
		k--

		if k > 0 {
			dfs(root.Right)
		}
	}

	dfs(root)

	return ans
}

func main() {
	tests := []struct {
		root     *TreeNode
		k        int
		expected int
	}{
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
					},
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			k:        1,
			expected: 1,
		},
		{
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 1,
						},
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
			k:        3,
			expected: 3,
		},
	}
	for _, test := range tests {
		actual := kthSmallest(test.root, test.k)
		if actual != test.expected {
			fmt.Println("[WRONG]", "test:", test, ", actual:", actual)
		}
	}
}
