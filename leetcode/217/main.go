// https://leetcode.com/problems/contains-duplicate/
package main

func containsDuplicate(nums []int) bool {
	numSet := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		_, ok := numSet[num]
		if ok {
			return true
		}
		numSet[num] = struct{}{}
	}

	return false
}
