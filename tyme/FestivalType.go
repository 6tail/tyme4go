package tyme

// FestivalType 节日类型
type FestivalType int

const (
	// DAY 日期
	DAY FestivalType = iota
	// TERM 节气
	TERM
	// EVE 除夕
	EVE
)

func (d FestivalType) GetCode() int {
	return int(d)
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

func (d FestivalType) GetName() string {
	switch d {
	case DAY:
		return "日期"
	case TERM:
		return "节气"
	default:
		return "除夕"
	}
}

func (d FestivalType) String() string {
	return d.GetName()
}
