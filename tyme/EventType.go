package tyme

// EventType 事件类型
type EventType int

const (
	// SOLAR_DAY 公历日期
	SOLAR_DAY EventType = iota
	// SOLAR_WEEK 几月第几个星期几
	SOLAR_WEEK
	// LUNAR_DAY 农历日期
	LUNAR_DAY
	// TERM_DAY 节气日期
	TERM_DAY
	// TERM_HS 节气天干
	TERM_HS
	// TERM_EB 节气地支
	TERM_EB
)

func (o EventType) GetCode() int {
	return int(o)
}

func NewEventType(code int) EventType {
	switch code {
	case 1:
		return SOLAR_WEEK
	case 2:
		return LUNAR_DAY
	case 3:
		return TERM_DAY
	case 4:
		return TERM_HS
	case 5:
		return TERM_EB
	default:
		return SOLAR_DAY
	}
}

func (o EventType) GetName() string {
	switch o {
	case SOLAR_WEEK:
		return "几月第几个星期几"
	case LUNAR_DAY:
		return "农历日期"
	case TERM_DAY:
		return "节气日期"
	case TERM_HS:
		return "节气天干"
	case TERM_EB:
		return "节气地支"
	default:
		return "公历日期"
	}
}

func (o EventType) String() string {
	return o.GetName()
}
