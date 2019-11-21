// Package valida implements JMBG, MBS, Municipal ID and OIB numbers validations.
package valida

// OIB validate OIB number.
func OIB(in string) bool {
	first10, last, ok := oldDigits(in, 11)
	if !ok {
		return false
	}

	o := 10
	for _, d := range first10 {
		o += d
		o %= 10

		if o == 0 {
			o = 10
		}

		o *= 2
		o %= 11
	}

	// calc control char
	ctrl := 11 - o
	if ctrl == 10 {
		ctrl = 0
	}

	return ctrl == last
}
