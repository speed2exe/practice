package main

import "fmt"

// 1,7,4,6,5,3,2,
//
//	^
//
// 1,7,4,6,5,3,2,
//
//	^
//
// 1,7,4,6,5,3,2,
//
//	^
//
// 1,7,4,6,5,3,2,
//
//	^
//
// 1,7,4,6,5,3,2, -> not increasing
//
//	^
//
// 1,7,4,6,5,3,2, ->  from most right, find the next number larger than 4, which is 5
//
//	^   ^
//
// 1,7,5,6,4,3,2, -> swap 4,5
//
//	^   ^
//
// 1,7,5,(reverse)
//
//	^
func nextPermutation(nums []int) {

	// find the next non-increasing from the right
	var i = len(nums) - 1
	for {
		if i <= 0 {
			reverse(nums)
			return
		}
		var curElem int = nums[i]

		var nextLeftElem int = nums[i-1]
		if nextLeftElem >= curElem {
			i--
			continue
		}

		i--
		break
	}

	var j int = len(nums) - 1
	for {
		if nums[j] > nums[i] {
			break
		}
		j--
	}

	// swap
	nums[i], nums[j] = nums[j], nums[i]

	reverse(nums[i+1:])
}

func reverse(nums []int) {
	var l int = 0
	var r int = len(nums) - 1
	for {
		if l >= r {
			break
		}
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

// 2,3,1
//
//	^
//
// 2,1,3
//
//	^
//
// 1,2,3
// 6
func main() {
	var p = []int{2, 3, 1}

	// [3,1,2]

	nextPermutation(p)
	fmt.Println(p)
}
