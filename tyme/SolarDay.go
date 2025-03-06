package tyme

import (
	"fmt"
	"math"
	"strconv"
)

var SolarDayNames = []string{"1日", "2日", "3日", "4日", "5日", "6日", "7日", "8日", "9日", "10日", "11日", "12日", "13日", "14日", "15日", "16日", "17日", "18日", "19日", "20日", "21日", "22日", "23日", "24日", "25日", "26日", "27日", "28日", "29日", "30日", "31日"}

// SolarDay 公历日
type SolarDay struct {
	AbstractTyme
	// 公历月
	month SolarMonth
	// 日
	day int
}

func (SolarDay) FromYmd(year int, month int, day int) (*SolarDay, error) {
	if day < 1 {
		return nil, fmt.Errorf(fmt.Sprintf("illegal solar day: %d-%d-%d", year, month, day))
	}
	m, err := SolarMonth{}.FromYm(year, month)
	if err != nil {
		return nil, err
	}
	if 1582 == year && 10 == month {
		if (day > 4 && day < 15) || day > 31 {
			return nil, fmt.Errorf(fmt.Sprintf("illegal solar day: %d-%d-%d", year, month, day))
		}
	} else if day > m.GetDayCount() {
		return nil, fmt.Errorf(fmt.Sprintf("illegal solar day: %d-%d-%d", year, month, day))
	}
	return &SolarDay{
		month: *m,
		day:   day,
	}, nil
}

// GetSolarMonth 公历月
func (o SolarDay) GetSolarMonth() SolarMonth {
	return o.month
}

// GetYear 年
func (o SolarDay) GetYear() int {
	return o.month.GetYear()
}

// GetMonth 月
func (o SolarDay) GetMonth() int {
	return o.month.GetMonth()
}

// GetDay 日
func (o SolarDay) GetDay() int {
	return o.day
}

// GetWeek 星期
func (o SolarDay) GetWeek() Week {
	return o.GetJulianDay().GetWeek()
}

// GetConstellation 星座
func (o SolarDay) GetConstellation() Constellation {
	index := 8
	y := o.GetMonth()*100 + o.day
	if y > 1221 || y < 120 {
		index = 9
	} else if y < 219 {
		index = 10
	} else if y < 321 {
		index = 11
	} else if y < 420 {
		index = 0
	} else if y < 521 {
		index = 1
	} else if y < 622 {
		index = 2
	} else if y < 723 {
		index = 3
	} else if y < 823 {
		index = 4
	} else if y < 923 {
		index = 5
	} else if y < 1024 {
		index = 6
	} else if y < 1123 {
		index = 7
	}
	return Constellation{}.FromIndex(index)
}

func (o SolarDay) GetName() string {
	return SolarDayNames[o.day-1]
}

func (o SolarDay) String() string {
	return fmt.Sprintf("%v%v", o.month, o.GetName())
}

func (o SolarDay) Next(n int) SolarDay {
	return o.GetJulianDay().Next(n).GetSolarDay()
}

// IsBefore 是否在指定公历日之前
func (o SolarDay) IsBefore(target SolarDay) bool {
	aYear := o.GetYear()
	bYear := target.GetYear()
	if aYear != bYear {
		return aYear < bYear
	}
	aMonth := o.GetMonth()
	bMonth := target.GetMonth()
	if aMonth != bMonth {
		return aMonth < bMonth
	}
	return o.day < target.GetDay()
}

// IsAfter 是否在指定公历日之后
func (o SolarDay) IsAfter(target SolarDay) bool {
	aYear := o.GetYear()
	bYear := target.GetYear()
	if aYear != bYear {
		return aYear > bYear
	}
	aMonth := o.GetMonth()
	bMonth := target.GetMonth()
	if aMonth != bMonth {
		return aMonth > bMonth
	}
	return o.day > target.GetDay()
}

// GetTerm 节气
func (o SolarDay) GetTerm() SolarTerm {
	return o.GetTermDay().GetSolarTerm()
}

