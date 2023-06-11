// https://leetcode.com/problems/container-with-most-water/?envType=featured-list&envId=top-interview-questions
package main

import "fmt"

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	curMax := min(height[l], height[r]) * (r - l)

	for r > l {
		if height[l] > height[r] {
			// move r to next highest
			curRight := height[r]
			r--
			for curRight > height[r] {
				r--
				if r == l {
					return curMax
				}
			}
		} else {
			// move l until next highest
			curLeft := height[l]
			l++
			for curLeft > height[l] {
				l++
				if l == r {
					return curMax
				}
			}
		}
		curMax = max(curMax, min(height[l], height[r])*(r-l))
	}

	return curMax
}

func main() {
	tests := []struct {
		height   []int
		expected int
	}{
		{
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
	}
	for _, test := range tests {
		actual := maxArea(test.height)
		if actual != test.expected {
			fmt.Println("[WRONG]", "test:", test, ", actual:", actual)
		}
	}
}
