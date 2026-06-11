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

// GetCompareIndex 用于比较大小的索引
func (o MonthUnit) GetCompareIndex() int {
	m := 0
	if o.month > 0 {
		m = o.month * 2
	} else {
		m = -o.month*2 + 1
	}
	return o.YearUnit.GetCompareIndex() + m*100
}
