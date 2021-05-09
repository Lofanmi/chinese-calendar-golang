package xl

import (
	"math"

	"github.com/Lofanmi/chinese-calendar-golang/base"
	"github.com/Lofanmi/chinese-calendar-golang/mathx"
	"github.com/Lofanmi/chinese-calendar-golang/utils"
)

func XL0_calc(xt, zn, t, n float64) float64 {
	return 0
}

func XL1_calc(zn, t, n float64) float64 {
	return 0
}

// ELon 星历函数(日月球面坐标计算) - 地球经度计算, 返回Date分点黄经, 传入世纪数, 取项数
func ELon(t, n float64) float64 {
	return 0
}

// MLon 星历函数(日月球面坐标计算) - 月球经度计算, 返回Date分点黄经, 传入世纪数, 取项数
func MLon(t, n float64) float64 {
	return 0
}

// Ev 地球速度, t是世纪数, 误差小于万分3
func Ev(t float64) float64 {
	f := 628.307585 * t
	return 628.332 + 21*math.Sin(1.527+f) + 0.44*math.Sin(1.48+f*2) + 0.129*math.Sin(5.82+f)*t + 0.00055*math.Sin(4.21+f)*t*t
}

// Mv 月球速度, 传入世经数
func Mv(t float64) float64 {
	// 误差小于5%
	v := 8399.71 - 914*math.Sin(0.7848+8328.691425*t+0.0001523*t*t)
	// 误差小于0.3%
	v -= 179*math.Sin(2.543+15542.7543*t) + 160*math.Sin(0.1874+7214.0629*t) + 62*math.Sin(3.14+16657.3828*t) + 34*math.Sin(4.827+16866.9323*t) + 22*math.Sin(4.9+23871.4457*t) + 12*math.Sin(2.59+14914.4523*t) + 7*math.Sin(0.23+6585.7609*t) + 5*math.Sin(0.9+25195.624*t) + 5*math.Sin(2.32-7700.3895*t) + 5*math.Sin(3.88+8956.9934*t) + 5*math.Sin(0.49+7771.3771*t)
	return v
}

// MSALon 月日视黄经的差值
func MSALon(t, Mn, Sn float64) float64 {
	return MLon(t, Mn) + base.GxcMoonLon(t) - (ELon(t, Sn) + base.GxcSunLon(t) + utils.PI)
}

// SALon 太阳视黄经
func SALon(t, n float64) float64 {
	// 注意, 这里的章动计算很耗时
	return ELon(t, n) + nutationLon2(t) + base.GxcSunLon(t) + utils.PI
}

// ELonT 已知地球真黄经求时间
func ELonT(W float64) float64 {
	v := 628.3319653318

	// v的精度0.03%, 详见原文
	t := (W - 1.75347) / v
	v = Ev(t)

	// 再算一次v有助于提高精度, 不算也可以
	t += (W - ELon(t, 10)) / v
	v = Ev(t)

	t += (W - ELon(t, -1)) / v

	return t
}

// MLonT 已知真月球黄经求时间
func MLonT(W float64) float64 {
	v := 8399.70911033384

	t := (W - 3.81034) / v

	// v的精度0.5%, 详见原文
	t += (W - MLon(t, 3)) / v
	v = Mv(t)

	t += (W - MLon(t, 20)) / v
	t += (W - MLon(t, -1)) / v

	return t
}

// MSALonT 已知月日视黄经差求时间
func MSALonT(W float64) float64 {
	v := 7771.37714500204

	t := (W + 1.08472) / v

	// v的精度0.5%, 详见原文
	t += (W - MSALon(t, 3, 3)) / v
	v = Mv(t) - Ev(t)

	t += (W - MSALon(t, 20, 10)) / v
	t += (W - MSALon(t, -1, 60)) / v

	return t
}

// SALonT 已知太阳视黄经反求时间
func SALonT(W float64) float64 {
	v := 628.3319653318

	// v的精度0.03%, 详见原文
	t := (W - 1.75347 - utils.PI) / v
	v = Ev(t)

	// 再算一次v有助于提高精度, 不算也可以
	t += (W - SALon(t, 10)) / v
	v = Ev(t)

	t += (W - SALon(t, -1)) / v

	return t
}

// MSALonT2 已知月日视黄经差求时间, 高速低精度, 误差不超过600秒 (只验算了几千年)
func MSALonT2(W float64) float64 {
	v := 7771.37714500204
	t := (W + 1.08472) / v
	t2 := t * t
	t -= (- 0.00003309*t2 + 0.10976*math.Cos(0.784758+8328.6914246*t+0.000152292*t2) + 0.02224*math.Cos(0.18740+7214.0628654*t-0.00021848*t2) - 0.03342*math.Cos(4.669257+628.307585*t)) / v
	L := MLon(t, 20) - (4.8950632 + 628.3319653318*t + 0.000005297*t*t + 0.0334166*math.Cos(4.669257+628.307585*t) + 0.0002061*math.Cos(2.67823+628.307585*t)*t + 0.000349*math.Cos(4.6261+1256.61517*t) - 20.5/utils.PI)
	v = 7771.38 - 914*math.Sin(0.7848+8328.691425*t+0.0001523*t*t) - 179*math.Sin(2.543+15542.7543*t) - 160*math.Sin(0.1874+7214.0629*t)
	t += (W - L) / v
	return t
}

