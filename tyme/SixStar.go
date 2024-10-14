package tyme

// SixStarNames 六曜名称
var SixStarNames = []string{"先胜", "友引", "先负", "佛灭", "大安", "赤口"}

// SixStar 六曜（孔明六曜星）
type SixStar struct {
	LoopTyme
}

func (SixStar) FromIndex(index int) SixStar {
	return SixStar{LoopTyme{}.FromIndex(SixStarNames, index)}
}

func (SixStar) FromName(name string) (*SixStar, error) {
	p, err := LoopTyme{}.FromName(SixStarNames, name)
	if err != nil {
		return nil, err
	}
	return &SixStar{*p}, nil
}

func (o SixStar) Next(n int) SixStar {
	return o.FromIndex(o.nextIndex(n))
}
