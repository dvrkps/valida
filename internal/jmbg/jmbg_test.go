package jmbg

import (
	"testing"

	"github.com/dvrkps/valida/internal/testutil"
)

func TestOK(t *testing.T) {
	var tests = []testutil.TestCase{
		{Name: "valid 1", Input: "0308964384007", Want: true},
		{Name: "valid 2", Input: "0123456788100", Want: true},
		{Name: "too short", Input: "123"},
		{Name: "invalid", Input: "1234567890123"},
		{Name: "not a number", Input: "123a567b90123"},
		{Name: "zeros", Input: "0000000000000"},
		{Name: "empty", Input: ""},
	}

	testutil.Run(t, OK, tests)
}
