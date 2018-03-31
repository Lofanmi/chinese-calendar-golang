package ganzhi

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/Lofanmi/chinese-calendar-golang/animal"
	"github.com/Lofanmi/chinese-calendar-golang/gan"
	"github.com/Lofanmi/chinese-calendar-golang/solarterm"
	"github.com/Lofanmi/chinese-calendar-golang/zhi"
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
func TestNewGanzhi(t *testing.T) {
	t1 := time.Date(solarterm.SolartermFromYear-1, 6, 1, 0, 0, 0, 0, loc())
	t2 := time.Date(solarterm.SolartermToYear+1, 6, 1, 0, 0, 0, 0, loc())
	t3 := time.Date(2018, 1, 1, 0, 0, 0, 0, loc())
	t4 := time.Date(2018, 2, 4, 5, 28, 34, 0, loc())
	t5 := time.Date(2018, 2, 4, 5, 28, 35, 0, loc())
	t6 := time.Date(2018, 2, 4, 5, 28, 36, 0, loc())
	maker := func(t *time.Time, yg, yz, mg, mz, dg, dz, hg, hz, p, n int64) *Ganzhi {
		return &Ganzhi{
			loc:           loc(),
			t:             t,
			YearGan:       gan.NewGan(yg),
			YearZhi:       zhi.NewZhi(yz),
			MonthGan:      gan.NewGan(mg),
			MonthZhi:      zhi.NewZhi(mz),
			DayGan:        gan.NewGan(dg),
			DayZhi:        zhi.NewZhi(dz),
			HourGan:       gan.NewGan(hg),
			HourZhi:       zhi.NewZhi(hz),
			PrevSolarterm: solarterm.NewSolarterm(p, loc()),
			NextSolarterm: solarterm.NewSolarterm(n, loc()),
		}
	}
	type args struct {
		t   *time.Time
		loc *time.Location
	}
	tests := []struct {
		name string
		args args
		want *Ganzhi
	}{
		{"test_1", args{&t1, loc()}, nil},
		{"test_2", args{&t2, loc()}, nil},
		{"test_3", args{&t3, loc()}, maker(&t3, 4, 10, 9, 1, 10, 6, 9, 1, 23, 0)},
		{"test_4", args{&t4, loc()}, maker(&t4, 4, 10, 10, 2, 4, 4, 10, 4, 1, 2)},
		{"test_5", args{&t5, loc()}, maker(&t5, 5, 11, 1, 3, 4, 4, 10, 4, 1, 3)},
		{"test_6", args{&t6, loc()}, maker(&t6, 5, 11, 1, 3, 4, 4, 10, 4, 2, 3)},
	}

	equals := func(a, b *Ganzhi) bool {
		if a == nil {
			return b == nil
		}
		return a.YearGan.Order() == b.YearGan.Order() &&
			a.YearZhi.Order() == b.YearZhi.Order() &&
			a.MonthGan.Order() == b.MonthGan.Order() &&
			a.MonthZhi.Order() == b.MonthZhi.Order() &&
			a.DayGan.Order() == b.DayGan.Order() &&
			a.DayZhi.Order() == b.DayZhi.Order() &&
			a.HourGan.Order() == b.HourGan.Order() &&
			a.HourZhi.Order() == b.HourZhi.Order() &&
			a.PrevSolarterm.Order() == b.PrevSolarterm.Order() &&
			a.NextSolarterm.Order() == b.NextSolarterm.Order()
	}

	errfunc := func(a, b *Ganzhi) (s1, s2 string) {
		if a == nil {
			s1 = "[nil]"
		} else {
			s1 = fmt.Sprintf("[%s%s %s%s %s%s %s%s, %s %s]",
				a.YearGan.Alias(),
				a.YearZhi.Alias(),
				a.MonthGan.Alias(),
				a.MonthZhi.Alias(),
				a.DayGan.Alias(),
				a.DayZhi.Alias(),
				a.HourGan.Alias(),
				a.HourZhi.Alias(),
				a.PrevSolarterm.Alias(),
				a.NextSolarterm.Alias(),
			)
		}
		s2 = fmt.Sprintf("[%s%s %s%s %s%s %s%s, %s %s]",
			b.YearGan.Alias(),
			b.YearZhi.Alias(),
			b.MonthGan.Alias(),
			b.MonthZhi.Alias(),
			b.DayGan.Alias(),
			b.DayZhi.Alias(),
			b.HourGan.Alias(),
			b.HourZhi.Alias(),
			b.PrevSolarterm.Alias(),
			b.NextSolarterm.Alias(),
		)
		return
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGanzhi(tt.args.t, tt.args.loc); !equals(got, tt.want) {
				s1, s2 := errfunc(got, tt.want)
				t.Errorf("NewGanzhi() = %s, want %s", s1, s2)
			}
		})
	}
}

