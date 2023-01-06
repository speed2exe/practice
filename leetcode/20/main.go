// https://leetcode.com/problems/valid-parentheses/
package main

var matchingParen = map[byte]byte{
	'}': '{',
	']': '[',
	')': '(',
}

func isValid(s string) bool {
	input := []byte(s)

	stack := []byte{}

	for _, v := range input {
		switch v {
		case '(', '{', '[':
			stack = append(stack, v)
		case '}', ']', ')':
			if len(stack) == 0 {
				return false
			}
			want := matchingParen[v]
			got := stack[len(stack)-1]
			if got != want {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
		}
	}

	return len(stack) == 0
}
