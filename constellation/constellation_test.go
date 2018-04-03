package constellation

import (
	"testing"
	"time"
)

func loc() *time.Location {
	loc, _ := time.LoadLocation("PRC")
	return loc
}
func now() *time.Time {
	t := time.Now()
	return &t
}

func testTime1() *time.Time {
	t := time.Date(2018, 1, 1, 0, 0, 0, 0, loc())
	return &t
}

func testTime2() *time.Time {
	t := time.Date(2018, 12, 1, 0, 0, 0, 0, loc())
	return &t
}

func TestNewConstellation(t *testing.T) {
	type args struct {
		t *time.Time
	}
	tests := []struct {
		name string
		args args
		want *Constellation
	}{
		{"test", args{now()}, &Constellation{now()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConstellation(tt.args.t); got.Alias() != tt.want.Alias() {
				t.Errorf("NewConstellation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstellation_Alias(t *testing.T) {
	tests := []struct {
		name          string
		constellation *Constellation
		want          string
	}{
		{"test_1", NewConstellation(testTime1()), "水瓶"},
		{"test_2", NewConstellation(testTime2()), "魔羯"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.constellation.Alias(); got != tt.want {
				t.Errorf("Constellation.Alias() = %v, want %v", got, tt.want)
			}
		})
	}
}
