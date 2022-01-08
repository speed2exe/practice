package practice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotationalCipher(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		factor int
		want   string
	}{
		{
			name:   "given nothing then 0",
			input:  "Zebra-493?",
			factor: 3,
			want:   "Cheud-726?",
		},
		{
			name:   "given max value uint64 then max uint64",
			input:  "abcdefghijklmNOPQRSTUVWXYZ0123456789",
			factor: 39,
			want:   "nopqrstuvwxyzABCDEFGHIJKLM9012345678",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := rotationalCipher(tt.input, tt.factor)
			assert.Equal(t, tt.want, fmt.Sprint(actual))
		})
	}
}
