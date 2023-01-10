// https://leetcode.com/problems/single-number/
package main

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}
