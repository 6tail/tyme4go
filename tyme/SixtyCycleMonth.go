package tyme

import (
	"fmt"
)

// SixtyCycleMonth 干支月
type SixtyCycleMonth struct {
	AbstractTyme
	// 干支年
	year SixtyCycleYear
	// 月柱
	month SixtyCycle
}

func (SixtyCycleMonth) New(year SixtyCycleYear, month SixtyCycle) SixtyCycleMonth {
	return SixtyCycleMonth{
		year:  year,
		month: month,
	}
}

func (SixtyCycleMonth) FromIndex(year int, index int) (*SixtyCycleMonth, error) {
	y, err := SixtyCycleYear{}.FromYear(year)
	if err != nil {
		return nil, err
	}
	m := y.GetFirstMonth().Next(index)
	return &m, nil
}

// GetSixtyCycleYear 干支年
func (o SixtyCycleMonth) GetSixtyCycleYear() SixtyCycleYear {
	return o.year
}

// GetYear 年柱
func (o SixtyCycleMonth) GetYear() SixtyCycle {
	return o.year.GetSixtyCycle()
}

// GetSixtyCycle 干支
func (o SixtyCycleMonth) GetSixtyCycle() SixtyCycle {
	return o.month
}

func (o SixtyCycleMonth) GetName() string {
	return o.month.String() + "月"
}

func (o SixtyCycleMonth) String() string {
	return fmt.Sprintf("%v%v", o.year, o.GetName())
}

func (o SixtyCycleMonth) Next(n int) SixtyCycleMonth {
	y, _ := SixtyCycleYear{}.FromYear((o.year.GetYear()*12 + o.GetIndexInYear() + n) / 12)
	return SixtyCycleMonth{}.New(*y, o.month.Next(n))
}

func (o SixtyCycleMonth) Equals(target SixtyCycleMonth) bool {
	return o.String() == target.String()
}

// GetIndexInYear 位于当年的索引(0-11)，寅月为0，依次类推
func (o SixtyCycleMonth) GetIndexInYear() int {
	return o.month.GetEarthBranch().Next(-2).GetIndex()
}

func (o SixtyCycleMonth) GetFirstDay() SixtyCycleDay {
	return SixtyCycleDay{}.FromSolarDay(SolarTerm{}.FromIndex(o.year.GetYear(), 3+o.GetIndexInYear()*2).GetSolarDay())
}

// GetDays 本月的干支日列表
func (o SixtyCycleMonth) GetDays() []SixtyCycleDay {
	var l []SixtyCycleDay
	d := o.GetFirstDay()
	for d.GetSixtyCycleMonth().Equals(o) {
		l = append(l, d)
		d = d.Next(1)
	}
	return l
}

// GetNineStar 九星
func (o SixtyCycleMonth) GetNineStar() NineStar {
	index := o.month.GetEarthBranch().GetIndex()
	if index < 2 {
		index += 3
	}
	return NineStar{}.FromIndex(27 - o.GetYear().GetEarthBranch().GetIndex()%3*3 - index)
}

// GetJupiterDirection 太岁方位
func (o SixtyCycleMonth) GetJupiterDirection() Direction {
	n := []int{7, -1, 1, 3}[o.month.GetEarthBranch().Next(-2).GetIndex()%4]
	if n == -1 {
		return o.month.GetHeavenStem().GetDirection()
	}
	return Direction{}.FromIndex(n)
}
