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
		{in: 0, want: nil},
		{in: -1, want: nil},
	}
	for _, tt := range tests {
		got := digits(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("digits(%v) = %v; want %v", tt.in, got, tt.want)
		}
	}
}

func TestDigitsString(t *testing.T) {
	tests := []struct {
		in   string
		want []int
	}{
		{in: "12345", want: []int{1, 2, 3, 4, 5}},
		{in: "", want: nil},
	}
	for _, tt := range tests {
		got := digitsString(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("digits(%v) = %v; want %v", tt.in, got, tt.want)
		}
	}
}

var result []int

func BenchmarkDigits(b *testing.B) {
	var r []int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = digits(12345678901234)
	}
	result = r
}

func BenchmarkDigitsString(b *testing.B) {
	var r []int
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = digitsString("12345678901234")
	}
	result = r
}
