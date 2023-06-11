package main

import "fmt"

// colors: abbbbaaacd
func minCost(colors string, neededTime []int) int {
	if len(colors) == 0 {
		return 0
	}

	totalCost := 0

	runningSum := neededTime[0]
	maxTime := neededTime[0]
	oldColor := colors[0]

	for i := range colors[1:] {
		i := i + 1

		color := colors[i]
		t := neededTime[i]

		if color == oldColor {
			runningSum += t
			maxTime = max1(t, maxTime)
		} else {
			totalCost += runningSum - maxTime
			runningSum = neededTime[i]
			maxTime = neededTime[i]
		}

		oldColor = color
	}

	totalCost += runningSum - maxTime

	return totalCost
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minCost("aabaa", []int{1, 2, 3, 4, 1}))
}
