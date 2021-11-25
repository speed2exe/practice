package practice

func MergeSort(nums []int) []int {
	if nums == nil {
		return nil
	}
	if len(nums) == 0 {
		return []int{}
	}
	if len(nums) <= 1 {
		return []int{nums[0]}
	}

	res := make([]int, len(nums))
	copy(res, nums)

	mid_i := len(res) / 2
	left, right := res[:mid_i], res[mid_i:]

	left = sort(left)
	right = sort(right)
	return merge(left, right)
}

func sort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	if len(nums) > 2 {
		return MergeSort(nums)
	}

	if nums[0] > nums[1] {
		nums[0], nums[1] = nums[1], nums[0]
	}

	return nums
}

// merge merges 2 sorted array of numbers
func merge(nums1 []int, nums2 []int) []int {
	res := make([]int, 0, len(nums1)+len(nums2))

	for {
		// if any of the slice is empty, just append everything else to
		// result and return
		if len(nums1) == 0 {
			return append(res, nums2...)
		}
		if len(nums2) == 0 {
			return append(res, nums1...)
		}

		// append first elem of nums1 if it is smaller or equal to
		// first elem of num2
		if nums1[0] < nums2[0] {
			res = append(res, nums1[0])
			nums1 = nums1[1:]
			continue
		}

		// append first elem of nums2 otherwise
		res = append(res, nums2[0])
		nums2 = nums2[1:]
	}
}
