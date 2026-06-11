package tyme

import (
	"fmt"
	"strconv"
)

// HijriYear 回历年
type HijriYear struct {
	YearUnit
}

func (HijriYear) Validate(year int) error {
	if year < -640 || year > 9666 {
		return fmt.Errorf("illegal hijri year: " + strconv.Itoa(year))
	}
	return nil
}

func (HijriYear) FromYear(year int) (*HijriYear, error) {
	err := HijriYear{}.Validate(year)
	if err != nil {
		return nil, err
	}
	return &HijriYear{
		YearUnit{
			year: year,
		},
	}, nil
}

// GetDayCount 天数（平年354天，闰年355天）
func (o HijriYear) GetDayCount() int {
	if o.IsLeap() {
		return 355
	}
	return 354
}

// IsLeap 是否闰年(1个闰周为30年，1个闰周中第2、5、7、10、13、16、18、21、24、26、29年为闰年)
func (o HijriYear) IsLeap() bool {
	i := o.IndexOf(o.year-1, 30)
	return i == 1 || i == 4 || i == 6 || i == 9 || i == 12 || i == 15 || i == 17 || i == 20 || i == 23 || i == 25 || i == 28
}

func (o HijriYear) GetName() string {
	return fmt.Sprintf("%d年", o.year)
}

func (o HijriYear) String() string {
	return o.GetName()
}

func (o HijriYear) Next(n int) HijriYear {
	y, _ := HijriYear{}.FromYear(o.year + n)
	return *y
}

// GetMonths 月份列表，1年有12个月。
func (o HijriYear) GetMonths() []HijriMonth {
	var l []HijriMonth
	for i := 1; i < 13; i++ {
		m, _ := HijriMonth{}.FromYm(o.year, i)
		l = append(l, *m)
	}
	return l
}

// GetFirstMonth 首月
func (o HijriYear) GetFirstMonth() HijriMonth {
	m, _ := HijriMonth{}.FromYm(o.year, 1)
	return *m
}
