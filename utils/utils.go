package utils

// OrderMod 取模函数, 等价于a % b, 但当结果为0时, 返回b
func OrderMod(a, b int64) (result int64) {
	result = a % b
	if result == 0 {
		result = b
	}
	return
}
