package mbs

import (
	"testing"
)

func TestOK(t *testing.T) {
	var tests = []struct {
		in   string
		want bool
	}{
		// valid
		{"01130234", true},
		{"011302340123", true},
		// too short
		{"123", false},
		// invalid
		{"12345678", false},
		// not number
		{"1a23b567", false},
		// all zeros
		{"00000000", false},
		{"000000000000", false},
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