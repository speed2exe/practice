// https://leetcode.com/problems/kth-largest-element-in-an-array/
package main

import "fmt"

// nums: [5,2,3,1,6,7], k: 6 => find 6th largest, returns 6
func findKthLargest(nums []int, k int) int {
	targetPos := len(nums) - k // position that we want assuming sorted

	l, r := 0, len(nums)-1
	for {
		piv := l + ((r - l) / 2)
		pivVal := nums[piv]

		// move piv val to end
		nums[piv], nums[r] = nums[r], nums[piv]

		// partition
		insertIndex := l
		eqlFlag := false
		for i := l; i < r; i++ {
			if nums[i] < pivVal {
				nums[insertIndex], nums[i] = nums[i], nums[insertIndex]
				insertIndex++
			}
			if nums[i] == pivVal { // 50% chance to swap if it's the same
				if eqlFlag {
					eqlFlag = false
				} else {
					nums[insertIndex], nums[i] = nums[i], nums[insertIndex]
					insertIndex++
					eqlFlag = true
				}
			}
		}

		// if found
		if targetPos == insertIndex {
			return pivVal
		}

		nums[insertIndex], nums[r] = nums[r], nums[insertIndex]

		if insertIndex < targetPos {
			l = insertIndex + 1
		} else {
			r = insertIndex - 1
		}
	}
}

func main() {
	fmt.Println(findKthLargest([]int{5, 2, 3, 1, 6, 7, 4}, 1))
	fmt.Println(findKthLargest([]int{5, 2, 3, 1, 6, 7, 4}, 2))
	fmt.Println(findKthLargest([]int{5, 2, 3, 1, 6, 7, 4}, 3))
	fmt.Println(findKthLargest([]int{5, 2, 3, 1, 6, 7, 4}, 4))
	fmt.Println(findKthLargest([]int{5, 2, 3, 1, 6, 7, 4}, 5))
	fmt.Println(findKthLargest([]int{5, 2, 3, 1, 6, 7, 4}, 6))
	fmt.Println(findKthLargest([]int{5, 2, 3, 1, 6, 7, 4}, 7))
}
