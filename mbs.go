package valida

import (
	"github.com/dvrkps/valida/internal/digits"
	"github.com/dvrkps/valida/internal/validate"
)

// MBS validate MBS number.
func MBS(in string) bool {
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

	return validate.IsValid(zzz, digs.Last())
}
