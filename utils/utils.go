package utils

import (
	"errors"
	"math"
	"regexp"
	"strconv"
)

// OrderMod 取模函数, 等价于a % b, 但当结果为0时, 返回b
func OrderMod(a, b int64) (result int64) {
	result = a % b
	if result == 0 {
		result = b
	}
	return
}

const (
	// PI 圆周率
	PI = 3.141592653589793
	// PI2 圆周率 * 2
	PI2 = PI * 2
	// PI05 圆周率 * 0.5
	PI05 = PI * 0.5
	// RAD 每弧度的角秒数
	RAD = 180 * 3600 / PI
	// CsREar 地球赤道半径(千米)
	CsREar = 6378.1366
	// CsREarA 平均半径
	CsREarA = 0.99834 * CsREar
	// CsBa 地球极赤半径比
	CsBa = 0.99664719
	// CsBa2 地球极赤半径比的平方
	CsBa2 = CsBa * CsBa
	// CsAU 天文单位长度(千米)
	CsAU = 1.49597870691e8
	// CsSinP sin(太阳视差)
	CsSinP = CsREar / CsAU
	// CsPI 太阳 CsPI视差
	// CsPI = math.Asin(CsSinP)
	CsPI = 4.26352097959108 / 100000
	// CsGS 光速(行米/秒)
	CsGS = 299792.458
	// CsAgx 每天文单位的光行时间(儒略世纪)
	CsAgx = CsAU / CsGS / 86400 / 36525

	// CsXxHH 行星会合周期
	// CsXxHH = [116, 584, 780, 399, 378, 370, 367, 367]
	// // 行星名称
	// xxName = ['地球', '水星', '金星', '火星', '木星', '土星', '天王星', '海王星', '冥王星']

	// RADd 每弧度的度数
	RADd = 180 / PI

	// J2000 儒略日期TT时 2451545.0
	J2000 = 2451545
	// CsK 月亮与地球的半径比(用于半影计算)
	CsK = 0.2725076
	// CsK2 月亮与地球的半径比(用于本影计算)
	CsK2 = 0.2722810
	// CsK0 太阳与地球的半径比(对应959.64)
	CsK0 = 109.1222
	// CsSMoon 用于月亮视半径计算
	CsSMoon = CsK * CsREar * 1.0000036 * RAD
	// CsSMoon2 用于月亮视半径计算
	CsSMoon2 = CsK2 * CsREar * 1.0000036 * RAD
	// CsSSun 用于太阳视半径计算
	CsSSun = 959.64
)

// Year2AYear 传入普通纪年或天文纪年, 返回天文纪年
func Year2AYear(year string) (Y int, err error) {
	y := regexp.MustCompile(`[^0-9Bb*-]`).ReplaceAllString(year, "")
	if len(y) <= 0 {
		err = errors.New("Year2Ayear invalid year: " + year)
		return
	}
	// 通用纪年法(公元前)
	if y[0] == 'B' || y[0] == 'b' || y[0] == '*' {
		var a int
		if a, err = strconv.Atoi(y[1:]); err != nil {
			return
		}
		Y = 1 - a
		if Y > 0 {
			Y = -10000
			err = errors.New("通用纪法的公元前纪法从B.C.1年开始。并且没有公元0年")
			return
		}
	} else {
		Y, err = strconv.Atoi(y)
	}
	if Y < -4712 {
		err = errors.New("超过B.C. 4713不准")
		return
	}
	if Y > 9999 {
		err = errors.New("超过9999年的农历计算很不准")
		return
	}
	return
}

// JD 公历转儒略日
// int2 (取整数部分) <=> math.Floor
func JD(yy, mm, dd int) float64 {
	var n float64
	y, m, d := float64(yy), float64(mm), float64(dd)
	// 判断是否为格里高利历日1582*372+10*31+15
	G := y*372+m*31+math.Floor(d) >= 588829
	if m <= 2 {
		m += 12
		y--
	}
	// 加百年闰
	if G {
		n = math.Floor(y / 100.0)
		n = 2 - n + math.Floor(n/4)
	}
	return math.Floor(365.25*(y+4716)) + math.Floor(30.6001*(m+1)) + d + n - 1524.5
}

// DD 儒略日数转公历
// int2 (取整数部分) <=> math.Floor
func DD(jd float64) (_Y, _M, _D, _h, _m, _s int) {
	// 取得日数的整数部份A及小数部分F
	D := math.Floor(jd + 0.5)
	F := jd + 0.5 - D
	c := 0.0

	if D >= 2299161 {
		c = math.Floor((D - 1867216.25) / 36524.25)
		D += 1 + c - math.Floor(c/4)
	}

	// 年数
	D += 1524
	_Y = int(math.Floor((D - 122.1) / 365.25))
	// 月数
	D -= math.Floor(365.25 * float64(_Y))
	_M = int(math.Floor(D / 30.601))
	// 日数
	D -= math.Floor(30.601 * float64(_M))
	_D = int(D)

	if _M > 13 {
		_M -= 13
		_Y -= 4715
	} else {
		_M -= 1
		_Y -= 4716
	}

	// 日的小数转为时分秒
	F *= 24
	_h = int(math.Floor(F))
	F -= float64(_h)

	F *= 60
	_m = int(math.Floor(F))
	F -= float64(_m)

	F *= 60
	_s = int(F)

	return
}
