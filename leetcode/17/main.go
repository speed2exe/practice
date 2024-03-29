// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
package main

var m = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	res := letterCombinationsBytes([]byte(digits))
	ret := make([]string, len(res))

	for i, r := range res {
		ret[i] = string(r)
	}
	return ret
}

func letterCombinationsBytes(digits []byte) [][]byte {
	if len(digits) == 0 {
		return [][]byte{}
	}
	firstCombi := m[digits[0]]                  // ['a', 'b', 'c']
	rest := letterCombinationsBytes(digits[1:]) // ['an', 'as', ...]
	if len(rest) == 0 {
		result := make([][]byte, len(firstCombi))
		for i, v := range firstCombi {
			result[i] = []byte{v}
		}
		return result
	}

	result := make([][]byte, 0, len(rest)*len(firstCombi))
	for _, v1 := range firstCombi {
		for _, v2 := range rest {
			result = append(result, append([]byte{v1}, v2...))
		}
	}
	return result
}
