package tyme

import (
	"fmt"
	"strconv"
)

var WeekUnitNames = []string{"第一周", "第二周", "第三周", "第四周", "第五周", "第六周"}

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
		return fmt.Errorf("illegal week index: " + strconv.Itoa(index))
	}
	if start < 0 || start > 6 {
		return fmt.Errorf("illegal week start: " + strconv.Itoa(start))
	}
	return nil
}
