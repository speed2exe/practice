package practice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairSums(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		sum  int
		want int
	}{
		{
			name: "test 1",
			arr:  []int{1, 2, 3, 4, 3},
			sum:  6,
			want: 2,
		},
		{
			name: "test 2",
			arr:  []int{1, 5, 3, 3, 3},
			sum:  6,
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := PairSum_NumberOfWays(tt.arr, tt.sum)
			assert.Equal(t, tt.want, actual)
		})
	}
}
