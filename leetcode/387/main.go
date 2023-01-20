// https://leetcode.com/problems/first-unique-character-in-a-string/
package main

func firstUniqChar(s string) int {
	bs := []byte(s)
	var m [26]int // 'a' to 'z'

	for _, c := range bs {
		c -= 'a'
		m[c]++
	}

	for i, c := range bs {
		c -= 'a'
		if m[c] == 1 {
			return i
		}
	}

	return -1
}
