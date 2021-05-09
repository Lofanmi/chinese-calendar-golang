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

func TestYear2AYear(t *testing.T) {
	type args struct {
		year string
	}
	tests := []struct {
		name    string
		args    args
		wantY   int
		wantErr bool
	}{
		{"(empty)", args{year: ""}, 0, true},
		{"2021", args{year: "2021"}, 2021, false},
		{"b2021", args{year: "b2021"}, -2020, false},
		{"*2021", args{year: "*2021"}, -2020, false},
		{"BC2021", args{year: "BC2021"}, -2020, false},
		{"BC2021-", args{year: "BC2021-"}, 0, true},
		{"BC-2021", args{year: "BC-2021"}, -10000, true},
		{"B9999", args{year: "B9999"}, -9998, true},
		{"20210", args{year: "20210"}, 20210, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotY, err := Year2AYear(tt.args.year)
			if (err != nil) != tt.wantErr {
				t.Errorf("Year2AYear() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotY != tt.wantY {
				t.Errorf("Year2AYear() gotY = %v, want %v", gotY, tt.wantY)
			}
		})
	}
}

func TestJD(t *testing.T) {
	type args struct {
		yy int
		mm int
		dd int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"1900-01-02", args{1900,1,2}, 2415021.5},
		{"2020-05-02", args{2020,5,2}, 2458971.5},
		{"2021-02-02", args{2021,2,2}, 2459247.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JD(tt.args.yy, tt.args.mm, tt.args.dd); got != tt.want {
				t.Errorf("JD() = %v, want %v", got, tt.want)
			}
		})
	}
}