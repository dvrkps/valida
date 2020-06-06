package oib

import (
	"testing"

	"github.com/dvrkps/valida/internal/testutil"
)

func TestOK(t *testing.T) {
	var tests = []testutil.TestCase{
		{Name: "valid", Input: "69435151530", Want: true},
		{Name: "too short", Input: "123"},
		{Name: "invalid", Input: "12345678901"},
		{Name: "not a number", Input: "123a5b7c901"},
		{Name: "empty"},
	}

	testutil.Run(t, OK, tests)
}
