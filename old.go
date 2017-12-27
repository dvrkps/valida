// Package valida implements JMBG, MBS, Municipal ID and OIB numbers validations.
package valida

import "strconv"

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

// MID validate Municipal ID number.
func MID(in string) bool {
	if len(in) != 4 || !isValidOld(in) {
		return false
	}

	ost, raz := midCalc(in)

	return midValid(in, ost, raz)

}

func midCalc(in string) (int, int) {
	first3 := in[:3]
	zzz := 0
	for i, r := range first3 {
		coef := 4 - i
		d := digit(string(r))
		zzz += d * coef
	}

	ost := zzz % 11
	raz := 11 - ost

	return ost, raz
}

func midValid(in string, ost, raz int) bool {
	last := digit(in[3:4])
	ok1 := ost == 1 && last == 0
	ok2 := ost > 1 && ost < 11 && last == raz
	if ok1 || ok2 {
		return true
	}
	return false
}

// OIB validate OIB number.
func OIB(in string) bool {
	if len(in) != 11 || !isValidOld(in) {
		return false
	}

	first10 := in[:len(in)-1]
	o := oibCalc(first10)

	last := digit(in[len(in)-1:])

	// calc control char
	ctrl := 11 - o
	if ctrl == 10 {
		ctrl = 0
	}
	return ctrl == last
}

func oibCalc(digits string) int {
	o := 10
	for _, r := range digits {
		// exit if char not digit
		d := digit(string(r))
		o += d
		o = o % 10
		if o == 0 {
			o = 10
		}
		o *= 2
		o = o % 11
	}
	return o
}

// digit returns one digit.
func digit(in string) int {
	d, err := strconv.Atoi(in)
	if d < 0 || err != nil {
		return 0
	}
	return d
}

// isValid validate input string.
func isValidOld(in string) bool {
	i, err := strconv.ParseInt(in, 10, 64)
	if i > 0 && err == nil {
		return true
	}
	return false
}