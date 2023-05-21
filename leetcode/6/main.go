// https://leetcode.com/problems/zigzag-conversion/
package main

import (
	"bytes"
	"fmt"
)

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	if len(s) < 2 {
		return s
	}

	bs := make([]byte, 0, len(s))

	// first row
	skips := numRows - 2
	for i := 0; i < len(s); i += skips + numRows {
		bs = append(bs, s[i])
	}

	// middle
	for i := 1; i <= skips; i++ {
		for j := i; j < len(s); j += skips + numRows {
			bs = append(bs, s[j])
			extra := j + 2*skips - 2*i + 2
			if extra >= len(s) {
				continue
			}
			bs = append(bs, s[extra])
		}
	}

	// last
	for i := numRows - 1; i < len(s); i += skips + numRows {
		bs = append(bs, s[i])
	}

	return string(bs)
}

func convertFormatted(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	if len(s) < 2 {
		return s
	}

	b := bytes.Buffer{}

	// first row
	skips := numRows - 2
	for i := 0; i < len(s); i += skips + numRows {
		b.WriteByte(s[i])
		b.Write(bytes.Repeat([]byte{' '}, skips))
	}
	b.Truncate(b.Len() - skips)
	b.WriteByte('\n')

	// middle
	for i := 1; i <= skips; i++ {
		for j := i; j < len(s); j += skips + numRows {
			b.WriteByte(s[j])
			extra := j + 2*skips - 2*i + 2
			if extra >= len(s) {
				continue
			}
			b.Write(bytes.Repeat([]byte{' '}, numRows-2-i))
			b.WriteByte(s[extra])
			b.Write(bytes.Repeat([]byte{' '}, i-1))
		}
		b.WriteByte('\n')
	}

	// last
	for i := numRows - 1; i < len(s); i += skips + numRows {
		b.WriteByte(s[i])
		b.Write(bytes.Repeat([]byte{' '}, skips))
	}

	return string(bytes.TrimRight(b.Bytes(), " "))
}

func main() {
	fmt.Println(convert("123", 3))
	fmt.Println(convertFormatted("123", 3))
}
