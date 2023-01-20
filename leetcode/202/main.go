// https://leetcode.com/problems/happy-number/
package main

func isHappy(n int) bool {
	digits := getDigits(n, nil)
	cache := make(map[int]bool)
	return isHappyWithDigits(n, digits, cache)
}

func isHappyWithDigits(n int, digits []byte, cache map[int]bool) bool {
	if v, ok := cache[n]; ok {
		return v
	}
	sum := 0
	for _, digit := range digits {
		sum += int(digit) * int(digit)
	}
	if sum == 1 {
		return true
	}
	cache[n] = false
	if isHappyWithDigits(sum, getDigits(sum, digits[:0]), cache) {
		cache[n] = true
		return true
	}
	return false
}

func getDigits(n int, digits []byte) []byte {
	count := 0
	for n > 0 {
		count++
		digits = append(digits, byte(n%10))
		n = n / 10
	}
	return digits
}
