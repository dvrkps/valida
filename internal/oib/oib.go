package oib

import "github.com/dvrkps/valida/internal/digits"

// OK validate OIB number.
func OK(in string) bool {
	const size = 11

	d, ok := digits.New(in, size)
	if !ok {
		return false
	}

	o := 10

	d.Range(func() {
		o += d.Current()
		o %= 10

		if o == 0 {
			o = 10
		}

		o *= 2
		o %= 11
	})

	// calc control char
	const (
		ctrl10 = 10
		ctrl11 = 11
	)

	ctrl := ctrl11 - o
	if ctrl == ctrl10 {
		ctrl = 0
	}

	return d.Control(ctrl)
}
