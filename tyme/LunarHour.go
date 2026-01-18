package tyme

import (
	"fmt"
)

// EightCharProvider 八字计算接口
var EightCharProvider IEightCharProvider = DefaultEightCharProvider{}

// LunarHour 农历时辰
type LunarHour struct {
	SecondUnit
}

func (LunarHour) Validate(year int, month int, day int, hour int, minute int, second int) error {
	err := SecondUnit{}.Validate(year, month, day, hour, minute, second)
	if err != nil {
		return err
	}
	return LunarDay{}.Validate(year, month, day)
}

func (LunarHour) FromYmdHms(year int, month int, day int, hour int, minute int, second int) (*LunarHour, error) {
	err := LunarHour{}.Validate(year, month, day, hour, minute, second)
	if err != nil {
		return nil, err
	}
	return &LunarHour{
		SecondUnit{
			DayUnit{
				MonthUnit{
					YearUnit{
						year: year,
					},
					month,
				},
				day,
			},
			hour,
			minute,
			second,
		},
	}, nil
}

// GetLunarDay 农历日
func (o LunarHour) GetLunarDay() LunarDay {
	d, _ := LunarDay{}.FromYmd(o.year, o.month, o.day)
	return *d
}

func (o LunarHour) GetName() string {
	return fmt.Sprintf("%v时", EarthBranch{}.FromIndex(o.GetIndexInDay()).GetName())
}

func (o LunarHour) String() string {
	return fmt.Sprintf("%v%v时", o.GetLunarDay(), o.GetSixtyCycle().GetName())
}

// GetIndexInDay 位于当天的索引
func (o LunarHour) GetIndexInDay() int {
	return (o.hour + 1) / 2
}

func (o LunarHour) Next(n int) LunarHour {
	if n == 0 {
		t, _ := LunarHour{}.FromYmdHms(o.year, o.month, o.day, o.hour, o.minute, o.second)
		return *t
	}
	h := o.hour + n*2
	diff := 1
	if h < 0 {
		diff = -1
	}
	hour := h
	if hour < 0 {
		hour = -hour
	}
	days := hour / 24 * diff
	hour = (hour % 24) * diff
	if hour < 0 {
		hour += 24
		days--
	}
	d := o.GetLunarDay().Next(days)
	t, _ := LunarHour{}.FromYmdHms(d.GetYear(), d.GetMonth(), d.GetDay(), hour, o.minute, o.second)
	return *t
}

// IsBefore 是否在指定农历时辰之前
func (o LunarHour) IsBefore(target LunarHour) bool {
	aDay := o.GetLunarDay()
	bDay := target.GetLunarDay()
	if !aDay.Equals(bDay) {
		return aDay.IsBefore(bDay)
	}
	if o.hour != target.GetHour() {
		return o.hour < target.GetHour()
	}
	if o.minute != target.GetMinute() {
		return o.minute < target.GetMinute()
	}
	return o.second < target.GetSecond()
}

// IsAfter 是否在指定农历时辰之后
func (o LunarHour) IsAfter(target LunarHour) bool {
	aDay := o.GetLunarDay()
	bDay := target.GetLunarDay()
	if !aDay.Equals(bDay) {
		return aDay.IsAfter(bDay)
	}
	if o.hour != target.GetHour() {
		return o.hour > target.GetHour()
	}
	if o.minute != target.GetMinute() {
		return o.minute > target.GetMinute()
	}
	return o.second > target.GetSecond()
}

// Deprecated: Use GetSixtyCycleHour.GetYear instead.
func (o LunarHour) GetYearSixtyCycle() SixtyCycle {
	return o.GetSixtyCycleHour().GetYear()
}

// Deprecated: Use GetSixtyCycleHour.GetMonth instead.
func (o LunarHour) GetMonthSixtyCycle() SixtyCycle {
	return o.GetSixtyCycleHour().GetMonth()
}

// Deprecated: Use GetSixtyCycleHour.GetDay instead.
func (o LunarHour) GetDaySixtyCycle() SixtyCycle {
	return o.GetSixtyCycleHour().GetDay()
}

// GetSixtyCycle 干支
func (o LunarHour) GetSixtyCycle() SixtyCycle {
	earthBranchIndex := o.GetIndexInDay() % 12
	d := o.GetLunarDay().GetSixtyCycle()
	if o.hour >= 23 {
		d = d.Next(1)
	}
	t, _ := SixtyCycle{}.FromName(HeavenStem{}.FromIndex(d.GetHeavenStem().GetIndex()%5*2+earthBranchIndex).GetName() + EarthBranch{}.FromIndex(earthBranchIndex).GetName())
	return *t
}

// GetTwelveStar 黄道黑道十二神
func (o LunarHour) GetTwelveStar() TwelveStar {
	return TwelveStar{}.FromIndex(o.GetSixtyCycle().GetEarthBranch().GetIndex() + (8-o.GetSixtyCycleHour().GetDay().GetEarthBranch().GetIndex()%6)*2)
}

// GetNineStar 九星
func (o LunarHour) GetNineStar() NineStar {
	d := o.GetLunarDay()
	solar := d.GetSolarDay()
	dongZhi := SolarTerm{}.FromIndex(solar.GetYear(), 0)
	earthBranchIndex := o.GetIndexInDay() % 12
	index := []int{8, 5, 2}[d.GetSixtyCycle().GetEarthBranch().GetIndex()%3]
	if !solar.IsBefore(dongZhi.GetJulianDay().GetSolarDay()) && solar.IsBefore(dongZhi.Next(12).GetJulianDay().GetSolarDay()) {
		index = 8 + earthBranchIndex - index
	} else {
		index -= earthBranchIndex
	}
	return NineStar{}.FromIndex(index)
}

// GetSolarTime 公历时刻
func (o LunarHour) GetSolarTime() SolarTime {
	d := o.GetLunarDay().GetSolarDay()
	t, _ := SolarTime{}.FromYmdHms(d.GetYear(), d.GetMonth(), d.GetDay(), o.hour, o.minute, o.second)
	return *t
}

// GetSixtyCycleHour 干支时辰
func (o LunarHour) GetSixtyCycleHour() SixtyCycleHour {
	return o.GetSolarTime().GetSixtyCycleHour()
}

// GetEightChar 八字
func (o LunarHour) GetEightChar() EightChar {
	return EightCharProvider.GetEightChar(o)
}

// GetRecommends 宜
func (o LunarHour) GetRecommends() ([]Taboo, error) {
	return Taboo{}.GetHourRecommends(o.GetSixtyCycleHour().GetDay(), o.GetSixtyCycle())
}

// GetAvoids 忌
func (o LunarHour) GetAvoids() ([]Taboo, error) {
	return Taboo{}.GetHourAvoids(o.GetSixtyCycleHour().GetDay(), o.GetSixtyCycle())
}

// GetMinorRen 小六壬
func (o LunarHour) GetMinorRen() MinorRen {
	return o.GetLunarDay().GetMinorRen().Next(o.GetIndexInDay())
}
