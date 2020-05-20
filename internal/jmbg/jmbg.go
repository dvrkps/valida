package jmbg

import (
	"github.com/dvrkps/valida/internal/check"
	"github.com/dvrkps/valida/internal/digits"
)

func OK(in string) bool {
	digs, ok := digits.ParseDigits([]byte(in))
	if !ok {
		return false
	}

	const (
		noDigits = 12
		coef2    = 2
		coef7    = 7
		coef8    = 8
	)

	coef := coef7

	zzz := 0
	for _, d := range digs.First(noDigits) {
		zzz += d * coef

		if coef == coef2 {
			coef = coef8
		}
		coef--
	}

	const wantRem = 0

	return check.OK(zzz, digs.Last(), wantRem)
}
