package valida

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDigits(t *testing.T) {
	tests := []struct {
		all    []int
		size   int
		woLast []int
		last   int
	}{
		{all: []int{1, 2, 3}, size: 3, woLast: []int{1, 2}, last: 3},
		{all: []int{1}, size: 1, woLast: []int{}, last: 1},
		{all: []int{}},
		{all: nil},
	}
	for _, tt := range tests {
		d := digits{all: tt.all}
		t.Run(fmt.Sprintf("%v", d.all), func(t *testing.T) {
			size := d.size()
			if d.size() != tt.size {
				t.Errorf("size() = %v, want %v",
					size, tt.size)
			}
			woLast := d.withoutLast()
			equal := true
			if len(woLast) != len(tt.woLast) {
				equal = false
			}
			for i, v := range woLast {
				if v != tt.woLast[i] {
					equal = false
				}
			}
			if !equal {
				t.Errorf("withoutLast() = %v, want %v",
					woLast, tt.woLast)
			}
			last := d.last()
			if last != tt.last {
				t.Errorf("last() = %v, want %v",
					last, tt.last)
			}
		})
	}
}

func TestParseDigits(t *testing.T) {
	tests := []struct {
		ok   bool
		in   string
		want digits
	}{
		{ok: true, in: "01234", want: digits{all: []int{0, 1, 2, 3, 4}}},
		{ok: true, in: "0", want: digits{all: []int{0}}},
		{ok: true, in: "12a34"},
		{in: "-0"},
		{in: "+4"},
		{in: "ab"},
	}
	for _, tt := range tests {
		in := []byte(tt.in)
		got, gotOK := parseDigits(in)
		if !tt.ok {
			if gotOK {
				t.Errorf("parseDigits(%q) = %v, %v; want %v, <false>",
					in, got, gotOK, tt.want)
			}

			continue
		}
		equal := true
		if len(got.all) != len(tt.want.all) {
			equal = false
		}
		for i, v := range got.all {
			if v != tt.want.all[i] {
				equal = false
			}
		}
		if !equal {
			t.Errorf("parseDigits(%q) = %v, %v; want %v, <false>",
				tt.in, got, gotOK, tt.want)
		}
	}
}

func TestOldDigits(t *testing.T) {
	tests := []struct {
		ok      bool
		in      string
		noChars int
		digits  []int
		last    int
	}{
		{ok: true, in: "12345", noChars: 5, digits: []int{1, 2, 3, 4}, last: 5},
		{in: "0000000000", noChars: 10, digits: []int{}, last: 0},
		{in: "12a", noChars: 3, digits: []int{}, last: 0},
		{in: "", noChars: 3, digits: []int{}, last: 0},
		{in: "-12", noChars: 3, digits: []int{}, last: 0},
	}
	for _, tt := range tests {
		gotDigits, gotLast, ok := oldDigits(tt.in, tt.noChars)
		if !tt.ok {
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
