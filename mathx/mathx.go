package mathx

import (
	"math"

	"github.com/Lofanmi/chinese-calendar-golang/utils"
)

// Int2 取整数部分 (int2)
func Int2(v float64) float64 {
	return math.Floor(v)
}

// Fmod math.Mod (fmod)
func Fmod(x, y float64) float64 {
	return math.Mod(x, y)
}

// Rad2MRad 对超过0-2PI的角度转为0-2PI (rad2mrad)
func Rad2MRad(v float64) float64 {
	v = Fmod(v, utils.PI2)
	if v < 0 {
		return v + utils.PI2
	}
	return v
}

// Rad2RRad 对超过-PI到PI的角度转为-PI到PI (rad2rrad)
func Rad2RRad(v float64) float64 {
	v = Fmod(v, utils.PI2)

	if v <= -utils.PI {
		return v + utils.PI2
	}
	if v > utils.PI {
		return v - utils.PI2
	}
	return v
}

// Mod2 临界余数(a与最近的整倍数b相差的距离) (mod2)
func Mod2(a, b float64) float64 {
	c := a / b
	c -= math.Floor(c)
	if c > 0.5 {
		c -= 1
	}
	return c * b
}

// LLR2XYZ 球面转直角坐标 (llr2xyz)
func LLR2XYZ(J, W, R float64) (x, y, z float64) {
	x = R * math.Cos(W) * math.Cos(J)
	y = R * math.Cos(W) * math.Sin(J)
	z = R * math.Sin(W)
	return
}

// XYZ2LLR 直角坐标转球 (xyz2llr)
func XYZ2LLR(x, y, z float64) (r0, r1, r2 float64) {
	r2 = math.Sqrt(x*x + y*y + z*z)
	r1 = math.Asin(z / r2)
	r0 = Rad2MRad(math.Atan2(y, x))
	return
}

// LLRConv 球面坐标旋转 (llrConv)
// 黄道赤道坐标变换, 赤到黄E取负
func LLRConv(J, W, R, E float64) (r0, r1, r2 float64) {
	r0 = math.Atan2(math.Sin(J)*math.Cos(E)-math.Tan(W)*math.Sin(E), math.Cos(J))
	r1 = math.Asin(math.Cos(E)*math.Sin(W) + math.Sin(E)*math.Cos(W)*math.Sin(J))
	r2 = R
	r0 = Rad2MRad(r0)
	return
}

// CD2DP 赤道坐标转为地平坐标 (CD2DP)
func CD2DP(z0, z1, z2, L, fa, gst float64) (a0, a1, a2 float64) {
	// 转到相对于地平赤道分点的赤道坐标
	a0 = z0 + utils.PI/2 - gst - L
	a1 = z1
	a2 = z2
	a0, a1, a2 = LLRConv(a0, a1, a2, utils.PI/2-fa)
	a0 = Rad2MRad(utils.PI/2 - a0)
	return
}

// J1J2 求角度差 (j1_j2)
func J1J2(J1, W1, J2, W2 float64) float64 {
	dJ := Rad2RRad(J1 - J2)
	dW := W1 - W2
	if math.Abs(dJ) < 1/1000 && math.Abs(dW) < 1/1000 {
		dJ *= math.Cos((W1 + W2) / 2)
		return math.Sqrt(dJ*dJ + dW*dW)
	}
	return math.Acos(math.Sin(W1)*math.Sin(W2) + math.Cos(W1)*math.Cos(W2)*math.Cos(dJ))
}

// H2G 日心球面转地心球面, Z星体球面坐标, A地球球面坐标
// 本函数是通用的球面坐标中心平移函数, 行星计算中将反复使用
func H2G(z0, z1, z2, a0, a1, a2 float64) (float64, float64, float64) {
	a0, a1, a2 = LLR2XYZ(a0, a1, a2) // 地球
	z0, z1, z2 = LLR2XYZ(z0, z1, z2) // 星体
	z0 -= a0
	z1 -= a1
	z2 -= a2
	return XYZ2LLR(z0, z1, z2)
}

// ShiChaJ 视差角(不是视差)
func ShiChaJ(gst, L, fa, J, W float64) float64 {
	H := gst + L - J // 天体的时角
	return Rad2MRad(math.Atan2(math.Sin(H), math.Tan(fa)*math.Cos(W)-math.Sin(W)*math.Cos(H)))
}
