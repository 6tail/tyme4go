package tyme

// TwelveStarNames 黄道黑道十二神名称
var TwelveStarNames = []string{"青龙", "明堂", "天刑", "朱雀", "金匮", "天德", "白虎", "玉堂", "天牢", "玄武", "司命", "勾陈"}

// TwelveStar 黄道黑道十二神
type TwelveStar struct {
	LoopTyme
}

func (TwelveStar) FromIndex(index int) TwelveStar {
	return TwelveStar{LoopTyme{}.FromIndex(TwelveStarNames, index)}
}

func (TwelveStar) FromName(name string) (*TwelveStar, error) {
	p, err := LoopTyme{}.FromName(TwelveStarNames, name)
	if err != nil {
		return nil, err
	}
	return &TwelveStar{*p}, nil
}

func (o TwelveStar) Next(n int) TwelveStar {
	return o.FromIndex(o.nextIndex(n))
}

// GetEcliptic 黄道黑道
func (o TwelveStar) GetEcliptic() Ecliptic {
	return Ecliptic{}.FromIndex([]int{0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 0, 1}[o.index])
}
