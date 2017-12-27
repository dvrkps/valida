package valida

// MBS validate MBS number.
func MBS(in string) bool {
	if len(in) == 12 {
		in = in[0:8]
	}
	first7, last, ok := digits(in, 8)
	if !ok {
		return false
	}

	zzz := 0
	for i, d := range first7 {
		coef := 8 - i
		zzz += d * coef
	}

	rem := zzz % 11
	diff := 11 - rem

	return validate(rem, 1, diff, last)
}
