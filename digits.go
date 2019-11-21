package valida

import "strconv"

type digits struct {
	all []int
}

func (d *digits) size() int {
	return len(d.all)
}

func (d *digits) withoutLast() []int {
	if len(d.all) < 2 {
		return []int{}
	}
	return d.all[:len(d.all)-1]
}

func (d *digits) last() int {
	if len(d.all) < 1 {
		return 0
	}
	return d.all[len(d.all)-1]
}

func parseDigits(in []byte) (digits, bool) {
	const (
		char0 = 48
		char9 = 57
	)

	all := make([]int, len(in))
	for i, c := range in {
		if c < char0 || c > char9 {
			return digits{}, false
		}
		all[i] = int(c) - char0
	}

	return digits{all: all}, true
}

func oldDigits(in string, noChars int) ([]int, int, bool) {
	empty := []int{}

	lenIn := len(in)
	if lenIn != noChars {
		return empty, 0, false
	}

	all := make([]int, 0, lenIn)

	for _, char := range in {
		d, err := strconv.Atoi(string(char))
		if err != nil {
			return empty, 0, false
		}

		all = append(all, d)
	}

	if isZero(all) {
		return empty, 0, false
	}

	digits := all[:len(all)-1]
	last := all[len(all)-1]

	return digits, last, true
}

func isZero(all []int) bool {
	n := 0

	for _, d := range all {
		if d == 0 {
			n++
		}
	}

	return n == len(all)
}
