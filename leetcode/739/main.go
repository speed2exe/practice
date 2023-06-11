// https://leetcode.com/problems/daily-temperatures/
package main

import (
	"fmt"
	"slices"
)

// temp:   [73, 74, 75, 71, 69, 72, 76, 73]
//
// result: [ 0,  0,  0,  0,  0,  0,  0,  0]
// stack : []
//
// result: [ 0,  0,  0,  0,  0,  0,  0,  0]
// stack : [(73, 1)]
//
// result: [ 1,  0,  0,  0,  0,  0,  0,  0]
// stack : [(73, 0), (74, 1)] => [(73, 0), (74, 1)]
// why: 74 is greater than 73, pop the stack
// 1-0 = 1, therefore days is 1
// index is 0 (73)
func dailyTemperatures(temperatures []int) []int {
	stack := make([]Item, 0, len(temperatures)/2)
	result := make([]int, len(temperatures))

	resolve := func(i int, temp int) {
		for {
			if len(stack) == 0 {
				return
			}
			top := stack[len(stack)-1]
			if temp <= top.temp {
				return
			}
			result[top.idx] = i - top.idx
			stack = stack[:len(stack)-1]
		}
	}

	for i, temp := range temperatures {
		resolve(i, temp)
		stack = append(stack, Item{temp: temp, idx: i})
	}

	return result
}

type Item struct {
	temp int
	idx  int
}

func main() {
	tests := []struct {
		temperatures []int
		expected     []int
	}{
		{
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			expected:     []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			temperatures: []int{30, 40, 50, 60},
			expected:     []int{1, 1, 1, 0},
		},
		{
			temperatures: []int{30, 60, 90},
			expected:     []int{1, 1, 0},
		},
	}
	for _, test := range tests {
		actual := dailyTemperatures(test.temperatures)
		if !slices.Equal(actual, test.expected) {
			fmt.Println("[WRONG]", "test:", test, ", actual:", actual)
		}
	}
}
