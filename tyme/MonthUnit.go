package tyme

// MonthUnit 月
type MonthUnit struct {
	YearUnit

	// 年
	month int
}

// GetMonth 月
func (o MonthUnit) GetMonth() int {
	return o.month
}
