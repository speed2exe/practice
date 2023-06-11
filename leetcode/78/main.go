// https://leetcode.com/problems/subsets/description/
package main

import "fmt"

func subsets(nums []int) [][]int {
	res := [][]int{{}}
	res = addSet(res, nums)
	return res
}

func addSet(acc [][]int, nums []int) [][]int {
	if len(nums) == 0 {
		return acc
	}

	for i, num := range nums {
		rest := make([]int, 0, len(nums)-1)

		for j, num2 := range nums {
			if i != j {
				rest = append(rest, num2)
			}
		}

		for _, ac := range acc {
			ref := append(ac, num)
			cp := make([]int, len(ref))
			for i, v := range ref {
				cp[i] = v
			}

			acc = append(acc, cp)
			addSet(acc, rest)
		}
	}

	return acc
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3, 4}))
}
