// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := 1
	last := nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num == last {
			continue
		}
		nums[result] = num
		result += 1
		last = num
	}

	return result
}
