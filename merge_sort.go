package practice

func MergeSort(nums []int) []int {
    if nums == nil {
        return nil
    }

    newNums := make([]int, len(nums))
    copy(newNums, nums)
    nums = newNums
    MergeSortNoCopy(nums)
    return nums
}

func MergeSortNoCopy(nums []int) {
	if len(nums) < 2 {
		return
	}

	mid_i := len(nums) / 2
	left, right := nums[:mid_i], nums[mid_i:]
    left_copy := make([]int, len(left))
    copy(left_copy, left)
    right_copy := make([]int, len(right))
    copy(right_copy, right)

	sort(left_copy)
	sort(right_copy)
    merge(nums, left_copy, right_copy)
	return
}

func sort(nums []int) {
	if len(nums) < 2 {
		return
	}
	if len(nums) > 2 {
        MergeSortNoCopy(nums)
        return
	}

	if nums[0] > nums[1] {
		nums[0], nums[1] = nums[1], nums[0]
	}

	return
}

// merge merges 2 sorted array of numbers
func merge(original, nums1, nums2 []int) {
    original = original[:0]

	for {
		// if any of the slice is empty, just append everything else to
		// result and return
		if len(nums1) == 0 {
			original = append(original, nums2...)
            return

		}
		if len(nums2) == 0 {
			original = append(original, nums1...)
            return
		}

		// append first elem of nums1 if it is smaller or equal to
		// first elem of num2
		if nums1[0] < nums2[0] {
			original = append(original, nums1[0])
			nums1 = nums1[1:]
			continue
		}

		// append first elem of nums2 otherwise
		original = append(original, nums2[0])
		nums2 = nums2[1:]
	}
}

// time complexity:



// space: 
