package valida

import "strconv"

func digits(in string, noChars int) ([]int, int, bool) {
	empty := []int{}

	lenIn := len(in)
	if lenIn != noChars {
		return empty, 0, false
	}

	all := make([]int, 0, lenIn)
	zeroes := 0
	for _, char := range in {
		d, err := strconv.Atoi(string(char))
		if err != nil {
			return empty, 0, false
		}
		if d == 0 {
			zeroes++
		}
		all = append(all, d)
	}

	if zeroes == lenIn {
		return empty, 0, false
	}

	digits := all[:len(all)-1]
	last := all[len(all)-1]
	return digits, last, true
}
