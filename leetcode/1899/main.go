package main

import "fmt"

// https://leetcode.com/problems/merge-triplets-to-form-target-triplet/
// Input: triplets = [[2,5,3],[1,8,4],[1,7,5]], target = [2,7,5]
// Output: true
// Explanation: Perform the following operations:
// - Choose the first and last triplets [[2,5,3],[1,8,4],[1,7,5]].
// Update the last triplet to be [max(2,1), max(5,7), max(3,5)] = [2,7,5]. triplets = [[2,5,3],[1,8,4],[2,7,5]]
// The target triplet [2,7,5] is now an element of triplets.
func mergeTriplets(tris [][]int, target []int) bool {
	var success0 bool
	var success1 bool
	var success2 bool

	// Iterate through all the triplets
	for _, tri := range tris {

		// filter invalid candidate
		if tri[0] > target[0] ||
			tri[1] > target[1] ||
			tri[2] > target[2] {
			continue
		}

		success0 = success0 || tri[0] == target[0]
		success1 = success1 || tri[1] == target[1]
		success2 = success2 || tri[2] == target[2]

		if success0 &&
			success1 &&
			success2 {
			return true
		}
	}

	return false
}

func main() {
	tris := [][]int{{2, 5, 3}, {1, 8, 4}, {1, 7, 5}}
	target := []int{2, 7, 5}
	fmt.Println(mergeTriplets(tris, target))
}
