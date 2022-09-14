package main

func nonDivisibleSubset(k int32, s []int32) int32 {
    // count is the result that will be returned
    var count int32

    // remainder stores the number of times a remainder is found
    // for each elem in s
    var remainder = make([]int32, k)
    for _, v := range s {
        remainder[v % k]++
    }

    // if there is a remainder of 0, then there can only exist
    // one of them because a pair of them will always be divisible
    // by k
    if remainder[0] > 0 {
        count++
    }

    // Example
    // remainder: [3, 4, 1, 2, 5, 6], k = 6
    // remainder of 0 occured 3 time(s)
    // remainder of 1 occured 4 time(s)
    // remainder of 2 occured 1 time(s)
    // remainder of 3 occured 2 time(s)
    // remainder of 4 occured 5 time(s)
    // remainder of 5 occured 6 time(s)
    // We cannot have more than one remainder of 1 and 4
    // because a pair of them will always be divisible by k
    // hence we take the max of the two
    // repeat for 2 and 4
    // note: if k is even, the mid case is not handled here
    left_index := 1
    right_index := int(k) - 1
    for left_index < right_index {
        left := remainder[left_index]
        right := remainder[right_index]
        if left > right {
            count += left
        } else {
            count += right
        }

        left_index += 1
        right_index -= 1
    }

    // handling mid elem
    // if k is even, then we can only have one remainder of k/2
    // because a pair of them will always be divisible by k
    if k % 2 == 0 && remainder[int(k/2)] > 0  {
        count++
    }
    return count
}
