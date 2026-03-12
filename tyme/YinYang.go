package tyme

// YinYang 阴阳
type YinYang int

const (
	// YIN 阴
	YIN YinYang = iota
	// YANG 阳
	YANG
)

func (o YinYang) GetCode() int {
	return int(o)
}

func NewYinYang(code int) YinYang {
	switch code {
	case 1:
		return YANG
	default:
		return YIN
	}
}

func (o YinYang) GetName() string {
	switch o {
	case YIN:
		return "阴"
	default:
		return "阳"
	}
}

func (o YinYang) String() string {
	return o.GetName()
}
