package tyme

// DipperNames 北斗九星名称
var DipperNames = []string{"天枢", "天璇", "天玑", "天权", "玉衡", "开阳", "摇光", "洞明", "隐元"}

// Dipper 北斗九星
type Dipper struct {
	LoopTyme
}

func (Dipper) FromIndex(index int) Dipper {
	return Dipper{LoopTyme{}.FromIndex(DipperNames, index)}
}

func (Dipper) FromName(name string) (*Dipper, error) {
	p, err := LoopTyme{}.FromName(DipperNames, name)
	if err != nil {
		return nil, err
	}
	return &Dipper{*p}, nil
}

func (o Dipper) Next(n int) Dipper {
	return o.FromIndex(o.nextIndex(n))
}
