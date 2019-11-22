package valida

// MID validate Municipal ID number.
func MID(in string) bool {
	digs, ok := parseDigits([]byte(in))
	if !ok {
		return false
	}

	zzz := 0

	const (
		noDigits = 3
		coefMax  = 4
	)

	for i, d := range digs.first(noDigits) {
		coef := coefMax - i
		zzz += d * coef
	}

	return isValid(zzz, digs.last())
}
