package tyme

// NineNames 数九名称
var NineNames = []string{"一九", "二九", "三九", "四九", "五九", "六九", "七九", "八九", "九九"}

// Nine 数九
type Nine struct {
	LoopTyme
}

func (Nine) FromIndex(index int) Nine {
	return Nine{LoopTyme{}.FromIndex(NineNames, index)}
}

func (Nine) FromName(name string) (*Nine, error) {
	p, err := LoopTyme{}.FromName(NineNames, name)
	if err != nil {
		return nil, err
	}
	return &Nine{*p}, nil
}

func (o Nine) Next(n int) Nine {
	return o.FromIndex(o.nextIndex(n))
}
