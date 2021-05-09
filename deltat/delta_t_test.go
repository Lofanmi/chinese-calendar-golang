package deltat

import (
	"testing"
)

func TestDtT(t *testing.T) {
	type args struct {
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"-4712-01-01 12:00:00,0", args{0}, 0.0007392361111111111},
		{"-4712-01-03,1.5", args{1.5}, 0.0007392456177326632},
		{"1900-01-02,2415021.5", args{2415021.5}, 1.654992013840947},
		{"2021-02-02,2459247.5", args{2459247.5}, 1.7145352866456673},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DtT(tt.args.y); got != tt.want {
				t.Errorf("DtT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDtCalc(t *testing.T) {
	type args struct {
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"1234.5", args{1234.5}, 640.557637851172},
		{"1950.5", args{1950.5}, 29.2895796875},
		{"2120.5", args{2120.5}, 256.816895},
		{"4000.5", args{4000.5}, 14719.198775},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DtCalc(tt.args.y); got != tt.want {
				t.Errorf("DtCalc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDtExt(t *testing.T) {
	type args struct {
		y   float64
		jsd float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"2021,31", args{2021, 31}, 105.24309999999998},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DtExt(tt.args.y, tt.args.jsd); got != tt.want {
				t.Errorf("DtExt() = %v, want %v", got, tt.want)
			}
		})
	}
}
