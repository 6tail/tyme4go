package tyme

// LunarSect1ChildLimitProvider Lunar的流派1童限计算（按天数和时辰数计算，3天1年，1天4个月，1时辰10天）
type LunarSect1ChildLimitProvider struct {
	AbstractChildLimitProvider
}

func (o LunarSect1ChildLimitProvider) GetInfo(birthTime SolarTime, term SolarTerm) ChildLimitInfo {
	termTime := term.GetJulianDay().GetSolarTime()
	end := termTime
	start := birthTime
	if birthTime.IsAfter(termTime) {
		end = birthTime
		start = termTime
	}
	endTimeZhiIndex := 11
	if end.GetHour() != 23 {
		endTimeZhiIndex = end.GetLunarHour().GetIndexInDay()
	}
	startTimeZhiIndex := 11
	if start.GetHour() != 23 {
		startTimeZhiIndex = start.GetLunarHour().GetIndexInDay()
	}
	// 时辰差
	hourDiff := endTimeZhiIndex - startTimeZhiIndex
	// 天数差
	dayDiff := end.GetSolarDay().Subtract(start.GetSolarDay())
	if hourDiff < 0 {
		hourDiff += 12
		dayDiff--
	}
	monthDiff := hourDiff * 10 / 30
	month := dayDiff*4 + monthDiff
	day := hourDiff*10 - monthDiff*30
	year := month / 12
	month = month - year*12

	return o.next(birthTime, year, month, day, 0, 0, 0)
}
