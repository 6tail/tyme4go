package tyme

import (
	"fmt"
)

var HijriDayNames = []string{"1日", "2日", "3日", "4日", "5日", "6日", "7日", "8日", "9日", "10日", "11日", "12日", "13日", "14日", "15日", "16日", "17日", "18日", "19日", "20日", "21日", "22日", "23日", "24日", "25日", "26日", "27日", "28日", "29日", "30日"}

// HijriDay 回历日
type HijriDay struct {
	DayUnit
}

func (HijriDay) Validate(year int, month int, day int) error {
	m, err := HijriMonth{}.FromYm(year, month)
	if err != nil {
		return err
	}
	if day < 1 || day > m.GetDayCount() {
		return fmt.Errorf("illegal hijri day: %d-%d-%d", year, month, day)
	}
	return nil
}

func (HijriDay) FromYmd(year int, month int, day int) (*HijriDay, error) {
	err := HijriDay{}.Validate(year, month, day)
	if err != nil {
		return nil, err
	}
	return &HijriDay{
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

// GetHijriMonth 回历月
func (o HijriDay) GetHijriMonth() HijriMonth {
	m, _ := HijriMonth{}.FromYm(o.year, o.month)
	return *m
}

func (o HijriDay) GetName() string {
	return HijriDayNames[o.day-1]
}

func (o HijriDay) String() string {
	return fmt.Sprintf("%v%v", o.GetHijriMonth(), o.GetName())
}

func (o HijriDay) Next(n int) HijriDay {
	return o.GetSolarDay().Next(n).GetHijriDay()
}

// IsBefore 是否在指定回历日之前
func (o HijriDay) IsBefore(target HijriDay) bool {
	return o.GetCompareIndex() < target.GetCompareIndex()
}

// IsAfter 是否在指定回历日之后
func (o HijriDay) IsAfter(target HijriDay) bool {
	return o.GetCompareIndex() > target.GetCompareIndex()
}

// GetIndexInYear 位于当年的索引
func (o HijriDay) GetIndexInYear() int {
	d, _ := HijriDay{}.FromYmd(o.year, 1, 1)
	return o.Subtract(*d)
}

// Subtract 回历日期相减，获得相差天数
func (o HijriDay) Subtract(target HijriDay) int {
	return int(o.GetJulianDay().Subtract(target.GetJulianDay()))
}

// GetJulianDay 儒略日
func (o HijriDay) GetJulianDay() JulianDay {
	return JulianDay{}.FromJulianDay(float64(o.FloorDiv(11*o.year+3, 30) + 354*o.year + 30*o.month - o.FloorDiv(o.month-1, 2) + o.day + 1948055))
}

// GetSolarDay 公历日
func (o HijriDay) GetSolarDay() SolarDay {
	d, _ := SolarDay{}.FromYmd(622, 7, 16)
	h, _ := HijriDay{}.FromYmd(1, 1, 1)
	return d.Next(o.Subtract(*h))
}

func (o HijriDay) Equals(target HijriDay) bool {
	return o.String() == target.String()
}
