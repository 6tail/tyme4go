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
