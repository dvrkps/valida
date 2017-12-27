package valida

import (
	"reflect"
	"testing"
)

func TestDigits(t *testing.T) {
	tests := []struct {
		fail    bool
		in      string
		noChars int
		digits  []int
		last    int
	}{
		{in: "12345", noChars: 5, digits: []int{1, 2, 3, 4}, last: 5},
		{fail: true, in: "0000000000", noChars: 10, digits: []int{}, last: 0},
		{fail: true, in: "12a", noChars: 3, digits: []int{}, last: 0},
		{fail: true, in: "", noChars: 3, digits: []int{}, last: 0},
		{fail: true, in: "-12", noChars: 3, digits: []int{}, last: 0},
	}
	for _, tt := range tests {
		gotDigits, gotLast, ok := digits(tt.in, tt.noChars)
		if tt.fail {
			if len(gotDigits) != 0 || gotLast != 0 || ok {
				t.Errorf("digits(%q, %v) = %v, %v, %v; want %v, %v, <false>",
					tt.in, tt.noChars, gotDigits, gotLast, ok, tt.digits, tt.last)
			}
			continue
		}

		if !reflect.DeepEqual(gotDigits, tt.digits) || gotLast != tt.last || !ok {
			t.Errorf("digits(%q, %v) = %v, %v, %v; want %v, %v, <false>",
				tt.in, tt.noChars, gotDigits, gotLast, ok, tt.digits, tt.last)
		}

	}
}
