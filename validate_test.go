package valida

import "testing"

func TestValidate(t *testing.T) {
	tests := []struct {
		ok      bool
		zzz     int
		last    int
		wantRem int
	}{
		{ok: true, zzz: 11, last: 0, wantRem: 0},
		{ok: true, zzz: 12, last: 0, wantRem: 1},
		{ok: true, zzz: 13, last: 9, wantRem: 0},
		{},
		{zzz: 12},
		{zzz: 14},
	}
	for _, tt := range tests {
		ok := validate(tt.zzz, tt.last, tt.wantRem)
		if ok != tt.ok {
			t.Errorf("validate(%v, %v, %v) = %v, want %v",
				tt.zzz, tt.last, tt.wantRem, ok, tt.ok)
		}
	}
}
