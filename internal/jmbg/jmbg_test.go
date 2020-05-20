package jmbg

import (
	"testing"
)

func TestOK(t *testing.T) {
	var tests = []struct {
		in   string
		want bool
	}{
		// valid
		{"0308964384007", true},
		{"0123456788100", true},
		// too short
		{"123", false},
		// invalid
		{"1234567890123", false},
		// not number
		{"123a567b90123", false},
		// all zeros
		{"0000000000000", false},
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
