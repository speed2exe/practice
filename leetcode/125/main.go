// https://leetcode.com/problems/valid-palindrome/
package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	left := -1
	right := len(s)
	for {
		nextLeftIdx, leftValue := findNextValid(s, left)
		nextRightIdx, rightValue := findPrevValid(s, right)

		if nextLeftIdx >= nextRightIdx {
			return true
		}
		if leftValue != rightValue {
			return false
		}

		left = nextLeftIdx
		right = nextRightIdx
	}
}

func findPrevValid(s string, i int) (int, rune) {
	i--
	for i > 0 {
		r := rune(s[i])
		if isValid(r) {
			return i, unicode.ToLower(r)
		}
		i--
	}
	return i, 0
}

func findNextValid(s string, i int) (int, rune) {
	i++
	for i < len(s) {
		r := rune(s[i])
		if isValid(r) {
			return i, unicode.ToLower(r)
		}
		i++
	}
	return i, 0
}

func isValid(r rune) bool {
	return unicode.IsNumber(r) || unicode.IsLetter(r)
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
