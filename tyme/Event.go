package tyme

import (
	"fmt"
	"regexp"
	"strings"
)

// Event 事件
type Event struct {
	AbstractCulture
	name string
	data string
}

func (Event) Validate(data string) error {
	if data == "" {
		return fmt.Errorf("illegal event data: empty")
	}
	if len(data) != 9 {
		return fmt.Errorf("illegal event data: " + data)
	}
	return nil
}

// Builder 事件构造器
func (Event) Builder() *EventBuilder {
	return newEventBuilder()
}

func NewEvent(name string, data string) (*Event, error) {
	err := Event{}.Validate(data)
	if err != nil {
		return nil, err
	}
	return &Event{name: name, data: data}, nil
}

func (o Event) getCharIndex(index int) int {
	return strings.IndexRune(EventManagerChars, rune(o.data[index]))
}

func (o Event) GetValue(index int) int {
	return o.getCharIndex(index) - 31
}

func (o Event) GetMonth(year int) []int {
	y := year
	m := o.GetValue(2)
	if m > 12 {
		m = 1
		y += 1
	}
	return []int{y, m}
}

// GetName 事件名称
func (o Event) GetName() string {
	return o.name
}

// GetData 事件数据
func (o Event) GetData() string {
	return o.data
}

// GetType 事件类型
func (o Event) GetType() EventType {
	return NewEventType(o.getCharIndex(1))
}

// GetStartYear 事件起始年
func (o Event) GetStartYear() int {
	n := 0
	size := len(EventManagerCharsRune)
	for i := 0; i < 3; i++ {
		n = n*size + o.getCharIndex(6+i)
	}
	return n
}

// FromName 从名称获取事件
func (Event) FromName(name string) *Event {
	re := regexp.MustCompile(fmt.Sprintf(EventManagerRegex, name))
	matches := re.FindStringSubmatch(EventManagerData)
	if len(matches) > 1 {
		e, err := NewEvent(name, matches[1])
		if err != nil {
			return nil
		}
		return e
	}
	return nil
}

// FromSolarDay 指定公历日的事件列表
func (Event) FromSolarDay(d SolarDay) []Event {
	var l []Event
	all := Event{}.All()
	for _, e := range all {
		t := e.GetSolarDay(d.year)
		if t != nil {
			if d.Equals(*t) {
				l = append(l, e)
			}
		}
	}
	return l
}

// All 所有事件
func (Event) All() []Event {
	re := regexp.MustCompile(fmt.Sprintf(EventManagerRegex, ".[^@]+"))
	matches := re.FindAllStringSubmatch(EventManagerData, -1)
	var l []Event
	for _, m := range matches {
		e, _ := NewEvent(m[2], m[1])
		l = append(l, *e)
	}
	return l
}

// GetSolarDay 公历日，如果当年没有该事件，返回nil
func (o Event) GetSolarDay(year int) *SolarDay {
	if year < o.GetStartYear() {
		return nil
	}
	var d *SolarDay
	switch o.GetType() {
	case SOLAR_DAY:
		d = o.getSolarDayBySolarDay(year)
		break
	case SOLAR_WEEK:
		d = o.getSolarDayByWeek(year)
		break
	case LUNAR_DAY:
		d = o.getSolarDayByLunarDay(year)
		break
	case TERM_DAY:
		d = o.getSolarDayByTerm(year)
		break
	case TERM_HS:
		d = o.getSolarDayByTermHeavenStem(year)
		break
	case TERM_EB:
		d = o.getSolarDayByTermEarthBranch(year)
		break
	}
	if d == nil {
		return nil
	}
	offset := o.GetValue(5)
	if offset == 0 {
		return d
	}
	t := d.Next(offset)
	return &t
}

func (o Event) getSolarDayBySolarDay(year int) *SolarDay {
	month := o.GetMonth(year)
	y := month[0]
	m := month[1]
	d := o.GetValue(3)
	delay := o.GetValue(4)

	sm, _ := SolarMonth{}.FromYm(y, m)
	lastDay := sm.GetDayCount()
	if d > lastDay {
		if delay == 0 {
			return nil
		}
		if delay < 0 {
			t, _ := SolarDay{}.FromYmd(y, m, d+delay)
			return t
		}
		t, _ := SolarDay{}.FromYmd(y, m, lastDay)
		n := t.Next(delay)
		return &n
	}
	t, _ := SolarDay{}.FromYmd(y, m, d)
	return t
}

func (o Event) getSolarDayByLunarDay(year int) *SolarDay {
	month := o.GetMonth(year)
	y := month[0]
	m := month[1]
	d := o.GetValue(3)
	delay := o.GetValue(4)

	lm, _ := LunarMonth{}.FromYm(y, m)
	lastDay := lm.GetDayCount()
	if d > lastDay {
		if delay == 0 {
			return nil
		}
		if delay < 0 {
			t, _ := LunarDay{}.FromYmd(y, m, d+delay)
			n := t.GetSolarDay()
			return &n
		}
		t, _ := LunarDay{}.FromYmd(y, m, lastDay)
		n := t.GetSolarDay().Next(delay)
		return &n
	}
	t, _ := LunarDay{}.FromYmd(y, m, d)
	n := t.GetSolarDay()
	return &n
}

func (o Event) getSolarDayByWeek(year int) *SolarDay {
	n := o.GetValue(3)
	if n == 0 {
		return nil
	}
	m, _ := SolarMonth{}.FromYm(year, o.GetValue(2))
	w := o.GetValue(4)

	if n > 0 {
		d := m.GetFirstDay()
		t := d.Next(d.GetWeek().StepsTo(w) + 7*n - 7)
		return &t
	}
	d, _ := SolarDay{}.FromYmd(year, m.month, m.GetDayCount())
	t := d.Next(d.GetWeek().StepsBackTo(w) + 7*n + 7)
	return &t
}

func (o Event) getSolarDayByTerm(year int) *SolarDay {
	d := SolarTerm{}.FromIndex(year, o.GetValue(2)).GetSolarDay()
	offset := o.GetValue(4)
	if offset == 0 {
		return &d
	}
	n := d.Next(offset)
	return &n
}

func (o Event) getSolarDayByTermHeavenStem(year int) *SolarDay {
	d := o.getSolarDayByTerm(year)
	n := d.Next(d.GetLunarDay().GetSixtyCycle().GetHeavenStem().StepsTo(o.GetValue(3)))
	return &n
}

func (o Event) getSolarDayByTermEarthBranch(year int) *SolarDay {
	d := o.getSolarDayByTerm(year)
	n := d.Next(d.GetLunarDay().GetSixtyCycle().GetEarthBranch().StepsTo(o.GetValue(3)))
	return &n
}
