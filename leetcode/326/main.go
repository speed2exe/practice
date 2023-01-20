// https://leetcode.com/problems/power-of-three/
package main

import (
	"math"
)

func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}
	power_f := math.Log2(float64(n)) / math.Log2(float64(3))
	power_f = math.RoundToEven(power_f)
	return int(math.Pow(3, power_f)) == n
}
