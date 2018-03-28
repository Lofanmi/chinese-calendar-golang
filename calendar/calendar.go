package calendar

import (
	"time"
)

type Calendar struct {
	loc *time.Location
	t   time.Time
}

func (this *Calendar) ByTimestamp(ts int64) *Calendar {
	return nil
}

// Current Current
// func (solarterm *Solarterm) Current() *Solarterm {
// 	t := time.Unix(getTimestamp(solarterm.timestampIndex), 0)
// 	after := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, solarterm.loc).Add(24 * time.Hour)
// 	if solarterm.t.Unix() < after.Unix() {
// 		return solarterm
// 	}
// 	return nil
// }
