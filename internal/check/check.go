package check

// OK checks number.
func OK(zzz int, control int, wantRem int) bool {
	if zzz == 0 {
		return false
	}

	const remModulo = 11
	rem := zzz % remModulo

	if rem == wantRem && control == 0 {
		return true
	}

	const remMin = 2
	if rem < remMin {
		return false
	}

	const diffMax = 11
	diff := diffMax - rem

	return control == diff
}
