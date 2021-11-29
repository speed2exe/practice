package practice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKPalindrome(t *testing.T) {
	tests := []struct {
		name        string
		givenString string
		givenK      int
		want        bool
	}{
		{
			name:        "given empty string then true",
			givenString: "",
			want:        true,
		},
		{
			name:        "given palindrome then true",
			givenString: "abcba",
			want:        true,
		},
		{
			name:        "given s=abxa and k=1 then want true",
			givenString: "abxa",
			givenK:      1,
			want:        true,
		},
		{
			name:        "given s=xaba and k=1 then want true",
			givenString: "xaba",
			givenK:      1,
			want:        true,
		},
		{
			name:        "given s=xacba and k=1 then want false",
			givenString: "xacba",
			givenK:      1,
			want:        false,
		},
		{
			name:        "given s=xacba and k=2 then want true",
			givenString: "xacba",
			givenK:      2,
			want:        true,
		},
		{
			name:        "given s=xabcYYba and k=1 then want false",
			givenString: "xacYYba",
			givenK:      1,
			want:        false,
		},
		{
			name:        "given s=xabYYcba and k=2 then want true",
			givenString: "xabYYcba",
			givenK:      2,
			want:        true,
		},
		{
			name:        "given s=abxa and k=0 then want true",
			givenString: "abxa",
			givenK:      0,
			want:        false,
		},
		{
			name:        "given s=abcd and k=4 then want true",
			givenString: "abcd",
			givenK:      4,
			want:        true,
		},
		{
			name:        "given s=abcd and k=8 then want true",
			givenString: "abcd",
			givenK:      8,
			want:        true,
		},
		{
			name:        "given s=xabas and k=2 then want true",
			givenString: "xabas",
			givenK:      2,
			want:        true,
		},
		{
			name:        "given s=xabas and k=1 then want false",
			givenString: "xabas",
			givenK:      1,
			want:        false,
		},
		{
			name:        "given s=saxyas and k=2 then want true",
			givenString: "xabas",
			givenK:      2,
			want:        true,
		},
		{
			name:        "given s=saxyas and k=1 then want false",
			givenString: "xabas",
			givenK:      1,
			want:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := KPalindrome([]rune(tt.givenString), tt.givenK)
			assert.Equal(t, tt.want, actual)
		})
	}
}
