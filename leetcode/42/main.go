// https://leetcode.com/problems/trapping-rain-water/description/
package main

import "fmt"

func trap(nums []int) int {
	if len(nums) <= 2 {
		return 0
	}
	pre := make([]int, len(nums)*2)
	leftHighest := pre[:len(nums)]
	rightHighest := pre[len(nums):]

	{
		curLeftHighest := nums[0]
		for i := range leftHighest {
			curLeftHighest = max(curLeftHighest, nums[i])
			leftHighest[i] = curLeftHighest
		}
	}

	{
		curRightHighest := nums[len(nums)-1]
		for i := range leftHighest {
			j := len(nums) - i - 1
			curRightHighest = max(curRightHighest, nums[j])
			rightHighest[j] = curRightHighest
		}
	}

	trapped := 0
	for i, num := range nums {
		left := leftHighest[i]
		right := rightHighest[i]
		pot := min(left, right)
		trapped += pot - num
	}

	return trapped
}

func main() {
	tests := []struct {
		height   []int
		expected int
	}{
		{
			height:   []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			expected: 6,
		},
		{
			height:   []int{4, 2, 0, 3, 2, 5},
			expected: 9,
		},
	}
	for _, test := range tests {
		actual := trap(test.height)
		if actual != test.expected {
			fmt.Println("[WRONG]", "test:", test, ", actual:", actual)
		}
	}
}
