package tyme

// HideHeavenStemType 藏干类型
type HideHeavenStemType int

const (
	// RESIDUAL 余气
	RESIDUAL HideHeavenStemType = iota
	// MIDDLE 中气
	MIDDLE
	// MAIN 本气
	MAIN
)

func (d HideHeavenStemType) GetCode() int {
	return int(d)
}

func NewHideHeavenStemType(code int) HideHeavenStemType {
	switch code {
	case 1:
		return MIDDLE
	case 2:
		return MAIN
	default:
		return RESIDUAL
	}
}

func (d HideHeavenStemType) GetName() string {
	switch d {
	case MIDDLE:
		return "中气"
	case MAIN:
		return "本气"
	default:
		return "余气"
	}
}

func (d HideHeavenStemType) String() string {
	return d.GetName()
}
