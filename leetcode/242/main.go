// https://leetcode.com/problems/valid-anagram/
package main

const s_len = 'z' + 1

func isAnagram(s, t string) bool {
	var letters [s_len]int
	for _, c := range s {
		letters[c]++
	}
	for _, c := range t {
		letters[c]--
	}
	for _, v := range letters {
		if v != 0 {
			return false
		}
	}
	return true
}
