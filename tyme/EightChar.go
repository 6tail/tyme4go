package tyme

import "fmt"

// EightChar 八字
type EightChar struct {
	AbstractCulture
	// 年柱
	year SixtyCycle
	// 月柱
	month SixtyCycle
	// 日柱
	day SixtyCycle
	// 时柱
	hour SixtyCycle
}

func (EightChar) New(year string, month string, day string, hour string) (*EightChar, error) {
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
	h, err := SixtyCycle{}.FromName(hour)
	if err != nil {
		return nil, err
	}
	return &EightChar{
		year:  *y,
		month: *m,
		day:   *d,
		hour:  *h,
	}, nil
}

// GetYear 年柱
func (o EightChar) GetYear() SixtyCycle {
	return o.year
}

// GetMonth 月柱
func (o EightChar) GetMonth() SixtyCycle {
	return o.month
}

// GetDay 日柱
func (o EightChar) GetDay() SixtyCycle {
	return o.day
}

// GetHour 时柱
func (o EightChar) GetHour() SixtyCycle {
	return o.hour
}

// GetFetalOrigin 胎元
func (o EightChar) GetFetalOrigin() SixtyCycle {
	t, _ := SixtyCycle{}.FromName(o.month.GetHeavenStem().Next(1).GetName() + o.month.GetEarthBranch().Next(3).GetName())
	return *t
}

// GetFetalBreath 胎息
func (o EightChar) GetFetalBreath() SixtyCycle {
	t, _ := SixtyCycle{}.FromName(o.day.GetHeavenStem().Next(5).GetName() + EarthBranch{}.FromIndex(13-o.day.GetEarthBranch().GetIndex()).GetName())
	return *t
}

// GetOwnSign 命宫
func (o EightChar) GetOwnSign() SixtyCycle {
	offset := o.month.GetEarthBranch().Next(-1).GetIndex() + o.hour.GetEarthBranch().Next(-1).GetIndex()
	if offset >= 14 {
		offset = 26 - offset
	} else {
		offset = 14 - offset
	}
	offset -= 1
	t, _ := SixtyCycle{}.FromName(HeavenStem{}.FromIndex((o.year.GetHeavenStem().GetIndex()+1)*2+offset).GetName() + EarthBranch{}.FromIndex(2+offset).GetName())
	return *t
}

// GetBodySign 身宫
func (o EightChar) GetBodySign() SixtyCycle {
	offset := o.month.GetEarthBranch().GetIndex() + o.hour.GetEarthBranch().GetIndex()
	offset %= 12
	offset -= 1
	t, _ := SixtyCycle{}.FromName(HeavenStem{}.FromIndex((o.year.GetHeavenStem().GetIndex()+1)*2+offset).GetName() + EarthBranch{}.FromIndex(2+offset).GetName())
	return *t
}

// GetDuty 建除十二值神
func (o EightChar) GetDuty() Duty {
	return Duty{}.FromIndex(o.day.GetEarthBranch().GetIndex() - o.month.GetEarthBranch().GetIndex())
}

func (o EightChar) GetName() string {
	return fmt.Sprintf("%v %v %v %v", o.year, o.month, o.day, o.hour)
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
	// 时辰地支转时刻，子时按零点算
	h := o.hour.GetEarthBranch().GetIndex() * 2
	baseYear := startYear - 1
	for y <= endYear {
		if y >= baseYear {
			// 立春为寅月的开始
			term := SolarTerm{}.FromIndex(y, 3)
			// 节令推移，年干支和月干支就都匹配上了
			if m > 0 {
				term = term.Next(m)
			}
			solarTime := term.GetJulianDay().GetSolarTime()
			if solarTime.GetYear() >= startYear {
				mi := 0
				s := 0
				// 日干支和节令干支的偏移值
				solarDay := solarTime.GetSolarDay()
				d := o.day.Next(-solarDay.GetLunarDay().GetSixtyCycle().GetIndex()).GetIndex()
				if d > 0 {
					// 从节令推移天数
					solarDay = solarDay.Next(d)
				} else if h == solarTime.GetHour() {
					// 如果正好是节令当天，且小时和节令的小时数相等的极端情况，把分钟和秒钟带上
					mi = solarTime.GetMinute()
					s = solarTime.GetSecond()
				}
				time, _ := SolarTime{}.FromYmdHms(solarDay.GetYear(), solarDay.GetMonth(), solarDay.GetDay(), h, mi, s)
				// 验证一下
				if time.GetLunarHour().GetEightChar().Equals(o) {
					l = append(l, *time)
				}
			}
		}
		y += 60
	}
	return l
}
