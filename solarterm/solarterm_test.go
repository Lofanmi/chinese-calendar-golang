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
