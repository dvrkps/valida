package testutil

import "testing"

// TestCase is test case data.
type TestCase struct {
	Name  string
	Input string
	Want  bool
}

func Run(t *testing.T, fn func(string) bool, testCases []TestCase) {
	for _, tt := range testCases {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			got := fn(tt.Input)
			if got != tt.Want {
				t.Fatalf("got %v; want %v", got, tt.Want)
			}
		})
	}
}
