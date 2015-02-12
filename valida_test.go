package valida

import (
	"fmt"
	"testing"
)

func ExampleOIB() {
	ok := "69435151530"
	wrong := "invalidOIB"
	fmt.Println(OIB(ok))
	fmt.Println(OIB(wrong))
	// Output:
	// true
	// false
}

func TestJMBG(t *testing.T) {
	var tests = []struct {
		in   string
		want bool
	}{
		// valid
		{"0308964384007", true},
		// too short
		{"123", false},
		// invalid
		{"1234567890123", false},
		// not number
		{"123a567b90123", false},
		// last invalid
		{"123456789012a", false},
		// all zeros
		{"0000000000000", false},
		// empty
		{"", false},
	}
	for _, tt := range tests {
		if got := JMBG(tt.in); got != tt.want {
			t.Errorf("JMBG(\"%v\") = %v; want %v",
				tt.in, got, tt.want)
		}
	}
}

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
		// all zeros
		{"00000000000", false},
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
