package tyme

// DayUnit 日
type DayUnit struct {
	MonthUnit

	// 日
	day int
}

// GetDay 日
func (o DayUnit) GetDay() int {
	return o.day
}

// GetCompareIndex 用于比较大小的索引
func (o DayUnit) GetCompareIndex() int {
	return o.MonthUnit.GetCompareIndex() + o.day
}
