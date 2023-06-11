// https://leetcode.com/problems/combination-sum/
package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	var ret [][]int
	backtrack(candidates, target, 0, []int{}, &ret)
	return ret
}

func backtrack(candidates []int, target, sum int, curCombi []int, ret *[][]int) {
	if sum == target {
		copiedCombi := append([]int{}, curCombi...)
		*ret = append(*ret, copiedCombi)
		return
	}

	for i := 0; i < len(candidates); i++ {
		if sum+candidates[i] > target {
			continue
		}
		nextCombi := append(curCombi, candidates[i])
		backtrack(candidates[i:], target, sum+candidates[i], nextCombi, ret)
	}
}

func main() {
	tests := []struct {
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			candidates: []int{2, 3},
			target:     5,
			expected:   [][]int{{5}},
		},
		{
			candidates: []int{2, 3, 6, 7},
			target:     7,
			expected:   [][]int{{2, 2, 3}, {7}},
		},
		{
			candidates: []int{2, 7},
			target:     7,
		},
	}
	for _, test := range tests {
		actual := combinationSum(test.candidates, test.target)
		fmt.Println(actual)
	}
}
