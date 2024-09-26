package tyme

import (
	"fmt"
	"math"
)

// J2000 2000年儒略日数(2000-1-1 12:00:00 UTC)
const J2000 = 2451545

// JulianDay 儒略日
type JulianDay struct {
	AbstractTyme
	// 儒略日
	day float64
}

func (JulianDay) FromJulianDay(day float64) JulianDay {
	return JulianDay{
		day: day,
	}
}

func (JulianDay) FromYmdHms(year int, month int, day int, hour int, minute int, second int) JulianDay {
	d := float64(day) + ((float64(second)/60+float64(minute))/60+float64(hour))/24
	n := 0
	g := year*372+month*31+int(d) >= 588829
	if month <= 2 {
		month += 12
		year--
	}
	if g {
		n = int(float64(year) / 100)
		n = 2 - n + int(float64(n)/4)
	}
	return JulianDay{}.FromJulianDay(float64(int(365.25*float64(year+4716))) + float64(int(30.6001*float64(month+1))) + d + float64(n) - 1524.5)
}

// GetDay 儒略日
func (o JulianDay) GetDay() float64 {
	return o.day
}

func (o JulianDay) GetName() string {
	return fmt.Sprintf("%v", o.day)
}

func (o JulianDay) String() string {
	return o.GetName()
}

func (o JulianDay) Next(n int) JulianDay {
	return JulianDay{}.FromJulianDay(o.day + float64(n))
}

// GetSolarDay 公历日
func (o JulianDay) GetSolarDay() SolarDay {
	return o.GetSolarTime().GetSolarDay()
}

// GetSolarTime 公历时刻
func (o JulianDay) GetSolarTime() SolarTime {
	d := int(o.day + 0.5)
	f := o.day + 0.5 - float64(d)

	if d >= 2299161 {
		c := int((float64(d) - 1867216.25) / 36524.25)
		d += 1 + c - (int)(float64(c)/4)
	}
	d += 1524
	year := int((float64(d) - 122.1) / 365.25)
	d -= int(365.25 * float64(year))
	month := int(float64(d) / 30.601)
	d -= int(30.601 * float64(month))
	day := d
	if month > 13 {
		month -= 13
		year -= 4715
	} else {
		month -= 1
		year -= 4716
	}
	f *= 24
	hour := int(f)

	f -= float64(hour)
	f *= 60
	minute := int(f)

	f -= float64(minute)
	f *= 60
	second := int(math.Round(f))
	if second > 59 {
		second -= 60
		minute += 1
	}
	if minute > 59 {
		minute -= 60
		hour += 1
	}
	if hour > 23 {
		hour -= 24
		day += 1
	}
	t, _ := SolarTime{}.FromYmdHms(year, month, day, hour, minute, second)
	return *t
}

// GetWeek 星期
func (o JulianDay) GetWeek() Week {
	return Week{}.FromIndex(int(o.day+0.5) + 7000001)
}

// Subtract 儒略日相减
func (o JulianDay) Subtract(target JulianDay) float64 {
	return o.day - target.GetDay()
}
