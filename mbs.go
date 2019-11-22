package valida

// MBS validate MBS number.
func MBS(in string) bool {
	if len(in) == 12 {
		in = in[:8]
	}

	digs, ok := parseDigits([]byte(in))
	if !ok {
		return false
	}

	zzz := 0

	for i, d := range digs.first(7) {
		coef := 8 - i
		zzz += d * coef
	}

	return isValid(zzz, digs.last())
}
