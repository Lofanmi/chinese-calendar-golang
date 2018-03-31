package lunar

import (
	"reflect"
	"testing"
	"time"

	"github.com/Lofanmi/chinese-calendar-golang/animal"
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

func TestFromSolarTimestamp(t *testing.T) {
	type args struct {
		ts  int64
		loc *time.Location
	}
	tests := []struct {
		name                 string
		args                 args
		wantLunarYear        int64
		wantLunarMonth       int64
		wantLunarDay         int64
		wantLunarMonthIsLeap bool
	}{
		{"test_1", args{1502769600, loc()}, 2017, 6, 24, true},
		{"test_2", args{1522422690, loc()}, 2018, 2, 14, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLunarYear, gotLunarMonth, gotLunarDay, gotLunarMonthIsLeap := FromSolarTimestamp(tt.args.ts, tt.args.loc)
			if gotLunarYear != tt.wantLunarYear {
				t.Errorf("FromSolarTimestamp() gotLunarYear = %v, want %v", gotLunarYear, tt.wantLunarYear)
			}
			if gotLunarMonth != tt.wantLunarMonth {
				t.Errorf("FromSolarTimestamp() gotLunarMonth = %v, want %v", gotLunarMonth, tt.wantLunarMonth)
			}
			if gotLunarDay != tt.wantLunarDay {
				t.Errorf("FromSolarTimestamp() gotLunarDay = %v, want %v", gotLunarDay, tt.wantLunarDay)
			}
			if gotLunarMonthIsLeap != tt.wantLunarMonthIsLeap {
				t.Errorf("FromSolarTimestamp() gotLunarMonthIsLeap = %v, want %v", gotLunarMonthIsLeap, tt.wantLunarMonthIsLeap)
			}
		})
	}
}

