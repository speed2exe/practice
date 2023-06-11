package main

import "fmt"

// Input: s1 = "ab", s2 = "eidbaooo"
// Output: true
// Explanation: s2 contains one permutation of s1 ("ba").
func checkInclusion(small, big string) bool {
	if len(small) > len(big) {
		return false
	}

	smallPerm := countMap(small)
	fmt.Println("countMap of smallPerm", smallPerm)

	var l int
	var r = len(small)
	bigSubStrPerm := countMap(big[l:r])

	for {
		if smallPerm == bigSubStrPerm {
			return true
		}
		{
			var c byte = big[l]
			var cIdx uint8 = c - 'a'
			bigSubStrPerm[cIdx]--
		}
		l++

		if r < len(big) {
			var c byte = big[r]
			var cIdx uint8 = c - 'a'
			bigSubStrPerm[cIdx]++
			r++
		} else {
			return false
		}
	}
}

func countMap(s string) [26]int {
	charCount := [26]int{}

	// fill in the charCount
	for i := range s {
		var c byte = s[i]
		var idx uint8 = c - 'a'
		charCount[idx]++
	}

	return charCount
}

func main() {
	println(checkInclusion("ab", "a"))
}
