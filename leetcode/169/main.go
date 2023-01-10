// https://leetcode.com/problems/majority-element/discussion/
package main

func majorityElement(nums []int) int {
	candidate := nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {
		next := nums[i]
		if next == candidate {
			count++
			continue
		}

		count--
		if count == 0 {
			candidate = next
			count = 1
		}
	}
	return candidate
}
