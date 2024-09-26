package tyme

// SevenStarNames 七曜名称
var SevenStarNames = []string{"日", "月", "火", "水", "木", "金", "土"}

// SevenStar 七曜（七政、七纬、七耀）
type SevenStar struct {
	LoopTyme
}

func (SevenStar) FromIndex(index int) SevenStar {
	return SevenStar{LoopTyme{}.FromIndex(SevenStarNames, index)}
}

func (SevenStar) FromName(name string) (*SevenStar, error) {
	p, err := LoopTyme{}.FromName(SevenStarNames, name)
	if err != nil {
		return nil, err
	}
	return &SevenStar{*p}, nil
}

func (o SevenStar) Next(n int) SevenStar {
	return o.FromIndex(o.nextIndex(n))
}

// GetWeek 星期
func (o SevenStar) GetWeek() Week {
	return Week{}.FromIndex(o.index)
}
