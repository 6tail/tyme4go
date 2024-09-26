package tyme

// WeekNames 星期名称
var WeekNames = []string{"日", "一", "二", "三", "四", "五", "六"}

// Week 星期
type Week struct {
	LoopTyme
}

func (Week) FromIndex(index int) Week {
	return Week{LoopTyme{}.FromIndex(WeekNames, index)}
}

func (Week) FromName(name string) (*Week, error) {
	p, err := LoopTyme{}.FromName(WeekNames, name)
	if err != nil {
		return nil, err
	}
	return &Week{*p}, nil
}

func (o Week) Next(n int) Week {
	return o.FromIndex(o.nextIndex(n))
}

// GetSevenStar 七曜
func (o Week) GetSevenStar() SevenStar {
	return SevenStar{}.FromIndex(o.index)
}

func (o Week) Equals(target Week) bool {
	return o.index == target.index
}
