package valida

// MBS validate MBS number.
func MBS(in string) bool {
	const maxDigits = 12
	if len(in) == maxDigits {
		in = in[:8]
	}

	digs, ok := parseDigits([]byte(in))
	if !ok {
		return false
	}

	zzz := 0

	const (
		noDigits = 7
		coefMax  = 8
	)

	for i, d := range digs.first(noDigits) {
		coef := coefMax - i
		zzz += d * coef
	}

	return isValid(zzz, digs.last())
}
