package tyme

import (
	"fmt"
	"math"
)

// SixtyCycleHour 干支时辰（立春换年，节令换月，23点换日）
type SixtyCycleHour struct {
	AbstractTyme
	// 公历时刻
	solarTime SolarTime
	// 干支日
	day SixtyCycleDay
	// 时柱
	hour SixtyCycle
}

func (SixtyCycleHour) FromSolarTime(solarTime SolarTime) SixtyCycleHour {
	solarYear := solarTime.GetYear()
	springSolarTime := SolarTerm{}.FromIndex(solarYear, 3).GetJulianDay().GetSolarTime()
	lunarHour := solarTime.GetLunarHour()
	lunarDay := lunarHour.GetLunarDay()
	lunarYear := lunarDay.GetLunarMonth().GetLunarYear()
	if lunarYear.GetYear() == solarYear {
		if solarTime.IsBefore(springSolarTime) {
			lunarYear = lunarYear.Next(-1)
		}
	} else if lunarYear.GetYear() < solarYear {
		if !solarTime.IsBefore(springSolarTime) {
			lunarYear = lunarYear.Next(1)
		}
	}

	term := solarTime.GetTerm()
	index := term.GetIndex() - 3
	if index < 0 && term.GetJulianDay().GetSolarTime().IsAfter(SolarTerm{}.FromIndex(solarYear, 3).GetJulianDay().GetSolarTime()) {
		index += 24
	}
	d := lunarDay.GetSixtyCycle()
	if solarTime.GetHour() == 23 {
		d = d.Next(1)
	}
	y, _ := SixtyCycleYear{}.FromYear(lunarYear.GetYear())
	m, _ := LunarMonth{}.FromYm(solarYear, 1)
	return SixtyCycleHour{
		solarTime: solarTime,
		day:       SixtyCycleDay{}.New(solarTime.GetSolarDay(), SixtyCycleMonth{}.New(*y, m.GetSixtyCycle().Next(int(math.Floor(float64(index)/2)))), d),
		hour:      lunarHour.GetSixtyCycle(),
	}
}

// GetYear 年柱
func (o SixtyCycleHour) GetYear() SixtyCycle {
	return o.day.GetYear()
}

// GetMonth 月柱
func (o SixtyCycleHour) GetMonth() SixtyCycle {
	return o.day.GetMonth()
}

// GetDay 日柱
func (o SixtyCycleHour) GetDay() SixtyCycle {
	return o.day.GetSixtyCycle()
}

// GetSixtyCycle 干支
func (o SixtyCycleHour) GetSixtyCycle() SixtyCycle {
	return o.hour
}

// GetSixtyCycleDay 干支日
func (o SixtyCycleHour) GetSixtyCycleDay() SixtyCycleDay {
	return o.day
}

// GetSolarTime 公历时刻
func (o SixtyCycleHour) GetSolarTime() SolarTime {
	return o.solarTime
}

func (o SixtyCycleHour) GetName() string {
	return fmt.Sprintf("%v时", o.hour)
}

func (o SixtyCycleHour) String() string {
	return fmt.Sprintf("%v%v", o.day, o.GetName())
}

// GetIndexInDay 位于当天的索引
func (o SixtyCycleHour) GetIndexInDay() int {
	h := o.solarTime.GetHour()
	if h == 23 {
		return 0
	}
	return (h + 1) / 2
}

func (o SixtyCycleHour) Next(n int) SixtyCycleHour {
	return SixtyCycleHour{}.FromSolarTime(o.solarTime.Next(n))
}

// GetTwelveStar 黄道黑道十二神
func (o SixtyCycleHour) GetTwelveStar() TwelveStar {
	return TwelveStar{}.FromIndex(o.hour.GetEarthBranch().GetIndex() + (8-o.GetDay().GetEarthBranch().GetIndex()%6)*2)
}

// GetNineStar 九星
func (o SixtyCycleHour) GetNineStar() NineStar {
	solar := o.solarTime.GetSolarDay()
	dongZhi := SolarTerm{}.FromIndex(solar.GetYear(), 0)
	xiaZhi := dongZhi.Next(12)
	asc := !solar.IsBefore(dongZhi.GetJulianDay().GetSolarDay()) && solar.IsBefore(xiaZhi.GetJulianDay().GetSolarDay())
	start := []int{8, 5, 2}[o.GetDay().GetEarthBranch().GetIndex()%3]
	if asc {
		start = 8 - start
	}
	earthBranchIndex := o.GetIndexInDay() % 12
	if !asc {
		earthBranchIndex = -earthBranchIndex
	}
	return NineStar{}.FromIndex(start + earthBranchIndex)
}

// GetEightChar 八字
func (o SixtyCycleHour) GetEightChar() EightChar {
	return EightChar{
		year:  o.GetYear(),
		month: o.GetMonth(),
		day:   o.GetDay(),
		hour:  o.hour,
	}
}

// GetRecommends 宜
func (o SixtyCycleHour) GetRecommends() ([]Taboo, error) {
	return Taboo{}.GetHourRecommends(o.GetDay(), o.hour)
}

// GetAvoids 忌
func (o SixtyCycleHour) GetAvoids() ([]Taboo, error) {
	return Taboo{}.GetHourAvoids(o.GetDay(), o.hour)
}