func TestToSolarTimestamp(t *testing.T) {
	type args struct {
		year        int64
		month       int64
		day         int64
		hour        int64
		minute      int64
		second      int64
		isLeapMonth bool
		loc         *time.Location
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"test_1", args{2017, 6, 24, 12, 0, 0, true, loc()}, 1502769600},
		{"test_2", args{2018, 2, 14, 23, 11, 30, true, loc()}, 1522422690},
		{"test_3", args{2018, 2, 14, 23, 11, 30, false, loc()}, 1522422690},
		{"test_4", args{1900, 1, 14, 23, 11, 30, false, loc()}, 0},
		{"test_5", args{2100, 12, 14, 23, 11, 30, false, loc()}, 0},
		{"test_6", args{1900 - 1, 1, 14, 23, 11, 30, false, loc()}, 0},
		{"test_7", args{2100 + 1, 12, 14, 23, 11, 30, false, loc()}, 0},
		{"test_8", args{1900, 2, 100000, 23, 11, 30, false, loc()}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToSolarTimestamp(tt.args.year, tt.args.month, tt.args.day, tt.args.hour, tt.args.minute, tt.args.second, tt.args.isLeapMonth, tt.args.loc); got != tt.want {
				t.Errorf("ToSolarTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewLunar(t *testing.T) {
	t1 := time.Date(2017, 8, 15, 12, 0, 0, 0, loc())
	t2 := time.Date(2018, 3, 30, 23, 11, 30, 0, loc())
	type args struct {
		t   *time.Time
		loc *time.Location
	}
	tests := []struct {
		name string
		args args
		want *Lunar
	}{
		{"test_1", args{&t1, loc()}, &Lunar{
			loc:         loc(),
			t:           &t1,
			year:        2017,
			month:       6,
			day:         24,
			monthIsLeap: true,
		}},
		{"test_2", args{&t2, loc()}, &Lunar{
			loc:         loc(),
			t:           &t2,
			year:        2018,
			month:       2,
			day:         14,
			monthIsLeap: false,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLunar(tt.args.t, tt.args.loc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLunar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLunar_LeapMonth(t *testing.T) {
	t1 := time.Date(2018, 6, 1, 0, 0, 0, 0, loc())
	t2 := time.Date(2017, 6, 1, 0, 0, 0, 0, loc())
	tests := []struct {
		name  string
		lunar *Lunar
		want  int64
	}{
		{"test_1", NewLunar(&t1, loc()), 0},
		{"test_2", NewLunar(&t2, loc()), 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lunar.LeapMonth(); got != tt.want {
				t.Errorf("Lunar.LeapMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLunar_IsLeap(t *testing.T) {
	t1 := time.Date(2018, 6, 1, 0, 0, 0, 0, loc())
	t2 := time.Date(2017, 6, 1, 0, 0, 0, 0, loc())
	tests := []struct {
		name  string
		lunar *Lunar
		want  bool
	}{
		{"test_1", NewLunar(&t1, loc()), false},
		{"test_2", NewLunar(&t2, loc()), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lunar.IsLeap(); got != tt.want {
				t.Errorf("Lunar.IsLeap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLunar_IsLeapMonth(t *testing.T) {
	t1 := time.Date(2018, 5, 1, 0, 0, 0, 0, loc())
	t2 := time.Date(2017, 6, 15, 0, 0, 0, 0, loc())
	t3 := time.Date(2017, 8, 15, 0, 0, 0, 0, loc())
	tests := []struct {
		name  string
		lunar *Lunar
		want  bool
	}{
		{"test_1", NewLunar(&t1, loc()), false},
		{"test_2", NewLunar(&t2, loc()), false},
		{"test_3", NewLunar(&t3, loc()), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lunar.IsLeapMonth(); got != tt.want {
				t.Errorf("Lunar.IsLeapMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLunar_Animal(t *testing.T) {
	t1 := time.Date(2018, 5, 1, 0, 0, 0, 0, loc())
	t2 := time.Date(2017, 6, 15, 0, 0, 0, 0, loc())
	t3 := time.Date(2017, 8, 15, 0, 0, 0, 0, loc())

	tests := []struct {
		name  string
		lunar *Lunar
		want  *animal.Animal
	}{
		{"test_1", NewLunar(&t1, loc()), animal.NewAnimal(11 - 1)},
		{"test_2", NewLunar(&t2, loc()), animal.NewAnimal(10 - 1)},
		{"test_3", NewLunar(&t3, loc()), animal.NewAnimal(10 - 1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lunar.Animal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lunar.Animal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLunar_YearAlias(t *testing.T) {
	t1 := time.Date(2018, 5, 1, 0, 0, 0, 0, loc())
	t2 := time.Date(2017, 6, 15, 0, 0, 0, 0, loc())
	t3 := time.Date(2017, 8, 15, 0, 0, 0, 0, loc())
	tests := []struct {
		name  string
		lunar *Lunar
		want  string
	}{
		{"test_1", NewLunar(&t1, loc()), "二零一八"},
		{"test_2", NewLunar(&t2, loc()), "二零一七"},
		{"test_3", NewLunar(&t3, loc()), "二零一七"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lunar.YearAlias(); got != tt.want {
				t.Errorf("Lunar.YearAlias() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLunar_MonthAlias(t *testing.T) {
	t1 := time.Date(2018, 5, 1, 0, 0, 0, 0, loc())
	t2 := time.Date(2017, 6, 15, 0, 0, 0, 0, loc())
	t3 := time.Date(2017, 8, 15, 0, 0, 0, 0, loc())
	tests := []struct {
		name  string
		lunar *Lunar
		want  string
	}{
		{"test_1", NewLunar(&t1, loc()), "三月"},
		{"test_2", NewLunar(&t2, loc()), "五月"},
		{"test_3", NewLunar(&t3, loc()), "闰六月"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lunar.MonthAlias(); got != tt.want {
				t.Errorf("Lunar.MonthAlias() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLunar_DayAlias(t *testing.T) {
	t1 := time.Date(2018, 5, 1, 0, 0, 0, 0, loc())
	t2 := time.Date(2017, 6, 4, 0, 0, 0, 0, loc())
	t3 := time.Date(2017, 6, 14, 0, 0, 0, 0, loc())
	t4 := time.Date(2017, 8, 21, 0, 0, 0, 0, loc())
	tests := []struct {
		name      string
		lunar     *Lunar
		wantAlias string
	}{
		{"test_1", NewLunar(&t1, loc()), "十六"},
		{"test_2", NewLunar(&t2, loc()), "初十"},
		{"test_3", NewLunar(&t3, loc()), "二十"},
		{"test_33", NewLunar(&t4, loc()), "三十"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAlias := tt.lunar.DayAlias(); gotAlias != tt.wantAlias {
				t.Errorf("Lunar.DayAlias() = %v, want %v", gotAlias, tt.wantAlias)
			}
		})
	}
}

func Test_lunarDays(t *testing.T) {
	type args struct {
		year  int64
		month int64
	}
	tests := []struct {
		name     string
		args     args
		wantDays int64
	}{
		{"test_1", args{2018, 1 - 1}, 0},
		{"test_2", args{2018, 12 + 1}, 0},
		{"test_3", args{2017, 6}, 29},
		{"test_3", args{2017, 8}, 30},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDays := lunarDays(tt.args.year, tt.args.month); gotDays != tt.wantDays {
				t.Errorf("lunarDays() = %v, want %v", gotDays, tt.wantDays)
			}
		})
	}
}
