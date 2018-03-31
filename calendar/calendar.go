package calendar

import (
	"encoding/json"
	"time"

	"github.com/Lofanmi/chinese-calendar-golang/ganzhi"
	"github.com/Lofanmi/chinese-calendar-golang/lunar"
	"github.com/Lofanmi/chinese-calendar-golang/solar"
)

// Calendar Calendar
type Calendar struct {
	loc    *time.Location
	t      *time.Time
	Solar  *solar.Solar
	Lunar  *lunar.Lunar
	Ganzhi *ganzhi.Ganzhi
}

var location *time.Location

// SetLocation SetLocation
func SetLocation(loc *time.Location) {
	location = loc
}

func loc() *time.Location {
	if location == nil {
		location, _ = time.LoadLocation("PRC")
	}
	return location
}

// ByTimestamp ByTimestamp
func ByTimestamp(ts int64) *Calendar {
	l := loc()
	t := time.Unix(ts, 0)
	sc := solar.NewSolar(&t, l)
	lc := lunar.NewLunar(&t, l)
	gz := ganzhi.NewGanzhi(&t, l)
	return &Calendar{
		loc:    l,
		t:      &t,
		Solar:  sc,
		Lunar:  lc,
		Ganzhi: gz,
	}
}

// BySolar BySolar
func BySolar(year, month, day, hour, minute, second int64) *Calendar {
	t := time.Date(int(year),
		time.Month(month),
		int(day),
		int(hour),
		int(minute),
		int(second),
		0,
		loc(),
	)
	return ByTimestamp(t.Unix())
}

// ByLunar ByLunar
func ByLunar(year, month, day, hour, minute, second int64, isLeapMonth bool) *Calendar {
	ts := lunar.ToSolarTimestamp(year, month, day, hour, minute, second, isLeapMonth, loc())
	return ByTimestamp(ts)
}

// ToJSON ToJSON
func (calendar *Calendar) ToJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m1 := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m3 := make(map[string]interface{})

	m1["year"] = calendar.Solar.GetYear()
	m1["month"] = calendar.Solar.GetMonth()
	m1["day"] = calendar.Solar.GetDay()
	m1["hour"] = calendar.Solar.GetHour()
	m1["minute"] = calendar.Solar.GetMinute()
	m1["second"] = calendar.Solar.GetSecond()
	m1["nanosecond"] = calendar.Solar.GetNanosecond()
	m1["is_leep"] = calendar.Solar.IsLeep()
	m1["week_number"] = calendar.Solar.WeekNumber()
	m1["week_alias"] = calendar.Solar.WeekAlias()
	m1["animal"] = calendar.Solar.Animal().Alias()
	m1["constellation"] = calendar.Solar.Constellation().Alias()

	m2["year"] = calendar.Lunar.GetYear()
	m2["month"] = calendar.Lunar.GetMonth()
	m2["day"] = calendar.Lunar.GetDay()
	m2["is_leap_month"] = calendar.Lunar.IsLeapMonth()
	m2["is_leap"] = calendar.Lunar.IsLeap()
	m2["leap_month"] = calendar.Lunar.LeapMonth()
	m2["is_leap_month"] = calendar.Lunar.IsLeapMonth()
	m2["animal"] = calendar.Lunar.Animal().Alias()
	m2["year_alias"] = calendar.Lunar.YearAlias()
	m2["month_alias"] = calendar.Lunar.MonthAlias()
	m2["day_alias"] = calendar.Lunar.DayAlias()

	m3["animal"] = calendar.Ganzhi.Animal().Alias()
	m3["year"] = calendar.Ganzhi.YearGanzhiAlias()
	m3["month"] = calendar.Ganzhi.MonthGanzhiAlias()
	m3["day"] = calendar.Ganzhi.DayGanzhiAlias()
	m3["hour"] = calendar.Ganzhi.HourGanzhiAlias()
	m3["year_order"] = calendar.Ganzhi.YearGanzhiOrder()
	m3["month_order"] = calendar.Ganzhi.MonthGanzhiOrder()
	m3["day_order"] = calendar.Ganzhi.DayGanzhiOrder()
	m3["hour_order"] = calendar.Ganzhi.HourGanzhiOrder()

	m["solar"] = m1
	m["lunar"] = m2
	m["ganzhi"] = m3

	return json.Marshal(m)
}
