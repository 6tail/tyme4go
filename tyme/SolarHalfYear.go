package tyme

import (
	"fmt"
)

var SolarHalfYearNames = []string{"上半年", "下半年"}

// SolarHalfYear 公历半年
type SolarHalfYear struct {
	AbstractTyme
	// 年
	year SolarYear
	// 索引，0-1
	index int
}

func (SolarHalfYear) FromIndex(year int, index int) (SolarHalfYear, error) {
	if index < 0 || index > 1 {
		return SolarHalfYear{}, fmt.Errorf(fmt.Sprintf("illegal solar half year index: %d", index))
	}
	y, err := SolarYear{}.FromYear(year)
	if err != nil {
		return SolarHalfYear{}, err
	}
	return SolarHalfYear{
		year:  y,
		index: index,
	}, nil
}

// GetSolarYear 公历年
func (o SolarHalfYear) GetSolarYear() SolarYear {
	return o.year
}

// GetYear 年
func (o SolarHalfYear) GetYear() int {
	return o.year.GetYear()
}

// GetIndex 索引，0-1
func (o SolarHalfYear) GetIndex() int {
	return o.index
}

func (o SolarHalfYear) GetName() string {
	return SolarHalfYearNames[o.index]
}

func (o SolarHalfYear) String() string {
	return fmt.Sprintf("%v%v", o.year, o.GetName())
}

func (o SolarHalfYear) Next(n int) SolarHalfYear {
	i := o.index
	y := o.GetYear()
	if n != 0 {
		i += n
		y += i / 2
		i %= 2
		if i < 0 {
			i += 2
			y -= 1
		}
	}
	obj, _ := SolarHalfYear{}.FromIndex(y, i)
	return obj
}

// GetMonths 月份列表，半年有6个月。
func (o SolarHalfYear) GetMonths() []SolarMonth {
	var l []SolarMonth
	y := o.GetYear()
	for i := 1; i < 7; i++ {
		m, _ := SolarMonth{}.FromYm(y, o.index*6+i)
		l[i-1] = m
	}
	return l
}

// GetSeasons 季度列表，半年有2个季度。
func (o SolarHalfYear) GetSeasons() []SolarSeason {
	var l []SolarSeason
	y := o.GetYear()
	for i := 0; i < 2; i++ {
		m, _ := SolarSeason{}.FromIndex(y, o.index*2+i)
		l[i] = m
	}
	return l
}
