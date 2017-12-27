package valida

// MBS validate MBS number.
func MBS(in string) bool {
	if (len(in) != 8 && len(in) != 12) || !isValidOld(in) {
		return false
	}

	ost, raz := mbsCalc(in)

	return mbsValid(in, ost, raz)

}

func mbsCalc(in string) (int, int) {
	digits := in[:7]
	zzz := 0
	for i, r := range digits {
		coef := 8 - i
		// exit if char not digit
		d := digit(string(r))
		zzz += d * coef
	}

	ost := zzz % 11
	raz := 11 - ost

	return ost, raz
}

func mbsValid(in string, ost, raz int) bool {
	control := digit(in[7:8])
	ok1 := ost == 1 && control == 0
	ok2 := ost > 1 && ost < 11 && control == raz
	if ok1 || ok2 {
		return true
	}
	return false
}
