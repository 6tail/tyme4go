package tyme

import (
	"fmt"
	"math"
)

var LunarDayNames = []string{"初一", "初二", "初三", "初四", "初五", "初六", "初七", "初八", "初九", "初十", "十一", "十二", "十三", "十四", "十五", "十六", "十七", "十八", "十九", "二十", "廿一", "廿二", "廿三", "廿四", "廿五", "廿六", "廿七", "廿八", "廿九", "三十"}

// LunarDay 农历日
type LunarDay struct {
	AbstractTyme
	// 农历月
	month LunarMonth
	// 日
	day int
}

func (LunarDay) FromYmd(year int, month int, day int) (*LunarDay, error) {
	m, err := LunarMonth{}.FromYm(year, month)
	if err != nil {
		return nil, err
	}
	if day < 1 || day > m.GetDayCount() {
		return nil, fmt.Errorf(fmt.Sprintf("illegal day %d in %v", day, m))
	}
	return &LunarDay{
		month: *m,
		day:   day,
	}, nil
}

// GetLunarMonth 农历月
func (o LunarDay) GetLunarMonth() LunarMonth {
	return o.month
}

// GetYear 年
func (o LunarDay) GetYear() int {
	return o.month.GetYear()
}

// GetMonth 月
func (o LunarDay) GetMonth() int {
	return o.month.GetMonthWithLeap()
}

// GetDay 日
func (o LunarDay) GetDay() int {
	return o.day
}

// GetWeek 星期
func (o LunarDay) GetWeek() Week {
	return o.GetSolarDay().GetWeek()
}

func (o LunarDay) GetName() string {
	return LunarDayNames[o.day-1]
}

func (o LunarDay) String() string {
	return fmt.Sprintf("%v%v", o.month, o.GetName())
}

func (o LunarDay) Next(n int) LunarDay {
	if 0 == n {
		d, _ := LunarDay{}.FromYmd(o.GetYear(), o.GetMonth(), o.day)
		return *d
	}
	return o.GetSolarDay().Next(n).GetLunarDay()
}

// IsBefore 是否在指定农历日之前
func (o LunarDay) IsBefore(target LunarDay) bool {
	aYear := o.GetYear()
	bYear := target.GetYear()
	if aYear != bYear {
		return aYear < bYear
	}
	aMonth := o.GetMonth()
	bMonth := target.GetMonth()
	if aMonth != bMonth {
		if aMonth < 0 {
			aMonth = -aMonth
		}
		if bMonth < 0 {
			bMonth = -bMonth
		}
		return aMonth < bMonth
	}
	return o.day < target.GetDay()
}

// IsAfter 是否在指定农历日之后
func (o LunarDay) IsAfter(target LunarDay) bool {
	aYear := o.GetYear()
	bYear := target.GetYear()
	if aYear != bYear {
		return aYear > bYear
	}
	aMonth := o.GetMonth()
	bMonth := target.GetMonth()
	if aMonth != bMonth {
		if aMonth < 0 {
			aMonth = -aMonth
		}
		if bMonth < 0 {
			bMonth = -bMonth
		}
		return aMonth > bMonth
	}
	return o.day > target.GetDay()
}

// GetYearSixtyCycle 当天的年干支（立春换）
func (o LunarDay) GetYearSixtyCycle() SixtyCycle {
	solarDay := o.GetSolarDay()
	solarYear := solarDay.GetYear()
	springSolarDay := SolarTerm{}.FromIndex(solarYear, 3).GetJulianDay().GetSolarDay()
	lunarYear := o.month.GetLunarYear()
	year := lunarYear.GetYear()
	sixtyCycle := lunarYear.GetSixtyCycle()
	if year == solarYear {
		if solarDay.IsBefore(springSolarDay) {
			sixtyCycle = sixtyCycle.Next(-1)
		}
	} else if year < solarYear {
		if !solarDay.IsBefore(springSolarDay) {
			sixtyCycle = sixtyCycle.Next(1)
		}
	}
	return sixtyCycle
}

// GetMonthSixtyCycle 当天的月干支（节气换）
func (o LunarDay) GetMonthSixtyCycle() SixtyCycle {
	solarDay := o.GetSolarDay()
	year := solarDay.GetYear()
	term := solarDay.GetTerm()
	index := term.GetIndex() - 3
	if index < 0 && term.GetJulianDay().GetSolarDay().IsAfter(SolarTerm{}.FromIndex(year, 3).GetJulianDay().GetSolarDay()) {
		index += 24
	}
	m, _ := LunarMonth{}.FromYm(year, 1)
	return m.GetSixtyCycle().Next(int(math.Floor(float64(index) / 2.0)))
}

// GetSixtyCycle 干支
func (o LunarDay) GetSixtyCycle() SixtyCycle {
	offset := int(o.month.GetFirstJulianDay().Next(o.day - 12).GetDay())
	t, _ := SixtyCycle{}.FromName(HeavenStem{}.FromIndex(offset).GetName() + EarthBranch{}.FromIndex(offset).GetName())
	return *t
}

// GetDuty 建除十二值神
func (o LunarDay) GetDuty() Duty {
	return Duty{}.FromIndex(o.GetSixtyCycle().GetEarthBranch().GetIndex() - o.GetMonthSixtyCycle().GetEarthBranch().GetIndex())
}

