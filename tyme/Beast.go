package tyme

var BeastNames = []string{"青龙", "玄武", "白虎", "朱雀"}

// Beast 神兽
type Beast struct {
	LoopTyme
}

func (Beast) FromIndex(index int) Beast {
	return Beast{LoopTyme{}.FromIndex(BeastNames, index)}
}

func (Beast) FromName(name string) (*Beast, error) {
	p, err := LoopTyme{}.FromName(BeastNames, name)
	if err != nil {
		return nil, err
	}
	return &Beast{*p}, nil
}

func (o Beast) Next(n int) Beast {
	return o.FromIndex(o.nextIndex(n))
}

// GetZone 宫
func (o Beast) GetZone() Zone {
	return Zone{}.FromIndex(o.index)
}
