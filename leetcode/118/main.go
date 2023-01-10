// https://leetcode.com/problems/pascals-triangle/
package main

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	result := make([][]int, numRows)
	result[0] = []int{1}
	for i := 1; i < numRows; i++ {
		result[i] = generateNext(result[i-1])
	}
	return result
}

func generateNext(prev []int) []int {
	result := make([]int, len(prev)+1)

	result[0] = prev[0]
	result[len(result)-1] = prev[len(prev)-1]

	for i := 1; i < len(result)-1; i++ {
		result[i] = prev[i-1] + prev[i]
	}

	return result
}
