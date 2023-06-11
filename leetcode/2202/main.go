package main

func maximumTop(nums []int, k int) int {
	if k == 0 {
		return nums[0]
	}
	if k == 1 {
		if len(nums) == 1 {
			return -1
		} else {
			return nums[1]
		}
	} else if len(nums) == 1 {
		if k%2 == 1 {
			return -1
		} else {
			return nums[0]
		}
	}
	if k > len(nums) {
		return maxArr(nums)
	}

	cur_max := maxArr(nums[0 : k-1])
	// [a,b,c,d,e], k=3
	// maxArr[a,b,c]
	if k == len(nums) {
		return cur_max
	} else {
		return maxArr([]int{cur_max, nums[k]})
	}

}

func maxArr(nums []int) int {
	var m int
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	return m
}

func main() {
	maximumTop([]int{99, 95, 68, 24, 18}, 69)
}
