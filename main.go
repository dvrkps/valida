package valida

func digits(n int) []int {
	var all []int
	for n > 0 {
		d := n % 10
		all = append([]int{d}, all...)
		n /= 10
	}
	return all
}
