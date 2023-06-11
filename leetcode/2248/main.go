package main

import "sort"

func intersection(numss [][]int) []int {
	if len(numss) == 0 {
		return []int{}
	} else if len(numss) == 1 {
		sort.Ints(numss[0])
		return numss[0]
	}

	curMap := mapFromSlice(numss[0])
	for _, nums := range numss[1:] {
		toRemove := []int{}
		m := mapFromSlice(nums)
		for k := range curMap {
			_, exists := m[k]
			if !exists {
				toRemove = append(toRemove, k)
			}
		}
		for _, r := range toRemove {
			delete(curMap, r)
		}
	}

	res := make([]int, 0, len(curMap))
	for k := range curMap {
		res = append(res, k)
	}

	sort.Ints(res)
	return res
}

func mapFromSlice(xs []int) map[int]struct{} {
	res := make(map[int]struct{})
	for _, x := range xs {
		res[x] = struct{}{}
	}
	return res
}
