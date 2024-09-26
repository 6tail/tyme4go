package tyme

import (
	"fmt"
	"math"
)

var LunarMonthNames = []string{"正月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"}
var lunarMonthCache map[string][]interface{}

// LunarMonth 农历月
type LunarMonth struct {
	AbstractTyme
	// 农历年
	year LunarYear
	// 月
	month int
	// 是否闰月
	leap bool
	// 天数
	dayCount int
	// 位于当年的索引，0-12
	indexInYear int
	// 初一的儒略日
	firstJulianDay JulianDay
}

func (LunarMonth) fromCache(cache []interface{}) LunarMonth {
	m := cache[1].(int)
	year, _ := LunarYear{}.FromYear(cache[0].(int))
	month := m
	if month < 0 {
		month = -month
	}
	leap := m < 0
	dayCount := cache[2].(int)
	indexInYear := cache[3].(int)
	firstJulianDay := JulianDay{}.FromJulianDay(cache[4].(float64))
	return LunarMonth{
		year:           *year,
		month:          month,
		leap:           leap,
		dayCount:       dayCount,
		indexInYear:    indexInYear,
		firstJulianDay: firstJulianDay,
	}
}

func (LunarMonth) New(year int, month int) (*LunarMonth, error) {
	currentYear, err := LunarYear{}.FromYear(year)
	if err != nil {
		return nil, err
	}
	currentLeapMonth := currentYear.GetLeapMonth()
	if month == 0 || month > 12 || month < -12 {
		return nil, fmt.Errorf(fmt.Sprintf("illegal lunar month: %d", month))
	}
	leap := month < 0
	m := month
	if m < 0 {
		m = -m
	}
	if leap && m != currentLeapMonth {
		return nil, fmt.Errorf(fmt.Sprintf("illegal leap month %d in lunar year %d", m, year))
	}

	// 冬至
	dongZhi := SolarTerm{}.FromIndex(year, 0)
	dongZhiJd := dongZhi.GetCursoryJulianDay()

	// 冬至前的初一，今年首朔的日月黄经差
	w := CalcShuo(dongZhiJd)
	if w > dongZhiJd {
		w -= 29.53
	}

	// 正常情况正月初一为第3个朔日，但有些特殊的
	offset := 2
	if year > 8 && year < 24 {
		offset = 1
	} else if year != 239 && year != 240 {
		y, _ := LunarYear{}.FromYear(year - 1)
		if y.GetLeapMonth() > 10 {
			offset = 3
		}
	}

	// 位于当年的索引
	index := m - 1
	if leap || (currentLeapMonth > 0 && m > currentLeapMonth) {
		index += 1
	}
	indexInYear := index

	// 本月初一
	w += 29.5306 * float64(offset+index)
	firstDay := CalcShuo(w)
	firstJulianDay := JulianDay{}.FromJulianDay(J2000 + firstDay)
	// 本月天数 = 下月初一 - 本月初一
	dayCount := int(CalcShuo(w+29.5306) - firstDay)
	return &LunarMonth{
		year:           *currentYear,
		month:          m,
		leap:           leap,
		dayCount:       dayCount,
		indexInYear:    indexInYear,
		firstJulianDay: firstJulianDay,
	}, nil
}

func (LunarMonth) FromYm(year int, month int) (*LunarMonth, error) {
	var m *LunarMonth = nil
	var err error = nil
	key := fmt.Sprintf("%d%d", year, month)
	if c, ok := lunarMonthCache[key]; ok {
		t := LunarMonth{}.fromCache(c)
		m = &t
	} else {
		m, err = LunarMonth{}.New(year, month)
		if err == nil {
			var l []interface{}
			l = append(l, m.GetYear())
			l = append(l, m.GetMonthWithLeap())
			l = append(l, m.GetDayCount())
			l = append(l, m.GetIndexInYear())
			l = append(l, m.GetFirstJulianDay().GetDay())
			lunarMonthCache[key] = l
		}
	}
	return m, err
}

// GetLunarYear 农历年
func (o LunarMonth) GetLunarYear() LunarYear {
	return o.year
}

// GetYear 年
func (o LunarMonth) GetYear() int {
	return o.year.GetYear()
}

// GetMonth 月
func (o LunarMonth) GetMonth() int {
	return o.month
}

