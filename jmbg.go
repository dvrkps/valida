package valida

// JMBG validate JMBG number.
func JMBG(in string) bool {
	first12, last, ok := digits(in, 13)
	if !ok {
		return false
	}

	coef := 7
	zzz := 0
	for _, d := range first12 {
		zzz += d * coef
		if coef == 2 {
			coef = 8
		}
		coef--
	}

	ost := zzz % 11
	raz := 11 - ost

	if ost == 0 && last == 0 {
		return true
	}
	if ost > 1 && ost < 11 && last == raz {
		return true
	}
	return false
}
