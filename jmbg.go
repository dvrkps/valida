package valida

// JMBG validate JMBG number.
func JMBG(in string) bool {
	digs, ok := parseDigits([]byte(in))
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
	for _, d := range digs.first(noDigits) {
		zzz += d * coef

		if coef == coef2 {
			coef = coef8
		}
		coef--
	}

	return isValidJMBG(zzz, digs.last())
}
