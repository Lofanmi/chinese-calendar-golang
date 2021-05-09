package base

import (
	"math"

	"github.com/Lofanmi/chinese-calendar-golang/deltat"
	"github.com/Lofanmi/chinese-calendar-golang/mathx"
	"github.com/Lofanmi/chinese-calendar-golang/utils"
)

// 天文基本问题

// SuoN 返回朔日的编号, jd应在朔日附近, 允许误差数天
func SuoN(jd float64) float64 {
	return math.Floor((jd + 8) / 29.5306)
}

// GxcSunLon 太阳光行差, t是世纪数
func GxcSunLon(t float64) float64 {
	// 平近点角
	v := -0.043126 + 628.301955*t - 0.000002732*t*t
	e := 0.016708634 - 0.000042037*t - 0.0000001267*t*t
	// 黄经光行差
	return (-20.49552 * (1 + e*math.Cos(v))) / utils.RAD
}

// GxcSunLat 黄纬光行差
func GxcSunLat(t float64) float64 {
	return 0
}

// GxcMoonLon 月球经度光行差, 误差0.07"
func GxcMoonLon(t float64) float64 {
	return -3.4E-6
}

// GxcMoonLat 月球纬度光行差, 误差0.006"
func GxcMoonLat(t float64) float64 {
	return 0.063 * math.Sin(0.057+8433.4662*t+0.000064*t*t) / utils.RAD
}

// PGst 传入T是2000年首起算的日数(UT), dt是deltaT(日), 精度要求不高时dt可取值为0
func PGst(T, dt float64) float64 {
	// 返回格林尼治平恒星时(不含赤经章动及非多项式部分), 即格林尼治子午圈的平春风点起算的赤经
	t := (T + dt) / 36525
	t2 := t * t
	t3 := t2 * t
	t4 := t3 * t
	// T是UT, 下一行的t是力学时(世纪数)
	return utils.PI2*(0.7790572732640+1.00273781191135448*T) + (0.014506+4612.15739966*t+1.39667721*t2-0.00009344*t3+0.00001882*t4)/utils.RAD
}

// PGst2 传入力学时J2000起算日数, 返回平恒星时
func PGst2(jd float64) float64 {
	dt := deltat.DtT(jd)
	return PGst(jd-dt, dt)
}

// SunShengJ 太阳升降计算
// jd儒略日(须接近L当地平午UT), L地理经度, fa地理纬度, sj=-1升, sj=1降
func SunShengJ(jd, L, fa, sj float64) float64 {
	jd = math.Floor(jd+0.5) - L/utils.PI2

	for i := 0; i < 2; i++ {
		// 黄赤交角
		T := jd / 36525
		E := (84381.4060 - 46.836769*T) / utils.RAD

		// 儒略世纪年数,力学时
		t := T + (32*(T+1.8)*(T+1.8)-20)/86400/36525

		J := (48950621.66 + 6283319653.318*t + 53*t*t - 994 + 334166*math.Cos(4.669257+628.307585*t) + 3489*math.Cos(4.6261+1256.61517*t) + 2060.6*math.Cos(2.67823+628.307585*t)*t) / 10000000

		// 太阳黄经以及它的正余弦值
		sinJ := math.Sin(J)
		cosJ := math.Cos(J)

		// 恒星时(子午圈位置)
		gst := (0.7790572732640+1.00273781191135448*jd)*utils.PI2 + (0.014506+4612.15739966*T+1.39667721*T*T)/utils.RAD

		// 太阳赤经
		A := math.Atan2(sinJ*math.Cos(E), cosJ)
		// 太阳赤纬
		D := math.Asin(math.Sin(E) * sinJ)

		// 太阳在地平线上的math.Cos(时角)计算
		cosH0 := (math.Sin(-50*60/utils.RAD) - math.Sin(fa)*math.Sin(D)) / (math.Cos(fa) * math.Cos(D))
		if math.Abs(cosH0) >= 1 {
			return 0
		}

		// (升降时角-太阳时角)/太阳速度
		jd += mathx.Rad2RRad(sj*math.Acos(cosH0)-(gst+L-A)) / 6.28
	}

	// 返回格林尼治UT
	return jd
}

// PtyZty 时差计算(高精度),t力学时儒略世纪数 (pty_zty)
func PtyZty(t float64) float64 {
	t2 := t * t
	t3 := t2 * t
	t4 := t3 * t
	t5 := t4 * t

	L := (1753470142+628331965331.8*t+5296.74*t2+0.432*t3-0.1124*t4-0.00009*t5)/1000000000 + utils.PI - 20.5/utils.RAD

	// 黄经章
	dL := -17.2 * math.Sin(2.1824-33.75705*t) / utils.RAD
	// 交角章
	dE := 9.2 * math.Cos(2.1824-33.75705*t) / utils.RAD
	// 真黄赤交角
	E := HCJJ(t) + dE

	// 地球坐标
	var z0, z1, z2 float64
	z0 = XL0_calc(0, 0, t, 50) + utils.PI + GxcSunLon(t) + dL
	z1 = - (2796*math.Cos(3.1987+8433.46616*t) + 1016*math.Cos(5.4225+550.75532*t) + 804*math.Cos(3.88+522.3694*t)) / 1000000000

	// z太阳地心赤道坐标
	z0, z1, z2 = mathx.LLRConv(z0, z1, z2, E)
	z0 -= dL * math.Cos(E)

	L = mathx.Rad2RRad(L - z0)

	// 单位是周(天)
	return L / utils.PI2
}

// PtyZty2 时差计算(低精度), 误差约在1秒以内, t力学时儒略世纪数 (pty_zty2)
func PtyZty2(t float64) float64 {
	L := (1753470142+628331965331.8*t+5296.74*t*t)/1000000000 + utils.PI
	var z0, z1, z2 float64
	E := (84381.4088 - 46.836051*t) / utils.RAD

	// 地球坐标
	z0 = XL0_calc(0, 0, t, 5) + utils.PI
	z1 = 0

	// z太阳地心赤道坐标
	z0, z1, z2 = mathx.LLRConv(z0, z1, z2, E)
	L = mathx.Rad2RRad(L - z0)

	// 单位是周(天)
	return L / utils.PI2
}

// HCJJ 黄赤交角, 返回P03黄赤交角, t是世纪数 (hcjj)
func HCJJ(t float64) float64 {
	t2 := t * t
	t3 := t2 * t
	t4 := t3 * t
	t5 := t4 * t
	return (84381.4060 - 46.836769*t - 0.0001831*t2 + 0.00200340*t3 - 5.76e-7*t4 - 4.34e-8*t5) / utils.RAD
}
