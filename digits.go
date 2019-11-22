package valida

type digits struct {
	all []int
}

func (d *digits) first(n int) []int {
	if len(d.all) < n {
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
