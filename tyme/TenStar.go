package tyme

// TenStarNames 十神名称
var TenStarNames = []string{"比肩", "劫财", "食神", "伤官", "偏财", "正财", "七杀", "正官", "偏印", "正印"}

// TenStar 十神
type TenStar struct {
	LoopTyme
}

func (TenStar) FromIndex(index int) TenStar {
	return TenStar{LoopTyme{}.FromIndex(TenStarNames, index)}
}

func (TenStar) FromName(name string) (*TenStar, error) {
	p, err := LoopTyme{}.FromName(TenStarNames, name)
	if err != nil {
		return nil, err
	}
	return &TenStar{*p}, nil
}

func (o TenStar) Next(n int) TenStar {
	return o.FromIndex(o.nextIndex(n))
}
