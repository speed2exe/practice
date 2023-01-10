// https://leetcode.com/problems/roman-to-integer/
package main

var symbolValues = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	result := 0

	prev := byte(' ')
	for _, c := range []byte(s) {
		switch c {
		case 'V', 'X':
			result += symbolValues[c]
			if prev == 'I' {
				result -= 2
			}
		case 'L', 'C':
			result += symbolValues[c]
			if prev == 'X' {
				result -= 20
			}
		case 'D', 'M':
			result += symbolValues[c]
			if prev == 'C' {
				result -= 200
			}
		default:
			result += symbolValues[c]
		}
		prev = c
	}

	return result

}
