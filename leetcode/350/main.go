// https://leetcode.com/problems/intersection-of-two-arrays-ii/
package main

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}

	var m [1001]int
	for _, num := range nums1 {
		m[num]++
	}

	insertPos := 0
	for _, num := range nums2 {
		val := m[num]
		if val > 0 {
			nums1[insertPos] = num
			insertPos++
			m[num] = val - 1
		}
	}

	return nums1[:insertPos]
}
