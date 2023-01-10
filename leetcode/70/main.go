// https://leetcode.com/problems/climbing-stairs/
package main

func climbStairs(n int) int {
	solution := make([]int, n)
	return climbStairsCached(n, solution)
}

func climbStairsCached(n int, solution []int) int {
	if n <= 2 {
		return n
	}
	ans := solution[n-1]
	if ans > 0 {
		return ans
	}

	result := climbStairsCached(n-1, solution) + climbStairsCached(n-2, solution)
	solution[n-1] = result
	return result
}
