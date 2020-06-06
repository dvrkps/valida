package mid

import (
	"testing"

	"github.com/dvrkps/valida/internal/testutil"
)

func TestOK(t *testing.T) {
	var tests = []testutil.TestCase{
		{Name: "valid", Input: "1333", Want: true},
		{Name: "too short", Input: "123"},
		{Name: "invalid", Input: "1234"},
		{Name: "not a number", Input: "1a23"},
		{Name: "empty"},
	}

	testutil.Run(t, OK, tests)
}
