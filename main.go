package valida

import "github.com/dvrkps/valida/internal/jmbg"

// JMBG validate JMBG number.
func JMBG(n string) bool {
	return jmbg.OK(n)
}
