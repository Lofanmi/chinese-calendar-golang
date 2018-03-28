package animal

// Animal Animal
type Animal struct {
	order int64
}

var animalAlias = [...]string{
	"鼠", "牛", "虎", "兔", "龙", "蛇",
	"马", "羊", "猴", "鸡", "狗", "猪",
}

// NewAnimal NewAnimal
func NewAnimal(order int64) *Animal {
	if !isSupported(order) {
		return nil
	}
	return &Animal{order: order}
}

// Alias Alias
func (animal *Animal) Alias() string {
	return animalAlias[(animal.order-1)%12]
}

func isSupported(order int64) bool {
	return 1 <= order && order <= 12
}
