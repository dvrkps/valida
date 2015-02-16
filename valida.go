// Package valida implements JMBG, MBS, Municipal ID and OIB numbers validations.
package valida

import "strconv"

// getDigit returns digit.
func getDigit(in string) int {
	d, err := strconv.Atoi(in)
	if d < 0 || err != nil {
		return 0
	}
	return d
}

// isValid validate input string.
func isValid(in string) bool {
	i, err := strconv.ParseInt(in, 10, 64)
	if i > 0 && err == nil {
		return true
	}
	return false
}

// JMBG validate JMBG number.
func JMBG(in string) bool {
	if len(in) != 13 || !isValid(in) {
		return false
	}
	// get first 12 chars
	digits := in[:12]
	coef := 7
	zzz := 0
	for _, r := range digits {
		// exit if char not digit
		d := getDigit(string(r))
		zzz += d * coef
		if coef == 2 {
			coef = 8
		}
		coef--
	}
	// get last digit
	last := getDigit(string(in[12:13]))

	ost := zzz % 11
	raz := 11 - ost
	ok1 := ost == 0 && last == 0
	ok2 := ost > 1 && ost < 11 && last == raz
	if ok1 || ok2 {
		return true
	}
	return false
}

// MBS validate MBS number.
func MBS(in string) bool {
	if (len(in) != 8 && len(in) != 12) || !isValid(in) {
		return false
	}
	// get first 7 chars
	digits := in[:7]
	zzz := 0
	for i, r := range digits {
		coef := 8 - i
		// exit if char not digit
		d := getDigit(string(r))
		zzz += d * coef
	}
	// exit if control char not digit
	control := getDigit(string(in[7:8]))
	ost := zzz % 11
	raz := 11 - ost
	ok1 := ost == 1 && control == 0
	ok2 := ost > 1 && ost < 11 && control == raz
	if ok1 || ok2 {
		return true
	}
	return false
}

// MID validate Municipal ID number.
func MID(in string) bool {
	if len(in) != 4 || !isValid(in) {
		return false
	}
	// get first 3 chars
	digits := in[:3]
	zzz := 0
	for i, r := range digits {
		coef := 4 - i
		d := getDigit(string(r))
		zzz += d * coef
	}
	last := getDigit(string(in[3:4]))
	ost := zzz % 11
	raz := 11 - ost
	ok1 := ost == 1 && last == 0
	ok2 := ost > 1 && ost < 11 && last == raz
	if ok1 || ok2 {
		return true
	}
	return false
}

// OIB validate OIB number.
func OIB(in string) bool {
	if len(in) != 11 {
		return false
	}
	// get first 10 chars
	digits := in[:len(in)-1]
	//
	o := 10
	for _, r := range digits {
		// exit if char not digit
		d, err := strconv.Atoi(string(r))
		if err != nil {
			return false
		}
		o += d
		o = o % 10
		if o == 0 {
			o = 10
		}
		o *= 2
		o = o % 11
	}
	// exit if last char not digit
	last, err := strconv.Atoi(string(in[len(in)-1:]))
	if err != nil {
		return false
	}
	// calc control char
	ctrl := 11 - o
	if ctrl == 10 {
		ctrl = 0
	}
	return ctrl == last
}
