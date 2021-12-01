package practice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestIncreasingSubSeq(t *testing.T) {
	tests := []struct {
		name      string
		givenNums []int
		wantLen   int
	}{
		{
			name:      "longest sub 1",
			givenNums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			wantLen:   4,
		},
		{
			name:      "longest sub 2",
			givenNums: []int{0, 1, 0, 3, 2, 3},
			wantLen:   4,
		},
		{
			name:      "longest sub 3",
			givenNums: []int{7, 7, 7, 7, 7},
			wantLen:   1,
		},
		{
			name:      "longest sub 1",
			givenNums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			wantLen:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := len(LongestIncreasingSubSeq(tt.givenNums))
			assert.Equal(t, tt.wantLen, actual)
		})
	}
}
