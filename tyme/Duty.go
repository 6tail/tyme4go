package tyme

// DutyNames 建除十二值神名称
var DutyNames = []string{"建", "除", "满", "平", "定", "执", "破", "危", "成", "收", "开", "闭"}

// Duty 建除十二值神
type Duty struct {
	LoopTyme
}

func (Duty) FromIndex(index int) Duty {
	return Duty{LoopTyme{}.FromIndex(DutyNames, index)}
}

func (Duty) FromName(name string) (*Duty, error) {
	p, err := LoopTyme{}.FromName(DutyNames, name)
	if err != nil {
		return nil, err
	}
	return &Duty{*p}, nil
}

func (o Duty) Next(n int) Duty {
	return o.FromIndex(o.nextIndex(n))
}
