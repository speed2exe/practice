// https://leetcode.com/problems/reverse-integer/
package main

import "fmt"

func reverse(x int) int {
	return int(reverse32(int32(x)))
}

func reverse32(x int32) int32 {
	if x < 0 {
		if x == -x {
			return 0
		}
		return -reverse32(-x)
	}
	var result int32
	for x > 0 {
		lastDigit := x % 10
		x = x / 10

		prev := result

		result *= 10
		result += lastDigit

		if ((result - lastDigit) / 10) != prev {
			return 0
		}
	}
	return result
}

func main() {
	// fmt.Println(reverse32(-1234))
	// fmt.Println(reverse32(1002103))
	// fmt.Println(reverse32(1230))
	fmt.Println(reverse32(-2147483648))
}
