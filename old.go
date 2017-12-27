// Package valida implements JMBG, MBS, Municipal ID and OIB numbers validations.
package valida

import "strconv"

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
