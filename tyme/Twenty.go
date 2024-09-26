package tyme

// TwentyNames 运名称
var TwentyNames = []string{"一运", "二运", "三运", "四运", "五运", "六运", "七运", "八运", "九运"}

// Twenty 运（20年=1运，3运=1元）
type Twenty struct {
	LoopTyme
}

func (Twenty) FromIndex(index int) Twenty {
	return Twenty{LoopTyme{}.FromIndex(TwentyNames, index)}
}

func (Twenty) FromName(name string) (*Twenty, error) {
	p, err := LoopTyme{}.FromName(TwentyNames, name)
	if err != nil {
		return nil, err
	}
	return &Twenty{*p}, nil
}

func (o Twenty) Next(n int) Twenty {
	return o.FromIndex(o.nextIndex(n))
}

// GetSixty 元
func (o Twenty) GetSixty() Sixty {
	return Sixty{}.FromIndex(o.index / 3)
}
