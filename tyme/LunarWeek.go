package tyme

import (
	"fmt"
)

var LunarWeekNames = []string{"第一周", "第二周", "第三周", "第四周", "第五周", "第六周"}

// LunarWeek 农历周
type LunarWeek struct {
	WeekUnit
}

func (LunarWeek) Validate(year int, month int, index int, start int) error {
	err := WeekUnit{}.Validate(year, month, index, start)
	if err != nil {
		return err
	}
	m, err := LunarMonth{}.FromYm(year, month)
	if err != nil {
		return err
	}
	if index >= m.GetWeekCount(start) {
		return fmt.Errorf(fmt.Sprintf("illegal lunar week index: %d in month: %v", index, m))
	}
	return nil
}

func (LunarWeek) FromYm(year int, month int, index int, start int) (*LunarWeek, error) {
	err := LunarWeek{}.Validate(year, month, index, start)
	if err != nil {
		return nil, err
	}
	return &LunarWeek{
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

// GetLunarMonth 农历月
func (o LunarWeek) GetLunarMonth() LunarMonth {
	m, _ := LunarMonth{}.FromYm(o.year, o.month)
	return *m
}

func (o LunarWeek) GetName() string {
	return LunarWeekNames[o.index]
}

func (o LunarWeek) String() string {
	return fmt.Sprintf("%v%v", o.GetLunarMonth(), o.GetName())
}

// GetFirstDay 本周第1天
func (o LunarWeek) GetFirstDay() LunarDay {
	firstDay, _ := LunarDay{}.FromYmd(o.year, o.month, 1)
	return firstDay.Next(o.index*7 - o.IndexOf(firstDay.GetWeek().GetIndex()-o.start, 7))
}

func (o LunarWeek) Next(n int) LunarWeek {
	d := o.index
	m := o.GetLunarMonth()
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
	t, _ := LunarWeek{}.FromYm(m.GetYear(), m.GetMonthWithLeap(), d, o.start)
	return *t
}

// GetDays 本周农历日列表
func (o LunarWeek) GetDays() []LunarDay {
	var l []LunarDay
	d := o.GetFirstDay()
	l = append(l, d)
	for i := 1; i < 7; i++ {
		l = append(l, d.Next(i))
	}
	return l
}

func (o LunarWeek) Equals(target LunarWeek) bool {
	return o.GetFirstDay().Equals(target.GetFirstDay())
}
