package valida

import (
	"fmt"
	"testing"
)

func ExampleOIB() {
	ok := OIB("69435151530")
	fmt.Println(ok)
	// Output:
	// true
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

func TestDigit(t *testing.T) {
	var tests = []struct {
		in   string
		want int
	}{
		// valid
		{"5", 5},
		// not number
		{"a", 0},
		// all zeros
		{"0", 0},
		// empty
		{"", 0},
	}
	for _, tt := range tests {
		if got := digit(tt.in); got != tt.want {
			t.Errorf("digit(\"%v\") = %v; want %v",
				tt.in, got, tt.want)
		}
	}
}

func TestIsValidOld(t *testing.T) {
	var tests = []struct {
		in   string
		want bool
	}{
		// valid
		{"0308964384007", true},
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
		if got := isValidOld(tt.in); got != tt.want {
			t.Errorf("isValidOld(\"%v\") = %v; want %v",
				tt.in, got, tt.want)
		}
	}
}
