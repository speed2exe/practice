// https://leetcode.com/problems/pacific-atlantic-water-flow/description/
package main

import "fmt"

type Key struct {
	row, col int
}

type Value struct {
	alantic, pacific bool
}

type Void struct{}

var void = Void{}

// 1 1

// 1 7 7
// 1 6 5
// 2 3 4
func pacificAtlantic(heights [][]int) [][]int {
	rows := len(heights)
	cols := len(heights[0])

	// get a set of pos that can reach pacific
	canReachPac := map[Key]Void{}
	{
		for i := 0; i < rows; i++ { // most left side
			canReachPac[Key{i, 0}] = void
		}
		for i := 0; i < cols; i++ { // most top side
			canReachPac[Key{0, i}] = void
		}
		bfs(canReachPac, heights)
	}

	canReachAt := map[Key]Void{}
	{
		for i := 0; i < rows; i++ { // most right side
			canReachAt[Key{i, cols - 1}] = void
		}
		for i := 0; i < cols; i++ { // most down side
			canReachAt[Key{rows - 1, i}] = void
		}
		bfs(canReachAt, heights)
	}

	result := [][]int{}

	var least, most map[Key]Void
	if len(canReachAt) > len(canReachPac) {
		least = canReachPac
		most = canReachAt
	} else {
		most = canReachPac
		least = canReachAt
	}
	for k := range least {
		if _, ok := most[k]; ok {
			result = append(result, []int{k.row, k.col})
		}
	}

	return result
}

func bfs(initialSet map[Key]Void, heights [][]int) {
	rows := len(heights)
	cols := len(heights[0])

	curSet := make([]Key, 0, len(initialSet))
	for k := range initialSet {
		curSet = append(curSet, k)
	}
	nextSet := []Key{}
	for len(curSet) > 0 {
		for _, key := range curSet {
			neighbors := make([]Key, 0, 4)
			if key.row > 0 {
				neighbors = append(neighbors, Key{key.row - 1, key.col})
			}
			if key.col > 0 {
				neighbors = append(neighbors, Key{key.row, key.col - 1})
			}
			if key.row < rows-1 {
				neighbors = append(neighbors, Key{key.row + 1, key.col})
			}
			if key.col < cols-1 {
				neighbors = append(neighbors, Key{key.row, key.col + 1})
			}

			curHeight := heights[key.row][key.col]
			for _, neighbor := range neighbors {
				if _, ok := initialSet[neighbor]; ok {
					continue
				}
				neighborHeight := heights[neighbor.row][neighbor.col]
				if neighborHeight >= curHeight {
					nextSet = append(nextSet, neighbor)
					initialSet[neighbor] = void
				}
			}
		}

		curSet, nextSet = nextSet, curSet
		nextSet = nextSet[:0]
	}
}

func main() {
	tests := []struct {
		heights [][]int
	}{
		// {
		// 	heights: [][]int{
		// 		{1, 2, 2, 3, 5},
		// 		{3, 2, 3, 4, 4},
		// 		{2, 4, 5, 3, 1},
		// 		{6, 7, 1, 4, 5},
		// 		{5, 1, 1, 2, 4},
		// 	}},
		// {
		// 	heights: [][]int{
		// 		{1},
		// 	}},
		{
			heights: [][]int{
				{2, 1},
				{1, 2},
			}},
	}
	for _, test := range tests {
		actual := pacificAtlantic(test.heights)
		fmt.Println("heights:", test.heights, ", actual:", actual)
	}
}