// GetTermDay 节气第几天
func (o SolarDay) GetTermDay() SolarTermDay {
	y := o.GetYear()
	i := o.GetMonth() * 2
	if i == 24 {
		y += 1
		i = 0
	}
	term := SolarTerm{}.FromIndex(y, i)
	day := term.GetJulianDay().GetSolarDay()
	for o.IsBefore(day) {
		term = term.Next(-1)
		day = term.GetJulianDay().GetSolarDay()
	}
	return SolarTermDay{}.New(term, o.Subtract(day))
}

// GetSolarWeek 公历周
// 参数 start 起始星期，1234560分别代表星期一至星期天
func (o SolarDay) GetSolarWeek(start int) SolarWeek {
	y := o.GetYear()
	m := o.GetMonth()
	d, _ := SolarDay{}.FromYmd(y, m, 1)
	w, _ := SolarWeek{}.FromYm(y, m, int(math.Ceil(float64(o.day+d.GetWeek().Next(-start).GetIndex())/7))-1, start)
	return *w
}

// GetPhenologyDay 七十二候
func (o SolarDay) GetPhenologyDay() PhenologyDay {
	term := o.GetTerm()
	dayIndex := o.Subtract(term.GetJulianDay().GetSolarDay())
	index := dayIndex / 5
	if index > 2 {
		index = 2
	}
	dayIndex -= index * 5
	return PhenologyDay{}.New(Phenology{}.FromIndex(term.GetIndex()*3+index), dayIndex)
}

// GetDogDay 三伏天
func (o SolarDay) GetDogDay() *DogDay {
	// 夏至
	xiaZhi := SolarTerm{}.FromIndex(o.GetYear(), 12)
	// 第1个庚日
	start := xiaZhi.GetJulianDay().GetSolarDay()
	// 第3个庚日，即初伏第1天
	start = start.Next(start.GetLunarDay().GetSixtyCycle().GetHeavenStem().StepsTo(6) + 20)
	days := o.Subtract(start)
	// 初伏以前
	if days < 0 {
		return nil
	}
	if days < 10 {
		d := DogDay{}.New(Dog{}.FromIndex(0), days)
		return &d
	}
	// 第4个庚日，中伏第1天
	start = start.Next(10)
	days = o.Subtract(start)
	if days < 10 {
		d := DogDay{}.New(Dog{}.FromIndex(1), days)
		return &d
	}
	// 第5个庚日，中伏第11天或末伏第1天
	start = start.Next(10)
	days = o.Subtract(start)
	// 立秋
	if xiaZhi.Next(3).GetJulianDay().GetSolarDay().IsAfter(start) {
		if days < 10 {
			d := DogDay{}.New(Dog{}.FromIndex(1), days+10)
			return &d
		}
		start = start.Next(10)
		days = o.Subtract(start)
	}
	if days >= 10 {
		return nil
	}
	d := DogDay{}.New(Dog{}.FromIndex(2), days)
	return &d
}

// GetNineDay 数九天
func (o SolarDay) GetNineDay() *NineDay {
	year := o.GetYear()
	start := SolarTerm{}.FromIndex(year+1, 0).GetJulianDay().GetSolarDay()
	if o.IsBefore(start) {
		start = SolarTerm{}.FromIndex(year, 0).GetJulianDay().GetSolarDay()
	}
	end := start.Next(81)
	if o.IsBefore(start) || !o.IsBefore(end) {
		return nil
	}
	days := o.Subtract(start)
	d := NineDay{}.New(Nine{}.FromIndex(days/9), days%9)
	return &d
}

