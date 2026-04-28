package tyme

import (
	"fmt"
	"strconv"
)

var LunarDayNames = []string{"初一", "初二", "初三", "初四", "初五", "初六", "初七", "初八", "初九", "初十", "十一", "十二", "十三", "十四", "十五", "十六", "十七", "十八", "十九", "二十", "廿一", "廿二", "廿三", "廿四", "廿五", "廿六", "廿七", "廿八", "廿九", "三十"}

// LunarDay 农历日
type LunarDay struct {
	DayUnit
}

func (LunarDay) Validate(year int, month int, day int) error {
	if day < 1 {
		return fmt.Errorf("illegal lunar day: " + strconv.Itoa(day))
	}
	m, err := LunarMonth{}.FromYm(year, month)
	if err != nil {
		return err
	}
	if day > m.GetDayCount() {
		return fmt.Errorf("illegal day %d in %v", day, m)
	}
	return nil
}

func (LunarDay) FromYmd(year int, month int, day int) (*LunarDay, error) {
	err := LunarDay{}.Validate(year, month, day)
	if err != nil {
		return nil, err
	}
	return &LunarDay{
		DayUnit{
			MonthUnit{
				YearUnit{
					year: year,
				},
				month,
			},
			day,
		},
	}, nil
}

// GetLunarMonth 农历月
func (o LunarDay) GetLunarMonth() LunarMonth {
	m, _ := LunarMonth{}.FromYm(o.year, o.month)
	return *m
}

// GetWeek 星期
func (o LunarDay) GetWeek() Week {
	return o.GetSolarDay().GetWeek()
}

func (o LunarDay) GetName() string {
	return LunarDayNames[o.day-1]
}

func (o LunarDay) String() string {
	return fmt.Sprintf("%v%v", o.GetLunarMonth(), o.GetName())
}

func (o LunarDay) Next(n int) LunarDay {
	return o.GetSolarDay().Next(n).GetLunarDay()
}

// IsBefore 是否在指定农历日之前
func (o LunarDay) IsBefore(target LunarDay) bool {
	aYear := o.year
	bYear := target.GetYear()
	if aYear != bYear {
		return aYear < bYear
	}
	aMonth := o.month
	bMonth := target.GetMonth()
	if aMonth != bMonth {
		b := bMonth
		if b < 0 {
			b = -b
		}
		a := aMonth
		if a < 0 {
			a = -a
		}
		return aMonth == b || a < b
	}
	return o.day < target.GetDay()
}

// IsAfter 是否在指定农历日之后
func (o LunarDay) IsAfter(target LunarDay) bool {
	aYear := o.year
	bYear := target.GetYear()
	if aYear != bYear {
		return aYear > bYear
	}
	aMonth := o.month
	bMonth := target.GetMonth()
	if aMonth != bMonth {
		a := aMonth
		if a < 0 {
			a = -a
		}
		b := bMonth
		if b < 0 {
			b = -b
		}
		return a == bMonth || a > b
	}
	return o.day > target.GetDay()
}

// Deprecated: Use SixtyCycleDay.GetYear instead.
func (o LunarDay) GetYearSixtyCycle() SixtyCycle {
	return o.GetSixtyCycleDay().GetYear()
}

// Deprecated: Use SixtyCycleDay.GetMonth instead.
func (o LunarDay) GetMonthSixtyCycle() SixtyCycle {
	return o.GetSixtyCycleDay().GetMonth()
}

// GetSixtyCycle 干支
func (o LunarDay) GetSixtyCycle() SixtyCycle {
	return SixtyCycle{}.FromIndex(int(o.GetLunarMonth().GetFirstJulianDay().Next(o.day - 12).GetDay()))
}

// GetDuty 建除十二值神
func (o LunarDay) GetDuty() Duty {
	return o.GetSixtyCycleDay().GetDuty()
}

// GetTwelveStar 黄道黑道十二神
func (o LunarDay) GetTwelveStar() TwelveStar {
	return o.GetSixtyCycleDay().GetTwelveStar()
}

