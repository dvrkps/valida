package testutil

import "testing"

func TestRun(t *testing.T) {
	tests := []TestCase{
		{Name: "valid",
			Input: "valid",
			Want:  true,
		},
		{Name: "invalid",
			Input: "invalid",
			Want:  false,
		},
	}

	testFn := func(in string) bool { return in == "valid" }

	Run(t, testFn, tests)
}
