// Package valida implements OIB and JMBG validations.
package valida

import "strconv"

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
