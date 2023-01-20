// https://leetcode.com/problems/move-zeroes/
package main

func moveZeroes(nums []int) {
	insertPos := 0
	for _, num := range nums {
		if num != 0 {
			nums[insertPos] = num
			insertPos++
		}
	}
	for insertPos < len(nums) {
		nums[insertPos] = 0
		insertPos++
	}
}
