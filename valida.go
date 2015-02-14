// Package valida implements OIB and JMBG validations.
package valida

import "strconv"

// JMBG validate JMBG number.
func JMBG(in string) bool {
	if len(in) != 13 || in == "0000000000000" {
		return false
	}
	// get first 12 chars
	digits := in[:len(in)-1]
	coef := 7
	zzz := 0
	for _, r := range digits {
		// exit if char not digit
		d, err := strconv.Atoi(string(r))
		if err != nil {
			return false
		}
		zzz += d * coef
		if coef == 2 {
			coef = 8
		}
		coef--
	}
	// exit if last char not digit
	last, err := strconv.Atoi(string(in[len(in)-1:]))
	if err != nil {
		return false
	}

	ost := zzz % 11
	raz := 11 - ost
	ok1 := ost == 0 && last == 0
	ok2 := ost > 1 && ost < 11 && last == raz
	if ok1 || ok2 {
		return true
	}
	return false
}

// MID validate Municipal ID number.
func MID(in string) bool {
	if len(in) != 4 || in == "0000" {
		return false
	}
	// get first 3 chars
	digits := in[:len(in)-1]
	coef := 4
	zzz := 0
	for _, r := range digits {
		// exit if char not digit
		d, err := strconv.Atoi(string(r))
		if err != nil {
			return false
		}
		zzz += d * coef
		if coef == 2 {
			coef = 8
		}
		coef--
	}
	// exit if last char not digit
	last, err := strconv.Atoi(string(in[len(in)-1:]))
	if err != nil {
		return false
	}
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
