// https://leetcode.com/problems/jump-game-ii/description/
package main

import (
	"fmt"
)

// nums: [2, 3, 1, 7, 4]
// cur:         ^
// end:               ^
// far:               ^
// min: 2
// returns 2
//
// nums: [1, 0, 0, 7]
// cur:      ^
// end:      ^
// far:      ^
// smallest: 3
// returns 2
//
// nums: [1, 2, 3]
// cur:      ^
// end:      ^
// far:      ^
// smallest: 1
// returns 2
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	last := len(nums) - 1

	smallest := 1
	cur := 0
	end := nums[cur]
	far := end

	for {
		if end >= last {
			return smallest
		}

		cur++

		// update furthest
		far = max(far, nums[cur]+cur)

		if cur == end {
			smallest++
			if far == end {
				panic("cannot jump to end") // no way to get to last elem
			}
			end = far
		}

		// fmt.Println("cur:", cur, "end:", end, "far:", far, "smallest:", smallest)
	}
}

func main() {
	tests := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{2, 3, 1, 1, 4},
			expected: 2,
		},
		{
			nums:     []int{2, 3, 0, 1, 4},
			expected: 2,
		},
		{
			nums:     []int{1, 2, 3},
			expected: 2,
		},
	}
	for _, test := range tests {
		actual := jump(test.nums)
		if actual != test.expected {
			fmt.Println("[WRONG]", "nums:", test.nums, "actual:", actual, "expected:", test.expected)
		}
	}
}
