// https://leetcode.com/problems/minimum-operations-to-make-all-array-elements-equal/
package main

import (
	"fmt"
	"sort"
)

// Example: nums: [3,5,1,7], queries: [4]
// plan: sort nums -> nums: [1,3,5,7]    O(nlog(n)), where n is len(nums)
//
// query: 4, sorted: [1,3,5,7] => mid (binary search) , O(m(log(n)))  where m is len(queries)
//
//	^
//
// sum up from start to mid, (num of elem * query) - sum (left)
// sum up from end to mid, sum(left) - (num of elem * query)
func minOperations(nums []int, queries []int) []int64 {
	sort.Ints(nums)

	res := make([]int64, len(queries))
	for i, q := range queries {
		res[i] = calculateOps(nums, q)
	}
	return res
}

// Example: nums: [1,3,5,7], queries: [4]
// s := 1
func calculateOps(nums []int, target int) int64 {
	s := binSearch(nums, target)
	// fmt.Printf("s: %v\n", s)
	res := 0
	{
		leftSum := 0
		for i := 0; i < s; i++ {
      leftSum += nums[i] // TODO: optimize
		}
		res += (target * s) - leftSum
		// fmt.Printf("leftSum: %v, res: %v\n", leftSum, res)
	}
	{
		rightSum := 0
		for i := s; i < len(nums); i++ {
      rightSum += nums[i] // TODO: optimize
		}
		res += rightSum - (target * (len(nums) - s))
		// fmt.Printf("rightSum: %v, res: %v\n", rightSum, res)
	}
	return int64(res)
}

// returns index such that all elems in nums[0:index] < target
// Example: nums: [1,3,5,7], target: 4
// returns: 2
func binSearch(nums []int, target int) int {
	// fmt.Printf("nums: %v\n", nums)

	if len(nums) == 0 {
		panic("len(nums) == 0")
	}

	l, r := 0, len(nums)-1
	// fmt.Printf("l: %v, r: %v\n", l, r)

	if target <= nums[l] {
		return l
	}
	if nums[r] <= target {
		return len(nums)
	}

	for l <= r {
		// fmt.Printf("loop: l: %v, r: %v\n", l, r)
		m := l + ((r - l) / 2)
		// fmt.Printf("loop: m: %v\n", m)
		if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	// fmt.Printf("l: %v\n", l)
	return l
}

func main() {
	nums := []int{7, 9, 17, 21, 26, 26, 28, 38, 41, 41, 47, 50, 51, 57, 58, 63, 66, 69, 72, 79, 83, 84, 87, 92, 97, 100}
	queries := []int{3}
	// nums := []int{2, 3, 6, 9}
	// queries := []int{10}
	fmt.Println(minOperations(nums, queries))
}
