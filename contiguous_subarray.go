package practice

// Contiguous Subarrays
//
// You are given an array arr of N integers. For each index i, you are required
// to determine the number of contiguous subarrays that fulfill the following
// conditions:
// - The value at index i must be the maximum element in the contiguous subarrays.
// - These contiguous subarrays must either start from or end on index i.
//
// Signature: int[] countSubarrays(int[] arr)
//
// Input
// Array arr is a non-empty list of unique integers that range between 1 to 1,000,000,000
// Size N is between 1 and 1,000,000
//
// Output
// An array where each index i contains an integer denoting the maximum number of contiguous subarrays of arr[i]
//
// Example:
// arr = [3, 4, 1, 6, 2]
// output = [1, 3, 1, 5, 1]
// Explanation:
// For index 0 - [3] is the only contiguous subarray that starts (or ends) with 3, and the maximum value in this subarray is 3.
// For index 1 - [4], [3, 4], [4, 1]
// For index 2 - [1]
// For index 3 - [6], [6, 2], [1, 6], [4, 1, 6], [3, 4, 1, 6]
// For index 4 - [2]
// So, the answer for the above input is [1, 3, 1, 5, 1]

func countSubarrays(array []int) []int {
	result := make([]int, len(array))

	for i := range array {
		result[i] = countSubarraysForIndex(i, array)
	}

	return result
}

func countSubarraysForIndex(i int, array []int) int {
	return countSubarraysForIndexLeft(array[i], array[:i]) +
		countSubarraysForIndexRight(array[i], array[i:])
}

// Example:
// v = 4
// array = [3]
func countSubarraysForIndexLeft(v int, array []int) int {
	var result int

	// start from the second last index to 0
	i := len(array) - 1
	for i >= 0 {
		last := array[i]
		if last > v {
			break
		}
		result++
		i--
	}

	return result
}

// Example:
// v = 4
// array = [4, 1, 6, 2]
func countSubarraysForIndexRight(v int, array []int) int {
	result := 1
	for result < len(array) {
		if array[result] > v {
			break
		}
		result++
	}

	return result
}
