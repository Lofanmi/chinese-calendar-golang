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