func TestGanzhi_Animal(t *testing.T) {
	t1 := time.Date(2018, 1, 1, 0, 0, 0, 0, loc())
	t2 := time.Date(2018, 2, 5, 0, 0, 0, 0, loc())
	tests := []struct {
		name string
		gz   *Ganzhi
		want *animal.Animal
	}{
		{"test_1", NewGanzhi(&t1, loc()), animal.NewAnimal(10)},
		{"test_2", NewGanzhi(&t2, loc()), animal.NewAnimal(11)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.gz.Animal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ganzhi.Animal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGanzhi_YearGanzhiAlias(t *testing.T) {
	t1 := time.Date(2018, 1, 1, 0, 0, 0, 0, loc())
	t2 := time.Date(2018, 2, 5, 0, 0, 0, 0, loc())
	tests := []struct {
		name string
		gz   *Ganzhi
		want string
	}{
		{"test_1", NewGanzhi(&t1, loc()), "丁酉"},
		{"test_2", NewGanzhi(&t2, loc()), "戊戌"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.gz.YearGanzhiAlias(); got != tt.want {
				t.Errorf("Ganzhi.YearGanzhiAlias() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGanzhi_MonthGanzhiAlias(t *testing.T) {
	t1 := time.Date(2018, 1, 1, 0, 0, 0, 0, loc())
	t2 := time.Date(2018, 2, 5, 0, 0, 0, 0, loc())
	tests := []struct {
		name string
		gz   *Ganzhi
		want string
	}{
		{"test_1", NewGanzhi(&t1, loc()), "壬子"},
		{"test_2", NewGanzhi(&t2, loc()), "甲寅"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.gz.MonthGanzhiAlias(); got != tt.want {
				t.Errorf("Ganzhi.MonthGanzhiAlias() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGanzhi_DayGanzhiAlias(t *testing.T) {
	t1 := time.Date(2018, 1, 1, 0, 0, 0, 0, loc())
	t2 := time.Date(2018, 2, 5, 0, 0, 0, 0, loc())
	tests := []struct {
		name string
		gz   *Ganzhi
		want string
	}{
		{"test_1", NewGanzhi(&t1, loc()), "癸巳"},
		{"test_2", NewGanzhi(&t2, loc()), "戊辰"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.gz.DayGanzhiAlias(); got != tt.want {
				t.Errorf("Ganzhi.DayGanzhiAlias() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGanzhi_HourGanzhiAlias(t *testing.T) {
	t1 := time.Date(2018, 1, 1, 0, 0, 0, 0, loc())
	t2 := time.Date(2018, 2, 5, 0, 0, 0, 0, loc())
	tests := []struct {
		name string
		gz   *Ganzhi
		want string
	}{
		{"test_1", NewGanzhi(&t1, loc()), "壬子"},
		{"test_2", NewGanzhi(&t2, loc()), "壬子"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.gz.HourGanzhiAlias(); got != tt.want {
				t.Errorf("Ganzhi.HourGanzhiAlias() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGanzhi_YearGanzhiOrder(t *testing.T) {
	t1 := time.Date(2018, 1, 1, 0, 0, 0, 0, loc())
	t2 := time.Date(2018, 2, 5, 0, 0, 0, 0, loc())
	tests := []struct {
		name string
		gz   *Ganzhi
		want int64
	}{
		{"test_1", NewGanzhi(&t1, loc()), 34},
		{"test_2", NewGanzhi(&t2, loc()), 35},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.gz.YearGanzhiOrder(); got != tt.want {
				t.Errorf("Ganzhi.YearGanzhiOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGanzhi_MonthGanzhiOrder(t *testing.T) {
	t1 := time.Date(2018, 1, 1, 0, 0, 0, 0, loc())
	t2 := time.Date(2018, 2, 5, 0, 0, 0, 0, loc())
	tests := []struct {
		name string
		gz   *Ganzhi
		want int64
	}{
		{"test_1", NewGanzhi(&t1, loc()), 49},
		{"test_2", NewGanzhi(&t2, loc()), 51},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.gz.MonthGanzhiOrder(); got != tt.want {
				t.Errorf("Ganzhi.MonthGanzhiOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGanzhi_DayGanzhiOrder(t *testing.T) {
	t1 := time.Date(2018, 1, 1, 0, 0, 0, 0, loc())
	t2 := time.Date(2018, 2, 5, 0, 0, 0, 0, loc())
	tests := []struct {
		name string
		gz   *Ganzhi
		want int64
	}{
		{"test_1", NewGanzhi(&t1, loc()), 30},
		{"test_2", NewGanzhi(&t2, loc()), 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.gz.DayGanzhiOrder(); got != tt.want {
				t.Errorf("Ganzhi.DayGanzhiOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGanzhi_HourGanzhiOrder(t *testing.T) {
	t1 := time.Date(2018, 1, 1, 0, 0, 0, 0, loc())
	t2 := time.Date(2018, 2, 5, 0, 0, 0, 0, loc())
	tests := []struct {
		name string
		gz   *Ganzhi
		want int64
	}{
		{"test_1", NewGanzhi(&t1, loc()), 49},
		{"test_2", NewGanzhi(&t2, loc()), 49},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.gz.HourGanzhiOrder(); got != tt.want {
				t.Errorf("Ganzhi.HourGanzhiOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
