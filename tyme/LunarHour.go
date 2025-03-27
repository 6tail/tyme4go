package tyme

import (
	"fmt"
	"math"
)

// EightCharProvider 八字计算接口
var EightCharProvider IEightCharProvider = DefaultEightCharProvider{}

// LunarHour 时辰
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

// GetYearSixtyCycle 当时的年干支（立春换）
func (o LunarHour) GetYearSixtyCycle() SixtyCycle {
	solarTime := o.GetSolarTime()
	solarYear := o.day.GetSolarDay().GetYear()
	springSolarTime := SolarTerm{}.FromIndex(solarYear, 3).GetJulianDay().GetSolarTime()
	lunarYear := o.day.GetLunarMonth().GetLunarYear()
	year := lunarYear.GetYear()
	sixtyCycle := lunarYear.GetSixtyCycle()
	if year == solarYear {
		if solarTime.IsBefore(springSolarTime) {
			sixtyCycle = sixtyCycle.Next(-1)
		}
	} else if year < solarYear {
		if !solarTime.IsBefore(springSolarTime) {
			sixtyCycle = sixtyCycle.Next(1)
		}
	}
	return sixtyCycle
}

// GetMonthSixtyCycle 当时的月干支（节气换）
func (o LunarHour) GetMonthSixtyCycle() SixtyCycle {
	solarTime := o.GetSolarTime()
	year := solarTime.GetYear()
	term := solarTime.GetTerm()
	index := term.GetIndex() - 3
	if index < 0 && term.GetJulianDay().GetSolarTime().IsAfter(SolarTerm{}.FromIndex(year, 3).GetJulianDay().GetSolarTime()) {
		index += 24
	}
	m, _ := LunarMonth{}.FromYm(year, 1)
	return m.GetSixtyCycle().Next(int(math.Floor(float64(index) / 2)))
}

// GetDaySixtyCycle 当时的日干支（23:00开始算做第二天）
func (o LunarHour) GetDaySixtyCycle() SixtyCycle {
	d := o.day.GetSixtyCycle()
	if o.hour < 23 {
		return d
	}
	return d.Next(1)
}

// GetSixtyCycle 干支
func (o LunarHour) GetSixtyCycle() SixtyCycle {
	earthBranchIndex := o.GetIndexInDay() % 12
	heavenStemIndex := o.GetDaySixtyCycle().GetHeavenStem().GetIndex()%5*2 + earthBranchIndex
	t, _ := SixtyCycle{}.FromName(HeavenStem{}.FromIndex(heavenStemIndex).GetName() + EarthBranch{}.FromIndex(earthBranchIndex).GetName())
	return *t
}

// GetTwelveStar 黄道黑道十二神
func (o LunarHour) GetTwelveStar() TwelveStar {
	return TwelveStar{}.FromIndex(o.GetSixtyCycle().GetEarthBranch().GetIndex() + (8-o.GetDaySixtyCycle().GetEarthBranch().GetIndex()%6)*2)
}

// GetNineStar 九星
func (o LunarHour) GetNineStar() NineStar {
	solar := o.day.GetSolarDay()
	dongZhi := SolarTerm{}.FromIndex(solar.GetYear(), 0)
	xiaZhi := dongZhi.Next(12)
	asc := !solar.IsBefore(dongZhi.GetJulianDay().GetSolarDay()) && solar.IsBefore(xiaZhi.GetJulianDay().GetSolarDay())
	start := []int{8, 5, 2}[o.day.GetSixtyCycle().GetEarthBranch().GetIndex()%3]
	if asc {
		start = 8 - start
	}
	earthBranchIndex := o.GetIndexInDay() % 12
	if !asc {
		earthBranchIndex = -earthBranchIndex
	}
	return NineStar{}.FromIndex(start + earthBranchIndex)
}

// GetSolarTime 获取时辰对应的公历时间
func (o LunarHour) GetSolarTime() SolarTime {
	d := o.day.GetSolarDay()
	t, _ := SolarTime{}.FromYmdHms(d.GetYear(), d.GetMonth(), d.GetDay(), o.hour, o.minute, o.second)
	return *t
}

// GetEightChar 八字
func (o LunarHour) GetEightChar() EightChar {
	return EightCharProvider.GetEightChar(o)
}

// GetRecommends 宜
func (o LunarHour) GetRecommends() ([]Taboo, error) {
	return Taboo{}.GetHourRecommends(o.GetDaySixtyCycle(), o.GetSixtyCycle())
}

// GetAvoids 忌
func (o LunarHour) GetAvoids() ([]Taboo, error) {
	return Taboo{}.GetHourAvoids(o.GetDaySixtyCycle(), o.GetSixtyCycle())
}

// GetMinorRen 小六壬
func (o LunarHour) GetMinorRen() MinorRen {
	return o.GetLunarDay().GetMinorRen().Next(o.GetIndexInDay())
}
