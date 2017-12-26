package valida

import (
	"reflect"
	"testing"
)

func TestDigits(t *testing.T) {
	tests := []struct {
		in   int
		want []int
	}{
		{in: 12345, want: []int{1, 2, 3, 4, 5}},
		{in: 0, want: []int{}},
		{in: -1, want: []int{}},
	}
	for _, tt := range tests {
		got := digits(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("digits(%v) = %v; want %v", tt.in, got, tt.want)
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

func TestIsValid(t *testing.T) {
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
		if got := isValid(tt.in); got != tt.want {
			t.Errorf("isValid(\"%v\") = %v; want %v",
				tt.in, got, tt.want)
		}
	}
}
