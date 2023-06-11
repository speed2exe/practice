package main

import "fmt"

func rearrangeArray(nums []int) []int {
	posPile := make([]int, 0, len(nums)/2)
	negPile := make([]int, 0, len(nums)/2)

	for _, num := range nums {
		if num > 0 {
			posPile = append(posPile, num)
		} else {
			negPile = append(negPile, num)
		}
	}
	fmt.Println(posPile)
	fmt.Println(negPile)

	var lenNums int = len(nums)
	nums = nums[:0]
	for i := 0; i < lenNums/2; i++ {
		nums = append(nums, posPile[i])
		nums = append(nums, negPile[i])
	}

	fmt.Println(nums)
	return nums
}

func main() {
	rearrangeArray([]int{3, 1, -2, -5, 2, -4})
}
