package main

func sumOfThree(num int64) []int64 {
	if num%3 > 0 {
		return []int64{}
	}
	base := num / 3
	return []int64{base - 1, base, base + 1}
}