// GetMonthWithLeap 月，当月为闰月时，返回负数
func (o LunarMonth) GetMonthWithLeap() int {
	if o.leap {
		return -o.month
	}
	return o.month
}

// GetDayCount 天数(大月30天，小月29天)
func (o LunarMonth) GetDayCount() int {
	return o.dayCount
}

// GetIndexInYear 位于当年的索引(0-12)
func (o LunarMonth) GetIndexInYear() int {
	return o.indexInYear
}

// GetSeason 农历季节
func (o LunarMonth) GetSeason() LunarSeason {
	return LunarSeason{}.FromIndex(o.month - 1)
}

// GetFirstJulianDay 初一的儒略日
func (o LunarMonth) GetFirstJulianDay() JulianDay {
	return o.firstJulianDay
}

// IsLeap 是否闰月
func (o LunarMonth) IsLeap() bool {
	return o.leap
}

// GetWeekCount 周数
func (o LunarMonth) GetWeekCount(start int) int {
	return int(math.Ceil(float64(o.IndexOf(o.firstJulianDay.GetWeek().GetIndex()-start, 7)+o.GetDayCount()) / 7))
}

func (o LunarMonth) GetName() string {
	s := LunarMonthNames[o.month-1]
	if o.leap {
		return "闰" + s
	}
	return s
}

func (o LunarMonth) String() string {
	return fmt.Sprintf("%v%v", o.year, o.GetName())
}

func (o LunarMonth) Next(n int) LunarMonth {
	if n == 0 {
		month, _ := LunarMonth{}.FromYm(o.GetYear(), o.GetMonthWithLeap())
		return *month
	}
	m := o.indexInYear + 1 + n
	y := o.year
	leapMonth := y.GetLeapMonth()
	if n > 0 {
		monthCount := 12
		if leapMonth > 0 {
			monthCount = 13
		}
		for m > monthCount {
			m -= monthCount
			y = y.Next(1)
			leapMonth = y.GetLeapMonth()
			monthCount = 12
			if leapMonth > 0 {
				monthCount = 13
			}
		}
	} else {
		for m <= 0 {
			y = y.Next(-1)
			leapMonth = y.GetLeapMonth()
			monthCount := 12
			if leapMonth > 0 {
				monthCount = 13
			}
			m += monthCount
		}
	}
	leap := false
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

// GetDays 获取本月的农历日列表
func (o LunarMonth) GetDays() []LunarDay {
	var l []LunarDay
	size := o.GetDayCount()
	y := o.GetYear()
	m := o.GetMonthWithLeap()
	for i := 0; i < size; i++ {
		d, _ := LunarDay{}.FromYmd(y, m, i+1)
		l = append(l, *d)
	}
	return l
}

// GetWeeks 获取本月的农历周列表
func (o LunarMonth) GetWeeks(start int) []LunarWeek {
	var l []LunarWeek
	size := o.GetWeekCount(start)
	y := o.GetYear()
	m := o.GetMonthWithLeap()
	for i := 0; i < size; i++ {
		w, _ := LunarWeek{}.FromYm(y, m, i, start)
		l = append(l, *w)
	}
	return l
}

// GetSixtyCycle 干支
func (o LunarMonth) GetSixtyCycle() SixtyCycle {
	t, _ := SixtyCycle{}.FromName(HeavenStem{}.FromIndex((o.year.GetSixtyCycle().GetHeavenStem().GetIndex()+1)*2+o.indexInYear).GetName() + EarthBranch{}.FromIndex(o.indexInYear+2).GetName())
	return *t
}

// GetNineStar 九星
func (o LunarMonth) GetNineStar() NineStar {
	return NineStar{}.FromIndex(27 - o.year.GetSixtyCycle().GetEarthBranch().GetIndex()%3*3 - o.GetSixtyCycle().GetEarthBranch().GetIndex())
}

// GetJupiterDirection 太岁方位
func (o LunarMonth) GetJupiterDirection() Direction {
	sixtyCycle := o.GetSixtyCycle()
	n := []int{7, -1, 1, 3}[sixtyCycle.GetEarthBranch().Next(-2).GetIndex()%4]
	if n == -1 {
		return sixtyCycle.GetHeavenStem().GetDirection()
	}
	return Direction{}.FromIndex(n)
}

// GetFetus 逐月胎神
func (o LunarMonth) GetFetus() FetusMonth {
	return *FetusMonth{}.FromLunarMonth(o)
}
