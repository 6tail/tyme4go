package tyme

// ZoneNames 宫名称
var ZoneNames = []string{"东", "北", "西", "南"}

// Zone 宫
type Zone struct {
	LoopTyme
}

func (Zone) FromIndex(index int) Zone {
	return Zone{LoopTyme{}.FromIndex(ZoneNames, index)}
}

func (Zone) FromName(name string) (*Zone, error) {
	p, err := LoopTyme{}.FromName(ZoneNames, name)
	if err != nil {
		return nil, err
	}
	return &Zone{*p}, nil
}

func (o Zone) Next(n int) Zone {
	return o.FromIndex(o.nextIndex(n))
}

// GetDirection 方位
func (o Zone) GetDirection() Direction {
	d, _ := Direction{}.FromName(o.GetName())
	return *d
}

// GetBeast 神兽
func (o Zone) GetBeast() Beast {
	return Beast{}.FromIndex(o.index)
}
