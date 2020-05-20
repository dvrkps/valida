package oib

import (
	"testing"
)

func TestOK(t *testing.T) {
	var tests = []struct {
		in   string
		want bool
	}{
		// valid
		{"69435151530", true},
		// too short
		{"123", false},
		// invalid
		{"12345678901", false},
		// not number
		{"123a5b7c901", false},
		// all zeros
		{"00000000000", false},
		// empty
		{"", false},
	}

	for _, tt := range tests {
		if got := OK(tt.in); got != tt.want {
			t.Errorf("OK(%q) = %v; want %v",
				tt.in, got, tt.want)
		}
	}
}
