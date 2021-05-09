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
	// // 地球赤道半径(千米)
	// cs_rEar = 6378.1366
	// // 平均半径
	// cs_rEarA = 0.99834 * self::cs_rEar
	// // 地球极赤半径比
	// cs_ba = 0.99664719
	// // 地球极赤半径比的平方
	// cs_ba2 = self::cs_ba * self::cs_ba
	// // 天文单位长度(千米)
	// cs_AU = 1.49597870691e8
	// // sin(太阳视差)
	// cs_sinP = self::cs_rEar / self::cs_AU
	// // 太阳视差
	// // cs_PI = asin(self::cs_sinP)
	// cs_PI = 4.26352097959108 / 100000
	// // 光速(行米/秒)
	// cs_GS = 299792.458
	// // 每天文单位的光行时间(儒略世纪)
	// cs_Agx = self::cs_AU / self::cs_GS / 86400 / 36525
	// // 行星会合周期
	// cs_xxHH = [116, 584, 780, 399, 378, 370, 367, 367]
	// // 行星名称
	// xxName = ['地球', '水星', '金星', '火星', '木星', '土星', '天王星', '海王星', '冥王星']

	// RADd 每弧度的度数
	RADd = 180 / PI

	// J2000 儒略日期TT时 2451545.0
	J2000 = 2451545
	// // 月亮与地球的半径比(用于半影计算)
	// cs_k = 0.2725076
	// // 月亮与地球的半径比(用于本影计算)
	// cs_k2 = 0.2722810
	// // 太阳与地球的半径比(对应959.64)
	// cs_k0 = 109.1222
	// // 用于月亮视半径计算
	// cs_sMoon = self::cs_k * self::cs_rEar * 1.0000036 * self::rad
	// // 用于月亮视半径计算
	// cs_sMoon2 = self::cs_k2 * self::cs_rEar * 1.0000036 * self::rad
	// // 用于太阳视半径计算
	// cs_sSun = 959.64
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
