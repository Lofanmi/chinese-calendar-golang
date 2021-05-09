package utils

import "testing"

func TestOrderMod(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name       string
		args       args
		wantResult int64
	}{
		{"test_1", args{0, 12}, 12},
		{"test_2", args{1, 12}, 1},
		{"test_3", args{12, 12}, 12},
		{"test_4", args{13, 12}, 1},
		{"test_5", args{2, 12}, 2},
		{"test_6", args{14, 12}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := OrderMod(tt.args.a, tt.args.b); gotResult != tt.wantResult {
				t.Errorf("OrderMod() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestDD(t *testing.T) {
	type args struct {
		jd float64
	}
	tests := []struct {
		name   string
		args   args
		want_Y int
		want_M int
		want_D int
		want_h int
		want_m int
		want_s int
	}{
		{"2018-02-04 05:28:29", args{2458153.72812377}, 2018, 2, 4, 5, 28, 29},
		{"50046-01-13 12:00:00", args{19999999}, 50046, 1, 13, 12, 00, 00},
		{"-4712-01-02 12:00:00", args{1}, -4712, 1, 2, 12, 00, 00},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got_Y, got_M, got_D, got_h, got_m, got_s := DD(tt.args.jd)
			if got_Y != tt.want_Y {
				t.Errorf("DD() got_Y = %v, want %v", got_Y, tt.want_Y)
			}
			if got_M != tt.want_M {
				t.Errorf("DD() got_M = %v, want %v", got_M, tt.want_M)
			}
			if got_D != tt.want_D {
				t.Errorf("DD() got_D = %v, want %v", got_D, tt.want_D)
			}
			if got_h != tt.want_h {
				t.Errorf("DD() got_h = %v, want %v", got_h, tt.want_h)
			}
			if got_m != tt.want_m {
				t.Errorf("DD() got_m = %v, want %v", got_m, tt.want_m)
			}
			if got_s != tt.want_s {
				t.Errorf("DD() got_s = %v, want %v", got_s, tt.want_s)
			}
		})
	}
}
