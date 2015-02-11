package valida

import "testing"

func TestOIB(t *testing.T) {
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
		// last invalid
		{"1234567890a", false},
		// empty
		{"", false},
	}
	for _, tt := range tests {
		if got := OIB(tt.in); got != tt.want {
			t.Errorf("OIB(\"%v\") = %v; want %v",
				tt.in, got, tt.want)
		}
	}
}
