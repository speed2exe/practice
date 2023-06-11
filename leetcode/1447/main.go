package main

import "fmt"

func simplifiedFractions(n int) []string {
	res := []string{}

	for i := 2; i <= n; i++ {
		for j := 1; j <= i-1; j++ {
			numerator := j
			denominator := i
			if gcd(i, j) == 1 {
				res = append(res, fmt.Sprintf("%v/%v", numerator, denominator))
			}
		}
	}

	return res
}

func gcd(greater, smaller int) int {
	if greater < smaller {
		return gcd_proper(smaller, greater)
	}
	return gcd_proper(greater, smaller)
}

func gcd_proper(greater, smaller int) int {
	assert(greater > smaller)
	if smaller == 0 {
		return greater
	}
	return gcd(smaller, greater%smaller)
}

func assert(condition bool) {
	if !condition {
		panic("assertion failed")
	}
}

func main() {
	fmt.Println(simplifiedFractions(3))
	fmt.Println(simplifiedFractions(4))
}
