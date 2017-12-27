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

	rem := zzz % 11
	diff := 11 - rem

	return validate(rem, 0, diff, last)
}
