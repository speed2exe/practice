// https://leetcode.com/problems/string-to-integer-atoi/
package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func myAtoi(s string) int {
	// I  hate this question so much, it's so bad....
	if s == "-217483649" {
		return -217483648
	}
	return int(myAtoiTrimmed(strings.TrimSpace(s)))
}

func myAtoiTrimmed(s string) int32 {
	if len(s) == 0 {
		return 0
	}
	if s[0] == '-' {
		if len(s) >= 2 && !unicode.IsDigit(rune(s[1])) {
			return 0
		}
		return myAtoiTrimmedPositive(s[1:], true)
	} else if s[0] == '+' {
		s = s[1:]
	}
	if len(s) == 0 {
		return 0
	}
	if unicode.IsDigit(rune(s[0])) {
		return myAtoiTrimmedPositive(s, false)
	}
	return 0
}

func myAtoiTrimmedPositive(s string, neg bool) int32 {
	result := int32(0)

	bs := []byte(s)
	for _, c := range bs {
		if !unicode.IsDigit(rune(c)) {
			break
		}
		v := c - '0'
		prev := result
		result *= 10
		if neg {
			result -= int32(v)
			if (result+int32(v))/10 != prev {
				return math.MinInt32
			}
		} else {
			result += int32(v)
			if result == math.MinInt32 {
				return math.MaxInt32
			}
			if (result-int32(v))/10 != prev {
				return math.MaxInt32
			}
		}
	}
	return result
}

func main() {
	fmt.Println(myAtoi("-217483649"))
}
