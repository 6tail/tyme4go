package tyme

import (
	"fmt"
)

// SolarTime 公历时刻
type SolarTime struct {
	AbstractTyme
	// 公历日
	day SolarDay
	// 时
	hour int
	// 分
	minute int
	// 秒
	second int
}

func (SolarTime) FromYmdHms(year int, month int, day int, hour int, minute int, second int) (*SolarTime, error) {
	if hour < 0 || hour > 23 {
		return nil, fmt.Errorf(fmt.Sprintf("illegal hour: %d", hour))
	}
	if minute < 0 || minute > 59 {
		return nil, fmt.Errorf(fmt.Sprintf("illegal minute: %d", minute))
	}
	if second < 0 || second > 59 {
		return nil, fmt.Errorf(fmt.Sprintf("illegal second: %d", second))
	}
	d, err := SolarDay{}.FromYmd(year, month, day)
	if err != nil {
		return nil, err
	}
	return &SolarTime{
		day:    *d,
		hour:   hour,
		minute: minute,
		second: second,
	}, nil
}

// GetSolarDay 公历日
func (o SolarTime) GetSolarDay() SolarDay {
	return o.day
}

// GetYear 年
func (o SolarTime) GetYear() int {
	return o.day.GetYear()
}

// GetMonth 月
func (o SolarTime) GetMonth() int {
	return o.day.GetMonth()
}

// GetDay 日
func (o SolarTime) GetDay() int {
	return o.day.GetDay()
}

// GetHour 时
func (o SolarTime) GetHour() int {
	return o.hour
}

// GetMinute 分
func (o SolarTime) GetMinute() int {
	return o.minute
}

// GetSecond 秒
func (o SolarTime) GetSecond() int {
	return o.second
}

func (o SolarTime) GetName() string {
	return fmt.Sprintf("%02d:%02d:%02d", o.hour, o.minute, o.second)
}

func (o SolarTime) String() string {
	return fmt.Sprintf("%v %v", o.day, o.GetName())
}

func (o SolarTime) Next(n int) SolarTime {
	if n == 0 {
		t, _ := SolarTime{}.FromYmdHms(o.GetYear(), o.GetMonth(), o.GetDay(), o.hour, o.minute, o.second)
		return *t
	}
	ts := o.second + n
	tm := o.minute + ts/60
	ts %= 60
	if ts < 0 {
		ts += 60
		tm -= 1
	}
	th := o.hour + tm/60
	tm %= 60
	if tm < 0 {
		tm += 60
		th -= 1
	}
	td := th / 24
	th %= 24
	if th < 0 {
		th += 24
		td -= 1
	}

	d := o.day.Next(td)
	t, _ := SolarTime{}.FromYmdHms(d.GetYear(), d.GetMonth(), d.GetDay(), th, tm, ts)
	return *t
}

// IsBefore 是否在指定公历时刻之前
func (o SolarTime) IsBefore(target SolarTime) bool {
	if !o.day.Equals(target.GetSolarDay()) {
		return o.day.IsBefore(target.GetSolarDay())
	}
	if o.hour != target.GetHour() {
		return o.hour < target.GetHour()
	}
	if o.minute != target.GetMinute() {
		return o.minute < target.GetMinute()
	}
	return o.second < target.GetSecond()
}

// IsAfter 是否在指定公历时刻之后
func (o SolarTime) IsAfter(target SolarTime) bool {
	if !o.day.Equals(target.GetSolarDay()) {
		return o.day.IsAfter(target.GetSolarDay())
	}
	if o.hour != target.GetHour() {
		return o.hour > target.GetHour()
	}
	if o.minute != target.GetMinute() {
		return o.minute > target.GetMinute()
	}
	return o.second > target.GetSecond()
}

// Subtract 公历时刻相减，获得相差秒数
func (o SolarTime) Subtract(target SolarTime) int {
	days := o.day.Subtract(target.GetSolarDay())
	cs := o.hour*3600 + o.minute*60 + o.second
	ts := target.GetHour()*3600 + target.GetMinute()*60 + target.GetSecond()
	seconds := cs - ts
	if seconds < 0 {
		seconds += 86400
		days--
	}
	seconds += days * 86400
	return seconds
}

// GetJulianDay 儒略日
func (o SolarTime) GetJulianDay() JulianDay {
	return JulianDay{}.FromYmdHms(o.GetYear(), o.GetMonth(), o.GetDay(), o.hour, o.minute, o.second)
}

// GetLunarHour 农历时辰
func (o SolarTime) GetLunarHour() LunarHour {
	d := o.day.GetLunarDay()
	h, _ := LunarHour{}.FromYmdHms(d.GetYear(), d.GetMonth(), d.GetDay(), o.hour, o.minute, o.second)
	return *h
}

// GetSixtyCycleHour 干支时辰
func (o SolarTime) GetSixtyCycleHour() SixtyCycleHour {
	return SixtyCycleHour{}.FromSolarTime(o)
}

// GetTerm 节气
func (o SolarTime) GetTerm() SolarTerm {
	y := o.GetYear()
	i := o.GetMonth() * 2
	if i == 24 {
		y += 1
		i = 0
	}
	term := SolarTerm{}.FromIndex(y, i)
	for o.IsBefore(term.GetJulianDay().GetSolarTime()) {
		term = term.Next(-1)
	}
	return term
}
