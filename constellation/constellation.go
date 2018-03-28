package constellation

import (
	"time"
)

// Constellation Constellation
type Constellation struct {
	t *time.Time
}

// NewConstellation NewConstellation
func NewConstellation(t *time.Time) *Constellation {
	return &Constellation{t: t}
}

// Alias Alias
func (constellation *Constellation) Alias() string {
	dates := [...]int{20, 19, 21, 21, 21, 22, 23, 23, 23, 23, 22, 22}
	constellations := []rune("水瓶双鱼白羊金牛双子巨蟹狮子处女天秤天蝎射手魔羯")
	from := constellation.t.Month() * 2
	if constellation.t.Day() < dates[constellation.t.Month()-1] {
		from -= 2
	}
	return string(constellations[from:][:2])
}
