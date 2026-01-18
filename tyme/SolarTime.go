package tyme

import (
	"fmt"
)

// SolarTime 公历时刻
type SolarTime struct {
	SecondUnit
}

func (SolarTime) Validate(year int, month int, day int, hour int, minute int, second int) error {
	err := SecondUnit{}.Validate(year, month, day, hour, minute, second)
	if err != nil {
		return err
	}
	return SolarDay{}.Validate(year, month, day)
}

func (SolarTime) FromYmdHms(year int, month int, day int, hour int, minute int, second int) (*SolarTime, error) {
	err := SolarTime{}.Validate(year, month, day, hour, minute, second)
	if err != nil {
		return nil, err
	}
	return &SolarTime{
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

// GetSolarDay 公历日
func (o SolarTime) GetSolarDay() SolarDay {
	d, _ := SolarDay{}.FromYmd(o.year, o.month, o.day)
	return *d
}

func (o SolarTime) GetName() string {
	return fmt.Sprintf("%02d:%02d:%02d", o.hour, o.minute, o.second)
}

func (o SolarTime) String() string {
	return fmt.Sprintf("%v %v", o.GetSolarDay(), o.GetName())
}

func (o SolarTime) Next(n int) SolarTime {
	if n == 0 {
		t, _ := SolarTime{}.FromYmdHms(o.year, o.month, o.day, o.hour, o.minute, o.second)
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

	d := o.GetSolarDay().Next(td)
	t, _ := SolarTime{}.FromYmdHms(d.GetYear(), d.GetMonth(), d.GetDay(), th, tm, ts)
	return *t
}

// IsBefore 是否在指定公历时刻之前
func (o SolarTime) IsBefore(target SolarTime) bool {
	aDay := o.GetSolarDay()
	bDay := target.GetSolarDay()
	if !aDay.Equals(bDay) {
		return aDay.IsBefore(bDay)
	}
	if o.hour != target.hour {
		return o.hour < target.hour
	}
	if o.minute != target.minute {
		return o.minute < target.minute
	}
	return o.second < target.second
}

// IsAfter 是否在指定公历时刻之后
func (o SolarTime) IsAfter(target SolarTime) bool {
	aDay := o.GetSolarDay()
	bDay := target.GetSolarDay()
	if !aDay.Equals(bDay) {
		return aDay.IsAfter(bDay)
	}
	if o.hour != target.hour {
		return o.hour > target.hour
	}
	if o.minute != target.minute {
		return o.minute > target.minute
	}
	return o.second > target.second
}

// Subtract 公历时刻相减，获得相差秒数
func (o SolarTime) Subtract(target SolarTime) int {
	days := o.GetSolarDay().Subtract(target.GetSolarDay())
	cs := o.hour*3600 + o.minute*60 + o.second
	ts := target.hour*3600 + target.minute*60 + target.second
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
	return JulianDay{}.FromYmdHms(o.year, o.month, o.day, o.hour, o.minute, o.second)
}

// GetLunarHour 农历时辰
func (o SolarTime) GetLunarHour() LunarHour {
	d := o.GetSolarDay().GetLunarDay()
	h, _ := LunarHour{}.FromYmdHms(d.GetYear(), d.GetMonth(), d.GetDay(), o.hour, o.minute, o.second)
	return *h
}

// GetSixtyCycleHour 干支时辰
func (o SolarTime) GetSixtyCycleHour() SixtyCycleHour {
	return SixtyCycleHour{}.FromSolarTime(o)
}

// GetTerm 节气
func (o SolarTime) GetTerm() SolarTerm {
	term := o.GetSolarDay().GetTerm()
	if o.IsBefore(term.GetJulianDay().GetSolarTime()) {
		term = term.Next(-1)
	}
	return term
}

func (o SolarTime) GetPhenology() Phenology {
	p := o.GetSolarDay().GetPhenology()
	if o.IsBefore(p.GetJulianDay().GetSolarTime()) {
		p = p.Next(-1)
	}
	return p
}

// GetPhase 月相
func (o SolarTime) GetPhase() Phase {
	month := o.GetLunarHour().GetLunarDay().GetLunarMonth().Next(1)
	p := Phase{}.FromIndex(month.GetYear(), month.GetMonthWithLeap(), 0)
	for p.GetSolarTime().IsAfter(o) {
		p = p.Next(-1)
	}
	return p
}
