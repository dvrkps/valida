package mid

import (
	"github.com/dvrkps/valida/internal/digits"
	"github.com/dvrkps/valida/internal/validate"
)

// OK validate Municipal ID number.
func OK(in string) bool {
	digs, ok := digits.ParseDigits([]byte(in))
	if !ok {
		return false
	}

	zzz := 0

	const (
		noDigits = 3
		coefMax  = 4
	)

	for i, d := range digs.First(noDigits) {
		coef := coefMax - i
		zzz += d * coef
	}

	return validate.IsValid(zzz, digs.Last())
}
