package mbs

import (
	"testing"

	"github.com/dvrkps/valida/internal/testutil"
)

func TestOK(t *testing.T) {
	var tests = []testutil.TestCase{
		{Name: "valid short", Input: "01130234", Want: true},
		{Name: "valid long", Input: "011302340123", Want: true},
		{Name: "too short", Input: "123"},
		{Name: "invalid", Input: "12345678"},
		{Name: "not a number", Input: "1a23b567"},
		{Name: "zeros short", Input: "00000000"},
		{Name: "zeros long", Input: "000000000000"},
		{Name: "empty", Input: ""},
	}

	testutil.Run(t, OK, tests)
}
