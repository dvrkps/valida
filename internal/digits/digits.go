package digits

// All holds number digits.
type All struct {
	exceptLast []int
	last       int
	current    int
}

// New creates all digits.
func New(in string, size int) (*All, bool) {
	if in == "" {
		return nil, false
	}

	if len(in) != size {
		return nil, false
	}

	const (
		char0 = 48
		char9 = 57
	)

	ds := make([]int, 0, len(in))

	for _, c := range []byte(in) {
		if c < char0 || c > char9 {
			return nil, false
		}

		ds = append(ds, int(c)-char0)
	}

	a := All{
		exceptLast: ds[:size-1],
		last:       ds[size-1],
	}

	return &a, true
}

// Range call function for all non control digits.
func (a *All) Range(fn func()) {
	a.current = 0

	for range a.exceptLast {
		fn()
		a.current++
	}
}

// Current returns current digit.
func (a *All) Current() int {
	return a.exceptLast[a.current]
}

// Control checks control digit.
func (a *All) Control(n int) bool {
	return n == a.last
}
