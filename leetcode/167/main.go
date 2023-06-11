// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
package main

func twoSum(numbers []int, target int) []int {
	leftIdx := 0
	rightIdx := len(numbers) - 1

	for {
		left := numbers[leftIdx]
		right := numbers[rightIdx]
		sum := left + right
		if sum == target {
			return []int{leftIdx + 1, rightIdx + 1}
		}
		if sum > target {
			rightIdx--
		} else {
			leftIdx++
		}
	}
}
