// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/?envType=study-plan-v2&id=top-interview-150
package main

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	insertIndex := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] == nums[insertIndex-2] {
			continue
		}
		nums[insertIndex] = nums[i]
		insertIndex++
	}
	return insertIndex
}
