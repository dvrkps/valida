package jmbg

import (
	"github.com/dvrkps/valida/internal/check"
	"github.com/dvrkps/valida/internal/digits"
)

// OK validate JMBG number.
func OK(in string) bool {
	const noDigits = 13

	d, ok := digits.New(in, noDigits)
	if !ok {
		return false
	}

	const (
		coef2 = 2
		coef7 = 7
		coef8 = 8
	)

	coef := coef7

	zzz := 0

	d.Range(func() {
		zzz += d.Current() * coef

		if coef == coef2 {
			coef = coef8
		}

		coef--
	})

	const wantRem = 0

	return check.OK(zzz, d.Control, wantRem)
}
