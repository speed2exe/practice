// from facebookrecruting

package practice

import (
	"bytes"
	"strconv"
	"unicode"
)

// Rotational Cipher
// One simple way to encrypt a string is to "rotate" every alphanumeric
// character by a certain amount. Rotating a character means replacing it with
// another character that is a certain number of steps away in normal
// alphabetic or numerical order.
//
// For example, if the string "Zebra-493?" is rotated 3 places, the resulting
// string is "Cheud-726?". Every alphabetic character is replaced with the
// character 3 letters higher (wrapping around from Z to A), and every numeric
// character replaced with the character 3 digits higher (wrapping around from
// 9 to 0). Note that the non-alphanumeric characters remain unchanged.
// Given a string and a rotation factor, return an encrypted string.
//
// Signature string rotationalCipher(string input, int rotationFactor)
//
// Input 1 <= |input| <= 1,000,000
// 0 <= rotationFactor <= 1,000,000
//
// Output
// Return the result of rotating input a number of times equal to rotationFactor.

// Example 1
// input = Zebra-493? rotationFactor = 3
// output = Cheud-726?

// Example 2 input =
// abcdefghijklmNOPQRSTUVWXYZ0123456789 rotationFactor = 39
// output = nopqrstuvwxyzABCDEFGHIJKLM9012345678

func rotationalCipher(input string, rotationFactor int) string {
	var result bytes.Buffer

	for _, c := range input {
		if unicode.IsNumber(c) {
			c = rotateNumber(c, rotationFactor)
		} else if unicode.IsLetter(c) {
			c = rotateLetter(c, rotationFactor)
		}
		result.WriteRune(c)
	}

	return result.String()
}

func rotateNumber(r rune, factor int) rune {
	number, _ := strconv.ParseUint(string(r), 10, 8)
	number += uint64(factor)
	number %= 10
	return rune([]rune(strconv.FormatUint(number, 10))[0])
}

const (
	minLowerChar rune = 'a'
	maxLowerChar rune = 'z'
	minUpperChar rune = 'A'
	maxUpperChar rune = 'Z'
)

func rotateLetter(r rune, factor int) rune {
	limit := maxUpperChar
	min := minUpperChar

	if unicode.IsLower(r) {
		limit = maxLowerChar
		min = minLowerChar
	}

	r += rune(factor)
	if r > limit {
		r = ((r - min) % 26) + min
	}
	return r
}
