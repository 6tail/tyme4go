package tyme

// EventBuilder 事件构造器
type EventBuilder struct {
	name string
	data []rune
}

func newEventBuilder() *EventBuilder {
	return &EventBuilder{
		name: "",
		data: []rune{'@', '_', '_', '_', '_', '_', '0', '0', '0'},
	}
}

// Name 设置事件名称
func (o EventBuilder) Name(name string) EventBuilder {
	o.name = name
	return o
}

func (o EventBuilder) getChar(index int) rune {
	return EventManagerCharsRune[index]
}

func (o EventBuilder) setValue(index int, n int) EventBuilder {
	o.data[index] = EventManagerCharsRune[31+n]
	return o
}

func (EventBuilder) encodeType(t EventType) rune {
	return EventManagerCharsRune[t.GetCode()]
}

func (o EventBuilder) content(t EventType, a int, b, c int) EventBuilder {
	o.data[1] = o.getChar(t.GetCode())
	return o.setValue(2, a).setValue(3, b).setValue(4, c)
}

// SolarDay 公历日期
func (o EventBuilder) SolarDay(solarMonth, solarDay, delayDays int) EventBuilder {
	return o.content(SOLAR_DAY, solarMonth, solarDay, delayDays)
}

// LunarDay 农历日期
func (o EventBuilder) LunarDay(lunarMonth, lunarDay, delayDays int) EventBuilder {
	return o.content(LUNAR_DAY, lunarMonth, lunarDay, delayDays)
}

// SolarWeek 几月第几个星期几
func (o EventBuilder) SolarWeek(solarMonth, weekIndex, week int) EventBuilder {
	return o.content(SOLAR_WEEK, solarMonth, weekIndex, week)
}

// TermDay 节气日期
func (o EventBuilder) TermDay(termIndex, delayDays int) EventBuilder {
	return o.content(TERM_DAY, termIndex, 0, delayDays)
}

// TermHeavenStem 节气天干
func (o EventBuilder) TermHeavenStem(termIndex, heavenStemIndex, delayDays int) EventBuilder {
	return o.content(TERM_HS, termIndex, heavenStemIndex, delayDays)
}

// TermEarthBranch 节气地支
func (o EventBuilder) TermEarthBranch(termIndex, earthBranchIndex, delayDays int) EventBuilder {
	return o.content(TERM_EB, termIndex, earthBranchIndex, delayDays)
}

// StartYear 起始年
func (o EventBuilder) StartYear(year int) EventBuilder {
	size := len(EventManagerCharsRune)
	n := year
	for i := 0; i < 3; i++ {
		o.data[8-i] = o.getChar(n % size)
		n /= size
	}
	return o
}

// Offset 偏移天数（最远支持-31至31天）
func (o EventBuilder) Offset(days int) EventBuilder {
	return o.setValue(5, days)
}

// Build 生成事件
func (o EventBuilder) Build() (*Event, error) {
	return NewEvent(o.name, string(o.data))
}
