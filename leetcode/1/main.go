// https://leetcode.com/problems/two-sum/
package main

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		panic("unreachable")
	}

	indexByDiff := make(map[int]int)
	firstDiff := target - nums[0]
	indexByDiff[firstDiff] = 0
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		j, ok := indexByDiff[num]
		if ok {
			return []int{j, i}
		}
		diff := target - num
		indexByDiff[diff] = i
	}

	panic("no solution")
}

// num: "56471", 2 -> "541"
// num: "199", 1 -> "19"
// num: "991", 1 -> "91"
// func main() {
// 	tests := []struct {
// 		nums     string
// 		target   int
// 		expected string
// 	}{
// 		{
// 			nums:     "56471",
// 			target:   2,
// 			expected: "541",
// 		},
// 		{
// 			nums:     "199",
// 			target:   1,
// 			expected: "19",
// 		},
// 		{
// 			nums:     "991",
// 			target:   1,
// 			expected: "91",
// 		},
// 	}
// 	for _, test := range tests {
// 		actual := removeKdigits(test.nums, test.target)
// 		if actual != test.expected {
//       fmt.Println("[WRONG]", "test:", test, ", actual:", actual)
// 		}
// 	}
// }
