// https://leetcode.com/problems/continuous-subarray-sum/description/

package main

import (
	"fmt"
)

// edge case: [36, ...]

// [23, 2, 4, ...], k = 6
// 0 loop: remainder of 5
// 1 loop: remainder of 25%6 = 1
// 2 loop: remainder of 29%6 = 5
func checkSubarraySum(nums []int, k int) bool {
	// remainder -> index
	m := map[int]int{0: -1}

	sum := 0

	for i, num := range nums {
		sum += num

		r := sum % k
		j, exists := m[r]
		if !exists {
			m[r] = i
			continue
		}
		if i-j > 1 {
			return true
		}
	}

	return false
}

func main() {
	// fmt.Println(checkSubarraySum([]int{23, 2, 4, 6, 7}, 6))
	// fmt.Println(checkSubarraySum([]int{23, 2, 4, 6, 6}, 7))
	fmt.Println(checkSubarraySum([]int{1, 0}, 7))
}
