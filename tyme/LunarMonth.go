package tyme

import (
	"fmt"
	"math"
)

var LunarMonthNames = []string{"正月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"}

// LunarMonth 农历月
type LunarMonth struct {
	MonthUnit
	// 是否闰月
	leap bool
}

func (LunarMonth) Validate(year int, month int) error {
	if month == 0 || month > 12 || month < -12 {
		return fmt.Errorf(fmt.Sprintf("illegal lunar month: %d", month))
	}
	if month < 0 {
		y, err := LunarYear{}.FromYear(year)
		if err != nil {
			return err
		}
		if -month != y.GetLeapMonth() {
			return fmt.Errorf(fmt.Sprintf("illegal leap month %d in lunar year %d", -month, year))
		}
	}
	return nil
}

func (LunarMonth) FromYm(year int, month int) (*LunarMonth, error) {
	err := LunarMonth{}.Validate(year, month)
	if err != nil {
		return nil, err
	}
	m := month
	leap := false
	if m < 0 {
		leap = true
		m = -m
	}
	return &LunarMonth{
		MonthUnit{
			YearUnit{
				year: year,
			},
			m,
		},
		leap,
	}, nil
}

// GetLunarYear 农历年
func (o LunarMonth) GetLunarYear() LunarYear {
	y, _ := LunarYear{}.FromYear(o.year)
	return *y
}

// GetMonthWithLeap 月，当月为闰月时，返回负数
func (o LunarMonth) GetMonthWithLeap() int {
	if o.leap {
		return -o.month
	}
	return o.month
}

func (o LunarMonth) getNewMoon() float64 {
	// 冬至
	dongZhiJd := SolarTerm{}.FromIndex(o.year, 0).GetCursoryJulianDay()

	// 冬至前的初一，今年首朔的日月黄经差
	w := CalcShuo(dongZhiJd)
	if w > dongZhiJd {
		w -= 29.53
	}

	// 正常情况正月初一为第3个朔日，但有些特殊的
	offset := 2
	if o.year > 8 && o.year < 24 {
		offset = 1
	} else if o.year != 239 && o.year != 240 {
		y, _ := LunarYear{}.FromYear(o.year - 1)
		if y.GetLeapMonth() > 10 {
			offset = 3
		}
	}

	// 本月初一
	return w + 29.5306*float64(offset+o.GetIndexInYear())
}

// GetDayCount 天数(大月30天，小月29天)
func (o LunarMonth) GetDayCount() int {
	w := o.getNewMoon()
	return int(CalcShuo(w+29.5306) - CalcShuo(w))
}

// GetIndexInYear 位于当年的索引(0-12)
func (o LunarMonth) GetIndexInYear() int {
	index := o.month - 1
	if o.IsLeap() {
		index += 1
	} else {
		leapMonth := o.GetLunarYear().GetLeapMonth()
		if leapMonth > 0 && o.month > leapMonth {
			index += 1
		}
	}
	return index
}

// GetSeason 农历季节
func (o LunarMonth) GetSeason() LunarSeason {
	return LunarSeason{}.FromIndex(o.month - 1)
}

// GetFirstJulianDay 初一的儒略日
func (o LunarMonth) GetFirstJulianDay() JulianDay {
	return JulianDay{}.FromJulianDay(J2000 + CalcShuo(o.getNewMoon()))
}

// IsLeap 是否闰月
func (o LunarMonth) IsLeap() bool {
	return o.leap
}

// GetWeekCount 周数
func (o LunarMonth) GetWeekCount(start int) int {
	return int(math.Ceil(float64(o.IndexOf(o.GetFirstJulianDay().GetWeek().GetIndex()-start, 7)+o.GetDayCount()) / 7))
}

func (o LunarMonth) GetName() string {
	s := LunarMonthNames[o.month-1]
	if o.leap {
		return "闰" + s
	}
	return s
}

func (o LunarMonth) String() string {
	return fmt.Sprintf("%v%v", o.GetLunarYear(), o.GetName())
}

func (o LunarMonth) Next(n int) LunarMonth {
	if n == 0 {
		month, _ := LunarMonth{}.FromYm(o.GetYear(), o.GetMonthWithLeap())
		return *month
	}
	m := o.GetIndexInYear() + 1 + n
	y := o.GetLunarYear()
	if n > 0 {
		monthCount := y.GetMonthCount()
		for m > monthCount {
			m -= monthCount
			y = y.Next(1)
			monthCount = y.GetMonthCount()
		}
	} else {
		for m <= 0 {
			y = y.Next(-1)
			m += y.GetMonthCount()
		}
	}
	leap := false
	leapMonth := y.GetLeapMonth()
	if leapMonth > 0 {
		if m == leapMonth+1 {
			leap = true
		}
		if m > leapMonth {
			m--
		}
	}
	if leap {
		m = -m
	}
	month, _ := LunarMonth{}.FromYm(y.GetYear(), m)
	return *month
}

// GetDays 本月的农历日列表
func (o LunarMonth) GetDays() []LunarDay {
	var l []LunarDay
	size := o.GetDayCount()
	m := o.GetMonthWithLeap()
	for i := 0; i < size; i++ {
		d, _ := LunarDay{}.FromYmd(o.year, m, i+1)
		l = append(l, *d)
	}
	return l
}

// GetFirstDay 初一
func (o LunarMonth) GetFirstDay() LunarDay {
	d, _ := LunarDay{}.FromYmd(o.year, o.GetMonthWithLeap(), 1)
	return *d
}

// GetWeeks 本月的农历周列表
func (o LunarMonth) GetWeeks(start int) []LunarWeek {
	var l []LunarWeek
	size := o.GetWeekCount(start)
	m := o.GetMonthWithLeap()
	for i := 0; i < size; i++ {
		w, _ := LunarWeek{}.FromYm(o.year, m, i, start)
		l = append(l, *w)
	}
	return l
}

// GetSixtyCycle 干支
func (o LunarMonth) GetSixtyCycle() SixtyCycle {
	t, _ := SixtyCycle{}.FromName(HeavenStem{}.FromIndex(o.GetLunarYear().GetSixtyCycle().GetHeavenStem().GetIndex()*2+o.month+1).GetName() + EarthBranch{}.FromIndex(o.month+1).GetName())
	return *t
}

// GetNineStar 九星
func (o LunarMonth) GetNineStar() NineStar {
	index := o.GetSixtyCycle().GetEarthBranch().GetIndex()
	if index < 2 {
		index += 3
	}
	return NineStar{}.FromIndex(27 - o.GetLunarYear().GetSixtyCycle().GetEarthBranch().GetIndex()%3*3 - index)
}

// GetJupiterDirection 太岁方位
func (o LunarMonth) GetJupiterDirection() Direction {
	sixtyCycle := o.GetSixtyCycle()
	n := []int{7, -1, 1, 3}[sixtyCycle.GetEarthBranch().Next(-2).GetIndex()%4]
	if n != -1 {
		return Direction{}.FromIndex(n)
	}
	return sixtyCycle.GetHeavenStem().GetDirection()
}

// GetFetus 逐月胎神
func (o LunarMonth) GetFetus() FetusMonth {
	return *FetusMonth{}.FromLunarMonth(o)
}

// GetMinorRen 小六壬
func (o LunarMonth) GetMinorRen() MinorRen {
	return MinorRen{}.FromIndex((o.month - 1) % 6)
}
