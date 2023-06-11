// https://leetcode.com/problems/two-sum/
package main

import "fmt"

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	result := make([]byte, 0, 16)

	var isNeg bool
	if num < 0 {
		isNeg = true
		num = -num
	}

	for {
		if num == 0 {
			break
		}

		quo := num / 7
		rem := num % 7
		result = append(result, byte(rem)+'0')
		num = quo
	}
	if isNeg {
		result = append(result, '-')
	}
	reverse(result)
	return string(result)
}

func reverse(input []byte) {
	for i := range input[:len(input)/2] {
		j := len(input) - i - 1
		input[i], input[j] = input[j], input[i]
	}
}

func main() {
	tests := []struct {
		num      int
		expected string
	}{
		{
			num:      100,
			expected: "202",
		},
		{
			num:      7,
			expected: "10",
		},
		{
			num:      -7,
			expected: "-10",
		},
	}
	for _, test := range tests {
		actual := convertToBase7(test.num)
		if actual != test.expected {
			fmt.Println("[WRONG]", "test:", test, ", actual:", actual)
		}
	}
}
