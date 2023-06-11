// https://leetcode.com/problems/remove-k-digits/description/
package main

import "fmt"

// Example:
// num: "56471", 2 -> "541"
// num: "199", 1 -> "19"
// num: "991", 1 -> "91"
func removeKdigits(num string, target int) string {
	if len(num) == 0 {
		return "0"
	}
	if target >= len(num) {
		return "0"
	}

	stack := make([]byte, 0, len(num)/2)
	stack = append(stack, num[0]) // stack: [5]
	num = num[1:]                 // num: [6471]

	for {
		if len(num) == 0 {
			// E.g.
			// stack: [123] / [991]
			// target: 2    / 1
			// num: []      / []
			// -> return [1]
			return trimZero(string(stack[:len(stack)-target]))
		}
		if target == 0 {
			// E.g.
			// stack: [654]
			// target: 0
			// num: [21]
			// -> return [65421]
			stack = append(stack, num...)
			return trimZero(string(stack))
		}

		// pop next
		popped := num[0]
		num = num[1:]

		// target: 1
		// stack: [123]
		// num: []
		// popped: 0
		// peeked: 3
		peeked := stack[len(stack)-1]
		if peeked <= popped {
			stack = append(stack, popped)
		} else if popped < peeked {
			stack[len(stack)-1] = popped
			target--

			for target > 0 && len(stack) > 1 && stack[len(stack)-1] < stack[len(stack)-2] {
				target--
				stack[len(stack)-2] = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		}
	}
}

func trimZero(s string) string {
	for i, c := range s {
		if c != '0' {
			return s[i:]
		}
	}
	return "0"
}

// num: "56471", 2 -> "541"
// num: "199", 1 -> "19"
// num: "991", 1 -> "91"
func main() {
	tests := []struct {
		nums     string
		target   int
		expected string
	}{
		{
			nums:     "56471",
			target:   2,
			expected: "541",
		},
		{
			nums:     "199",
			target:   1,
			expected: "19",
		},
		{
			nums:     "991",
			target:   1,
			expected: "91",
		},
		{
			nums:     "1432219",
			target:   3,
			expected: "1219",
		},
		{
			nums:     "1230",
			target:   3,
			expected: "0",
		},
	}
	for _, test := range tests {
		actual := removeKdigits(test.nums, test.target)
		if actual != test.expected {
			fmt.Println("[WRONG]", "test:", test, ", actual:", actual)
		}
	}
}
