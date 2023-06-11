package main

import (
	"strconv"
)

func numDecodings(s string) int {
	dp := make(map[int]int)
	dp[len(s)] = 1 // if we have empty string, there's only 1 way to decode(nothing)

	var dfs func(i int) int
	dfs = func(i int) int {
		if ans, ok := dp[i]; ok {
			return ans
		}
		if s[i] == '0' {
			return 0
		}
		res := dfs(i + 1) //case where decode first digit as char

		// check condition for 2nd case
		if (i + 1) < len(s) {
			doubleDigStr := s[i : i+2]
			num, err := strconv.Atoi(doubleDigStr)
			if err != nil {
				panic(err)
			}
			if num >= 10 && num <= 26 {
				res += dfs(i + 2)
			}
		}

		dp[i] = res
		return res

	}
	return dfs(0)
}

func main() {
	println(numDecodings("12"))
}
