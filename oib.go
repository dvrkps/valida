// Package valida implements JMBG, MBS, Municipal ID and OIB numbers validations.
package valida

// OIB validate OIB number.
func OIB(in string) bool {
	digs, ok := parseDigits([]byte(in))
	if !ok {
		return false
	}
	if digs.size() != 11 {
		return false
	}

	o := 10
	for _, d := range digs.withoutLast() {
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

	return ctrl == digs.last()
}
