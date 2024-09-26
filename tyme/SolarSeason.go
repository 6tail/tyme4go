package tyme

import (
	"fmt"
)

var SolarSeasonNames = []string{"一季度", "二季度", "三季度", "四季度"}

// SolarSeason 公历季度
type SolarSeason struct {
	AbstractTyme
	// 公历年
	year SolarYear
	// 索引，0-1
	index int
}

func (SolarSeason) FromIndex(year int, index int) (SolarSeason, error) {
	if index < 0 || index > 3 {
		return SolarSeason{}, fmt.Errorf(fmt.Sprintf("illegal solar season index: %d", index))
	}
	y, err := SolarYear{}.FromYear(year)
	if err != nil {
		return SolarSeason{}, err
	}
	return SolarSeason{
		year:  y,
		index: index,
	}, nil
}

// GetSolarYear 公历年
func (o SolarSeason) GetSolarYear() SolarYear {
	return o.year
}

// GetYear 年
func (o SolarSeason) GetYear() int {
	return o.year.GetYear()
}

// GetIndex 索引，0-1
func (o SolarSeason) GetIndex() int {
	return o.index
}

func (o SolarSeason) GetName() string {
	return SolarSeasonNames[o.index]
}

func (o SolarSeason) String() string {
	return fmt.Sprintf("%v%v", o.year, o.GetName())
}

func (o SolarSeason) Next(n int) SolarSeason {
	i := o.index
	y := o.GetYear()
	if n != 0 {
		i += n
		y += i / 4
		i %= 4
		if i < 0 {
			i += 4
			y -= 1
		}
	}
	obj, _ := SolarSeason{}.FromIndex(y, i)
	return obj
}

// GetMonths 月份列表，1季度有3个月。
func (o SolarSeason) GetMonths() []SolarMonth {
	var l []SolarMonth
	y := o.GetYear()
	for i := 1; i < 4; i++ {
		m, _ := SolarMonth{}.FromYm(y, o.index*3+i)
		l[i-1] = m
	}
	return l
}
