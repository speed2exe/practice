package practice

import "fmt"

/*
A K-palidrome is a string which transform into a palindrome on
removing at most k characters.

Given a string S, and an integer K, print "YES" if S is a K-palidrome;
otherwise print "NO"

Constraints: S has at most 20,000 characters
*/

//Method
// 1. Get a reverse of the string
// 2. Iterate both reverse version and forward version
//      - Find if characters matches

// printIfKPanlindrome answers the actual question
func printIfKPanlindrome(s string, k int) {
	if KPalindrome([]rune(s), k) {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}

func reversedRunes(main []rune) []rune {
	l := len(main)
	res := make([]rune, l)

	for i := range main {
		res[l-i-1] = main[i]
	}
	return res
}

func KPalindrome(main []rune, k int) bool {
	reversed := reversedRunes(main)
	mid := len(main) / 2
	return kPalindrome(main[:mid], reversed[:mid], k)
}

func kPalindrome(main, reversed []rune, k int) bool {
	for i := range main {
		if main[i] != reversed[i] {
			if k == 0 {
				return false
			}
			if len(main) == cap(main) {
				return k >= len(main)
			}
			if len(reversed) == cap(reversed) {
				return k >= len(reversed)
			}

			return kPalindrome(main[i+1:len(main)+1], reversed[i:], k-1) ||
				kPalindrome(main[i:], reversed[i+1:len(reversed)+1], k-1)
		}
	}

	return true
}
