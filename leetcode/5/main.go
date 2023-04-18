// https://leetcode.com/problems/longest-palindromic-substring/description/
package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	// s: "baba"
	//     0123
	// expanded: "|b|a|b|a|"
	//            012345678
	//if maxIdx = 3, maxLen = 5
	//then the longest palindrome is "bab"

	expanded := make([]byte, len(s)*2+1)
	for i := 0; i < len(s); i++ {
		expanded[i*2+1] = s[i]
		expanded[i*2] = '|'
	}
	expanded[len(expanded)-1] = '|'

	// fmt.Printf("expanded len: %d\n", len(expanded))

	solution := make([]int, len(expanded))
	solution[0] = 1
	solution[1] = 3

	nextMid := 2

	for {
		// fmt.Printf("\n\ncurrent solution: %v\n", solution)
		// fmt.Printf("\n")
		// fmt.Printf("nextMid: %d\n", nextMid)

		maxLen := maxPalindromeLen(nextMid, expanded)

		// fmt.Printf("maxLen: %d\n", maxLen)
		// fmt.Printf("expanded: %v\n", string(expanded))
		// fmt.Printf("expanded len: %d\n", len(expanded))

		// fmt.Printf("nextMid+maxLen/2: %d\n", nextMid+maxLen/2)
		if nextMid+maxLen/2 == len(expanded)-1 {
			solution[nextMid] = maxLen
			break
		}

		solution[nextMid] = maxLen

		targetPrefixIdx := nextMid - maxLen/2
		// fmt.Printf("targetPrefixIdx: %d\n", targetPrefixIdx)

		// traverse back
		var i int
		for i = 1; i < maxLen/2; i++ {
			prevMid := nextMid - i
			prevLen := solution[prevMid]

			gotPrefix := prevMid - prevLen/2

			if gotPrefix == targetPrefixIdx {
				break
			}

			if gotPrefix < targetPrefixIdx {
				solution[nextMid+i] = prevLen - 2*(targetPrefixIdx-gotPrefix)
				continue
			}

			if gotPrefix > targetPrefixIdx {
				solution[nextMid+i] = prevLen
				continue
			}
		}

		nextMid = nextMid + i
	}

	maxLen := 0
	maxIdx := 0
	for i := 0; i < len(solution); i++ {
		if solution[i] > maxLen {
			maxLen = solution[i]
			maxIdx = i
		}
	}
	// fmt.Printf("solution: %v\n", solution)
	// fmt.Printf("maxLen: %d, maxIdx: %d\n", maxLen, maxIdx)
	start := maxIdx - maxLen/2
	end := maxIdx + maxLen/2
	// fmt.Printf("start: %d, end: %d\n", start, end)
	realStart := start / 2
	realEnd := end / 2
	// fmt.Printf("realStart: %d, realEnd: %d\n", realStart, realEnd)
	return s[realStart:realEnd]
}

func maxPalindromeLen(i int, bs []byte) int {
	result := 1

	next := i + 1
	prev := i - 1
	for {
		if next >= len(bs) || prev < 0 {
			return result
		}
		if bs[next] != bs[prev] {
			return result
		}
		result += 2
		next++
		prev--
	}
}

func main() {
	// input := "abaxabaxabybaxabyb"
	input := "abaxabaxabybaxabyb"
	// input := "abababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababa"
	// input := "ccc"

	fmt.Printf("len(s): %v", len(input))
	fmt.Println(longestPalindrome(input))
}
