package tyme

import (
	"fmt"
	"math"
)

// ThreePillars 三柱（年柱、月柱、日柱）
type ThreePillars struct {
	AbstractCulture
	// 年柱
	year SixtyCycle
	// 月柱
	month SixtyCycle
	// 日柱
	day SixtyCycle
}

func (ThreePillars) New(year string, month string, day string) (*ThreePillars, error) {
	y, err := SixtyCycle{}.FromName(year)
	if err != nil {
		return nil, err
	}
	m, err := SixtyCycle{}.FromName(month)
	if err != nil {
		return nil, err
	}
	d, err := SixtyCycle{}.FromName(day)
	if err != nil {
		return nil, err
	}
	return &ThreePillars{
		year:  *y,
		month: *m,
		day:   *d,
	}, nil
}

func (ThreePillars) FromSixtyCycle(year SixtyCycle, month SixtyCycle, day SixtyCycle) ThreePillars {
	return ThreePillars{
		year:  year,
		month: month,
		day:   day,
	}
}

// GetYear 年柱
func (o ThreePillars) GetYear() SixtyCycle {
	return o.year
}

// GetMonth 月柱
func (o ThreePillars) GetMonth() SixtyCycle {
	return o.month
}

// GetDay 日柱
func (o ThreePillars) GetDay() SixtyCycle {
	return o.day
}

func (o ThreePillars) GetName() string {
	return fmt.Sprintf("%v %v %v", o.year, o.month, o.day)
}

func (o ThreePillars) String() string {
	return o.GetName()
}

func (o ThreePillars) Equals(target ThreePillars) bool {
	return o.String() == target.String()
}

// GetSolarDays 公历日列表(支持1-9999年)
func (o ThreePillars) GetSolarDays(startYear int, endYear int) []SolarDay {
	var l []SolarDay
	// 月地支距寅月的偏移值
	m := o.month.GetEarthBranch().Next(-2).GetIndex()
	// 月天干要一致
	if (!HeavenStem{}.FromIndex((o.year.GetHeavenStem().GetIndex()+1)*2 + m).Equals(o.month.GetHeavenStem())) {
		return l
	}
	// 1年的立春是辛酉，序号57
	y := o.year.Next(-57).GetIndex() + 1
	// 节令偏移值
	m *= 2
	baseYear := startYear - 1
	if baseYear > y {
		y += 60 * int(math.Ceil(float64(baseYear-y)/60.0))
	}
	for y <= endYear {
		// 立春为寅月的开始
		term := SolarTerm{}.FromIndex(y, 3)
		// 节令推移，年干支和月干支就都匹配上了
		if m > 0 {
			term = term.Next(m)
		}
		solarDay := term.GetSolarDay()
		if solarDay.GetYear() >= startYear {
			// 日干支和节令干支的偏移值
			d := o.day.Next(-solarDay.GetLunarDay().GetSixtyCycle().GetIndex()).GetIndex()
			if d > 0 {
				// 从节令推移天数
				solarDay = solarDay.Next(d)
			}
			// 验证一下
			if solarDay.GetSixtyCycleDay().GetThreePillars().Equals(o) {
				l = append(l, solarDay)
			}
		}
		y += 60
	}
	return l
}
