package tyme

// EclipticNames 黄道黑道名称
var EclipticNames = []string{"黄道", "黑道"}

// Ecliptic 黄道黑道
type Ecliptic struct {
	LoopTyme
}

func (Ecliptic) FromIndex(index int) Ecliptic {
	return Ecliptic{LoopTyme{}.FromIndex(EclipticNames, index)}
}

func (Ecliptic) FromName(name string) (*Ecliptic, error) {
	p, err := LoopTyme{}.FromName(EclipticNames, name)
	if err != nil {
		return nil, err
	}
	return &Ecliptic{*p}, nil
}

func (o Ecliptic) Next(n int) Ecliptic {
	return o.FromIndex(o.nextIndex(n))
}

// GetLuck 吉凶
func (o Ecliptic) GetLuck() Luck {
	return Luck{}.FromIndex(o.index)
}
