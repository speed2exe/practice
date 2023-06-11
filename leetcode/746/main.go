// https://leetcode.com/problems/min-cost-climbing-stairs/description/
package main

import "fmt"

// cost:  [1,9,1,1,9,1,1,9,1]
//
// cache: [0,0,1,2,3,3,4,5,5],6
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n == 0 {
		return 0
	} else if n == 1 {
		return cost[0]
	}

	cache := make([]int, len(cost))

	for i := range cache[2:] {
		i := i + 2
		oneBack := cache[i-1]
		twoBack := cache[i-2]
		oneCost := oneBack + cost[i-1]
		twoCost := twoBack + cost[i-2]
		// fmt.Println("i:", i, "oneBack:", oneBack, "twoBack:", twoBack, "oneCost:", oneCost, "twoCost:", twoCost)
		if oneCost < twoCost {
			cache[i] = oneCost
		} else {
			cache[i] = twoCost
		}
		// fmt.Println("cache:", cache)
		// fmt.Println()
	}

	// fmt.Println("cache:", cache)

	return min(
		cost[len(cost)-1]+cache[len(cache)-1],
		cost[len(cost)-2]+cache[len(cache)-2],
	)
}

func main() {
	tests := []struct {
		cost     []int
		expected int
	}{
		{
			cost:     []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			expected: 6,
		},
	}
	for _, test := range tests {
		actual := minCostClimbingStairs(test.cost)
		if actual != test.expected {
			fmt.Println("[WRONG]", "test:", test, ", actual:", actual)
		}
	}
}
