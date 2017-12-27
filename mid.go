package valida

// MID validate Municipal ID number.
func MID(in string) bool {

	first3, last, ok := digits(in, 4)
	if !ok {
		return false
	}

	zzz := 0
	for i, d := range first3 {
		coef := 4 - i
		zzz += d * coef
	}

	rem := zzz % 11
	diff := 11 - rem

	return validate(rem, 1, diff, last)
}
