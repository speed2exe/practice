// https://leetcode.com/problems/subarray-sum-equals-k/description/
package main

import "fmt"

func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] == k {
			return 1
		} else {
			return 0
		}
	}

	m := map[int]int{0: 1}

	// map[prefixSum]count = {0:1}

	// [1, -1, 1, 1], k = 3
	//  ^             {0:1, 1:1}
	//      ^         {0:2, 1:1}
	//         ^      {0:2, 1:2}
	//            ^   {0:2, 1:2, 3:3} , k - current sum is 0,  there exist 0 in the map, which value is 2, therefore res += 2

	count := 0
	curSum := 0
	for i, v := range nums {
		println("iter", i)
		curSum += v
		println("curSum:", curSum)
		{ // increment res
			diff := curSum - k
			curP, ok := m[diff]
			if ok {
				count += curP
			}
		}
		println("count:", count)

		// fmt.Println("beforem:", m)
		{ // update m
			curP, ok := m[curSum]
			if ok {
				m[curSum] = curP + 1
			} else {
				m[curSum] = 1
			}
		}
		// fmt.Println("m:", m)
		// fmt.Println()
	}

	return count
}

func main() {
	fmt.Println(subarraySum([]int{1, -1, 0}, 0))
}
