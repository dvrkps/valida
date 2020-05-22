package mid

import (
	"github.com/dvrkps/valida/internal/check"
	"github.com/dvrkps/valida/internal/digits"
)

// OK validate Municipal ID number.
func OK(in string) bool {
	const size = 4

	d, ok := digits.New(in, size)
	if !ok {
		return false
	}

	const coefMax = 4

	var (
		zzz int
		i   int
	)

	d.Range(func() {
		coef := coefMax - i
		zzz += d.Current() * coef
		i++
	})

	const wantRem = 1

	return check.OK(zzz, d.Control, wantRem)
}
