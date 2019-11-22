package valida

// MID validate Municipal ID number.
func MID(in string) bool {
	digs, ok := parseDigits([]byte(in))
	if !ok {
		return false
	}

	zzz := 0

	for i, d := range digs.first(3) {
		coef := 4 - i
		zzz += d * coef
	}

	return isValid(zzz, digs.last())
}
