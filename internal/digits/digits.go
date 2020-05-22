package digits

// All holds number digits.
type All struct {
	exceptLast []int
	last       int
	current    int
}

// New creates all digits.
func New(in string, size int) (*All, bool) {
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

/*
func (d *Digits) First(n int) []int {
	if len(d.all) < n {
		return []int{}
	}

	return d.all[:len(d.all)-1]
}

func (d *Digits) Last() int {
	const minLen = 1

	lenAll := len(d.all)
	if lenAll < minLen {
		return 0
	}

	return d.all[lenAll-1]
}

func ParseDigits(in []byte) (Digits, bool) {
	const (
		char0 = 48
		char9 = 57
	)

	all := make([]int, len(in))

	for i, c := range in {
		if c < char0 || c > char9 {
			return Digits{}, false
		}

		all[i] = int(c) - char0
	}

	return Digits{all: all}, true
}
*/