// SALonT2 已知太阳视黄经反求时间, 高速低精度, 最大误差不超过600秒
func SALonT2(W float64) float64 {
	v := 628.3319653318
	t := (W - 1.75347 - utils.PI) / v
	t -= (0.000005297*t*t + 0.0334166*math.Cos(4.669257+628.307585*t) + 0.0002061*math.Cos(2.67823+628.307585*t)*t) / v
	t += (W - ELon(t, 8) - utils.PI + (20.5+17.2*math.Sin(2.1824-33.75705*t))/utils.PI) / v
	return t
}

// MoonIll 月亮被照亮部分的比例
func MoonIll(t float64) float64 {
	t2 := t * t
	t3 := t2 * t
	t4 := t3 * t
	dm := utils.PI / 180
	// 日月平距角
	D := (297.8502042 + 445267.1115168*t - 0.0016300*t2 + t3/545868 - t4/113065000) * dm
	// 太阳平近点
	M := (357.5291092 + 35999.0502909*t - 0.0001536*t2 + t3/24490000) * dm
	// 月亮平近点
	m := (134.9634114 + 477198.8676313*t + 0.0089970*t2 + t3/69699 - t4/14712000) * dm

	a := utils.PI - D + (-6.289*math.Sin(m)+2.100*math.Sin(M)-1.274*math.Sin(D*2-m)-0.658*math.Sin(D*2)-0.214*math.Sin(m*2)-0.110*math.Sin(D))*dm

	return (1 + math.Cos(a)) / 2
}

// MoonRad 转入地平纬度及地月质心距离, 返回站心视半径(角秒)
func MoonRad(r, h float64) float64 {
	return utils.CsSMoon / r * (1 + math.Sin(h)*utils.CsREar/r)
}

// MoonMinR 求月亮近点时间和距离, t为儒略世纪数力学时
func MoonMinR(t float64, min bool) (float64, float64) {
	a, b := 27.55454988/36525, 0.0
	if min {
		b = -10.3302 / 36525
	} else {
		b = 3.4471 / 36525
	}
	// 平近(远)点时间
	t = b + a*mathx.Int2((t-b)/a+0.5)
	// 初算二次
	dt := 2.0 / 36525
	r1 := XL1_calc(2, t-dt, 10)
	r2 := XL1_calc(2, t, 10)
	r3 := XL1_calc(2, t+dt, 10)
	t += (r1 - r3) / (r1 + r3 - 2*r2) * dt / 2
	dt = 0.5 / 36525
	r1 = XL1_calc(2, t-dt, 20)
	r2 = XL1_calc(2, t, 20)
	r3 = XL1_calc(2, t+dt, 20)
	t += (r1 - r3) / (r1 + r3 - 2*r2) * dt / 2
	// 精算
	dt = 1200 / 86400 / 36525
	r1 = XL1_calc(2, t-dt, -1)
	r2 = XL1_calc(2, t, -1)
	r3 = XL1_calc(2, t+dt, -1)
	t += (r1 - r3) / (r1 + r3 - 2*r2) * dt / 2
	r2 += (r1 - r3) / (r1 + r3 - 2*r2) * (r3 - r1) / 8

	return t, r2
}

// MoonNode 月亮升交点
func MoonNode(t float64, asc bool) (float64, float64) {
	a, b := 27.21222082/36525.0, 0.0
	if asc {
		b = 21 / 36525.0
	} else {
		b = 35 / 36525.0
	}
	// 平升(降)交点时间
	t = b + a*mathx.Int2((t-b)/a+0.5)

	dt := 0.5 / 36525
	w := XL1_calc(1, t, 10)
	w2 := XL1_calc(1, t+dt, 10)
	v := (w2 - w) / dt
	t -= w / v

	dt = 0.05 / 36525
	w = XL1_calc(1, t, 40)
	w2 = XL1_calc(1, t+dt, 40)
	v = (w2 - w) / dt
	t -= w / v

	w = XL1_calc(1, t, -1)
	t -= w / v

	return t, XL1_calc(0, t, -1)
}

// 地球近远点
func earthMinR(t float64, min bool) (float64, float64) {
	a, b := 365.25963586/36525, 0.0
	if min {
		b = 1.7 / 36525
	} else {
		b = 184.5 / 36525
	}
	// 平近(远)点时间
	t = b + a*mathx.Int2((t-b)/a+0.5)
	// 初算二次
	dt := 3 / 36525.0
	r1 := XL0_calc(0, 2, t-dt, 10)
	r2 := XL0_calc(0, 2, t, 10)
	r3 := XL0_calc(0, 2, t+dt, 10)
	// 误差几个小时
	t += (r1 - r3) / (r1 + r3 - 2*r2) * dt / 2

	dt = 0.2 / 36525
	r1 = XL0_calc(0, 2, t-dt, 80)
	r2 = XL0_calc(0, 2, t, 80)
	r3 = XL0_calc(0, 2, t+dt, 80)
	// 误差几分钟
	t += (r1 - r3) / (r1 + r3 - 2*r2) * dt / 2

	// 精算
	dt = 0.01 / 36525
	r1 = XL0_calc(0, 2, t-dt, -1)
	r2 = XL0_calc(0, 2, t, -1)
	r3 = XL0_calc(0, 2, t+dt, -1)
	// 误差小于秒
	t += (r1 - r3) / (r1 + r3 - 2*r2) * dt / 2

	r2 += (r1 - r3) / (r1 + r3 - 2*r2) * (r3 - r1) / 8

	return t, r2
}
