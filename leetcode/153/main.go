// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
package main

import (
	"math"
)

// nums: []int{4,5,6,7,0,1,2}
func findMin(nums []int) int {
	if len(nums) == 0 {
		return math.MaxInt
	} else if len(nums) == 1 {
		return nums[0]
	} else if nums[0] < nums[len(nums)-1] { // rotated 0 times
		return nums[0]
	}

	// 4,5,6,7,   0,1,2,3
	// l                r
	l, r := 0, len(nums)-1

	for {
		if r-l <= 1 {
			return min(nums[l], nums[r])
		}

		// 4,5,6,7,0,1,2
		// l     m     r
		m := l + ((r - l) / 2)

		// are we in the right or left sort section
		inLeftSorted := nums[m] > nums[l]

		if inLeftSorted {
			l = m
		} else {
			// 5,6,7,0,1,2,3,4
			// l       m     r
			// l     r
			r = m
		}

	}

}

func main() {
	println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	//                    0                 6
	//                    l        m        r
	//                                l  m  r
}
