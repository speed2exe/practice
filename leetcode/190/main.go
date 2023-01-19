// https://leetcode.com/problems/reverse-bits/
package main

import (
	"unsafe"
)

func reverseBits(num uint32) uint32 {
	bs := *(*[4]uint8)(unsafe.Pointer(&num))
	bs[0], bs[1], bs[2], bs[3] = bs[3], bs[2], bs[1], bs[0]
	for i, b := range bs {
        b = b >> 4 | b << 4
        b = (b & 0b11001100) >> 2 | (b & 0b00110011) << 2
        b = (b & 0b10101010) >> 1 | (b & 0b01010101) << 1
        bs[i] = b
	}
    return *(*uint32)(unsafe.Pointer(&bs))
}
