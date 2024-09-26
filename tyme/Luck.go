package tyme

// LuckNames 吉凶名称
var LuckNames = []string{"吉", "凶"}

// Luck 吉凶
type Luck struct {
	LoopTyme
}

func (Luck) FromIndex(index int) Luck {
	return Luck{LoopTyme{}.FromIndex(LuckNames, index)}
}

func (Luck) FromName(name string) (*Luck, error) {
	p, err := LoopTyme{}.FromName(LuckNames, name)
	if err != nil {
		return nil, err
	}
	return &Luck{*p}, nil
}

func (o Luck) Next(n int) Luck {
	return o.FromIndex(o.nextIndex(n))
}
