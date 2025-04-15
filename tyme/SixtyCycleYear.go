package tyme

import (
	"fmt"
	"math"
)

// SixtyCycleYear 干支年
type SixtyCycleYear struct {
	AbstractTyme
	// 年
	year int
}

func (SixtyCycleYear) FromYear(year int) (*SixtyCycleYear, error) {
	if year < -1 || year > 9999 {
		return nil, fmt.Errorf(fmt.Sprintf("illegal sixty cycle year: %d", year))
	}
	return &SixtyCycleYear{
		year: year,
	}, nil
}

// GetYear 年
func (o SixtyCycleYear) GetYear() int {
	return o.year
}

// GetName 名称
func (o SixtyCycleYear) GetName() string {
	return fmt.Sprintf("%s年", o.GetSixtyCycle())
}

func (o SixtyCycleYear) String() string {
	return o.GetName()
}

func (o SixtyCycleYear) Next(n int) SixtyCycleYear {
	y, _ := SixtyCycleYear{}.FromYear(o.year + n)
	return *y
}

// GetSixtyCycle 干支
func (o SixtyCycleYear) GetSixtyCycle() SixtyCycle {
	return SixtyCycle{}.FromIndex(o.year - 4)
}

// GetTwenty 运
func (o SixtyCycleYear) GetTwenty() Twenty {
	return Twenty{}.FromIndex(int(math.Floor(float64(o.year-1864) / 20)))
}

// GetNineStar 九星
func (o SixtyCycleYear) GetNineStar() NineStar {
	return NineStar{}.FromIndex(63 + o.GetTwenty().GetSixty().GetIndex()*3 - o.GetSixtyCycle().GetIndex())
}

// GetJupiterDirection  太岁方位
func (o SixtyCycleYear) GetJupiterDirection() Direction {
	return Direction{}.FromIndex([]int{0, 7, 7, 2, 3, 3, 8, 1, 1, 6, 0, 0}[o.GetSixtyCycle().GetEarthBranch().GetIndex()])
}

// GetFirstMonth 首月（依据五虎遁和正月起寅的规律）
func (o SixtyCycleYear) GetFirstMonth() SixtyCycleMonth {
	h := HeavenStem{}.FromIndex((o.GetSixtyCycle().GetHeavenStem().GetIndex() + 1) * 2)
	m, _ := SixtyCycle{}.FromName(h.GetName() + "寅")
	return SixtyCycleMonth{}.New(o, *m)
}

// GetMonths 干支月列表
func (o SixtyCycleYear) GetMonths() []SixtyCycleMonth {
	var l []SixtyCycleMonth
	m := o.GetFirstMonth()
	l = append(l, m)
	for i := 1; i < 12; i++ {
		l = append(l, m.Next(i))
	}
	return l
}
