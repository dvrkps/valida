package mbs

import (
	"github.com/dvrkps/valida/internal/check"
	"github.com/dvrkps/valida/internal/digits"
)

// OK validate MBS number.
func OK(in string) bool {
	const (
		longDigits  = 12
		shortDigits = 8
	)

	if len(in) == longDigits {
		in = in[:shortDigits]
	}

	d, ok := digits.New(in, shortDigits)
	if !ok {
		return false
	}

	var (
		zzz = 0
		i   = 0
	)

	const coefMax = 8

	d.Range(func() {
		coef := coefMax - i
		zzz += d.Current() * coef
		i++
	})

	const wantRem = 1

	return check.OK(zzz, d.Control, wantRem)
}
