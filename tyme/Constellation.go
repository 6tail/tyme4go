package tyme

var ConstellationNames = []string{"白羊", "金牛", "双子", "巨蟹", "狮子", "处女", "天秤", "天蝎", "射手", "摩羯", "水瓶", "双鱼"}

// Constellation 星座
type Constellation struct {
	LoopTyme
}

func (Constellation) FromIndex(index int) Constellation {
	return Constellation{LoopTyme{}.FromIndex(ConstellationNames, index)}
}

func (Constellation) FromName(name string) (*Constellation, error) {
	p, err := LoopTyme{}.FromName(ConstellationNames, name)
	if err != nil {
		return nil, err
	}
	return &Constellation{*p}, nil
}

func (o Constellation) Next(n int) Constellation {
	return o.FromIndex(o.nextIndex(n))
}
