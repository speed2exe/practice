package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(median([]float32{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Println(median([]float32{1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(median([]float32{1, 2, 3, 4, 4, 5, 6, 7}))
}

func median(arr []float32) float32 {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}
	if len(arr) == 2 {
		return (arr[0] + arr[1]) / 2
	}

	mid := len(arr) / 2
	k := findKthLargest(arr, mid)
	if len(arr)%2 == 1 {
		return k
	}
	k2 := findKthLargest(arr, mid-1)
	return (k + k2) / 2
}

func findKthLargest(arr []float32, k int) float32 {
	if len(arr) == 1 {
		return arr[0]
	}

	p := partition(arr)
	if p == k {
		return arr[p]
	}

	if p < k {
		return findKthLargest(arr[p+1:], k-(p+1))
	}

	return findKthLargest(arr[:p], k)
}

func partition(arr []float32) int {
	rand_idx := rand.Intn(len(arr))
	rand_value := arr[rand_idx]
	running_idx := 0
	for i, v := range arr {
		if v < rand_value {
			arr[i] = arr[running_idx]
			arr[running_idx] = v
			running_idx++
		}
	}
	return running_idx
}
