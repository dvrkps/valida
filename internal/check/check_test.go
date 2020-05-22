package check

import "testing"

func TestOK(t *testing.T) {
	const (
		zzz11    = 11
		zzz12    = 12
		zzz13    = 13
		wantRem1 = 1
		control9 = 9
	)

	returnTrue := func(n int) bool { return true }
	returnFalse := func(n int) bool { return false }

	tests := []struct {
		ok        bool
		zzz       int
		wantRem   int
		controlFn func(int) bool
	}{
		{ok: true, zzz: zzz11, controlFn: returnTrue},
		{ok: true, zzz: zzz12, controlFn: returnTrue, wantRem: wantRem1},
		{ok: true, zzz: zzz13, controlFn: func(n int) bool { return n == control9 }},
		{controlFn: returnFalse},
		{zzz: zzz12, controlFn: returnFalse},
	}

	for _, tt := range tests {
		ok := OK(tt.zzz, tt.controlFn, tt.wantRem)
		if ok != tt.ok {
			t.Errorf("ok(%v, fn, %v) = %v, want %v",
				tt.zzz, tt.wantRem, ok, tt.ok)
		}
	}
}
