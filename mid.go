package valida

import "github.com/dvrkps/valida/internal/digits"

// MID validate Municipal ID number.
func MID(in string) bool {
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

	return isValid(zzz, digs.Last())
}
