package tyme

// FestivalType 节日类型
// Deprecated
type FestivalType int

const (
	// DAY 日期
	DAY FestivalType = iota
	// TERM 节气
	TERM
	// EVE 除夕
	EVE
)

func (o FestivalType) GetCode() int {
	return int(o)
}

func NewFestivalType(code int) FestivalType {
	switch code {
	case 1:
		return TERM
	case 2:
		return EVE
	default:
		return DAY
	}
}

func (o FestivalType) GetName() string {
	switch o {
	case DAY:
		return "日期"
	case TERM:
		return "节气"
	default:
		return "除夕"
	}
}

func (o FestivalType) String() string {
	return o.GetName()
}