// GetNineStar 九星
func (o LunarDay) GetNineStar() NineStar {
	return o.GetSolarDay().GetNineStar()
}

// GetJupiterDirection 太岁方位
func (o LunarDay) GetJupiterDirection() Direction {
	index := o.GetSixtyCycle().GetIndex()
	if index%12 < 6 {
		return Element{}.FromIndex(index / 12).GetDirection()
	}
	return o.GetLunarMonth().GetLunarYear().GetJupiterDirection()
}

// GetFetusDay 逐日胎神
func (o LunarDay) GetFetusDay() FetusDay {
	return FetusDay{}.FromLunarDay(o)
}

// GetPhaseDay 月相第几天
func (o LunarDay) GetPhaseDay() PhaseDay {
	today := o.GetSolarDay()
	m := o.GetLunarMonth().Next(1)
	p := Phase{}.FromIndex(m.GetYear(), m.GetMonthWithLeap(), 0)
	d := p.GetSolarDay()
	for d.IsAfter(today) {
		p = p.Next(-1)
		d = p.GetSolarDay()
	}
	return PhaseDay{}.New(p, today.Subtract(d))
}

// GetPhase 月相
func (o LunarDay) GetPhase() Phase {
	return o.GetPhaseDay().GetPhase()
}

// GetSixStar 六曜
func (o LunarDay) GetSixStar() SixStar {
	m := o.month
	if m < 0 {
		m = -m
	}
	return SixStar{}.FromIndex((m + o.day - 2) % 6)
}

// GetSolarDay 公历日
func (o LunarDay) GetSolarDay() SolarDay {
	return o.GetLunarMonth().GetFirstJulianDay().Next(o.day - 1).GetSolarDay()
}

// GetSixtyCycleDay 干支日
func (o LunarDay) GetSixtyCycleDay() SixtyCycleDay {
	return o.GetSolarDay().GetSixtyCycleDay()
}

// GetTwentyEightStar 二十八宿
func (o LunarDay) GetTwentyEightStar() TwentyEightStar {
	return TwentyEightStar{}.FromIndex([]int{10, 18, 26, 6, 14, 22, 2}[o.GetSolarDay().GetWeek().GetIndex()]).Next(-7 * o.GetSixtyCycle().GetEarthBranch().GetIndex())
}

// GetFestival 农历传统节日，如果当天不是农历传统节日，返回nil
func (o LunarDay) GetFestival() *LunarFestival {
	return LunarFestival{}.FromYmd(o.year, o.month, o.day)
}

func (o LunarDay) Equals(target LunarDay) bool {
	return o.String() == target.String()
}

// GetHours 当天的农历时辰列表
func (o LunarDay) GetHours() []LunarHour {
	var l []LunarHour
	t, _ := LunarHour{}.FromYmdHms(o.year, o.month, o.day, 0, 0, 0)
	l = append(l, *t)
	for i := 0; i < 24; i += 2 {
		t, _ = LunarHour{}.FromYmdHms(o.year, o.month, o.day, i+1, 0, 0)
		l = append(l, *t)
	}
	return l
}

// GetGods 神煞列表(吉神宜趋，凶神宜忌)
func (o LunarDay) GetGods() ([]God, error) {
	return o.GetSixtyCycleDay().GetGods()
}

// GetRecommends 宜
func (o LunarDay) GetRecommends() ([]Taboo, error) {
	return o.GetSixtyCycleDay().GetRecommends()
}

// GetAvoids 忌
func (o LunarDay) GetAvoids() ([]Taboo, error) {
	return o.GetSixtyCycleDay().GetAvoids()
}

// GetMinorRen 小六壬
func (o LunarDay) GetMinorRen() MinorRen {
	return o.GetLunarMonth().GetMinorRen().Next(o.day - 1)
}

// GetThreePillars 三柱
func (o LunarDay) GetThreePillars() ThreePillars {
	return o.GetSixtyCycleDay().GetThreePillars()
}
