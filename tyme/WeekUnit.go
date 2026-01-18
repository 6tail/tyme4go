package tyme

import "fmt"

// WeekUnit 周
type WeekUnit struct {
	MonthUnit

	// 索引，0-5
	index int
	// 起始星期，1234560分别代表星期一至星期天
	start int
}

// GetIndex 索引，0-5
func (o WeekUnit) GetIndex() int {
	return o.index
}

// GetStart 起始星期
func (o WeekUnit) GetStart() Week {
	return Week{}.FromIndex(o.start)
}

func (WeekUnit) Validate(_ int, _ int, index int, start int) error {
	if index < 0 || index > 5 {
		return fmt.Errorf(fmt.Sprintf("illegal week index: %d", index))
	}
	if start < 0 || start > 6 {
		return fmt.Errorf(fmt.Sprintf("illegal week start: %d", start))
	}
	return nil
}
