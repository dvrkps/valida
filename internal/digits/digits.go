package digits

type Digits struct {
	all []int
}

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
