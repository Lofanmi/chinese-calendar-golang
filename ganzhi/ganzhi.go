package ganzhi

import (
	"time"

	"github.com/Lofanmi/chinese-calendar-golang/animal"
	"github.com/Lofanmi/chinese-calendar-golang/gan"
	"github.com/Lofanmi/chinese-calendar-golang/solarterm"
	"github.com/Lofanmi/chinese-calendar-golang/utils"
	"github.com/Lofanmi/chinese-calendar-golang/zhi"
)

// Ganzhi Ganzhi
type Ganzhi struct {
	loc           *time.Location
	t             *time.Time
	YearGan       *gan.Gan
	YearZhi       *zhi.Zhi
	MonthGan      *gan.Gan
	MonthZhi      *zhi.Zhi
	DayGan        *gan.Gan
	DayZhi        *zhi.Zhi
	HourGan       *gan.Gan
	HourZhi       *zhi.Zhi
	PrevSolarterm *solarterm.Solarterm
	NextSolarterm *solarterm.Solarterm
}

// NewGanzhi NewGanzhi
func NewGanzhi(t *time.Time, loc *time.Location) *Ganzhi {
	year := int64(t.Year())
	if !isSupported(year) {
		return nil
	}
	if year < solarterm.SpringTimestamp(year) {
		year--
	}
	yearGan := gan.NewGan(utils.OrderMod(year-3, 10))
	yearZhi := zhi.NewZhi(utils.OrderMod(year-3, 12))

	p, n := solarterm.CalcSolarterm(t, loc)

	i := p.Index()
	if n.Index()-p.Index() == 2 {
		i++
	}
	i = utils.OrderMod(int64((i%24)/2), 12)

	monthZhi := zhi.NewZhi(utils.OrderMod(i+2, 12))
	monthGan := gan.NewGan(utils.OrderMod(i+yearGan.Order()*2, 10))

	begin := time.Date(solarterm.SolartermFromYear, 1, 1, 0, 0, 0, 0, loc)
	seconds := t.Sub(begin).Seconds()
	dayOrder := utils.OrderMod(int64(seconds/86400)+31, 60)

	dayGan := gan.NewGan(utils.OrderMod(dayOrder, 10))
	dayZhi := zhi.NewZhi(utils.OrderMod(dayOrder, 12))

	hourZhi := zhi.NewZhi(utils.OrderMod(int64(((t.Hour()+1)/2)+1), 12))
	hourGan := gan.NewGan(utils.OrderMod(hourZhi.Order()-2+dayGan.Order()*2, 10))

	return &Ganzhi{
		loc:           loc,
		t:             t,
		YearGan:       yearGan,
		YearZhi:       yearZhi,
		MonthGan:      monthGan,
		MonthZhi:      monthZhi,
		DayGan:        dayGan,
		DayZhi:        dayZhi,
		HourGan:       hourGan,
		HourZhi:       hourZhi,
		PrevSolarterm: p,
		NextSolarterm: n,
	}
}

// Animal Animal
func (gz *Ganzhi) Animal() *animal.Animal {
	return animal.NewAnimal(gz.YearZhi.Order())
}

// YearGanzhiAlias YearGanzhiAlias
func (gz *Ganzhi) YearGanzhiAlias() string {
	return gz.YearGan.Alias() + gz.YearZhi.Alias()
}

// MonthGanzhiAlias MonthGanzhiAlias
func (gz *Ganzhi) MonthGanzhiAlias() string {
	return gz.MonthGan.Alias() + gz.MonthZhi.Alias()
}

// DayGanzhiAlias DayGanzhiAlias
func (gz *Ganzhi) DayGanzhiAlias() string {
	return gz.DayGan.Alias() + gz.DayZhi.Alias()
}

// HourGanzhiAlias HourGanzhiAlias
func (gz *Ganzhi) HourGanzhiAlias() string {
	return gz.HourGan.Alias() + gz.HourZhi.Alias()
}

// YearGanzhiOrder YearGanzhiOrder
func (gz *Ganzhi) YearGanzhiOrder() int64 {
	return ganzhiOrder(gz.YearGan.Order(), gz.YearZhi.Order())
}

// MonthGanzhiOrder MonthGanzhiOrder
func (gz *Ganzhi) MonthGanzhiOrder() int64 {
	return ganzhiOrder(gz.MonthGan.Order(), gz.MonthZhi.Order())
}

// DayGanzhiOrder DayGanzhiOrder
func (gz *Ganzhi) DayGanzhiOrder() int64 {
	return ganzhiOrder(gz.DayGan.Order(), gz.DayZhi.Order())
}

// HourGanzhiOrder HourGanzhiOrder
func (gz *Ganzhi) HourGanzhiOrder() int64 {
	return ganzhiOrder(gz.HourGan.Order(), gz.HourZhi.Order())
}

func isSupported(year int64) bool {
	return solarterm.SolartermFromYear <= year && year < solarterm.SolartermToYear
}

func ganzhiOrder(ganOrder, zhiOrder int64) int64 {
	return utils.OrderMod(((ganOrder+10-zhiOrder)%10)/2*12+zhiOrder, 60)
}
