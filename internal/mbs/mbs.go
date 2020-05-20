package mbs

import (
	"github.com/dvrkps/valida/internal/check"
	"github.com/dvrkps/valida/internal/digits"
)

// OK validate MBS number.
func OK(in string) bool {
	const maxDigits = 12
	if len(in) == maxDigits {
		in = in[:8]
	}

	digs, ok := digits.ParseDigits([]byte(in))
	if !ok {
		return false
	}

	zzz := 0

	const (
		noDigits = 7
		coefMax  = 8
	)

	for i, d := range digs.First(noDigits) {
		coef := coefMax - i
		zzz += d * coef
	}

	const wantRem = 1

	return check.OK(zzz, digs.Last(), wantRem)
}
