// https://leetcode.com/problems/rotate-array/

package main

func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	l := 0
	r := len(nums) - 1
	for r > l {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
