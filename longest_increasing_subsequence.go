package practice

//#Given an integer array nums, return the length of the longest strictly
//increasing subsequence.
//
//A subsequence is a sequence that can be derived from an array by deleting some
//or no elements without changing the order of the remaining elements. For
//example, [3,6,2,7] is a subsequence of the array [0,3,1,6,2,2,7].

//Example 1:
//
//Input: nums = [10,9,2,5,3,7,101,18]
//Output: 4
//Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
//Example 2:
//
//Input: nums = [0,1,0,3,2,3]
//Output: 4
//Example 3:
//
//Input: nums = [7,7,7,7,7,7,7]
//Output: 1
//
//
//Constraints:
//
//1 <= nums.length <= 2500
//-104 <= nums[i] <= 104

func LongestIncreasingSubSeq(src []int) []int { // call stack: [10, 9, ..., 18], [10, ..., 101], ...,  [10, 9] ,[10]
	if len(src) < 2 { // return: [2,3,7,18],[2,3,7,101],[2,3,7],[2, 3],[2,5], [2],[9],[10]
		return src
	}

	prevSolution := LongestIncreasingSubSeq(src[:len(src)-1])
	lastNum := prevSolution[len(prevSolution)-1]

	srcLast := src[len(src)-1]

	// if last number of previous solution
	// is lower than the last number of src
	// append the last num to the previous solution and return
	if lastNum < srcLast {
		return append(prevSolution, srcLast)
	}

	// check if we are able to replace the last
	// prevSolution: [1, 2, 8],   srcLast: 7
	// should return [1, 2, 7]

	// check if the prevSolution has only 1 element
	if len(prevSolution) == 1 {
		prevSolution[len(prevSolution)-1] = srcLast
		return prevSolution
	}
	secondLast := prevSolution[len(prevSolution)-2]
	if secondLast < srcLast {
		prevSolution[len(prevSolution)-1] = srcLast
	}

	return prevSolution
}
