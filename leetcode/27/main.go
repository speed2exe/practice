// https://leetcode.com/problems/remove-element/?envType=study-plan-v2&envId=top-interview-150
package main

func removeElement(nums []int, val int) int {
	i := 0
	for i < len(nums) {
		if nums[i] != val {
			i++
			continue
		}

		nums[i] = nums[len(nums)-1]
		nums = nums[:len(nums)-1]
	}
	return i
}
