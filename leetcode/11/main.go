// https://leetcode.com/problems/container-with-most-water/?envType=featured-list&envId=top-interview-questions
package main

import "fmt"

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1

	result := 0
	for r > l {
		result = max(result, (r-l)*min(height[l], height[r]))
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
