package valida

func isValid(zzz int, last int) bool {
	return validate(zzz, last, 1)
}

func isValidJMBG(zzz int, last int) bool {
	return validate(zzz, last, 0)
}

func validate(zzz int, last int, wantRem int) bool {
	if zzz == 0 {
		return false
	}
	rem := zzz % 11

	if rem == wantRem && last == 0 {
		return true
	}

	if rem < 2 {
		return false
	}

	if rem > 10 {
		return false
	}

	diff := 11 - rem

	return last == diff
}
