package valida

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

func digitsString(in string) []int {
	var all []int
	for _, r := range in {
		d := digit(string(r))
		all = append(all, d)
	}
	return all
}
