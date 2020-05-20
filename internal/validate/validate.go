package validate

func IsValid(zzz int, last int) bool {
	return validate(zzz, last, 1)
}

func IsValidJMBG(zzz int, last int) bool {
	return validate(zzz, last, 0)
}

func validate(zzz int, last int, wantRem int) bool {
	if zzz == 0 {
		return false
	}

	const remModulo = 11
	rem := zzz % remModulo

	if rem == wantRem && last == 0 {
		return true
	}

	const remMin = 2
	if rem < remMin {
		return false
	}

	const diffMax = 11
	diff := diffMax - rem

	return last == diff
}
