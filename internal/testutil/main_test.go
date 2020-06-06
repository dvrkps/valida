package testutil

import "testing"

func TestRun(t *testing.T) {
	tests := []TestCase{
		{Name: "ok",
			Input: "valid",
			Want:  true,
		},
		{Name: "not-ok",
			Input: "invalid",
			Want:  false,
		},
	}

	testFn := func(in string) bool { return in == "valid" }

	Run(t, testFn, tests)
}
