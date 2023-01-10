// https://leetcode.com/problems/merge-sorted-array/
package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	ptr1 := m - 1
	ptr2 := n - 1
	insertPtr := len(nums1) - 1

	for {
		if ptr1 < 0 {
			copy(nums1[:ptr2+1], nums2[:ptr2+1])
			return
		}
		if ptr2 < 0 {
			return
		}

		num1 := nums1[ptr1]
		num2 := nums2[ptr2]

		if num1 > num2 {
			nums1[insertPtr] = num1
			ptr1 -= 1
		} else {
			nums1[insertPtr] = num2
			ptr2 -= 1
		}
		insertPtr -= 1
	}

}

// Case 1
// nums1: [4,5,6,0,0,0]
// insPtr: 5
// ptr1: 2
// nums2: [1,2,3]
// ptr2: 2

// nums1: [4,5,6,4,5,6]
// insPtr: 2
// ptr1: -1
// nums2: [1,2,3]
// ptr2: 2

// Case 2
// nums1: [1,2,3,0,0,0]
// m: 3
// nums2: [2,5,6]
// n: 3

// nums1: [1,2,3,0,0,0]
// insIdx: 5
// ptr1: 5
// nums2: [2,5,6]
// ptr2: 2

// nums1: [1,2,3,0,0,6]
// insIdx: 4
// ptr1: 2
// nums2: [2,5,6]
// ptr2: 1

// ..

// nums1: [1,2,3,4,5,6]
// insIdx: -1
// ptr1: -1
// nums2: [2,5,6]
// ptr2: -1
