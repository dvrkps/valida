package oib

import "github.com/dvrkps/valida/internal/digits"

// OK validate OIB number.
func OK(in string) bool {
	digs, ok := digits.ParseDigits([]byte(in))
	if !ok {
		return false
	}

	o := 10

	const noDigits = 10

	for _, d := range digs.First(noDigits) {
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

	return ctrl == digs.Last()
}
