package practice

// worst case complexity:
// let n = len(nums)
// lg(n)
func MergeSort(nums []int) []int { // nums: [1, 2, 3, 4, 5, 6, 7, 8] | [1, 2, 3, 4]
	if len(nums) < 2 {
		return nums
	}

	mid_i := len(nums) / 2
	left, right := nums[:mid_i], nums[mid_i:]

	left = MergeSort(left)  // left = [1, 2, 3, 4] | [1, 2]
	right = MergeSort(right) // right = [5, 6, 7, 8] | [3, 4]
	return merge(left, right)
}

// worst case complexity: len(left) + len(right)
// merge merges 2 sorted array of numbers
func merge(left, right []int) []int {
	res := make([]int, 0, len(left)+len(right))

	for {
		// if any of the slice is empty, just append everything else to
		// result and return
		if len(left) == 0 {
			return append(res, right...)
		}
		if len(right) == 0 {
			return append(res, left...)
		}

		// append first elem of nums1 if it is smaller or equal to
		// first elem of num2
		if left[0] < right[0] {
			res = append(res, left[0])
			left = left[1:]
			continue
		}

		// append first elem of nums2 otherwise
		res = append(res, right[0])
		right = right[1:]
	}
}

// time complexity:

// space:
