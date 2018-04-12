package lunar

import (
	"fmt"
	"strings"
	"time"

	"github.com/Lofanmi/chinese-calendar-golang/animal"
	"github.com/Lofanmi/chinese-calendar-golang/utils"
)

// Lunar 农历
type Lunar struct {
	loc              *time.Location
	t                *time.Time
	year, month, day int64
	monthIsLeap      bool
}

var numberAlias = [...]string{
	"零", "一", "二", "三", "四",
	"五", "六", "七", "八", "九",
}

var dateAlias = [...]string{
	"初", "十", "廿", "卅",
}

var lunarMonthAlias = [...]string{
	"正", "二", "三", "四", "五", "六",
	"七", "八", "九", "十", "冬", "腊",
}

var lunars = [...]int64{
	0x04bd8, 0x04ae0, 0x0a570, 0x054d5, 0x0d260, 0x0d950, 0x16554, 0x056a0, 0x09ad0, 0x055d2, // 1900-1909
	0x04ae0, 0x0a5b6, 0x0a4d0, 0x0d250, 0x1d255, 0x0b540, 0x0d6a0, 0x0ada2, 0x095b0, 0x14977, // 1910-1919
	0x04970, 0x0a4b0, 0x0b4b5, 0x06a50, 0x06d40, 0x1ab54, 0x02b60, 0x09570, 0x052f2, 0x04970, // 1920-1929
	0x06566, 0x0d4a0, 0x0ea50, 0x06e95, 0x05ad0, 0x02b60, 0x186e3, 0x092e0, 0x1c8d7, 0x0c950, // 1930-1939
	0x0d4a0, 0x1d8a6, 0x0b550, 0x056a0, 0x1a5b4, 0x025d0, 0x092d0, 0x0d2b2, 0x0a950, 0x0b557, // 1940-1949
	0x06ca0, 0x0b550, 0x15355, 0x04da0, 0x0a5b0, 0x14573, 0x052b0, 0x0a9a8, 0x0e950, 0x06aa0, // 1950-1959
	0x0aea6, 0x0ab50, 0x04b60, 0x0aae4, 0x0a570, 0x05260, 0x0f263, 0x0d950, 0x05b57, 0x056a0, // 1960-1969
	0x096d0, 0x04dd5, 0x04ad0, 0x0a4d0, 0x0d4d4, 0x0d250, 0x0d558, 0x0b540, 0x0b6a0, 0x195a6, // 1970-1979
	0x095b0, 0x049b0, 0x0a974, 0x0a4b0, 0x0b27a, 0x06a50, 0x06d40, 0x0af46, 0x0ab60, 0x09570, // 1980-1989
	0x04af5, 0x04970, 0x064b0, 0x074a3, 0x0ea50, 0x06b58, 0x055c0, 0x0ab60, 0x096d5, 0x092e0, // 1990-1999
	0x0c960, 0x0d954, 0x0d4a0, 0x0da50, 0x07552, 0x056a0, 0x0abb7, 0x025d0, 0x092d0, 0x0cab5, // 2000-2009
	0x0a950, 0x0b4a0, 0x0baa4, 0x0ad50, 0x055d9, 0x04ba0, 0x0a5b0, 0x15176, 0x052b0, 0x0a930, // 2010-2019
	0x07954, 0x06aa0, 0x0ad50, 0x05b52, 0x04b60, 0x0a6e6, 0x0a4e0, 0x0d260, 0x0ea65, 0x0d530, // 2020-2029
	0x05aa0, 0x076a3, 0x096d0, 0x04afb, 0x04ad0, 0x0a4d0, 0x1d0b6, 0x0d250, 0x0d520, 0x0dd45, // 2030-2039
	0x0b5a0, 0x056d0, 0x055b2, 0x049b0, 0x0a577, 0x0a4b0, 0x0aa50, 0x1b255, 0x06d20, 0x0ada0, // 2040-2049
	0x14b63, 0x09370, 0x049f8, 0x04970, 0x064b0, 0x168a6, 0x0ea50, 0x06b20, 0x1a6c4, 0x0aae0, // 2050-2059
	0x0a2e0, 0x0d2e3, 0x0c960, 0x0d557, 0x0d4a0, 0x0da50, 0x05d55, 0x056a0, 0x0a6d0, 0x055d4, // 2060-2069
	0x052d0, 0x0a9b8, 0x0a950, 0x0b4a0, 0x0b6a6, 0x0ad50, 0x055a0, 0x0aba4, 0x0a5b0, 0x052b0, // 2070-2079
	0x0b273, 0x06930, 0x07337, 0x06aa0, 0x0ad50, 0x14b55, 0x04b60, 0x0a570, 0x054e4, 0x0d160, // 2080-2089
	0x0e968, 0x0d520, 0x0daa0, 0x16aa6, 0x056d0, 0x04ae0, 0x0a9d4, 0x0a2d0, 0x0d150, 0x0f252, // 2090-2099
	0x0d520, // 2100
}

// NewLunar 创建农历对象
func NewLunar(t *time.Time, loc *time.Location) *Lunar {
	year, month, day, isLeap := FromSolarTimestamp(t.Unix(), loc)
	return &Lunar{
		loc:         loc,
		t:           t,
		year:        year,
		month:       month,
		day:         day,
		monthIsLeap: isLeap,
	}
}

// FromSolarTimestamp 通过时间戳创建
func FromSolarTimestamp(ts int64, loc *time.Location) (lunarYear, lunarMonth, lunarDay int64, lunarMonthIsLeap bool) {
	var (
		i, offset, leap         int64
		daysOfYear, daysOfMonth int64
		isLeap                  bool
	)
	// 与 1900-01-31 相差多少天
	t := time.Unix(ts, 0)
	t1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc)
	t2 := time.Date(1900, 1, 31, 0, 0, 0, 0, loc)
	offset = (t1.Unix() - t2.Unix()) / 86400

	for i = 1900; i < 2101 && offset > 0; i++ {
		daysOfYear = daysOfLunarYear(i)
		offset -= daysOfYear
	}
	if offset < 0 {
		offset += daysOfYear
		i--
	}

	// 农历年
	lunarYear = i
	// 闰哪个月
	leap = leapMonth(i)

	isLeap = false

	// 用当年的天数 offset, 逐个减去每月(农历)的天数, 求出当天是本月的第几天
	for i = 1; i < 13 && offset > 0; i++ {
		// 闰月
		if leap > 0 && i == (leap+1) && !isLeap {
			i--
			isLeap = true
			// 计算农历月天数
			daysOfMonth = leapDays(lunarYear)
		} else {
			// 计算农历普通月天数
			daysOfMonth = lunarDays(lunarYear, i)
		}
		// 解除闰月
		if true == isLeap && i == (leap+1) {
			isLeap = false
		}
		offset -= daysOfMonth
	}
	// offset 为 0 时, 并且刚才计算的月份是闰月, 要校正
	if 0 == offset && leap > 0 && i == leap+1 {
		if isLeap {
			isLeap = false
		} else {
			isLeap = true
			i--
		}
	}
	if offset < 0 {
		offset += daysOfMonth
		i--
	}
	// 农历月
	lunarMonth = i
	// 农历日
	lunarDay = offset + 1
	// 农历是否为闰月
	lunarMonthIsLeap = isLeap

	return
}

