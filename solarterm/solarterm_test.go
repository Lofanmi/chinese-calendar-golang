package solarterm

import (
	"reflect"
	"testing"
	"time"
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

func minIndex() int64 {
	return 0
}

func maxIndex() int64 {
	return lenTimestamp() - 1
}

func TestNewSolarterm(t *testing.T) {
	type args struct {
		index int64
		loc   *time.Location
	}
	tests := []struct {
		name string
		args args
		want *Solarterm
	}{
		{"nil_min", args{minIndex() - 1, loc()}, nil},
		{"nil_max", args{maxIndex() + 1, loc()}, nil},
		{"test_min", args{minIndex(), loc()}, &Solarterm{loc(), minIndex()}},
		{"test_max", args{maxIndex(), loc()}, &Solarterm{loc(), maxIndex()}},
		{"test", args{100, loc()}, &Solarterm{loc(), 100}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSolarterm(tt.args.index, tt.args.loc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSolarterm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolarterm_Alias(t *testing.T) {
	tests := []struct {
		name      string
		solarterm *Solarterm
		want      string
	}{
		{"test_1", NewSolarterm(minIndex(), loc()), "小寒"},
		{"test_2", NewSolarterm(maxIndex(), loc()), "冬至"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.solarterm.Alias(); got != tt.want {
				t.Errorf("Solarterm.Alias() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolarterm_Timestamp(t *testing.T) {
	tests := []struct {
		name      string
		solarterm *Solarterm
		want      int64
	}{
		{"test_1", NewSolarterm(minIndex(), loc()), getTimestamp(minIndex())},
		{"test_2", NewSolarterm(maxIndex(), loc()), getTimestamp(maxIndex())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.solarterm.Timestamp(); got != tt.want {
				t.Errorf("Solarterm.Timestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolarterm_Time(t *testing.T) {
	tests := []struct {
		name      string
		solarterm *Solarterm
		want      time.Time
	}{
		{"test_1", NewSolarterm(minIndex(), loc()), time.Unix(getTimestamp(minIndex()), 0)},
		{"test_2", NewSolarterm(maxIndex(), loc()), time.Unix(getTimestamp(maxIndex()), 0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.solarterm.Time(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solarterm.Time() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolarterm_Prev(t *testing.T) {
	tests := []struct {
		name      string
		solarterm *Solarterm
		want      *Solarterm
	}{
		{"test_1", NewSolarterm(minIndex(), loc()), nil},
		{"test_2", NewSolarterm(maxIndex(), loc()), NewSolarterm(maxIndex()-1, loc())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.solarterm.Prev(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solarterm.Prev() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolarterm_Next(t *testing.T) {
	tests := []struct {
		name      string
		solarterm *Solarterm
		want      *Solarterm
	}{
		{"test_1", NewSolarterm(minIndex(), loc()), NewSolarterm(minIndex()+1, loc())},
		{"test_2", NewSolarterm(maxIndex(), loc()), nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.solarterm.Next(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solarterm.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolarterm_IsToday(t *testing.T) {
	tests := []struct {
		name      string
		solarterm *Solarterm
		want      bool
	}{
		{"test_1", NewSolarterm(minIndex(), loc()), false},
		{"test_2", NewSolarterm(maxIndex(), loc()), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.solarterm.IsToday(); got != tt.want {
				t.Errorf("Solarterm.IsToday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolarterm_Index(t *testing.T) {
	tests := []struct {
		name      string
		solarterm *Solarterm
		want      int64
	}{
		{"test_1", NewSolarterm(minIndex(), loc()), minIndex()},
		{"test_2", NewSolarterm(maxIndex(), loc()), maxIndex()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.solarterm.Index(); got != tt.want {
				t.Errorf("Solarterm.Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolarterm_Order(t *testing.T) {
	tests := []struct {
		name      string
		solarterm *Solarterm
		want      int64
	}{
		{"test_1", NewSolarterm(2, loc()), 1},
		{"test_2", NewSolarterm(25, loc()), 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.solarterm.Order(); got != tt.want {
				t.Errorf("Solarterm.Order() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpringTimestamp(t *testing.T) {
	type args struct {
		year int64
	}
	tests := []struct {
		name     string
		args     args
		wantTime int64
	}{
		{"zero_1", args{SolartermFromYear - 1}, 0},
		{"zero_2", args{SolartermToYear + 1}, 0},
		{"test_1", args{2017}, 1486136072},
		{"test_2", args{2018}, 1517693315},
		{"test_3", args{2019}, 1549250026},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTime := SpringTimestamp(tt.args.year); gotTime != tt.wantTime {
				t.Errorf("SpringTimestamp() = %v, want %v", gotTime, tt.wantTime)
			}
		})
	}
}

func TestCalcSolarterm(t *testing.T) {
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-02-05 00:00:00", loc())
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-03-21 00:15:26", loc())
	type args struct {
		t   *time.Time
		loc *time.Location
	}
	tests := []struct {
		name  string
		args  args
		wantP *Solarterm
		wantN *Solarterm
	}{
		{"test_1", args{&t1, loc()}, NewSolarterm(2738, loc()), NewSolarterm(2739, loc())},
		{"test_2", args{&t2, loc()}, NewSolarterm(2740, loc()), NewSolarterm(2742, loc())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP, gotN := CalcSolarterm(tt.args.t, tt.args.loc)
			if !reflect.DeepEqual(gotP, tt.wantP) {
				t.Errorf("CalcSolarterm() gotP = %v, want %v", gotP, tt.wantP)
			}
			if !reflect.DeepEqual(gotN, tt.wantN) {
				t.Errorf("CalcSolarterm() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestSolarterm_IsInDay(t *testing.T) {
	now := time.Now()
	type args struct {
		t *time.Time
	}
	tests := []struct {
		name      string
		solarterm *Solarterm
		args      args
		want      bool
	}{
		{"test_1", NewSolarterm(minIndex(), loc()), args{&now}, false},
		{"test_2", NewSolarterm(maxIndex(), loc()), args{&now}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.solarterm.IsInDay(tt.args.t); got != tt.want {
				t.Errorf("Solarterm.IsInDay() = %v, want %v", got, tt.want)
			}
		})
	}
}
