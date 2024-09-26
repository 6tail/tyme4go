package tyme

import (
	"fmt"
)

var LunarWeekNames = []string{"第一周", "第二周", "第三周", "第四周", "第五周", "第六周"}

// LunarWeek 农历周
type LunarWeek struct {
	AbstractTyme
	// 农历月
	month LunarMonth
	// 索引，0-5
	index int
	// 起始星期
	start Week
}

func (LunarWeek) FromYm(year int, month int, index int, start int) (*LunarWeek, error) {
	if index < 0 || index > 5 {
		return nil, fmt.Errorf(fmt.Sprintf("illegal lunar week index: %d", index))
	}
	if start < 0 || start > 6 {
		return nil, fmt.Errorf(fmt.Sprintf("illegal lunar week start: %d", start))
	}
	m, err := LunarMonth{}.FromYm(year, month)
	if err != nil {
		return nil, err
	}
	if index >= m.GetWeekCount(start) {
		return nil, fmt.Errorf(fmt.Sprintf("illegal lunar week index: %d in month: %v", index, m))
	}
	return &LunarWeek{
		month: *m,
		index: index,
		start: Week{}.FromIndex(start),
	}, nil
}

// GetLunarMonth 农历月
func (o LunarWeek) GetLunarMonth() LunarMonth {
	return o.month
}

// GetYear 年
func (o LunarWeek) GetYear() int {
	return o.month.GetYear()
}

// GetMonth 月
func (o LunarWeek) GetMonth() int {
	return o.month.GetMonthWithLeap()
}

// GetIndex 周索引(0-5)
func (o LunarWeek) GetIndex() int {
	return o.index
}

// GetStart 起始星期
func (o LunarWeek) GetStart() Week {
	return o.start
}

func (o LunarWeek) GetName() string {
	return LunarWeekNames[o.index]
}

func (o LunarWeek) String() string {
	return fmt.Sprintf("%v%v", o.month, o.GetName())
}

// GetFirstDay 本周第1天
func (o LunarWeek) GetFirstDay() LunarDay {
	firstDay, _ := LunarDay{}.FromYmd(o.GetYear(), o.GetMonth(), 1)
	return firstDay.Next(o.index*7 - o.IndexOf(firstDay.GetWeek().GetIndex()-o.start.GetIndex(), 7))
}

func (o LunarWeek) Next(n int) LunarWeek {
	startIndex := o.start.GetIndex()
	d := o.index
	m := o.month
	if n > 0 {
		d += n
		weekCount := m.GetWeekCount(startIndex)
		for d >= weekCount {
			d -= weekCount
			m = m.Next(1)
			day, _ := LunarDay{}.FromYmd(m.GetYear(), m.GetMonthWithLeap(), 1)
			if !day.GetWeek().Equals(o.start) {
				d += 1
			}
			weekCount = m.GetWeekCount(startIndex)
		}
	} else if n < 0 {
		d += n
		for d < 0 {
			day, _ := LunarDay{}.FromYmd(m.GetYear(), m.GetMonthWithLeap(), 1)
			if !day.GetWeek().Equals(o.start) {
				d -= 1
			}
			m = m.Next(-1)
			d += m.GetWeekCount(startIndex)
		}
	}
	t, _ := LunarWeek{}.FromYm(m.GetYear(), m.GetMonthWithLeap(), d, startIndex)
	return *t
}

// GetDays 本周农历日列表
func (o LunarWeek) GetDays() []LunarDay {
	var l []LunarDay
	d := o.GetFirstDay()
	l[0] = d
	for i := 1; i < 7; i++ {
		l[i] = d.Next(i)
	}
	return l
}

func (o LunarWeek) Equals(target LunarWeek) bool {
	return o.GetFirstDay().Equals(target.GetFirstDay())
}
