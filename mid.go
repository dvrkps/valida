package valida

// MID validate Municipal ID number.
func MID(in string) bool {
	first3, last, ok := oldDigits(in, 4)
	if !ok {
		return false
	}

	zzz := 0

	for i, d := range first3 {
		coef := 4 - i
		zzz += d * coef
	}

	return isValid(zzz, last)
}
