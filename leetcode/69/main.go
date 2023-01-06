// https://leetcode.com/problems/sqrtx/
package main

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	left := 1
	right := x

	for {
		if left+1 >= right {
			return left
		}

		mid := (left + right) / 2
		midSquared := mid * mid
		if midSquared == x {
			return mid
		}

		if midSquared < x {
			left = mid
		} else {
			right = mid
		}
	}
}
