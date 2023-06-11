// https://leetcode.com/problems/two-sum/
package main

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		panic("unreachable")
	}

	indexByDiff := make(map[int]int)
	firstDiff := target - nums[0]
	indexByDiff[firstDiff] = 0
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		j, ok := indexByDiff[num]
		if ok {
			return []int{j, i}
		}
		diff := target - num
		indexByDiff[diff] = i
	}

	panic("no solution")
}
