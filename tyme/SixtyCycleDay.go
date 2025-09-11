package tyme

import (
	"fmt"
	"math"
)

// SixtyCycleDay 干支日（立春换年，节令换月）
type SixtyCycleDay struct {
	AbstractTyme
	// 公历日
	solarDay SolarDay
	// 干支月
	month SixtyCycleMonth
	// 日柱
	day SixtyCycle
}

func (SixtyCycleDay) New(solarDay SolarDay, month SixtyCycleMonth, day SixtyCycle) SixtyCycleDay {
	return SixtyCycleDay{
		solarDay: solarDay,
		month:    month,
		day:      day,
	}
}

func (SixtyCycleDay) FromSolarDay(solarDay SolarDay) SixtyCycleDay {
	solarYear := solarDay.GetYear()
	springSolarDay := SolarTerm{}.FromIndex(solarYear, 3).GetJulianDay().GetSolarDay()
	lunarDay := solarDay.GetLunarDay()
	lunarYear := lunarDay.GetLunarMonth().GetLunarYear()
	if lunarYear.GetYear() == solarYear {
		if solarDay.IsBefore(springSolarDay) {
			lunarYear = lunarYear.Next(-1)
		}
	} else if lunarYear.GetYear() < solarYear {
		if !solarDay.IsBefore(springSolarDay) {
			lunarYear = lunarYear.Next(1)
		}
	}
	term := solarDay.GetTerm()
	index := term.GetIndex() - 3
	if index < 0 && term.GetJulianDay().GetSolarDay().IsAfter(springSolarDay) {
		index += 24
	}
	y, _ := SixtyCycleYear{}.FromYear(lunarYear.GetYear())
	m, _ := LunarMonth{}.FromYm(solarYear, 1)
	return SixtyCycleDay{
		solarDay: solarDay,
		month:    SixtyCycleMonth{}.New(*y, m.GetSixtyCycle().Next(int(math.Floor(float64(index)*0.5)))),
		day:      lunarDay.GetSixtyCycle(),
	}
}

// GetSolarDay 公历日
func (o SixtyCycleDay) GetSolarDay() SolarDay {
	return o.solarDay
}

// GetSixtyCycleMonth 干支月
func (o SixtyCycleDay) GetSixtyCycleMonth() SixtyCycleMonth {
	return o.month
}

// GetYear 年柱
func (o SixtyCycleDay) GetYear() SixtyCycle {
	return o.month.GetYear()
}

// GetMonth 月柱
func (o SixtyCycleDay) GetMonth() SixtyCycle {
	return o.month.GetSixtyCycle()
}

// GetSixtyCycle 干支
func (o SixtyCycleDay) GetSixtyCycle() SixtyCycle {
	return o.day
}

func (o SixtyCycleDay) GetName() string {
	return o.day.String() + "日"
}

func (o SixtyCycleDay) String() string {
	return fmt.Sprintf("%v%v", o.month, o.GetName())
}

func (o SixtyCycleDay) Next(n int) SixtyCycleDay {
	return SixtyCycleDay{}.FromSolarDay(o.solarDay.Next(n))
}

// GetDuty 建除十二值神
func (o SixtyCycleDay) GetDuty() Duty {
	return Duty{}.FromIndex(o.day.GetEarthBranch().GetIndex() - o.GetMonth().GetEarthBranch().GetIndex())
}

// GetTwelveStar 黄道黑道十二神
func (o SixtyCycleDay) GetTwelveStar() TwelveStar {
	return TwelveStar{}.FromIndex(o.day.GetEarthBranch().GetIndex() + (8-o.GetMonth().GetEarthBranch().GetIndex()%6)*2)
}

// GetNineStar 九星
func (o SixtyCycleDay) GetNineStar() NineStar {
	d := o.solarDay
	dongZhi := SolarTerm{}.FromIndex(d.GetYear(), 0)
	dongZhiSolar := dongZhi.GetJulianDay().GetSolarDay()
	xiaZhiSolar := dongZhi.Next(12).GetJulianDay().GetSolarDay()
	dongZhiSolar2 := dongZhi.Next(24).GetJulianDay().GetSolarDay()
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
	if !d.IsBefore(solarShunBai) && d.IsBefore(solarNiZi) {
		offset = d.Subtract(solarShunBai)
	} else if !d.IsBefore(solarNiZi) && d.IsBefore(solarShunBai2) {
		offset = 8 - d.Subtract(solarNiZi)
	} else if !d.IsBefore(solarShunBai2) {
		offset = d.Subtract(solarShunBai2)
	} else if d.IsBefore(solarShunBai) {
		offset = 8 + solarShunBai.Subtract(d)
	}
	return NineStar{}.FromIndex(offset)
}

// GetJupiterDirection 太岁方位
func (o SixtyCycleDay) GetJupiterDirection() Direction {
	index := o.day.GetIndex()
	if index%12 < 6 {
		return Element{}.FromIndex(index / 12).GetDirection()
	}
	return o.month.GetSixtyCycleYear().GetJupiterDirection()
}

// GetFetusDay 逐日胎神
func (o SixtyCycleDay) GetFetusDay() FetusDay {
	return FetusDay{}.FromSixtyCycleDay(o)
}

// GetTwentyEightStar 二十八宿
func (o SixtyCycleDay) GetTwentyEightStar() TwentyEightStar {
	return TwentyEightStar{}.FromIndex([]int{10, 18, 26, 6, 14, 22, 2}[o.solarDay.GetWeek().GetIndex()]).Next(-7 * o.day.GetEarthBranch().GetIndex())
}

func (o SixtyCycleDay) Equals(target SixtyCycleDay) bool {
	return o.String() == target.String()
}

// GetHours 当天的干支时辰列表
func (o SixtyCycleDay) GetHours() []SixtyCycleHour {
	var l []SixtyCycleHour
	d := o.solarDay.Next(-1)
	t, _ := SolarTime{}.FromYmdHms(d.GetYear(), d.GetMonth(), d.GetDay(), 23, 0, 0)
	h := SixtyCycleHour{}.FromSolarTime(*t)
	l = append(l, h)
	for i := 0; i < 11; i++ {
		h = h.Next(7200)
		l = append(l, h)
	}
	return l
}

// GetGods 神煞列表(吉神宜趋，凶神宜忌)
func (o SixtyCycleDay) GetGods() ([]God, error) {
	return God{}.GetDayGods(o.GetMonth(), o.day)
}

// GetRecommends 宜
func (o SixtyCycleDay) GetRecommends() ([]Taboo, error) {
	return Taboo{}.GetDayRecommends(o.GetMonth(), o.day)
}

// GetAvoids 忌
func (o SixtyCycleDay) GetAvoids() ([]Taboo, error) {
	return Taboo{}.GetDayAvoids(o.GetMonth(), o.day)
}
