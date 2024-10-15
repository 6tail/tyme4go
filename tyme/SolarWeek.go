package tyme

import (
	"fmt"
)

var SolarWeekNames = []string{"第一周", "第二周", "第三周", "第四周", "第五周", "第六周"}

// SolarWeek 公历周
type SolarWeek struct {
	AbstractTyme
	// 公历月
	month SolarMonth
	// 索引，0-5
	index int
	// 起始星期
	start Week
}

func (SolarWeek) FromYm(year int, month int, index int, start int) (*SolarWeek, error) {
	if index < 0 || index > 5 {
		return nil, fmt.Errorf(fmt.Sprintf("illegal solar week index: %d", index))
	}
	if start < 0 || start > 6 {
		return nil, fmt.Errorf(fmt.Sprintf("illegal solar week start: %d", start))
	}
	m, err := SolarMonth{}.FromYm(year, month)
	if err != nil {
		return nil, err
	}
	if index >= m.GetWeekCount(start) {
		return nil, fmt.Errorf(fmt.Sprintf("illegal solar week index: %d in month: %v", index, m))
	}
	return &SolarWeek{
		month: m,
		index: index,
		start: Week{}.FromIndex(start),
	}, nil
}

// GetSolarMonth 公历月
func (o SolarWeek) GetSolarMonth() SolarMonth {
	return o.month
}

// GetYear 年
func (o SolarWeek) GetYear() int {
	return o.month.GetYear()
}

// GetMonth 月
func (o SolarWeek) GetMonth() int {
	return o.month.GetMonth()
}

// GetIndex 周索引(0-5)
func (o SolarWeek) GetIndex() int {
	return o.index
}

// GetStart 起始星期
func (o SolarWeek) GetStart() Week {
	return o.start
}

// GetIndexInYear 位于当年的索引(0-11)
func (o SolarWeek) GetIndexInYear() int {
	i := 0
	firstDay := o.GetFirstDay()
	// 今年第1周
	t, _ := SolarWeek{}.FromYm(o.GetYear(), 1, 0, o.start.GetIndex())
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
	return fmt.Sprintf("%v%v", o.month, o.GetName())
}

// GetFirstDay 本周第1天
func (o SolarWeek) GetFirstDay() SolarDay {
	firstDay, _ := SolarDay{}.FromYmd(o.GetYear(), o.GetMonth(), 1)
	return firstDay.Next(o.index*7 - o.IndexOf(firstDay.GetWeek().GetIndex()-o.start.GetIndex(), 7))
}

func (o SolarWeek) Next(n int) SolarWeek {
	startIndex := o.start.GetIndex()
	d := o.index
	m := o.month
	if n > 0 {
		d += n
		weekCount := m.GetWeekCount(startIndex)
		for d >= weekCount {
			d -= weekCount
			m = m.Next(1)
			day, _ := SolarDay{}.FromYmd(m.GetYear(), m.GetMonth(), 1)
			if !day.GetWeek().Equals(o.start) {
				d += 1
			}
			weekCount = m.GetWeekCount(startIndex)
		}
	} else if n < 0 {
		d += n
		for d < 0 {
			day, _ := SolarDay{}.FromYmd(m.GetYear(), m.GetMonth(), 1)
			if !day.GetWeek().Equals(o.start) {
				d -= 1
			}
			m = m.Next(-1)
			d += m.GetWeekCount(startIndex)
		}
	}
	t, _ := SolarWeek{}.FromYm(m.GetYear(), m.GetMonth(), d, startIndex)
	return *t
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
