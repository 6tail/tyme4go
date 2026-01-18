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
