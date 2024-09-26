package tyme

import "math"

// China95ChildLimitProvider 元亨利贞的童限计算
type China95ChildLimitProvider struct {
	ChildLimitProvider
}

func (o China95ChildLimitProvider) GetInfo(birthTime SolarTime, term SolarTerm) ChildLimitInfo {
	// 出生时刻和节令时刻相差的分钟数
	minutes := int(math.Abs(float64(term.GetJulianDay().GetSolarTime().Subtract(birthTime)))) / 60
	year := minutes / 4320
	minutes %= 4320
	month := minutes / 360
	minutes %= 360
	day := minutes / 12

	birthday := birthTime.GetSolarDay()
	sm, _ := SolarMonth{}.FromYm(birthday.GetYear()+year, birthday.GetMonth())
	sm = sm.Next(month)

	d := birthday.GetDay() + day
	dc := sm.GetDayCount()
	for d > dc {
		d -= dc
		sm = sm.Next(1)
		dc = sm.GetDayCount()
	}

	t, _ := SolarTime{}.FromYmdHms(sm.GetYear(), sm.GetMonth(), d, birthTime.GetHour(), birthTime.GetMinute(), birthTime.GetSecond())
	return ChildLimitInfo{}.New(birthTime, *t, year, month, day, 0, 0)
}
