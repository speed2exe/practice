// https://leetcode.com/problems/excel-sheet-column-number/discussion/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(titleToNumber("BA"))
}

func titleToNumber(s string) int {
	result := 0

	for _, c := range []byte(s) {
		result *= 26
		result += int(c - 'A' + 1)
	}

	return result
}
