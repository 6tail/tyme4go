package tyme

import "fmt"

// SolarYear 公历年
type SolarYear struct {
	AbstractTyme
	// 年
	year int
}

func (SolarYear) FromYear(year int) (SolarYear, error) {
	if year < 1 || year > 9999 {
		return SolarYear{}, fmt.Errorf(fmt.Sprintf("illegal solar year: %d", year))
	}
	return SolarYear{
		year: year,
	}, nil
}

// GetYear 年
func (o SolarYear) GetYear() int {
	return o.year
}

// GetDayCount 天数（1582年355天，平年365天，闰年366天）
func (o SolarYear) GetDayCount() int {
	if 1582 == o.year {
		return 355
	}
	if o.IsLeap() {
		return 366
	}
	return 365
}

// IsLeap 是否闰年(1582年以前，使用儒略历，能被4整除即为闰年。以后采用格里历，四年一闰，百年不闰，四百年再闰。)
func (o SolarYear) IsLeap() bool {
	if o.year < 1600 {
		return o.year%4 == 0
	}
	return (o.year%4 == 0 && o.year%100 != 0) || (o.year%400 == 0)
}

func (o SolarYear) GetName() string {
	return fmt.Sprintf("%d年", o.year)
}

func (o SolarYear) String() string {
	return o.GetName()
}

func (o SolarYear) Next(n int) SolarYear {
	y, _ := SolarYear{}.FromYear(o.year + n)
	return y
}

// GetMonths 月份列表，1年有12个月。
func (o SolarYear) GetMonths() []SolarMonth {
	var l []SolarMonth
	for i := 1; i < 13; i++ {
		m, _ := SolarMonth{}.FromYm(o.year, i)
		l[i-1] = m
	}
	return l
}

// GetSeasons 季度列表，半年有2个季度。
func (o SolarYear) GetSeasons() []SolarSeason {
	var l []SolarSeason
	y := o.GetYear()
	for i := 0; i < 4; i++ {
		m, _ := SolarSeason{}.FromIndex(y, i)
		l[i] = m
	}
	return l
}

// GetHalfYears 半年列表，1年有2个半年。
func (o SolarYear) GetHalfYears() []SolarHalfYear {
	var l []SolarHalfYear
	y := o.GetYear()
	for i := 0; i < 2; i++ {
		m, _ := SolarHalfYear{}.FromIndex(y, i)
		l[i] = m
	}
	return l
}
