// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	currentMin := prices[0]
	maxProfit := 0

	prevPrice := prices[0]
	isIncreasing := false
	for i := 1; i < len(prices); i++ {
		currentPrice := prices[i]
		if currentPrice == prevPrice {
			continue
		}
		currentlyIncreasing := currentPrice > prevPrice

		if !isIncreasing && currentlyIncreasing {
			isIncreasing = true
			if prevPrice < currentMin {
				currentMin = prevPrice
			}
		} else if isIncreasing && !currentlyIncreasing {
			isIncreasing = false
			profit := prevPrice - currentMin
			if profit > maxProfit {
				maxProfit = profit
			}
		}

		prevPrice = currentPrice
	}

	if isIncreasing {
		profit := prevPrice - currentMin
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0, 9}))
}
