package tyme

import (
	"math"
)

var SolarTermNames = []string{"冬至", "小寒", "大寒", "立春", "雨水", "惊蛰", "春分", "清明", "谷雨", "立夏", "小满", "芒种", "夏至", "小暑", "大暑", "立秋", "处暑", "白露", "秋分", "寒露", "霜降", "立冬", "小雪", "大雪"}

// SolarTerm 节气
type SolarTerm struct {
	LoopTyme
	// 年
	year int
	// 粗略的儒略日
	cursoryJulianDay float64
}

func (SolarTerm) FromIndex(year int, index int) SolarTerm {
	parent := LoopTyme{}.FromIndex(SolarTermNames, index)
	size := parent.GetSize()
	y := (year*size + index) / size
	o := SolarTerm{parent, y, 0}
	o.cursoryJulianDay = o.initByYear(y, parent.GetIndex())
	return o
}

func (SolarTerm) FromName(year int, name string) (*SolarTerm, error) {
	p, err := LoopTyme{}.FromName(SolarTermNames, name)
	if err != nil {
		return nil, err
	}
	o := SolarTerm{*p, year, 0}
	o.cursoryJulianDay = o.initByYear(year, o.index)
	return &o, nil
}

func (o SolarTerm) initByYear(year int, offset int) float64 {
	jd := math.Floor(float64(year-2000)*365.2422 + 180)
	// 355是2000.12冬至，得到较靠近jd的冬至估计值
	w := math.Floor((jd-355+183)/365.2422)*365.2422 + 355
	if CalcQi(w) > jd {
		w -= 365.2422
	}
	return CalcQi(w + 15.2184*float64(offset))
}

func (o SolarTerm) GetName() string {
	return SolarTermNames[o.index]
}

func (o SolarTerm) String() string {
	return o.GetName()
}

func (o SolarTerm) Next(n int) SolarTerm {
	size := o.GetSize()
	i := o.index + n
	return SolarTerm{}.FromIndex((o.year*size+i)/size, o.indexOf(i))
}

// IsJie 是否节令
func (o SolarTerm) IsJie() bool {
	return o.index%2 == 1
}

// IsQi 是否气令
func (o SolarTerm) IsQi() bool {
	return o.index%2 == 0
}

// GetJulianDay 儒略日
func (o SolarTerm) GetJulianDay() JulianDay {
	return JulianDay{}.FromJulianDay(QiAccurate2(o.cursoryJulianDay) + J2000)
}

// GetYear 年
func (o SolarTerm) GetYear() int {
	return o.year
}

// GetCursoryJulianDay 粗略的儒略日
func (o SolarTerm) GetCursoryJulianDay() float64 {
	return o.cursoryJulianDay
}
