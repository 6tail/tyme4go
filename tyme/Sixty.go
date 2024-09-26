package tyme

// SixtyNames 元名称
var SixtyNames = []string{"上元", "中元", "下元"}

// Sixty 元（60年=1元）
type Sixty struct {
	LoopTyme
}

func (Sixty) FromIndex(index int) Sixty {
	return Sixty{LoopTyme{}.FromIndex(SixtyNames, index)}
}

func (Sixty) FromName(name string) (*Sixty, error) {
	p, err := LoopTyme{}.FromName(SixtyNames, name)
	if err != nil {
		return nil, err
	}
	return &Sixty{*p}, nil
}

func (o Sixty) Next(n int) Sixty {
	return o.FromIndex(o.nextIndex(n))
}
