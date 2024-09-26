package tyme

// LandNames 九野名称
var LandNames = []string{"玄天", "朱天", "苍天", "阳天", "钧天", "幽天", "颢天", "变天", "炎天"}

// Land 九野
type Land struct {
	LoopTyme
}

func (Land) FromIndex(index int) Land {
	return Land{LoopTyme{}.FromIndex(LandNames, index)}
}

func (Land) FromName(name string) (*Land, error) {
	p, err := LoopTyme{}.FromName(LandNames, name)
	if err != nil {
		return nil, err
	}
	return &Land{*p}, nil
}

func (o Land) Next(n int) Land {
	return o.FromIndex(o.nextIndex(n))
}

// GetDirection 方位
func (o Land) GetDirection() Direction {
	return Direction{}.FromIndex(o.index)
}
