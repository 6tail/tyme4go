package tyme

import (
	"fmt"
)

// SolarFestivalNames 公历现代节日名称
var SolarFestivalNames = []string{"元旦", "妇女节", "植树节", "劳动节", "青年节", "儿童节", "建党节", "建军节", "教师节", "国庆节"}

// SolarFestivalData 公历现代节日数据
var SolarFestivalData = "0VV__0Ux0Xc__0Ux0Xg__0_Q0ZV__0Ux0ZY__0Ux0aV__0Ux0bV__0Uo0cV__0Ug0de__0_V0eV__0Ux"

// SolarFestival 公历现代节日
type SolarFestival struct {
	AbstractFestival
}

func (SolarFestival) New(festivalType FestivalType, index int, event Event, day SolarDay) SolarFestival {
	return SolarFestival{
		AbstractFestival{
			festivalType: festivalType,
			index:        index,
			event:        event,
			day:          day.DayUnit,
		},
	}
}

func (SolarFestival) FromIndex(year int, index int) *SolarFestival {
	if index < 0 || index >= len(SolarFestivalNames) {
		return nil
	}
	start := index * 8
	e, _ := NewEvent(SolarFestivalNames[index], "@"+SolarFestivalData[start:start+8])
	if year < e.GetStartYear() {
		return nil
	}
	d, err := SolarDay{}.FromYmd(year, e.GetValue(2), e.GetValue(3))
	if err != nil {
		return nil
	}
	f := SolarFestival{}.New(DAY, index, *e, *d)
	return &f
}

func (SolarFestival) FromYmd(year int, month int, day int) *SolarFestival {
	d, err := SolarDay{}.FromYmd(year, month, day)
	if err != nil {
		return nil
	}
	for i, name := range SolarFestivalNames {
		start := i * 8
		e, _ := NewEvent(name, "@"+SolarFestivalData[start:start+8])
		if d.year >= e.GetStartYear() && d.month == e.GetValue(2) && d.day == e.GetValue(3) {
			f := SolarFestival{}.New(DAY, i, *e, *d)
			return &f
		}
	}
	return nil
}

// GetType 节日类型
func (o SolarFestival) GetType() FestivalType {
	return o.festivalType
}

// GetDay 公历日
func (o SolarFestival) GetDay() SolarDay {
	d, _ := SolarDay{}.FromYmd(o.day.year, o.day.month, o.day.day)
	return *d
}

// GetStartYear 起始年
func (o SolarFestival) GetStartYear() int {
	return o.event.GetStartYear()
}

func (o SolarFestival) Next(n int) *SolarFestival {
	size := len(SolarFestivalNames)
	i := o.index + n
	return SolarFestival{}.FromIndex((o.day.GetYear()*size+i)/size, o.IndexOf(i, size))
}

func (o SolarFestival) String() string {
	return fmt.Sprintf("%v %v", o.GetDay(), o.GetName())
}