// GetTwelveStar 黄道黑道十二神
func (o LunarDay) GetTwelveStar() TwelveStar {
	return TwelveStar{}.FromIndex(o.GetSixtyCycle().GetEarthBranch().GetIndex() + (8-o.GetMonthSixtyCycle().GetEarthBranch().GetIndex()%6)*2)
}

// GetNineStar 九星
func (o LunarDay) GetNineStar() NineStar {
	solar := o.GetSolarDay()
	dongZhi := SolarTerm{}.FromIndex(solar.GetYear(), 0)
	xiaZhi := dongZhi.Next(12)
	dongZhi2 := dongZhi.Next(24)
	dongZhiSolar := dongZhi.GetJulianDay().GetSolarDay()
	xiaZhiSolar := xiaZhi.GetJulianDay().GetSolarDay()
	dongZhiSolar2 := dongZhi2.GetJulianDay().GetSolarDay()
	dongZhiIndex := dongZhiSolar.GetLunarDay().GetSixtyCycle().GetIndex()
	xiaZhiIndex := xiaZhiSolar.GetLunarDay().GetSixtyCycle().GetIndex()
	dongZhiIndex2 := dongZhiSolar2.GetLunarDay().GetSixtyCycle().GetIndex()
	index := -dongZhiIndex
	if dongZhiIndex > 29 {
		index = 60 - dongZhiIndex
	}
	solarShunBai := dongZhiSolar.Next(index)
	index = -dongZhiIndex2
	if dongZhiIndex2 > 29 {
		index = 60 - dongZhiIndex2
	}
	solarShunBai2 := dongZhiSolar2.Next(index)
	index = -xiaZhiIndex
	if xiaZhiIndex > 29 {
		index = 60 - xiaZhiIndex
	}
	solarNiZi := xiaZhiSolar.Next(index)
	offset := 0
	if !solar.IsBefore(solarShunBai) && solar.IsBefore(solarNiZi) {
		offset = solar.Subtract(solarShunBai)
	} else if !solar.IsBefore(solarNiZi) && solar.IsBefore(solarShunBai2) {
		offset = 8 - solar.Subtract(solarNiZi)
	} else if !solar.IsBefore(solarShunBai2) {
		offset = solar.Subtract(solarShunBai2)
	} else if solar.IsBefore(solarShunBai) {
		offset = 8 + solarShunBai.Subtract(solar)
	}
	return NineStar{}.FromIndex(offset)
}

// GetJupiterDirection 太岁方位
func (o LunarDay) GetJupiterDirection() Direction {
	index := o.GetSixtyCycle().GetIndex()
	if index%12 < 6 {
		return Element{}.FromIndex(index / 12).GetDirection()
	}
	return o.month.GetLunarYear().GetJupiterDirection()
}

// GetFetusDay 逐日胎神
func (o LunarDay) GetFetusDay() FetusDay {
	return FetusDay{}.FromLunarDay(o)
}

// GetPhase 月相
func (o LunarDay) GetPhase() Phase {
	return Phase{}.FromIndex(o.day - 1)
}

// GetSixStar 六曜
func (o LunarDay) GetSixStar() SixStar {
	return SixStar{}.FromIndex((o.month.GetMonth() + o.day - 2) % 6)
}

// GetSolarDay 公历日
func (o LunarDay) GetSolarDay() SolarDay {
	return o.month.GetFirstJulianDay().Next(o.day - 1).GetSolarDay()
}

// GetTwentyEightStar 二十八宿
func (o LunarDay) GetTwentyEightStar() TwentyEightStar {
	return TwentyEightStar{}.FromIndex([]int{10, 18, 26, 6, 14, 22, 2}[o.GetSolarDay().GetWeek().GetIndex()]).Next(-7 * o.GetSixtyCycle().GetEarthBranch().GetIndex())
}

// GetFestival 农历传统节日，如果当天不是农历传统节日，返回nil
func (o LunarDay) GetFestival() *LunarFestival {
	f, _ := LunarFestival{}.FromYmd(o.GetYear(), o.GetMonth(), o.day)
	return f
}

func (o LunarDay) Equals(target LunarDay) bool {
	return o.String() == target.String()
}

// GetHours 当天的时辰列表
func (o LunarDay) GetHours() []LunarHour {
	var l []LunarHour
	y := o.GetYear()
	m := o.GetMonth()
	t, _ := LunarHour{}.FromYmdHms(y, m, o.day, 0, 0, 0)
	l = append(l, *t)
	for i := 0; i < 24; i += 2 {
		t, _ = LunarHour{}.FromYmdHms(y, m, o.day, i+1, 0, 0)
		l = append(l, *t)
	}
	return l
}

// GetGods 神煞列表(吉神宜趋，凶神宜忌)
func (o LunarDay) GetGods() ([]God, error) {
	return God{}.GetDayGods(o.GetMonthSixtyCycle(), o.GetSixtyCycle())
}

// GetRecommends 宜
func (o LunarDay) GetRecommends() ([]Taboo, error) {
	return Taboo{}.GetDayRecommends(o.GetMonthSixtyCycle(), o.GetSixtyCycle())
}

// GetAvoids 忌
func (o LunarDay) GetAvoids() ([]Taboo, error) {
	return Taboo{}.GetDayAvoids(o.GetMonthSixtyCycle(), o.GetSixtyCycle())
}
