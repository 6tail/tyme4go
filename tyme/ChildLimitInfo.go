package tyme

// ChildLimitInfo 童限信息
type ChildLimitInfo struct {
	// 开始(即出生)的公历时刻
	startTime SolarTime
	// 结束(即开始起运)的公历时刻
	endTime SolarTime
	// 年数
	yearCount int
	// 月数
	monthCount int
	// 日数
	dayCount int
	// 小时数
	hourCount int
	// 分钟数
	minuteCount int
}

func (ChildLimitInfo) New(startTime SolarTime, endTime SolarTime, yearCount int, monthCount int, dayCount int, hourCount int, minuteCount int) ChildLimitInfo {
	return ChildLimitInfo{
		startTime:   startTime,
		endTime:     endTime,
		yearCount:   yearCount,
		monthCount:  monthCount,
		dayCount:    dayCount,
		hourCount:   hourCount,
		minuteCount: minuteCount,
	}
}

// GetStartTime 开始(即出生)的公历时刻
func (o ChildLimitInfo) GetStartTime() SolarTime {
	return o.startTime
}

// GetEndTime 结束(即开始起运)的公历时刻
func (o ChildLimitInfo) GetEndTime() SolarTime {
	return o.endTime
}

// GetYearCount 年数
func (o ChildLimitInfo) GetYearCount() int {
	return o.yearCount
}

// GetMonthCount 月数
func (o ChildLimitInfo) GetMonthCount() int {
	return o.monthCount
}

// GetDayCount 日数
func (o ChildLimitInfo) GetDayCount() int {
	return o.dayCount
}

// GetHourCount 小时数
func (o ChildLimitInfo) GetHourCount() int {
	return o.hourCount
}

// GetMinuteCount 分钟数
func (o ChildLimitInfo) GetMinuteCount() int {
	return o.minuteCount
}
