package practice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name  string
		given []int
		want  []int
	}{
		{
			name:  "given nil then want nil",
			given: nil,
			want:  nil,
		},
		{
			name:  "given empty slice then want empty slice",
			given: []int{},
			want:  []int{},
		},
		{
			name:  "given single elem then return single elem",
			given: []int{8},
			want:  []int{8},
		},
		{
			name:  "given 2 elems then return sorted",
			given: []int{8, 2},
			want:  []int{2, 8},
		},
		{
			name:  "given 7 elems then return sorted",
			given: []int{8, 2, 3, 4, 1, 5, 6},
			want:  []int{1, 2, 3, 4, 5, 6, 8},
		},
		{
			name:  "given 7 same elems with 1 different elem then return sorted",
			given: []int{8, 8, 8, 1},
			want:  []int{1, 8, 8, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+", with copy", func(t *testing.T) {
			actual := MergeSort(tt.given)
			assert.Equal(t, tt.want, actual)
		})
	}
}
