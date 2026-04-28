package tyme

import "fmt"

// LunarFestivalNames 农历传统节日名称
var LunarFestivalNames = []string{"春节", "元宵节", "龙头节", "上巳节", "清明节", "端午节", "七夕节", "中元节", "中秋节", "重阳节", "冬至节", "腊八节", "除夕"}

// LunarFestivalData 农历传统节日数据
var LunarFestivalData = "2VV__0002Vj__0002WW__0002XX__0003b___0002ZZ__0002bb__0002bj__0002cj__0002dd__0003s___0002gc__0002hV_U000"

// LunarFestival 农历传统节日（依据国家标准《农历的编算和颁行》GB/T 33661-2017）
type LunarFestival struct {
	AbstractFestival
}

func (LunarFestival) New(festivalType FestivalType, index int, event Event, day LunarDay) LunarFestival {
	return LunarFestival{
		AbstractFestival{
			festivalType: festivalType,
			index:        index,
			event:        event,
			day:          day.DayUnit,
		},
	}
}

func (LunarFestival) FromIndex(year int, index int) *LunarFestival {
	if index < 0 || index >= len(LunarFestivalNames) {
		return nil
	}
	start := index * 8
	e, _ := NewEvent(LunarFestivalNames[index], "@"+LunarFestivalData[start:start+8])
	switch e.GetType() {
	case LUNAR_DAY:
		m := e.GetMonth(year)
		d, _ := LunarDay{}.FromYmd(m[0], m[1], e.GetValue(3))
		offset := e.GetValue(5)
		if 0 == offset {
			f := LunarFestival{}.New(DAY, index, *e, *d)
			return &f
		}
		f := LunarFestival{}.New(DAY, index, *e, d.Next(offset))
		return &f
	case TERM_DAY:
		f := LunarFestival{}.New(TERM, index, *e, SolarTerm{}.FromIndex(year, e.GetValue(2)).GetSolarDay().GetLunarDay())
		return &f
	default:
	}
	return nil
}

func (LunarFestival) FromYmd(year int, month int, day int) *LunarFestival {
	d, err := LunarDay{}.FromYmd(year, month, day)
	if err != nil {
		return nil
	}
	for i, name := range LunarFestivalNames {
		start := i * 8
		e, _ := NewEvent(name, "@"+LunarFestivalData[start:start+8])
		switch e.GetType() {
		case LUNAR_DAY:
			offset := e.GetValue(5)
			if 0 == offset {
				if d.month == e.GetValue(2) && d.day == e.GetValue(3) {
					f := LunarFestival{}.New(DAY, i, *e, *d)
					return &f
				}
			} else {
				m := e.GetMonth(d.year)
				next := d.Next(-offset)
				if next.year == m[0] && next.month == m[1] && next.day == e.GetValue(3) {
					f := LunarFestival{}.New(DAY, i, *e, *d)
					return &f
				}
			}
			break
		case TERM_DAY:
			term := d.GetSolarDay().GetTermDay()
			if term.dayIndex == 0 && term.GetSolarTerm().index == e.GetValue(2)%24 {
				f := LunarFestival{}.New(TERM, i, *e, *d)
				return &f
			}
			break
		default:
		}
	}
	return nil
}

// GetDay 农历日
func (o LunarFestival) GetDay() LunarDay {
	d, _ := LunarDay{}.FromYmd(o.day.year, o.day.month, o.day.day)
	return *d
}

// GetSolarTerm 节气
func (o LunarFestival) GetSolarTerm() *SolarTerm {
	t := o.GetDay().GetSolarDay().GetTermDay()
	if 0 == t.dayIndex {
		term := t.GetSolarTerm()
		return &term
	}
	return nil
}

func (o LunarFestival) Next(n int) *LunarFestival {
	size := len(LunarFestivalNames)
	i := o.index + n
	return LunarFestival{}.FromIndex((o.day.GetYear()*size+i)/size, o.IndexOf(i, size))
}

func (o LunarFestival) String() string {
	return fmt.Sprintf("%v %v", o.GetDay(), o.GetName())
}
