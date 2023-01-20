// https://leetcode.com/problems/basic-calculator/
package main

import "fmt"

func calculate(s string) int {
	if len(s) == 0 {
		return 0
	}

	result := 0
	sign := 1
	num := 0

	var signStack []int
	signStack = append(signStack, sign)

	for _, c := range s {
		fmt.Printf("BEFORE: c: %s, result: %d, signStack:%v, sign: %d, num: %d\n", string(c), result, signStack, sign, num)
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num = num*10 + int(c-'0')
		case '-', '+':
			// offload result
			result += sign * num

			sign = signStack[len(signStack)-1]
			if c == '-' {
				sign *= -1
			}
			num = 0
		case '(':
			signStack = append(signStack, sign)
		case ')':
			signStack = signStack[:len(signStack)-1]
		}
		fmt.Printf("AFTER: c: %s, result: %d, signStack:%v, sign: %d, num: %d\n", string(c), result, signStack, sign, num)
	}

	result += sign * num
	return result
}

func main() {
	fmt.Println(calculate("-(1-2+3)"))
}
