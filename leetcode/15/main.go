// https://leetcode.com/problems/3sum/
package main

import (
	"fmt"
)

func threeSum(nums []int) [][]int {
	if len(nums) <= 2 {
		return [][]int{}
	}
	quickSort(nums)

	var result [][]int
	i := 0

top:
	for {
		if nums[i] > 0 {
			break
		}

		lo, hi := i+1, len(nums)-1

	outer:
		for hi > lo {
			hiVal := nums[hi]
			loVal := nums[lo]
			iVal := nums[i]
			sum := hiVal + loVal + iVal
			if sum == 0 {
				result = append(result, []int{iVal, hiVal, loVal})
				lo++
				for {
					if lo >= hi {
						break outer
					} else if nums[lo] == loVal {
						lo++
					} else {
						break
					}
				}
				hi--
				for {
					if lo >= hi {
						break outer
					} else if nums[hi] == hiVal {
						hi--
					} else {
						break
					}
				}
			} else if sum > 0 {
				hi--
				for {
					if lo >= hi {
						break outer
					} else if nums[hi] == hiVal {
						hi--
					} else {
						break
					}
				}
			} else {
				lo++
				for {
					if lo >= hi {
						break outer
					} else if nums[lo] == loVal {
						lo++
					} else {
						break
					}
				}
			}
		}

		i++
		for {
			if i == len(nums)-2 {
				break top
			}
			if nums[i] == nums[i-1] {
				i++
			} else {
				break
			}
		}
	}

	return result
}

// Incorrect solution
// func deduplicate(nums []int) []int {
// 	ins := 0
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] == nums[ins] {
// 			continue
// 		}
// 		ins++
// 		nums[ins] = nums[i]
// 	}
// 	return nums[:ins+1]
// }

// [1 2 3 5 9 | 7 4 6 8]
//        ^
//              ^

//	[ 2 3 4 7 6 1 0]
//
// ins:        ^
// piv:  ^
func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	piv := len(nums) / 2 // or random index within range of num
	pivVal := nums[piv]

	last := len(nums) - 1
	nums[piv], nums[last] = nums[last], nums[piv]

	ins := 0
	for i := 0; i < last; i++ {
		if nums[i] < pivVal {
			nums[i], nums[ins] = nums[ins], nums[i]
			ins++
		}
	}

	nums[last], nums[ins] = nums[ins], nums[last]

	quickSort(nums[0:ins])
	quickSort(nums[ins+1:])
}

func main() {
	// test quicksort
	tests := []struct {
		nums []int
	}{
		{
			nums: []int{-1, 0, 1, 2, -1, -4},
		},
	}
	for _, test := range tests {
		fmt.Println(threeSum(test.nums))
	}
}
