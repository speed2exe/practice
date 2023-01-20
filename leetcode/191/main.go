// https://leetcode.com/problems/number-of-1-bits/
package main

import (
	"unsafe"
)

func hammingWeight(num uint32) int {
	result := 0
	bs := *(*[4]uint8)(unsafe.Pointer(&num))
	for _, b := range bs {
		for i := 0; i < 8; i++ {
			if (b & (1 << i)) > 0 {
				result++
			}
		}
	}
	return result
}
