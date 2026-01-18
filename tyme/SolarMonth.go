package tyme

import (
	"fmt"
	"math"
)

var SolarMonthNames = []string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"}
var SolarMonthDays = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

// SolarMonth 公历月
type SolarMonth struct {
	MonthUnit
}

func (SolarMonth) Validate(year int, month int) error {
	if month < 1 || month > 12 {
		return fmt.Errorf(fmt.Sprintf("illegal solar month: %d", month))
	}
	return SolarYear{}.Validate(year)
}

func (SolarMonth) FromYm(year int, month int) (*SolarMonth, error) {
	err := SolarMonth{}.Validate(year, month)
	if err != nil {
		return nil, err
	}
	return &SolarMonth{
		MonthUnit{
			YearUnit{
				year: year,
			},
			month,
		},
	}, nil
}

// GetSolarYear 公历年
func (o SolarMonth) GetSolarYear() SolarYear {
	y, _ := SolarYear{}.FromYear(o.year)
	return *y
}

// GetDayCount 天数（1582年10月只有21天)
func (o SolarMonth) GetDayCount() int {
	if 1582 == o.GetYear() && 10 == o.month {
		return 21
	}
	d := SolarMonthDays[o.GetIndexInYear()]
	//公历闰年2月多一天
	if 2 == o.month && o.GetSolarYear().IsLeap() {
		d++
	}
	return d
}

// GetIndexInYear 位于当年的索引(0-11)
func (o SolarMonth) GetIndexInYear() int {
	return o.month - 1
}

// GetWeekCount 周数
func (o SolarMonth) GetWeekCount(start int) int {
	d, _ := SolarDay{}.FromYmd(o.GetYear(), o.month, 1)
	return int(math.Ceil(float64(o.IndexOf(d.GetWeek().GetIndex()-start, 7)+o.GetDayCount()) / 7))
}

func (o SolarMonth) GetName() string {
	return SolarMonthNames[o.GetIndexInYear()]
}

func (o SolarMonth) String() string {
	return fmt.Sprintf("%v%v", o.GetSolarYear(), o.GetName())
}

func (o SolarMonth) Next(n int) SolarMonth {
	i := o.month - 1 + n
	m, _ := SolarMonth{}.FromYm((o.year*12+i)/12, o.IndexOf(i, 12)+1)
	return *m
}

// GetWeeks 本月公历周列表
func (o SolarMonth) GetWeeks(start int) []SolarWeek {
	var l []SolarWeek
	size := o.GetWeekCount(start)
	for i := 0; i < size; i++ {
		w, _ := SolarWeek{}.FromYm(o.year, o.month, i, start)
		l = append(l, *w)
	}
	return l
}

// GetDays 本月公历日列表
func (o SolarMonth) GetDays() []SolarDay {
	var l []SolarDay
	size := o.GetDayCount()
	for i := 1; i <= size; i++ {
		d, _ := SolarDay{}.FromYmd(o.year, o.month, i)
		l = append(l, *d)
	}
	return l
}

// GetFirstDay 本月第1天
func (o SolarMonth) GetFirstDay() SolarDay {
	d, _ := SolarDay{}.FromYmd(o.year, o.month, 1)
	return *d
}
