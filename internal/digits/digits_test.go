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
			if ok == false {
				return
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
		{
			name:  "empty",
			input: "",
		},
		{
			name:  "size not equal",
			input: "123",
			size:  5,
		},
		{
			name:  "not a number",
			input: "1a2b3",
			size:  5,
		},
	}

	return tests
}
