package valida

import (
	"github.com/dvrkps/valida/internal/jmbg"
	"github.com/dvrkps/valida/internal/mbs"
	"github.com/dvrkps/valida/internal/mid"
	"github.com/dvrkps/valida/internal/oib"
)

// JMBG validate JMBG number.
func JMBG(n string) bool {
	return jmbg.OK(n)
}

// MBS validate MBS number.
func MBS(n string) bool {
	return mbs.OK(n)
}

// MID validate MID number.
func MID(n string) bool {
	return mid.OK(n)
}

// OIB validate OIB number.
func OIB(n string) bool {
	return oib.OK(n)
}
