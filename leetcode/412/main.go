// https://leetcode.com/problems/fizz-buzz/
package main

import "strconv"

const (
	FIZZ     = "Fizz"
	BUZZ     = "Buzz"
	FIZZBUZZ = "FizzBuzz"
)

func fizzBuzz(n int) []string {
	result := make([]string, 0, n)

	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			result = append(result, FIZZBUZZ)
		case i%5 == 0:
			result = append(result, BUZZ)
		case i%3 == 0:
			result = append(result, FIZZ)
		default:
			result = append(result, strconv.Itoa(i))
		}
	}

	return result
}
