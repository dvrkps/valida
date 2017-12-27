package valida

import (
	"fmt"
	"testing"
)

func ExampleMBS() {
	okShort := MBS("01130234")
	okLong := MBS("011302340123")
	fmt.Println(okShort)
	fmt.Println(okLong)
	// Output:
	// true
	// true
}

func TestMBS(t *testing.T) {
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
		// last invalid
		{"1234567x", false},
		// all zeros
		{"00000000", false},
		// empty
		{"", false},
	}
	for _, tt := range tests {
		if got := MBS(tt.in); got != tt.want {
			t.Errorf("MBS(\"%v\") = %v; want %v",
				tt.in, got, tt.want)
		}
	}
}
