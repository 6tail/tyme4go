package tyme

import (
	"fmt"
	"math"
)

// EightChar 八字
type EightChar struct {
	AbstractCulture
	// 三柱
	threePillars ThreePillars
	// 时柱
	hour SixtyCycle
}

func (EightChar) New(year string, month string, day string, hour string) (*EightChar, error) {
	t, err := ThreePillars{}.New(year, month, day)
	if err != nil {
		return nil, err
	}
	h, err := SixtyCycle{}.FromName(hour)
	if err != nil {
		return nil, err
	}
	return &EightChar{
		threePillars: *t,
		hour:         *h,
	}, nil
}

func (EightChar) FromSixtyCycle(year SixtyCycle, month SixtyCycle, day SixtyCycle, hour SixtyCycle) EightChar {
	t := ThreePillars{}.FromSixtyCycle(year, month, day)
	return EightChar{
		threePillars: t,
		hour:         hour,
	}
}

// GetYear 年柱
func (o EightChar) GetYear() SixtyCycle {
	return o.threePillars.GetYear()
}

// GetMonth 月柱
func (o EightChar) GetMonth() SixtyCycle {
	return o.threePillars.GetMonth()
}

// GetDay 日柱
func (o EightChar) GetDay() SixtyCycle {
	return o.threePillars.GetDay()
}

// GetHour 时柱
func (o EightChar) GetHour() SixtyCycle {
	return o.hour
}

// GetFetalOrigin 胎元
func (o EightChar) GetFetalOrigin() SixtyCycle {
	m := o.GetMonth()
	t, _ := SixtyCycle{}.FromName(m.GetHeavenStem().Next(1).GetName() + m.GetEarthBranch().Next(3).GetName())
	return *t
}

// GetFetalBreath 胎息
func (o EightChar) GetFetalBreath() SixtyCycle {
	d := o.GetDay()
	t, _ := SixtyCycle{}.FromName(d.GetHeavenStem().Next(5).GetName() + EarthBranch{}.FromIndex(13-d.GetEarthBranch().GetIndex()).GetName())
	return *t
}

// GetOwnSign 命宫
func (o EightChar) GetOwnSign() SixtyCycle {
	m := o.GetMonth().GetEarthBranch().GetIndex() - 1
	if m < 1 {
		m += 12
	}
	h := o.hour.GetEarthBranch().GetIndex() - 1
	if h < 1 {
		h += 12
	}
	offset := m + h
	if offset >= 14 {
		offset = 26 - offset
	} else {
		offset = 14 - offset
	}
	t, _ := SixtyCycle{}.FromName(HeavenStem{}.FromIndex((o.GetYear().GetHeavenStem().GetIndex()+1)*2+offset-1).GetName() + EarthBranch{}.FromIndex(offset+1).GetName())
	return *t
}

// GetBodySign 身宫
func (o EightChar) GetBodySign() SixtyCycle {
	offset := o.GetMonth().GetEarthBranch().GetIndex() - 1
	if offset < 1 {
		offset += 12
	}
	offset += o.hour.GetEarthBranch().GetIndex() + 1
	if offset > 12 {
		offset -= 12
	}
	t, _ := SixtyCycle{}.FromName(HeavenStem{}.FromIndex((o.GetYear().GetHeavenStem().GetIndex()+1)*2+offset-1).GetName() + EarthBranch{}.FromIndex(offset+1).GetName())
	return *t
}

// Deprecated: Use SixtyCycleDay.GetDuty instead.
func (o EightChar) GetDuty() Duty {
	return Duty{}.FromIndex(o.GetDay().GetEarthBranch().GetIndex() - o.GetMonth().GetEarthBranch().GetIndex())
}

func (o EightChar) GetName() string {
	return fmt.Sprintf("%v %v", o.threePillars, o.hour)
}

func (o EightChar) String() string {
	return o.GetName()
}

func (o EightChar) Equals(target EightChar) bool {
	return o.String() == target.String()
}

// GetSolarTimes 公历时刻列表(支持1-9999年)
func (o EightChar) GetSolarTimes(startYear int, endYear int) []SolarTime {
	var l []SolarTime
	year := o.GetYear()
	month := o.GetMonth()
	day := o.GetDay()
	// 月地支距寅月的偏移值
	m := month.GetEarthBranch().Next(-2).GetIndex()
	// 月天干要一致
	if (!HeavenStem{}.FromIndex((year.GetHeavenStem().GetIndex()+1)*2 + m).Equals(month.GetHeavenStem())) {
		return l
	}
	// 1年的立春是辛酉，序号57
	y := year.Next(-57).GetIndex() + 1
	// 节令偏移值
	m *= 2
	// 时辰地支转时刻
	h := o.hour.GetEarthBranch().GetIndex() * 2
	hours := []int{h}
	if h == 0 {
		hours = []int{0, 23}
	}
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
		solarTime := term.GetJulianDay().GetSolarTime()
		if solarTime.GetYear() >= startYear {
			// 日干支和节令干支的偏移值
			solarDay := solarTime.GetSolarDay()
			d := day.Next(-solarDay.GetLunarDay().GetSixtyCycle().GetIndex()).GetIndex()
			if d > 0 {
				// 从节令推移天数
				solarDay = solarDay.Next(d)
			}
			for _, hour := range hours {
				mi := 0
				s := 0
				if d == 0 && hour == solarTime.GetHour() {
					// 如果正好是节令当天，且小时和节令的小时数相等的极端情况，把分钟和秒钟带上
					mi = solarTime.GetMinute()
					s = solarTime.GetSecond()
				}
				time, _ := SolarTime{}.FromYmdHms(solarDay.GetYear(), solarDay.GetMonth(), solarDay.GetDay(), hour, mi, s)
				t := *time
				if d == 30 {
					t = t.Next(-3600)
				}
				// 验证一下
				if t.GetLunarHour().GetEightChar().Equals(o) {
					l = append(l, t)
				}
			}
		}
		y += 60
	}
	return l
}
