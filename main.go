package valida

import "strconv"

func digits(n int) []int {
	raw := make([]int, 0, 20)
	var d int
	for n > 0 {
		d = n % 10
		raw = append(raw, d)
		n /= 10
	}

	lenRaw := len(raw)
	all := make([]int, 0, lenRaw)
	for i := lenRaw - 1; i >= 0; i-- {
		all = append(all, raw[i])
	}

	return all
}

// digit returns one digit.
func digit(in string) int {
	d, err := strconv.Atoi(in)
	if d < 0 || err != nil {
		return 0
	}
	return d
}

// isValid validate input string.
func isValid(in string) bool {
	i, err := strconv.ParseInt(in, 10, 64)
	if i > 0 && err == nil {
		return true
	}
	return false
}
