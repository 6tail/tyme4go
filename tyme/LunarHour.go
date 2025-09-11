package tyme

import (
	"fmt"
)

// EightCharProvider 八字计算接口
var EightCharProvider IEightCharProvider = DefaultEightCharProvider{}

// LunarHour 农历时辰
type LunarHour struct {
	AbstractTyme
	// 农历日
	day LunarDay
	// 时
	hour int
	// 分
	minute int
	// 秒
	second int
	// 公历时刻（第一次使用时才会初始化）
	solarTime *SolarTime
	// 干支时辰（第一次使用时才会初始化）
	sixtyCycleHour *SixtyCycleHour
}

func (LunarHour) FromYmdHms(year int, month int, day int, hour int, minute int, second int) (*LunarHour, error) {
	if hour < 0 || hour > 23 {
		return nil, fmt.Errorf(fmt.Sprintf("illegal hour: %d", hour))
	}
	if minute < 0 || minute > 59 {
		return nil, fmt.Errorf(fmt.Sprintf("illegal minute: %d", minute))
	}
	if second < 0 || second > 59 {
		return nil, fmt.Errorf(fmt.Sprintf("illegal second: %d", second))
	}
	d, err := LunarDay{}.FromYmd(year, month, day)
	if err != nil {
		return nil, err
	}
	return &LunarHour{
		day:    *d,
		hour:   hour,
		minute: minute,
		second: second,
	}, nil
}

// GetLunarDay 农历日
func (o LunarHour) GetLunarDay() LunarDay {
	return o.day
}

// GetYear 年
func (o LunarHour) GetYear() int {
	return o.day.GetYear()
}

// GetMonth 月
func (o LunarHour) GetMonth() int {
	return o.day.GetMonth()
}

// GetDay 日
func (o LunarHour) GetDay() int {
	return o.day.GetDay()
}

// GetHour 时
func (o LunarHour) GetHour() int {
	return o.hour
}

// GetMinute 分
func (o LunarHour) GetMinute() int {
	return o.minute
}

// GetSecond 秒
func (o LunarHour) GetSecond() int {
	return o.second
}

func (o LunarHour) GetName() string {
	return fmt.Sprintf("%v时", EarthBranch{}.FromIndex(o.GetIndexInDay()).GetName())
}

func (o LunarHour) String() string {
	return fmt.Sprintf("%v%v时", o.day, o.GetSixtyCycle().GetName())
}

// GetIndexInDay 位于当天的索引
func (o LunarHour) GetIndexInDay() int {
	return (o.hour + 1) / 2
}

func (o LunarHour) Next(n int) LunarHour {
	if n == 0 {
		t, _ := LunarHour{}.FromYmdHms(o.GetYear(), o.GetMonth(), o.GetDay(), o.hour, o.minute, o.second)
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
	d := o.day.Next(days)
	t, _ := LunarHour{}.FromYmdHms(d.GetYear(), d.GetMonth(), d.GetDay(), hour, o.minute, o.second)
	return *t
}

// IsBefore 是否在指定农历时辰之前
func (o LunarHour) IsBefore(target LunarHour) bool {
	if !o.day.Equals(target.GetLunarDay()) {
		return o.day.IsBefore(target.GetLunarDay())
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
	if !o.day.Equals(target.GetLunarDay()) {
		return o.day.IsAfter(target.GetLunarDay())
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
	d := o.day.GetSixtyCycle()
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
	solar := o.day.GetSolarDay()
	dongZhi := SolarTerm{}.FromIndex(solar.GetYear(), 0)
	earthBranchIndex := o.GetIndexInDay() % 12
	index := []int{8, 5, 2}[o.day.GetSixtyCycle().GetEarthBranch().GetIndex()%3]
	if !solar.IsBefore(dongZhi.GetJulianDay().GetSolarDay()) && solar.IsBefore(dongZhi.Next(12).GetJulianDay().GetSolarDay()) {
		index = 8 + earthBranchIndex - index
	} else {
		index -= earthBranchIndex
	}
	return NineStar{}.FromIndex(index)
}

// GetSolarTime 公历时刻
func (o LunarHour) GetSolarTime() SolarTime {
	if o.solarTime == nil {
		d := o.day.GetSolarDay()
		t, _ := SolarTime{}.FromYmdHms(d.GetYear(), d.GetMonth(), d.GetDay(), o.hour, o.minute, o.second)
		o.solarTime = t
	}
	return *o.solarTime
}

// GetSixtyCycleHour 干支时辰
func (o LunarHour) GetSixtyCycleHour() SixtyCycleHour {
	if o.sixtyCycleHour == nil {
		h := o.GetSolarTime().GetSixtyCycleHour()
		o.sixtyCycleHour = &h
	}
	return *o.sixtyCycleHour
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
