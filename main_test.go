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
