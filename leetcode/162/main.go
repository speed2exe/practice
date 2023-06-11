// https://leetcode.com/problems/find-peak-element/description/

package main

import "fmt"

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	} else if len(nums) == 2 {
		return biggerIdx(nums[0], nums[1])
	}

	// if left is peak or right is peak
	if nums[0] > nums[1] {
		return 0
	} else if nums[len(nums)-1] > nums[len(nums)-2] {
		return len(nums) - 1
	}

	leftIdx, rightIdx := 1, len(nums)-2

	for {
		midIdx := leftIdx + ((rightIdx - leftIdx) / 2)

		// case where peak is on the right
		// 1,2,3,4,5,4,3,2,1
		// ^               ^ leftIdx, rightIdx
		//     ^ ^      mid
		if nums[midIdx] < nums[midIdx+1] {
			leftIdx = midIdx + 1
		} else if nums[midIdx-1] > nums[midIdx] {
			// case where peak is on the left
			// 1,2,3,4,5,4,3,2,1
			// ^               ^ leftIdx, rightIdx
			//           ^ ^      mid
			rightIdx = midIdx - 1
		} else {
			return midIdx
		}
	}
}

func biggerIdx(a, b int) int {
	if a > b {
		return 0
	}
	return 1
}

func main() {
	fmt.Println(findPeakElement([]int{6, 5, 4, 3, 2, 3, 2}))
}
