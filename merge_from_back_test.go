package practice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeFromBack(t *testing.T) {
	tests := []struct {
		name   string
		first  []int
		second []int
		want   []int
	}{
		{
			name: "given nil then want nil",
		},
		{
			name:   "given empty slice then want empty slice",
			first:  []int{},
			second: []int{},
			want:   []int{},
		},
		{
			name:   "given no second elem then return first",
			first:  []int{1, 2},
			second: []int{},
			want:   []int{1, 2},
		},
		{
			name:   "given no first elem then return first",
			first:  []int{-1, -1},
			second: []int{1, 2},
			want:   []int{1, 2},
		},
		{
			name:   "given second higher than first then merge",
			first:  []int{1, 2, 3, -1, -1},
			second: []int{4, 5},
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			name:   "given first higher than second then merge",
			first:  []int{100, 200, 300, -1, -1},
			second: []int{1, 2},
			want:   []int{1, 2, 100, 200, 300},
		},
		{
			name:   "given 7 same elems with 1 different elem then return sorted",
			first:  []int{1, 3, 5, 7, 9, -1, -1, -1, -1, -1},
			second: []int{2, 4, 6, 8, 10},
			want:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeBackwards(tt.first, tt.second)
			assert.Equal(t, tt.want, tt.first)
		})
	}
}
