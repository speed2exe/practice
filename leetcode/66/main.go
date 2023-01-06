// https://leetcode.com/problems/plus-one/
package main

func plusOne(digits []int) []int {
	extra := plusOneInplace(digits)
	if extra {
		return append([]int{1}, digits...)
	}
	return digits
}

func plusOneInplace(digits []int) bool {
	if len(digits) == 0 {
		return false
	}
	lastIdx := len(digits) - 1
	last := digits[lastIdx]
	if last != 9 {
		digits[lastIdx] = last + 1
		return false
	}

	digits[lastIdx] = 0
	if len(digits) == 1 {
		return true
	}
	return plusOneInplace(digits[:lastIdx])
}
