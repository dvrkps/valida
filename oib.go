// Package valida implements JMBG, MBS, Municipal ID and OIB numbers validations.
package valida

// OIB validate OIB number.
func OIB(in string) bool {
	digs, ok := parseDigits([]byte(in))
	if !ok {
		return false
	}

	o := 10

	const noDigits = 10

	for _, d := range digs.first(noDigits) {
		o += d
		o %= 10

		if o == 0 {
			o = 10
		}

		o *= 2
		o %= 11
	}

	// calc control char
	const (
		ctrl10 = 10
		ctrl11 = 11
	)

	ctrl := ctrl11 - o
	if ctrl == ctrl10 {
		ctrl = 0
	}

	return ctrl == digs.last()
}
