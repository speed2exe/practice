package practice

import (
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
			name:   "can rotate 3",
			input:  "Zebra-493?",
			factor: 3,
			want:   "Cheud-726?",
		},
		{
			name:   "can rotate 39",
			input:  "abcdefghijklmNOPQRSTUVWXYZ0123456789",
			factor: 39,
			want:   "nopqrstuvwxyzABCDEFGHIJKLM9012345678",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := rotationalCipher(tt.input, tt.factor)
			assert.Equal(t, tt.want, actual)
		})
	}
}
