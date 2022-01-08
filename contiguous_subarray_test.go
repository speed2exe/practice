package practice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContiguousSubarrays(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "test 5 elem",
			arr:  []int{3, 4, 1, 6, 2},
			want: []int{1, 3, 1, 5, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := countSubarrays(tt.arr)
			assert.Equal(t, tt.want, actual)
		})
	}
}
