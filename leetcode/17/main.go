// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
package main

import "fmt"

var m = [8][]byte{
	{'a', 'b', 'c'},
	{'d', 'e', 'f'},
	{'g', 'h', 'i'},
	{'j', 'k', 'l'},
	{'m', 'n', 'o'},
	{'p', 'q', 'r', 's'},
	{'t', 'u', 'v'},
	{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	retN := 1 // worst case combination
	for range digits {
		retN *= 4
	}

	ret := make([]string, 0, retN)
	backtrack(digits, &ret)
	return ret
}

func backtrack(digits string, ret *[]string) {
	if len(digits) == 0 {
		return
	} else if len(digits) == 1 {
		for _, c := range m[digits[0]-'2'] {
			*ret = append(*ret, string(c))
		}
		return
	}

	backtrack(digits[1:], ret)
	n := len(*ret)
	for i := 0; i < n; i++ {
		for _, c := range m[digits[0]-'2'] {
			*ret = append(*ret, (*ret)[i]+string(c))
		}
	}

	// remove all previous
	for i := 0; i < n; i++ {
		j := len(*ret) - i - 1
		(*ret)[i] = (*ret)[j]
	}
	*ret = (*ret)[:len(*ret)-n]
}

func main() {
	tests := []struct {
		digits string
	}{
		{
			digits: "23",
		},
	}
	for _, test := range tests {
		actual := letterCombinations(test.digits)
		fmt.Println(actual)
	}
}
