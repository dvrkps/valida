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

	tests := []struct {
		ok      bool
		zzz     int
		control int
		wantRem int
	}{
		{ok: true, zzz: zzz11},
		{ok: true, zzz: zzz12, wantRem: wantRem1},
		{ok: true, zzz: zzz13, control: control9},
		{},
		{zzz: zzz12},
	}

	for _, tt := range tests {
		ok := OK(tt.zzz, tt.control, tt.wantRem)
		if ok != tt.ok {
			t.Errorf("validate(%v, %v, %v) = %v, want %v",
				tt.zzz, tt.control, tt.wantRem, ok, tt.ok)
		}
	}
}
