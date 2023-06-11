// https://leetcode.com/problems/happy-number/
package main

import "fmt"

func isHappy(n int) bool {
	seen := map[int]struct{}{}

	for {
		if _, ok := seen[n]; ok {
			return false
		} else {
			seen[n] = struct{}{}
		}
		sum := 0
		for n > 0 {
			r := n % 10
			sum += r * r
			n /= 10
		}
		if sum == 1 {
			return true
		}
		n = sum
	}
}

func main() {
	tests := []struct {
		n        int
		expected bool
	}{
		{
			n:        19,
			expected: true,
		},
		{
			n:        2,
			expected: false,
		},
	}
	for _, test := range tests {
		actual := isHappy(test.n)
		if actual != test.expected {
			fmt.Println("[WRONG]", "test:", test, ", actual:", actual)
		}
	}
}
