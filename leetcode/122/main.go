// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
package main

import "fmt"

// prices: [7, 1, 5, 3, 6, 4]
// output: 7
// prices: [7, 1, 5, 3, 6, 4] (switch if there is change in direction)
//
//	B  S          => profit += 4
//
// prices: [7, 1, 5, 3, 6, 4]
//
//	B  S    => profit += 3
func maxProfit(prices []int) int {
	profit := 0

	curPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		nextPrice := prices[i]
		if nextPrice > curPrice {
			profit += (nextPrice - curPrice)
		}
		curPrice = nextPrice
	}

	return profit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

