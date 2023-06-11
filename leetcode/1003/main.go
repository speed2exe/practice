// https://leetcode.com/problems/check-if-word-is-valid-after-substitutions/
package main

import "fmt"

// target: abc
// s = "ab abc c"
// stack: ""
// push into the stack until non match
// s = "abc c"
// stack = "ab"

// prefix of s contains abc is complete, skip
// s = "c"
// stack = "ab"

// if it does not match, push into stack
// s = ""
// stack = "abc"

// if stack ends with "abc", pop it
// s = ""
// stack = ""

// check that stack is empty
var target = "abc"

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%len(target) != 0 {
		return false
	}

	stack := []byte{}
	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])
		// // fmt.Println(string(stack))

		for {
			if bytesHasSuffix(stack, target) {
				stack = stack[:len(stack)-len(target)]
				// // fmt.Println("after trim: ", string(stack))

				continue
			}
			break
		}
	}
  // fmt.Println("stack is: ",stack)
	return len(stack) == 0
}

func bytesHasSuffix(bs []byte, suffix string) bool {
	// fmt.Println("bytesHasSuffix", string(bs), suffix)
	if len(suffix) > len(bs) {
		// fmt.Println("false")
		return false
	}
	for i := 0; i < len(suffix); i++ {
		sufEnd := suffix[len(suffix)-1-i]
		sufBs := bs[len(bs)-1-i]
		if sufBs != sufEnd {
			// fmt.Println("false")
			return false
		}
	}
	// fmt.Println("true")
	return true
}

func main() {
	res := isValid("abccba")
	// "aabcbabcc"
	// "a abc b abc c"
	fmt.Println(res)
}
