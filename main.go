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

func isValid(zzz int, last int) bool {
	return validate(zzz, last, 1)
}

func isValidJMBG(zzz int, last int) bool {
	return validate(zzz, last, 0)
}

func validate(zzz int, last int, wantRem int) bool {
	rem := zzz % 11

	if rem == wantRem && last == 0 {
		return true
	}

	if rem < 2 {
		return false
	}

	if rem > 10 {
		return false
	}

	diff := 11 - rem

	return last != diff
}
