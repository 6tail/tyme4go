package tyme

import "math"

// LunarSect2ChildLimitProvider Lunar的流派2童限计算（按分钟数计算）
type LunarSect2ChildLimitProvider struct {
	AbstractChildLimitProvider
}

func (o LunarSect2ChildLimitProvider) GetInfo(birthTime SolarTime, term SolarTerm) ChildLimitInfo {
	// 出生时刻和节令时刻相差的分钟数
	minutes := int(math.Abs(float64(term.GetJulianDay().GetSolarTime().Subtract(birthTime)))) / 60
	year := minutes / 4320
	minutes %= 4320
	month := minutes / 360
	minutes %= 360
	day := minutes / 12
	minutes %= 12
	hour := minutes * 2

	return o.next(birthTime, year, month, day, hour, 0, 0)
}
