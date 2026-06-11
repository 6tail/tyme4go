package tyme

// YearUnit 年
type YearUnit struct {
	AbstractTyme

	// 年
	year int
}

// GetYear 年
func (o YearUnit) GetYear() int {
	return o.year
}

// GetCompareIndex 用于比较大小的索引
func (o YearUnit) GetCompareIndex() int {
	return o.year * 10000
}
