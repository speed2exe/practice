package main

import "strings"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//	    1    <--- root
//	  /   \
//	 2     3
//	/ \   / \
//
// 4  5  6   7
// startValue: 4
// destValue: 6
// return: UURL
func getDirections(root *TreeNode, startValue int, destValue int) string {
	pathToStart, ok1 := findPath(root, startValue, nil) // [f, f]
	pathToDest, ok2 := findPath(root, destValue, nil)   // [t, t]
	if !ok1 || !ok2 {
		panic("unreachable")
	}

	// [..., f, f] => [f, f] => ['U', 'U']
	// [..., t, t] => [t, t] => ['R', 'R']
	var builder strings.Builder
	for {
		if len(pathToDest) == 0 {
			break
		}
		if len(pathToStart) == 0 {
			break
		}
		if pathToDest[0] != pathToStart[0] {
			break
		}
		pathToStart = pathToStart[1:]
		pathToDest = pathToDest[1:]
	}

	for i := 0; i < len(pathToStart); i++ {
		builder.WriteByte('U')
	}
	for i := 0; i < len(pathToDest); i++ {
		if pathToDest[i] {
			builder.WriteByte('R')
		} else {
			builder.WriteByte('L')
		}
	}

	return builder.String()
}

func findPath(root *TreeNode, target int, cur_path []bool) ([]bool, bool) {
	if root == nil {
		return nil, false
	}
	if root.Val == target {
		return cur_path, true
	}
	res1, found1 := findPath(root.Left, target, append(cur_path, false))
	if found1 {
		return res1, true
	}
	res2, found2 := findPath(root.Right, target, append(cur_path, true))
	if found2 {
		return res2, true
	}
	return nil, false
}
