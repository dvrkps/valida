package valida

import (
	"fmt"
	"testing"
)

func ExampleJMBG() {
	ok := JMBG("0308964384007")
	fmt.Println(ok)
	// Output:
	// true
}

func TestJMBG(t *testing.T) {
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
