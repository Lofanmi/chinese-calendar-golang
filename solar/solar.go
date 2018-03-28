package solar

import (
	"time"

	"github.com/Lofanmi/chinese-calendar-golang/animal"
	"github.com/Lofanmi/chinese-calendar-golang/constellation"
	"github.com/Lofanmi/chinese-calendar-golang/solarterm"
	"github.com/Lofanmi/chinese-calendar-golang/utils"
)

// Solar Solar
type Solar struct {
	loc              *time.Location
	t                *time.Time
	CurrentSolarterm *solarterm.Solarterm
	PrevSolarterm    *solarterm.Solarterm
	NextSolarterm    *solarterm.Solarterm
}

var weekAlias = [...]string{
	"日", "一", "二", "三", "四", "五", "六",
}

// NewSolar NewSolar
func NewSolar(t *time.Time, loc *time.Location) *Solar {
	var c *solarterm.Solarterm
	p, n := solarterm.CalcSolarterm(t, loc)
	if n.Index()-p.Index() == 1 {
		if p.IsInDay(t) {
			c = p
			p = p.Prev()
		}
		if n.IsInDay(t) {
			c = n
			p = c.Prev()
			n = c.Next()
		}
	}
	return &Solar{
		loc:              loc,
		t:                t,
		CurrentSolarterm: c,
		PrevSolarterm:    p,
		NextSolarterm:    n,
	}
}

// IsLeep IsLeep
func (solar *Solar) IsLeep() bool {
	year := solar.t.Year()
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}

// WeekNumber WeekNumber
func (solar *Solar) WeekNumber() int64 {
	return int64(solar.t.Weekday())
}

// WeekAlias WeekAlias
func (solar *Solar) WeekAlias() string {
	return weekAlias[solar.WeekNumber()]
}

// Animal Animal
func (solar *Solar) Animal() *animal.Animal {
	return animal.NewAnimal(utils.OrderMod(int64(solar.t.Year()-3), 12))
}

// Constellation Constellation
func (solar *Solar) Constellation() *constellation.Constellation {
	return constellation.NewConstellation(solar.t)
}
