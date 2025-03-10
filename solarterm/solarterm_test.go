package solarterm

import (
	"reflect"
	"testing"
	"time"
)

func minIndex() int64 {
	return 0
}

func maxIndex() int64 {
	return lenJ2000() - 1
}

func TestNewSolarterm(t *testing.T) {
	type args struct {
		index int64
	}
	tests := []struct {
		name string
		args args
		want *Solarterm
	}{
		{"nil_min", args{minIndex() - 1}, nil},
		{"nil_max", args{maxIndex() + 1}, nil},
		{"test_min", args{minIndex()}, &Solarterm{minIndex()}},
		{"test_max", args{maxIndex()}, &Solarterm{maxIndex()}},
		{"test", args{100}, &Solarterm{100}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSolarterm(tt.args.index); !reflect.DeepEqual(got, tt.want) {
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
		{"test_1", NewSolarterm(minIndex()), "小寒"},
		{"test_2", NewSolarterm(maxIndex()), "冬至"},
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
		{"test_1", NewSolarterm(minIndex()), getTimestamp(minIndex())},
		{"test_2", NewSolarterm(maxIndex()), getTimestamp(maxIndex())},
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
		{"test_1", NewSolarterm(minIndex()), time.Unix(getTimestamp(minIndex()), 0)},
		{"test_2", NewSolarterm(maxIndex()), time.Unix(getTimestamp(maxIndex()), 0)},
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
		{"test_1", NewSolarterm(minIndex()), nil},
		{"test_2", NewSolarterm(maxIndex()), NewSolarterm(maxIndex() - 1)},
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
		{"test_1", NewSolarterm(minIndex()), NewSolarterm(minIndex() + 1)},
		{"test_2", NewSolarterm(maxIndex()), nil},
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
		{"test_1", NewSolarterm(minIndex()), false},
		{"test_2", NewSolarterm(maxIndex()), false},
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
		{"test_1", NewSolarterm(minIndex()), minIndex()},
		{"test_2", NewSolarterm(maxIndex()), maxIndex()},
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
		{"test_1", NewSolarterm(2), 1},
		{"test_2", NewSolarterm(25), 24},
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
		{"test_1", args{2017}, 1486136044},
		{"test_2", args{2018}, 1517693305},
		{"test_3", args{2019}, 1549250054},
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
	t1 := time.Date(2018, 2, 5, 0, 0, 0, 0, time.Local)
	t2 := time.Date(2018, 3, 21, 0, 15, 28, 0, time.Local)
	type args struct {
		t *time.Time
	}
	tests := []struct {
		name  string
		args  args
		wantP *Solarterm
		wantN *Solarterm
	}{
		{"test_1", args{&t1}, NewSolarterm(2738), NewSolarterm(2739)},
		{"test_2", args{&t2}, NewSolarterm(2741), NewSolarterm(2742)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP, gotN := CalcSolarterm(tt.args.t)
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
		{"test_1", NewSolarterm(minIndex()), args{&now}, false},
		{"test_2", NewSolarterm(maxIndex()), args{&now}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.solarterm.IsInDay(tt.args.t); got != tt.want {
				t.Errorf("Solarterm.IsInDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolarterm_Equals(t *testing.T) {
	tests := []struct {
		name       string
		solarterm  *Solarterm
		solarterm2 *Solarterm
		want       bool
	}{
		{"test_1", NewSolarterm(minIndex()), NewSolarterm(minIndex()), true},
		{"test_2", NewSolarterm(maxIndex()), NewSolarterm(minIndex()), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.solarterm.Equals(tt.solarterm2) != tt.want {
				t.Errorf("Solarterm.Equals() failed")
			}
		})
	}
}
