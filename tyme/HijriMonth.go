package tyme

import (
	"fmt"
	"strconv"
)

var HijriMonthNames = []string{"穆哈兰姆月", "色法尔月", "赖比尔·敖外鲁月", "赖比尔·阿色尼月", "主马达·敖外鲁月", "主马达·阿色尼月", "赖哲卜月", "舍尔邦月", "赖买丹月", "闪瓦鲁月", "都尔喀尔德月", "都尔黑哲月"}

// HijriMonth 回历月
type HijriMonth struct {
	MonthUnit
}

func (HijriMonth) Validate(year int, month int) error {
	if month < 1 || month > 12 {
		return fmt.Errorf("illegal hijri month: " + strconv.Itoa(month))
	}
	return HijriYear{}.Validate(year)
}

func (HijriMonth) FromYm(year int, month int) (*HijriMonth, error) {
	err := HijriMonth{}.Validate(year, month)
	if err != nil {
		return nil, err
	}
	return &HijriMonth{
		MonthUnit{
			YearUnit{
				year: year,
			},
			month,
		},
	}, nil
}

// GetHijriYear 回历年
func (o HijriMonth) GetHijriYear() HijriYear {
	y, _ := HijriYear{}.FromYear(o.year)
	return *y
}

// GetDayCount 天数（单数月30天，双数月29天，闰年第12月30天)
func (o HijriMonth) GetDayCount() int {
	d := 30
	if o.month%2 == 0 {
		d = 29
	}
	if o.month == 12 && o.GetHijriYear().IsLeap() {
		d++
	}
	return d
}

// GetIndexInYear 位于当年的索引(0-11)
func (o HijriMonth) GetIndexInYear() int {
	return o.month - 1
}

func (o HijriMonth) GetName() string {
	return HijriMonthNames[o.GetIndexInYear()]
}

func (o HijriMonth) String() string {
	return fmt.Sprintf("%v%v", o.GetHijriYear(), o.GetName())
}

func (o HijriMonth) Next(n int) HijriMonth {
	i := o.month - 1 + n
	m, _ := HijriMonth{}.FromYm((o.year*12+i)/12, o.IndexOf(i, 12)+1)
	return *m
}

// GetDays 本月回历日列表
func (o HijriMonth) GetDays() []HijriDay {
	var l []HijriDay
	size := o.GetDayCount()
	for i := 1; i <= size; i++ {
		d, _ := HijriDay{}.FromYmd(o.year, o.month, i)
		l = append(l, *d)
	}
	return l
}

// GetFirstDay 本月第1天
func (o HijriMonth) GetFirstDay() HijriDay {
	d, _ := HijriDay{}.FromYmd(o.year, o.month, 1)
	return *d
}
