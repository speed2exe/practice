// https://leetcode.com/problems/search-in-rotated-sorted-array
package main

import "fmt"

// Example: nums: 4,5,6,7,0,1,2,3, target: 6
// return 2 (index of 6)
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	// nums: 4,5,6,7,0,1,2,3
	//       l             r
	l, r := 0, len(nums)-1

	// need to handle case when nums = []int{1}
	for l <= r {
		m := l + ((r - l) / 2)
		if nums[m] == target {
			return m
		}

		// check if mid is in left or right sorted portion
		// middle value could be the same as left value
		if nums[l] <= nums[m] {
			// in the left portion
			// nums: 4,5,6,7,0,1,2,3
			//       l             r
			//           m
			if target > nums[m] { // target: 7
				l = m + 1
			} else if target < nums[l] { // target: 2
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			// in the right portion
			// nums: 4,5,6,7,0,1,2,3
			//       l             r
			//                 m
			if target > nums[r] { // target: 6
				r = m - 1
			} else if target < nums[m] { // target: 0
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}
