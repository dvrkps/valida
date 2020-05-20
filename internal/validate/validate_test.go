package validate

import "testing"

func TestValidate(t *testing.T) {
	const (
		zzz11    = 11
		zzz12    = 12
		zzz13    = 13
		wantRem1 = 1
		last9    = 9
	)

	tests := []struct {
		ok      bool
		zzz     int
		last    int
		wantRem int
	}{
		{ok: true, zzz: zzz11},
		{ok: true, zzz: zzz12, wantRem: wantRem1},
		{ok: true, zzz: zzz13, last: last9},
		{},
		{zzz: zzz12},
	}

	for _, tt := range tests {
		ok := validate(tt.zzz, tt.last, tt.wantRem)
		if ok != tt.ok {
			t.Errorf("validate(%v, %v, %v) = %v, want %v",
				tt.zzz, tt.last, tt.wantRem, ok, tt.ok)
		}
	}
}