// GetHideHeavenStemDay 人元司令分野
func (o SolarDay) GetHideHeavenStemDay() HideHeavenStemDay {
	dayCounts := []int{3, 5, 7, 9, 10, 30}
	term := o.GetTerm()
	if term.IsQi() {
		term = term.Next(-1)
	}
	dayIndex := o.Subtract(term.GetJulianDay().GetSolarDay())
	startIndex := (term.GetIndex() - 1) * 3
	data := "93705542220504xx1513904541632524533533105544806564xx7573304542018584xx95"[startIndex : startIndex+6]
	days := 0
	heavenStemIndex := 0
	typeIndex := 0
	for typeIndex < 3 {
		i := typeIndex * 2
		d := data[i : i+1]
		count := 0
		if d != "x" {
			heavenStemIndex, _ = strconv.Atoi(d)
			dayCountIndex, _ := strconv.Atoi(data[i+1 : i+2])
			count = dayCounts[dayCountIndex]
			days += count
		}
		if dayIndex <= days {
			dayIndex -= days - count
			break
		}
		typeIndex++
	}
	hideHeavenStemType := RESIDUAL
	if typeIndex == 1 {
		hideHeavenStemType = MIDDLE
	} else if typeIndex == 2 {
		hideHeavenStemType = MAIN
	}
	return HideHeavenStemDay{}.New(HideHeavenStem{}.FromIndex(heavenStemIndex, hideHeavenStemType), dayIndex)
}

// GetPlumRainDay 梅雨天（芒种后的第1个丙日入梅，小暑后的第1个未日出梅）
func (o SolarDay) GetPlumRainDay() *PlumRainDay {
	// 芒种
	grainInEar := SolarTerm{}.FromIndex(o.GetYear(), 11)
	start := grainInEar.GetJulianDay().GetSolarDay()
	// 芒种后的第1个丙日
	start = start.Next(start.GetLunarDay().GetSixtyCycle().GetHeavenStem().StepsTo(2))

	// 小暑
	slightHeat := grainInEar.Next(2)
	end := slightHeat.GetJulianDay().GetSolarDay()
	// 小暑后的第1个未日
	end = end.Next(end.GetLunarDay().GetSixtyCycle().GetEarthBranch().StepsTo(7))

	if o.IsBefore(start) || o.IsAfter(end) {
		return nil
	}
	if o.Equals(end) {
		t := PlumRainDay{}.New(PlumRain{}.FromIndex(1), 0)
		return &t
	}
	t := PlumRainDay{}.New(PlumRain{}.FromIndex(0), o.Subtract(start))
	return &t
}

// GetIndexInYear 位于当年的索引
func (o SolarDay) GetIndexInYear() int {
	d, _ := SolarDay{}.FromYmd(o.GetYear(), 1, 1)
	return o.Subtract(*d)
}

// Subtract 公历日期相减，获得相差天数
func (o SolarDay) Subtract(target SolarDay) int {
	return int(o.GetJulianDay().Subtract(target.GetJulianDay()))
}

// GetJulianDay 儒略日
func (o SolarDay) GetJulianDay() JulianDay {
	return JulianDay{}.FromYmdHms(o.GetYear(), o.GetMonth(), o.day, 0, 0, 0)
}

// GetLunarDay 农历日
func (o SolarDay) GetLunarDay() LunarDay {
	t, _ := LunarMonth{}.FromYm(o.GetYear(), o.GetMonth())
	m := *t
	days := o.Subtract(m.GetFirstJulianDay().GetSolarDay())
	for days < 0 {
		m = m.Next(-1)
		days += m.GetDayCount()
	}
	d, _ := LunarDay{}.FromYmd(m.GetYear(), m.GetMonthWithLeap(), days+1)
	return *d
}

// GetLegalHoliday 法定假日，如果当天不是法定假日，返回nil
func (o SolarDay) GetLegalHoliday() *LegalHoliday {
	f, _ := LegalHoliday{}.FromYmd(o.GetYear(), o.GetMonth(), o.day)
	return f
}

// GetFestival 公历现代节日，如果当天不是公历现代节日，返回nil
func (o SolarDay) GetFestival() *SolarFestival {
	f, _ := SolarFestival{}.FromYmd(o.GetYear(), o.GetMonth(), o.day)
	return f
}

func (o SolarDay) Equals(target SolarDay) bool {
	return o.String() == target.String()
}
