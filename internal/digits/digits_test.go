package digits

import "testing"

func TestNew(t *testing.T) {
	for _, tt := range testCases() {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			a, ok := New(tt.input, tt.size)
			if ok != tt.ok {
				t.Fatalf("ok: got %v; want %v", ok, tt.ok)
			}

			if !a.Control(tt.wantControl) {
				t.Fatalf("control: not equal %v", tt.wantControl)
			}

			var (
				wantIterations = len(tt.input) - 1
				i              int
				sum            int
			)

			a.Range(func() {
				sum += a.Current()
				i++
			})

			if i != wantIterations {
				t.Fatalf("i: got %v;want %v", i, wantIterations)
			}

			if sum != tt.wantSum {
				t.Fatalf("sum: got %v;want %v", sum, tt.wantSum)
			}
		})
	}
}

type testCase struct {
	name        string
	input       string
	size        int
	wantControl int
	wantSum     int
	ok          bool
}

func testCases() []testCase {
	tests := []testCase{
		{
			name:        "valid",
			input:       "01234",
			size:        5,
			wantControl: 4, // last digit
			wantSum:     6, // 0+1+2+3
			ok:          true,
		},
	}

	return tests
}

/*
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
*/
