// https://leetcode.com/problems/longest-substring-without-repeating-characters/
package main

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	charSet := map[byte]struct{}{s[0]: {}}
	max := 1
	left := 0
	right := 1
	for right < len(s) {
		nextChar := s[right]
		for {
			_, exists := charSet[nextChar]
			if !exists {
				break
			}
			delete(charSet, s[left])
			left++
		}

		charSet[nextChar] = struct{}{}
		if currentMax := len(charSet); currentMax > max {
			max = currentMax
		}
		right++
	}

	return max
}
