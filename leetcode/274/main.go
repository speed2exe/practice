// https://leetcode.com/problems/h-index/description

package main

import "fmt"

// Example: [3, 0, 6, 1, 5] => 3
//
//	^
//
// table    [0  1  0  1  0  2]
// index:    0  1  2  3  4  5
// len(citations) ... 0, keep track of
func hIndex(citations []int) int {
	n := len(citations)
	table := make([]int, n+1)
	for _, c := range citations {
		i := min2(c, n)
		table[i]++
	}

	count := 0
	for i := len(citations); i > 0; i-- {
		count += table[i]
		if count >= i {
			return i
		}
	}
	return 0
}

func min2(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	type Test struct {
		citations []int
		expected  int
	}
	tests := []Test{
		{
			citations: []int{3, 0, 6, 1, 5},
			expected:  3,
		},
		{
			citations: []int{1, 3, 1},
			expected:  1,
		},
	}
	for _, test := range tests {
		actual := hIndex(test.citations)
		if actual != test.expected {
			fmt.Println("[WRONG] Test:", test, ", Your ans:", actual)
		}
	}
}
