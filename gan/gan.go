package gan

// Gan Gan
type Gan struct {
	order int64
}

var ganAlias = [...]string{
	"甲", "乙", "丙", "丁", "戊",
	"己", "庚", "辛", "壬", "癸",
}

// NewGan NewGan
func NewGan(order int64) *Gan {
	if !isSupported(order) {
		return nil
	}
	return &Gan{order: order}
}

// Alias Alias
func (gan *Gan) Alias() string {
	return ganAlias[(gan.order-1)%10]
}

// Order Order
func (gan *Gan) Order() int64 {
	return gan.order
}

func isSupported(order int64) bool {
	return 1 <= order && order <= 10
}
