package valida

// JMBG validate JMBG number.
func JMBG(in string) bool {
	digs, ok := parseDigits([]byte(in))
	if !ok {
		return false
	}
	if digs.size() != 13 {
		return false
	}

	coef := 7

	zzz := 0
	for _, d := range digs.withoutLast() {
		zzz += d * coef

		if coef == 2 {
			coef = 8
		}
		coef--
	}

	return isValidJMBG(zzz, digs.last())
}
