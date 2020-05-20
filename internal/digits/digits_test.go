package digits

import (
	"testing"
)

func TestParseDigits(t *testing.T) {
	tests := []struct {
		ok   bool
		in   string
		want Digits
	}{
		{ok: true, in: "01234", want: Digits{all: []int{0, 1, 2, 3, 4}}},
		{ok: true, in: "0", want: Digits{all: []int{0}}},
		{in: "12a34"},
		{in: "-0"},
		{in: "+4"},
		{in: "ab"},
	}
	for _, tt := range tests {
		in := []byte(tt.in)
		got, gotOK := ParseDigits(in)

		if !tt.ok {
			if gotOK {
				t.Errorf("ParseDigits(%q) = %v, %v; want %v, <false>",
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
			t.Errorf("ParseDigits(%q) = %v, %v; want %v, <false>",
				tt.in, got, gotOK, tt.want)
		}
	}
}
