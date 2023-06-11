// https://leetcode.com/problems/rotting-oranges/description/
package main

import "fmt"

// 2 1 1
// 1 1 0
// 0 1 1
//
// cycle 1
// list: [(0,0)]
// 2 1 1
// 1 1 0
// 0 1 1
//
// list: [(0,1),(1,0)]
// cycle 2
// 2 2 1
// 2 1 0
// 0 1 1
//
// list: [(0,2),(1,1)]
// cycle 3
// 2 2 2
// 2 2 0
// 0 1 1
//
// list: [(2, 1)]
// cycle 4
// 2 2 2
// 2 2 0
// 0 2 1
//
// list: [(2, 2)]
// cycle 5
// 2 2 2
// 2 2 0
// 0 2 2
//
// list: []
// cycle 6
// 2 2 2
// 2 2 0
// 0 2 2
func orangesRotting(grid [][]int) int {
	rowCount := len(grid)
	columnCount := len(grid[0])

	steps := 0

	// count the num of initRottenCount apple
	initRottenCount := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
				initRottenCount++
			}
		}
	}

	// initialize index and populate them with rotten apple
	rottenList := make([]Index, 0, initRottenCount)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
				rottenList = append(rottenList, Index{i, j})
			}
		}
	}

	newRottenList := []Index{}
	for {
		newRottenList = newRottenList[:0]

		// get all four corner of every rotten orange
		// filter only those oranges that can be rotten
		for _, rottenIdx := range rottenList {
			rottenIdx := Index{rottenIdx.Row, rottenIdx.Column}
			pots := [4]Index{
				{rottenIdx.Row + 1, rottenIdx.Column},
				{rottenIdx.Row, rottenIdx.Column + 1},
				{rottenIdx.Row - 1, rottenIdx.Column},
				{rottenIdx.Row, rottenIdx.Column - 1},
			}
			for _, pot := range pots {
				if pot.Column >= 0 && pot.Row >= 0 &&
					pot.Column < columnCount && pot.Row < rowCount &&
					grid[pot.Row][pot.Column] == 1 {
					newRottenList = append(newRottenList, pot)
				}
			}
		}

		if len(newRottenList) == 0 {
			break
		}
		steps++

		for _, newRot := range newRottenList {
			grid[newRot.Row][newRot.Column] = 2
		}

		// update rotten list for next loop
		rottenList, newRottenList = newRottenList, rottenList
	}

	// check that there are no fresh oranges
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return steps
}

type Index struct {
	Row    int
	Column int
}

func main() {
	tests := []struct {
		grid     [][]int
		expected int
	}{
		// {
		// 	grid: [][]int{
		// 		{2, 1, 1},
		// 		{1, 1, 0},
		// 		{0, 1, 1},
		// 	},
		// 	expected: 4,
		// },
		// {
		// 	grid: [][]int{
		// 		{2, 1, 1},
		// 		{0, 1, 1},
		// 		{1, 0, 1},
		// 	},
		// 	expected: -1,
		// },
		// {
		// 	grid: [][]int{
		// 		{0, 2},
		// 	},
		// 	expected: 0,
		// },
		{
			grid: [][]int{
				{0, 2},
				{0, 1},
				{0, 1},
				{1, 1},
				{1, 1},
				{1, 1},
			},
			expected: 6,
		},
	}
	for _, test := range tests {
		actual := orangesRotting(test.grid)
		if actual != test.expected {
			fmt.Println("[WRONG]", "test:", test, ", actual:", actual)
		}
	}
}

func prettyPrintGrid(ss [][]int) {
	for _, s := range ss {
		for _, e := range s {
			fmt.Print(e)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

// 0,2
// 0,1
// 0,1
// 1,1
// 1,1
// 1,1
