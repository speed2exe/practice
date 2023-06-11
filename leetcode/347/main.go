// https://leetcode.com/problems/top-k-frequent-elements/
package main

import "fmt"

// nums: [1,1,1,2,2,3], k = 2
// ans:  [1,2], reason: 1 and 2 appear the top 2 most times
//
// Plan:
// collect into map: nums => {1:3, 2:2, 3:1}
//
// frequency: [0,0,0,0,0,0,0] => idx: frequency, value: num that appears frequency times
func topKFrequent(nums []int, k int) []int {
	freqMap := map[int]int{}
	for _, num := range nums { // populate the frequency map
		prev, ok := freqMap[num]
		if ok {
			freqMap[num] = prev + 1
		} else {
			freqMap[num] = 1
		}
	}

	freqTab := make([][]int, len(nums)+1) // index represent the freq, value is the num that happens that many times
	for k, v := range freqMap {
		freqTab[v] = append(freqTab[v], k)
	}

	// fmt.Println("freqTab:", freqTab)

	ans := []int{}
	i := len(nums)
	for {

		// fmt.Println("i:", i)
		if len(ans) >= k {
			break
		}
		val := freqTab[i]
		ans = append(ans, val...)
		i--
	}
	return ans
}

func main() {
	// fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	// fmt.Println(topKFrequent([]int{-1, -1}, 1))
	fmt.Println(topKFrequent([]int{1, 2}, 2))
}
