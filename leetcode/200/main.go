package main

import "fmt"

// Cache: map[{byte,byte}]struct{}
// Plan: interate all the elements of grid
func numIslands(grid [][]byte) int {
	// create cache
	// interate all elements of grid, filter only 1 (land)
	// if visited, skip
	// else, incr number of numIslands
	// do a dfs, mark all adjacent position as visited
	// return the num of island

	cache := make([][]bool, len(grid))
	for i, row := range grid {
		cache[i] = make([]bool, len(row))
	}

	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if markIsland(cache, grid, i, j) {
				count++
			}
		}
	}
	return count
}

func markIsland(cache [][]bool, grid [][]byte, i, j int) bool {
	if grid[i][j] == '0' {
		return false
	}
	if cache[i][j] {
		return false
	}
	cache[i][j] = true

	if i > 0 {
		markIsland(cache, grid, i-1, j)
	}
	if len(grid)-1 > i {
		markIsland(cache, grid, i+1, j)
	}
	if j > 0 {
		markIsland(cache, grid, i, j-1)
	}
	if len(grid[i])-1 > j {
		markIsland(cache, grid, i, j+1)
	}

	return true
}

func main() {
	fmt.Println(
		numIslands([][]byte{
			// {1, 1, 1, 1, 0},
			// {1, 1, 0, 1, 0},
			// {1, 1, 0, 0, 0},
			// {0, 0, 0, 0, 0},

			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		}))
}
