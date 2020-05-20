package valida

import (
	"testing"
)

func TestMID(t *testing.T) {
	var tests = []struct {
		in   string
		want bool
	}{
		// valid
		{"1333", true},
		{"4880", true},
		// too short
		{"123", false},
		// invalid
		{"1234", false},
		// not number
		{"1a23", false},
		// all zeros
		{"0000", false},
		// empty
		{"", false},
	}

	for _, tt := range tests {
		if got := MID(tt.in); got != tt.want {
			t.Errorf("MID(%q) = %v; want %v",
				tt.in, got, tt.want)
		}
	}
}
