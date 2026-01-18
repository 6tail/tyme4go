package tyme

import (
	"fmt"
)

var SolarWeekNames = []string{"第一周", "第二周", "第三周", "第四周", "第五周", "第六周"}

// SolarWeek 公历周
type SolarWeek struct {
	WeekUnit
}

func (SolarWeek) Validate(year int, month int, index int, start int) error {
	err := WeekUnit{}.Validate(year, month, index, start)
	if err != nil {
		return err
	}
	m, err := SolarMonth{}.FromYm(year, month)
	if err != nil {
		return err
	}
	if index >= m.GetWeekCount(start) {
		return fmt.Errorf(fmt.Sprintf("illegal solar week index: %d in month: %v", index, m))
	}
	return nil
}

func (SolarWeek) FromYm(year int, month int, index int, start int) (*SolarWeek, error) {
	err := SolarWeek{}.Validate(year, month, index, start)
	if err != nil {
		return nil, err
	}
	return &SolarWeek{
		WeekUnit{
			MonthUnit{
				YearUnit{
					year: year,
				},
				month,
			},
			index,
			start,
		},
	}, nil
}

// GetSolarMonth 公历月
func (o SolarWeek) GetSolarMonth() SolarMonth {
	m, _ := SolarMonth{}.FromYm(o.year, o.month)
	return *m
}

// GetIndexInYear 位于当年的索引(0-11)
func (o SolarWeek) GetIndexInYear() int {
	i := 0
	firstDay := o.GetFirstDay()
	// 今年第1周
	t, _ := SolarWeek{}.FromYm(o.year, 1, 0, o.start)
	w := *t
	for ; !w.GetFirstDay().Equals(firstDay); w = w.Next(1) {
		i += 1
	}
	return i
}

func (o SolarWeek) GetName() string {
	return SolarWeekNames[o.index]
}

func (o SolarWeek) String() string {
	return fmt.Sprintf("%v%v", o.GetSolarMonth(), o.GetName())
}

// GetFirstDay 本周第1天
func (o SolarWeek) GetFirstDay() SolarDay {
	firstDay, _ := SolarDay{}.FromYmd(o.year, o.month, 1)
	return firstDay.Next(o.index*7 - o.IndexOf(firstDay.GetWeek().GetIndex()-o.start, 7))
}

func (o SolarWeek) Next(n int) SolarWeek {
	d := o.index
	m := o.GetSolarMonth()
	if n > 0 {
		d += n
		weekCount := m.GetWeekCount(o.start)
		for d >= weekCount {
			d -= weekCount
			m = m.Next(1)
			if m.GetFirstDay().GetWeek().index != o.start {
				d += 1
			}
			weekCount = m.GetWeekCount(o.start)
		}
	} else if n < 0 {
		d += n
		for d < 0 {
			if m.GetFirstDay().GetWeek().index != o.start {
				d -= 1
			}
			m = m.Next(-1)
			d += m.GetWeekCount(o.start)
		}
	}
	w, _ := SolarWeek{}.FromYm(m.GetYear(), m.GetMonth(), d, o.start)
	return *w
}

// GetDays 本周公历日列表
func (o SolarWeek) GetDays() []SolarDay {
	var l []SolarDay
	d := o.GetFirstDay()
	l = append(l, d)
	for i := 1; i < 7; i++ {
		l = append(l, d.Next(i))
	}
	return l
}

func (o SolarWeek) Equals(target SolarWeek) bool {
	return o.GetFirstDay().Equals(target.GetFirstDay())
}
