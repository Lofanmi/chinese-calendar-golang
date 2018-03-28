package zhi

// Zhi Zhi
type Zhi struct {
	order int64
}

var zhiAlias = [...]string{
	"子", "丑", "寅", "卯", "辰", "巳",
	"午", "未", "申", "酉", "戌", "亥",
}

// NewZhi NewZhi
func NewZhi(order int64) *Zhi {
	if !isSupported(order) {
		return nil
	}
	return &Zhi{order: order}
}

// Alias Alias
func (zhi *Zhi) Alias() string {
	return zhiAlias[(zhi.order-1)%12]
}

// Order Order
func (zhi *Zhi) Order() int64 {
	return zhi.order
}

func isSupported(order int64) bool {
	return 1 <= order && order <= 12
}
