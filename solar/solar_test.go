package solar

import (
	"reflect"
	"testing"
	"time"

	"github.com/Lofanmi/chinese-calendar-golang/animal"
	"github.com/Lofanmi/chinese-calendar-golang/constellation"
	"github.com/Lofanmi/chinese-calendar-golang/solarterm"
)

var (
	defaultLoc *time.Location
)

func loc() *time.Location {
	if defaultLoc == nil {
		defaultLoc, _ = time.LoadLocation("PRC")
	}
	return defaultLoc
}

func TestNewSolar(t *testing.T) {
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-03-21 00:00:26", loc())
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-03-21 00:15:26", loc())
	t3, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-03-21 00:15:27", loc())
	t4, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-03-21 00:16:26", loc())
	t5, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-04-01 00:00:00", loc())
	type args struct {
		t   *time.Time
		loc *time.Location
	}
	tests := []struct {
		name string
		args args
		want *Solar
	}{
		{"test_1", args{&t1, loc()}, &Solar{
			loc:              loc(),
			t:                &t1,
			CurrentSolarterm: solarterm.NewSolarterm(2741, loc()),
			PrevSolarterm:    solarterm.NewSolarterm(2740, loc()),
			NextSolarterm:    solarterm.NewSolarterm(2742, loc()),
		}},
		{"test_2", args{&t2, loc()}, &Solar{
			loc:              loc(),
			t:                &t2,
			CurrentSolarterm: solarterm.NewSolarterm(2741, loc()),
			PrevSolarterm:    solarterm.NewSolarterm(2740, loc()),
			NextSolarterm:    solarterm.NewSolarterm(2742, loc()),
		}},
		{"test_3", args{&t3, loc()}, &Solar{
			loc:              loc(),
			t:                &t3,
			CurrentSolarterm: solarterm.NewSolarterm(2741, loc()),
			PrevSolarterm:    solarterm.NewSolarterm(2740, loc()),
			NextSolarterm:    solarterm.NewSolarterm(2742, loc()),
		}},
		{"test_4", args{&t4, loc()}, &Solar{
			loc:              loc(),
			t:                &t4,
			CurrentSolarterm: solarterm.NewSolarterm(2741, loc()),
			PrevSolarterm:    solarterm.NewSolarterm(2740, loc()),
			NextSolarterm:    solarterm.NewSolarterm(2742, loc()),
		}},
		{"test_5", args{&t5, loc()}, &Solar{
			loc:              loc(),
			t:                &t5,
			CurrentSolarterm: nil,
			PrevSolarterm:    solarterm.NewSolarterm(2741, loc()),
			NextSolarterm:    solarterm.NewSolarterm(2742, loc()),
		}},
	}
	equal := func(a, b *Solar) (result bool) {
		if a.t.Nanosecond() != b.t.Nanosecond() {
			result = false
		} else if a.CurrentSolarterm != nil && b.CurrentSolarterm != nil && a.CurrentSolarterm.Index() != b.CurrentSolarterm.Index() {
			result = false
		} else if a.PrevSolarterm != nil && b.PrevSolarterm != nil && a.PrevSolarterm.Index() != b.PrevSolarterm.Index() {
			result = false
		} else if a.NextSolarterm != nil && b.NextSolarterm != nil && a.NextSolarterm.Index() != b.NextSolarterm.Index() {
			result = false
		} else {
			result = true
		}
		return
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSolar(tt.args.t, tt.args.loc); !equal(got, tt.want) {
				t.Errorf("NewSolar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolar_IsLeep(t *testing.T) {
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-03-21 00:00:00", loc())
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-03-21 00:00:00", loc())
	tests := []struct {
		name  string
		solar *Solar
		want  bool
	}{
		{"test_2018", NewSolar(&t1, loc()), false},
		{"test_2020", NewSolar(&t2, loc()), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.solar.IsLeep(); got != tt.want {
				t.Errorf("Solar.IsLeep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolar_WeekNumber(t *testing.T) {
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-03-21 00:00:00", loc())
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-03-25 00:01:00", loc())
	tests := []struct {
		name  string
		solar *Solar
		want  int64
	}{
		{"test_1", NewSolar(&t1, loc()), 3},
		{"test_2", NewSolar(&t2, loc()), 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.solar.WeekNumber(); got != tt.want {
				t.Errorf("Solar.WeekNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolar_WeekAlias(t *testing.T) {
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-03-21 00:00:00", loc())
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-03-25 00:01:00", loc())
	tests := []struct {
		name  string
		solar *Solar
		want  string
	}{
		{"test_1", NewSolar(&t1, loc()), "三"},
		{"test_2", NewSolar(&t2, loc()), "日"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.solar.WeekAlias(); got != tt.want {
				t.Errorf("Solar.WeekAlias() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolar_Animal(t *testing.T) {
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-03-21 00:00:00", loc())
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-03-21 00:00:00", loc())
	tests := []struct {
		name  string
		solar *Solar
		want  *animal.Animal
	}{
		{"test_1", NewSolar(&t1, loc()), animal.NewAnimal(11)},
		{"test_1", NewSolar(&t2, loc()), animal.NewAnimal(12)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.solar.Animal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solar.Animal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolar_Constellation(t *testing.T) {
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-03-21 00:00:00", loc())
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-11-21 00:00:00", loc())
	tests := []struct {
		name  string
		solar *Solar
		want  *constellation.Constellation
	}{
		{"test_1", NewSolar(&t1, loc()), constellation.NewConstellation(&t1)},
		{"test_1", NewSolar(&t2, loc()), constellation.NewConstellation(&t2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.solar.Constellation(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solar.Constellation() = %v, want %v", got, tt.want)
			}
		})
	}
}
