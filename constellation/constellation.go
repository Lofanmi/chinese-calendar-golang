package constellation

import (
	"time"
)

// Constellation 星座
type Constellation struct {
	t *time.Time
}

// NewConstellation 创建星座对象
func NewConstellation(t *time.Time) *Constellation {
	return &Constellation{t: t}
}

// Alias 返回星座名称
func (constellation *Constellation) Alias() string {
	dates := [...]int{20, 19, 21, 21, 21, 22, 23, 23, 23, 23, 22, 22}
	constellations := []rune("水瓶双鱼白羊金牛双子巨蟹狮子处女天秤天蝎射手魔羯")
	from := constellation.t.Month() * 2
	if constellation.t.Day() < dates[constellation.t.Month()-1] {
		from -= 2
	}
	return string(constellations[from:][:2])
}
