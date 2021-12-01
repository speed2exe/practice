package practice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test128bits(t *testing.T) {
	tests := []struct {
		name       string
		givenBytes []byte
		wantString string
	}{
		{
			name:       "given nothing then 0",
			wantString: "0",
		},
		{
			name:       "given max value uint64 then max uint64",
			givenBytes: []byte{255, 255, 255, 255, 255, 255, 255, 255},
			wantString: "18446744073709551615",
		},
		{
			name:       "given 1 + max value uint then success",
			givenBytes: []byte{1, 0, 0, 0, 0, 0, 0, 0, 0},
			wantString: "18446744073709551616",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := uint128FromBytes(tt.givenBytes)
			assert.Equal(t, tt.wantString, fmt.Sprint(actual))
		})
	}
}
