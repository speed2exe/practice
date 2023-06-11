// https://leetcode.com/problems/3sum/
package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := len(nums) - 1
		for k > j {
			sum := nums[i] + nums[j] + nums[k]
			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				for k > j && nums[j-1] == nums[j] {
					j++
				}
			}
		}
	}
	return result
}
