// https://leetcode.com/problems/longest-common-prefix/
package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	result := []byte{}
	firstStr := strs[0]
	restStrs := strs[1:]
	for i, c := range []byte(firstStr) {
		for _, str := range restStrs {
			if len(str) == i {
				return string(result)
			}
			if str[i] != c {
				return string(result)
			}
		}
		result = append(result, c)
	}

	return string(result)
}