// ToSolarTimestamp 转换为时间戳
func ToSolarTimestamp(year, month, day, hour, minute, second int64, isLeapMonth bool, loc *time.Location) int64 {
	var (
		i, offset int64
	)
	// 参数合法性效验
	if year < 1900 || year > 2100 {
		return 0
	}
	// 参数区间 1900.1.31~2100.12.1
	m := leapMonth(year)
	// 传参要求计算该闰月公历 但该年得出的闰月与传参的月份并不同
	if isLeapMonth && (m != month) {
		isLeapMonth = false
	}
	// 超出了最大极限值
	if 2100 == year && 12 == month && day > 1 || 1900 == year && 1 == month && day < 31 {
		return 0
	}
	days := lunarDays(year, month)
	maxDays := days
	// if month is leap, _day use leapDays method
	if isLeapMonth {
		maxDays = leapDays(year)
	}
	// 参数合法性效验
	if day > maxDays {
		return 0
	}
	// 计算农历的时间差
	offset = 0
	for i = 1900; i < year; i++ {
		offset += daysOfLunarYear(i)
	}
	isAdd := false
	for i = 1; i < month; i++ {
		leap := leapMonth(year)
		if !isAdd {
			// 处理闰月
			if leap <= i && leap > 0 {
				offset += leapDays(year)
				isAdd = true
			}
		}
		offset += lunarDays(year, i)
	}
	// 转换闰月农历 需补充该年闰月的前一个月的时差
	if isLeapMonth {
		offset += days
	}
	// 1900 年农历正月初一的公历时间为 1900年1月30日0时0分0秒 (该时间也是本农历的最开始起始点)
	// startTimestamp := time.Date(1900, 1, 30, 0, 0, 0, 0, loc).Unix()
	var startTimestamp int64 = -2206512000

	return (offset+day)*86400 + startTimestamp + hour*3600 + minute*60 + second
}

// LeapMonth 获取闰月(0表示不闰, 5表示闰五月)
func (lunar *Lunar) LeapMonth() int64 {
	return leapMonth(lunar.year)
}

// IsLeap 是否闰年
func (lunar *Lunar) IsLeap() bool {
	return lunar.LeapMonth() != 0
}

// IsLeapMonth 是否闰月
func (lunar *Lunar) IsLeapMonth() bool {
	return lunar.monthIsLeap
}

// Animal 返回年份生肖
func (lunar *Lunar) Animal() *animal.Animal {
	return animal.NewAnimal(utils.OrderMod(lunar.year-3, 12))
}

// YearAlias 汉字表示年(二零一八)
func (lunar *Lunar) YearAlias() string {
	s := fmt.Sprintf("%d", lunar.year)
	for i, replace := range numberAlias {
		s = strings.Replace(s, fmt.Sprintf("%d", i), replace, -1)
	}
	return s
}

// MonthAlias 汉字表示月(八月, 闰六月)
func (lunar *Lunar) MonthAlias() string {
	pre := ""
	if lunar.monthIsLeap {
		pre = "闰"
	}
	return pre + lunarMonthAlias[lunar.month-1] + "月"
}

// DayAlias 汉字表示日(初一, 初十...)
func (lunar *Lunar) DayAlias() (alias string) {
	switch lunar.day {
	case 10:
		alias = "初十"
	case 20:
		alias = "二十"
	case 30:
		alias = "三十"
	default:
		alias = dateAlias[(int)(lunar.day/10)] + numberAlias[lunar.day%10]
	}
	return
}

// GetYear 年
func (lunar *Lunar) GetYear() int64 {
	return lunar.year
}

// GetMonth 月
func (lunar *Lunar) GetMonth() int64 {
	return lunar.month
}

// GetDay 日
func (lunar *Lunar) GetDay() int64 {
	return lunar.day
}

func daysOfLunarYear(year int64) int64 {
	var (
		i, sum int64
	)
	sum = 29 * 12
	for i = 0x8000; i > 0x8; i >>= 1 {
		if (lunars[year-1900] & i) != 0 {
			sum++
		}
	}
	return sum + leapDays(year)
}

func leapMonth(year int64) int64 {
	return lunars[year-1900] & 0xf
}

func leapDays(year int64) (days int64) {
	if leapMonth(year) == 0 {
		days = 0
	} else if (lunars[year-1900] & 0x10000) != 0 {
		days = 30
	} else {
		days = 29
	}
	return
}

func lunarDays(year, month int64) (days int64) {
	if month > 12 || month < 1 {
		days = 0
	} else if (lunars[year-1900] & (0x10000 >> uint64(month))) != 0 {
		days = 30
	} else {
		days = 29
	}
	return
}
