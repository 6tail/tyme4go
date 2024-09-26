package tyme

// YinYang 阴阳
type YinYang int

const (
	// YIN 阴
	YIN YinYang = iota
	// YANG 阳
	YANG
)

func (d YinYang) GetCode() int {
	return int(d)
}

func NewYinYang(code int) YinYang {
	switch code {
	case 1:
		return YANG
	default:
		return YIN
	}
}

func (d YinYang) GetName() string {
	switch d {
	case YIN:
		return "阴"
	default:
		return "阳"
	}
}

func (d YinYang) String() string {
	return d.GetName()
}
