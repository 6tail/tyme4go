package tyme

import "math"

// China95ChildLimitProvider 元亨利贞的童限计算
type China95ChildLimitProvider struct {
	AbstractChildLimitProvider
}

func (o China95ChildLimitProvider) GetInfo(birthTime SolarTime, term SolarTerm) ChildLimitInfo {
	// 出生时刻和节令时刻相差的分钟数
	minutes := int(math.Abs(float64(term.GetJulianDay().GetSolarTime().Subtract(birthTime)))) / 60
	year := minutes / 4320
	minutes %= 4320
	month := minutes / 360
	minutes %= 360
	day := minutes / 12

	return o.next(birthTime, year, month, day, 0, 0, 0)
}
