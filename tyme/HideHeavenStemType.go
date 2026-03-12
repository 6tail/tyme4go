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

func (o HideHeavenStemType) GetCode() int {
	return int(o)
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

func (o HideHeavenStemType) GetName() string {
	switch o {
	case MIDDLE:
		return "中气"
	case MAIN:
		return "本气"
	default:
		return "余气"
	}
}

func (o HideHeavenStemType) String() string {
	return o.GetName()
}
