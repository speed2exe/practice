// https://leetcode.com/problems/missing-number/
package main

func missingNumber(nums []int) int {
	xor := 0
	for i, num := range nums {
		xor ^= i
		xor ^= num
	}
	return xor ^ len(nums)
}
