package tyme

// TenNames 旬名称
var TenNames = []string{"甲子", "甲戌", "甲申", "甲午", "甲辰", "甲寅"}

// Ten 旬
type Ten struct {
	LoopTyme
}

func (Ten) FromIndex(index int) Ten {
	return Ten{LoopTyme{}.FromIndex(TenNames, index)}
}

func (Ten) FromName(name string) (*Ten, error) {
	p, err := LoopTyme{}.FromName(TenNames, name)
	if err != nil {
		return nil, err
	}
	return &Ten{*p}, nil
}

func (o Ten) Next(n int) Ten {
	return o.FromIndex(o.nextIndex(n))
}
