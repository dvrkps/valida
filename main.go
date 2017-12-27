package valida

import "strconv"

func digits(in string, noChars int) ([]int, int, bool) {
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

	digits := all[:len(all)-1]
	last := all[len(all)-1]
	return digits, last, true
}
